package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/curio-research/keystone/keystone/ecs"
	"github.com/gorilla/websocket"
)

const (
	ReadBufferSize  = 1024
	WriteBufferSize = 1024
	StreamInterval  = 100 * time.Millisecond
)

// event types
const (
	Diff = "DIFF"
)

type StreamServer struct {
	// lock for data packets
	DataPacketsMutex sync.Mutex

	// list of data packets that need to be published by the websocket streamer to clients
	DataPackets UpdatePackets

	// a pool of connections
	Conns map[*websocket.Conn]bool // Connection Pool
}

type UpdatePackets []UpdatePacket

// update packet container a group of ecs updates with additional metadata
// that helps clients create the scene

type UpdatePacket struct {
	// array of ECS packets that need to be broadcasted to clients
	EcsUpdates []ecs.ECSUpdate `json:"ecsUpdates"`

	// id of package that corresponds with the HTTP requests's returned UUID (similar to a transaction hash)
	Uuid string `json:"uuid"`

	// timestamp
	Time int64 `json:"time"`

	// error message string returned from a request
	Message string `json:"message"`

	// metadata that helps clients determine what to do with state data
	ClientEvents ClientEvents `json:"clientEvents"`
}

type ClientEvents []ClientEvent

// used once per-request
type EventCtx struct {
	ClientEvents ClientEvents
}

// adds a client event to the event context
func (e *EventCtx) EmitEvent(eventType string, data any) {
	e.ClientEvents = append(e.ClientEvents, ClientEvent{
		Type: eventType,
		Data: data,
	})
}

type ClientEvent struct {
	Type string `json:"type"`

	Data any `json:"data"`
}

// Start WS Server
func NewStreamServer() (*StreamServer, error) {
	ws := StreamServer{}
	ws.Conns = make(map[*websocket.Conn]bool)
	ws.DataPackets = NewUpdatePackets()
	ws.DataPacketsMutex = sync.Mutex{}

	upgrader := websocket.Upgrader{
		ReadBufferSize:  ReadBufferSize,
		WriteBufferSize: WriteBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade normal http protocol to websocket
		websocket, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		ws.Conns[websocket] = true

		for {
			_, _, err := websocket.ReadMessage()

			if err != nil {
				delete(ws.Conns, websocket)
				break
			}
		}
	})

	indexerPort := os.Getenv("INDEXER_PORT")
	if indexerPort == "" {
		return nil, errors.New("Indexer port not set in .env file")
	}

	port, err := strconv.Atoi(indexerPort)
	if err != nil {
		return nil, errors.New("Indexer port is not a number")
	}

	ws.publishMessage()

	go func() {
		http.ListenAndServe(fmt.Sprintf("%s%d", ":", port), nil)
	}()

	return &ws, nil
}

type WSMessage struct {
	EventType string `json:"type"`
	Payload   any    `json:"message"`
}

func (ws *StreamServer) publishMessage() {

	ticker := time.NewTicker(StreamInterval)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:

				packetsToPublish := ws.FetchClearDataPackets()

				if len(packetsToPublish) == 0 {
					continue
				}

				payload := &WSMessage{
					EventType: Diff,
					Payload:   packetsToPublish}

				for conn := range ws.Conns {

					// TODO: in the future encrypt string
					err := conn.WriteJSON(payload)
					if err != nil {
						log.Println(err)
					}

				}

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

// this is similar to broadcasting events in Solidity.
// we broadcast state changes to client along with additional useful metadata, for clients, data pipelines down the line, etc
func (ws *StreamServer) PublishStateChanges(ecsUpdates ecs.ECSUpdateArray, uuid string, message string, clientEvents ClientEvents) {
	if ws == nil {
		return
	}

	ws.DataPacketsMutex.Lock()
	defer ws.DataPacketsMutex.Unlock()

	ws.DataPackets = append(ws.DataPackets, UpdatePacket{
		EcsUpdates:   ecsUpdates,
		Uuid:         uuid,
		Time:         time.Now().Unix(),
		Message:      message,
		ClientEvents: clientEvents,
	})
}

// TODO: test this
func filterEcsUpdatesWithoutLocal(tableUpdates ecs.ECSUpdateArray) ecs.ECSUpdateArray {
	// if the component starts with the word local, then filter it out
	filteredUpdates := ecs.ECSUpdateArray{}
	for _, ecsUpdate := range tableUpdates {
		// TODO: local prefix
		if !strings.HasPrefix(ecsUpdate.Table, "local") {
			filteredUpdates = append(filteredUpdates, ecsUpdate)
		}
	}

	return filteredUpdates
}

func (ws *StreamServer) FetchClearDataPackets() []UpdatePacket {
	ws.DataPacketsMutex.Lock()

	res := []UpdatePacket{}
	for _, item := range ws.DataPackets {
		res = append(res, item)
	}

	ws.DataPackets = NewUpdatePackets()

	ws.DataPacketsMutex.Unlock()

	return res
}

func NewUpdatePackets() UpdatePackets {
	return UpdatePackets{}
}
