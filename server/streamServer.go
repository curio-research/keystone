package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/curio-research/go-backend/engine"
	"github.com/gorilla/websocket"
)

const (
	Diff            = "DIFF"
	ReadBufferSize  = 1024
	WriteBufferSize = 1024
	StreamInterval  = 200 * time.Millisecond
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

// update packet container a group of ecs updates with additional metadata helping the client
type UpdatePacket struct {
	// array of ECS packets that need to be broadcasted to clients
	EcsUpdates []engine.ECSData `json:"ecsUpdates"`

	// id of package
	Uuid string `json:"uuid"`

	// timestamp
	Time int64 `json:"time"`

	// TODO: to be filled into the future
	// ex: it might contain instructions for the client to understand events like "Tower{1}-explode-fireball"
	Instructions any `json:"instructions"`
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
		// Upgrade normal http protocl to websocket
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
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (ws *StreamServer) publishMessage() {

	ticker := time.NewTicker(StreamInterval)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:

				packetsToPublish := ws.FetchClearDataPackets()

				if len(packetsToPublish) > 0 {

					dataPacketsStr, err := json.Marshal(packetsToPublish)

					if err != nil {
						log.Println(err)
						return
					}

					diff := &WSMessage{Diff, string(dataPacketsStr)}
					diffStr, parseDiffError := json.Marshal(diff)

					if parseDiffError != nil {
						log.Println(parseDiffError)
						return
					}

					for conn := range ws.Conns {
						if err := conn.WriteMessage(1, []byte(string(diffStr))); err != nil {
							log.Println(err)
							return
						}
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
func (ws *StreamServer) PublishStateChanges(ecsUpdates engine.ECSUpdateArray, uuid string) {
	if ecsUpdates == nil || len(ecsUpdates) == 0 {
		return
	}

	ws.DataPacketsMutex.Lock()
	defer ws.DataPacketsMutex.Unlock()

	ws.DataPackets = append(ws.DataPackets, UpdatePacket{EcsUpdates: ecsUpdates, Uuid: uuid, Time: time.Now().Unix()})
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
