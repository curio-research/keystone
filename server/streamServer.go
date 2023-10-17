package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/curio-research/keystone/keystone/state"
	"github.com/gorilla/websocket"
)

const (
	ReadBufferSize  = 1024
	WriteBufferSize = 1024
	StreamInterval  = 100 * time.Millisecond
)

type StreamServer struct {
	// lock for protobuf packets
	ProtoBufPacketsMutex sync.Mutex

	// list of client message data packets to be sent to client
	ClientEventsQueue []ClientEvent

	TableUpdatesQueue []state.TableUpdate

	// a pool of connections
	Conns map[*websocket.Conn]bool

	// admin connections Pool
	AdminConns map[*websocket.Conn]bool

	PlayerIdToConnection map[int]*websocket.Conn
}

// ProtoBuf Packets
// type ProtoBufMessagePackets []*NetworkMessage

// update packet container a group of table updates with additional metadata
// that helps clients create the scene

// TODO: unused. remove in future
type UpdatePacket struct {
	// array of ECS packets that need to be broadcasted to clients
	TableUpdates []state.TableUpdate `json:"tableUpdates"`

	// id of package that corresponds with the HTTP requests's returned UUID (similar to a transaction hash)
	Uuid string `json:"uuid"`

	// timestamp
	Time int64 `json:"time"`

	// error message string returned from a request
	Message string `json:"message"`
}

type NetworkMessages []*NetworkMessage

type ClientEvent struct {
	NetworkMessage *NetworkMessage
	PlayerIds      []int
}

// used once per-request
type EventCtx struct {
	ClientEvents []ClientEvent
}

// adds a client event to the event context
func (e *EventCtx) AddEvent(msg *NetworkMessage, playerIds []int) {
	clientMessage := ClientEvent{
		NetworkMessage: msg,
		PlayerIds:      playerIds,
	}

	e.ClientEvents = append(e.ClientEvents, clientMessage)
}

type ISocketRequestRouter func(ctx *EngineCtx, requestMsg *NetworkMessage, socketConnection *websocket.Conn)

// Start WS Server
func NewStreamServer(s *gin.Engine, ctx *EngineCtx, router ISocketRequestRouter, websocketPort int) (*StreamServer, error) {
	ws := StreamServer{}
	ws.Conns = make(map[*websocket.Conn]bool)
	ws.ClientEventsQueue = make([]ClientEvent, 0)
	ws.PlayerIdToConnection = make(map[int]*websocket.Conn)
	ws.ProtoBufPacketsMutex = sync.Mutex{}
	ws.AdminConns = make(map[*websocket.Conn]bool)

	upgrader := websocket.Upgrader{
		ReadBufferSize:  ReadBufferSize,
		WriteBufferSize: WriteBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	s.GET("/", func(context *gin.Context) {
		websocket, err := upgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}

		ws.Conns[websocket] = true

		for {
			_, msg, err := websocket.ReadMessage()

			// if the transaction type is "connect to game", then we apply the connection mapping

			if err != nil {
				delete(ws.Conns, websocket)
				break
			}

			// deserialize from bytes
			requestMsg := NewMessageFromBuffer(msg)

			router(ctx, requestMsg, websocket)
		}
	})

	s.GET("/subscribeAllTableUpdates", func(context *gin.Context) {
		websocket, err := upgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		ws.AddAdminConnection(websocket)

		for {
			_, _, err := websocket.ReadMessage()

			if err != nil {
				ws.RemoveAdminConnection(websocket)
				break
			}
		}
	})

	ws.PublishMessage()

	go func() {
		http.ListenAndServe(fmt.Sprintf("%s%d", ":", websocketPort), s)
	}()

	return &ws, nil
}

func (ws *StreamServer) SetPlayerIdToConnection(playerId int, conn *websocket.Conn) {
	ws.PlayerIdToConnection[playerId] = conn
}

func (ws *StreamServer) AddAdminConnection(conn *websocket.Conn) {
	ws.AdminConns[conn] = true
}

func (ws *StreamServer) RemoveAdminConnection(conn *websocket.Conn) {
	delete(ws.AdminConns, conn)
}

type WSMessage struct {
	EventType string `json:"type"`
	Payload   any    `json:"message"`
}

func (ws *StreamServer) PublishMessage() {
	ticker := time.NewTicker(StreamInterval)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:

				packets := ws.FetchEventsFromQueue()
				ws.ClearClientMessageQueue()

				tableUpdates := ws.FetchTableUpdatesFromQueue()
				ws.ClearTableUpdatesQueue()

				tableUpdateBytes, _ := state.EncodeTableUpdateArrayToBytes(tableUpdates)

				// broadcast all state updates to admins
				for conn := range ws.AdminConns {
					conn.WriteMessage(websocket.TextMessage, tableUpdateBytes)
				}

				// loop through players that have playerIds that are negative
				if len(packets) == 0 {
					continue
				}

				ws.ProtoBufPacketsMutex.Lock()

				// loop through packets and broadcast to user
				for _, packet := range packets {

					// broadcast to all players
					if packet.PlayerIds == nil {
						for conn := range ws.Conns {
							// Send probuf packet data back to client

							buffer := packet.NetworkMessage.ParseToBuffer()
							conn.WriteMessage(websocket.BinaryMessage, buffer)
						}
					} else {
						// only broadcast to select players
						for _, playerId := range packet.PlayerIds {
							conn := ws.PlayerIdToConnection[playerId]

							if conn != nil {
								buffer := packet.NetworkMessage.ParseToBuffer()
								conn.WriteMessage(2, buffer)
							}

						}
					}

				}

				ws.ProtoBufPacketsMutex.Unlock()

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

// this is similar to broadcasting events in Solidity.
// we broadcast state changes to client along with additional useful metadata, for clients, data pipelines down the line, etc
// TODO: NOTE: currently we do not broadcast table updates
func (ws *StreamServer) PublishStateChanges(tableUpdates state.TableUpdateArray, clientEvents []ClientEvent) {
	if ws == nil {
		return
	}

	ws.ProtoBufPacketsMutex.Lock()

	// Example from event to NetworkMessage, should be modified for real use case
	ws.ClientEventsQueue = append(ws.ClientEventsQueue, clientEvents...)

	// push state updates to queue
	ws.TableUpdatesQueue = append(ws.TableUpdatesQueue, tableUpdates...)

	ws.ProtoBufPacketsMutex.Unlock()
}

func (ws *StreamServer) FetchEventsFromQueue() []ClientEvent {
	ws.ProtoBufPacketsMutex.Lock()

	res := []ClientEvent{}
	res = append(res, ws.ClientEventsQueue...)

	ws.ProtoBufPacketsMutex.Unlock()

	return res
}

// clear all messages in the client message queue
func (ws *StreamServer) ClearClientMessageQueue() {
	ws.ClientEventsQueue = make([]ClientEvent, 0)
}

// fetch all table updates from queue
func (ws *StreamServer) FetchTableUpdatesFromQueue() []state.TableUpdate {
	ws.ProtoBufPacketsMutex.Lock()

	res := []state.TableUpdate{}
	res = append(res, ws.TableUpdatesQueue...)

	ws.ProtoBufPacketsMutex.Unlock()

	return res
}

// clear all table updates from queue
func (ws *StreamServer) ClearTableUpdatesQueue() {
	ws.TableUpdatesQueue = make([]state.TableUpdate, 0)
}
