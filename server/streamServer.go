package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/curio-research/keystone/state"
	"github.com/gorilla/websocket"
)

const (
	readBufferSize        = 1024
	writeBufferSize       = 1024
	defaultStreamInterval = 100 // ms
)

type ConnectionType struct {
	SubscribeAllStateUpdates bool
}

type StreamServer struct {
	// Websocket port
	Port int

	// Stream interval (milliseconds)
	StreamInterval int

	// Socket request router
	SocketRequestRouter ISocketRequestRouter

	// Client message data packets to be broadcasted
	ClientEventCh chan []ClientEvent

	// Table updates to be broadcasted
	TableUpdatesCh chan []state.TableUpdate

	// A pool of connections
	Conns      map[*websocket.Conn]ConnectionType
	ConnsMutex sync.Mutex

	// PlayerID to websocket connection
	PlayerIdToConnection      map[int]*websocket.Conn
	PlayerIdToConnectionMutex sync.Mutex
}

// update packet container a group of table updates with additional metadata
// that helps clients create the scene

// TODO: unused. remove in future
type UpdatePacket struct {
	// array of table update packets that need to be broadcasted to clients
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

// Initialize new stream server
func NewStreamServer() *StreamServer {
	s := &StreamServer{}

	s.Port = DefaultWebsocketPort
	s.Conns = make(map[*websocket.Conn]ConnectionType)
	s.PlayerIdToConnection = make(map[int]*websocket.Conn)
	s.StreamInterval = int(defaultStreamInterval)
	s.ClientEventCh = make(chan []ClientEvent, DefaultChannelBuffer)
	s.TableUpdatesCh = make(chan []state.TableUpdate, DefaultChannelBuffer)

	return s
}

// Set socket request router
func (s *StreamServer) SetSocketRequestRouter(router ISocketRequestRouter) {
	s.SocketRequestRouter = router
}

// Start websocket server
// TODO: have this return an error?
func (s *StreamServer) Start(ctx *EngineCtx) {

	upgrader := websocket.Upgrader{
		ReadBufferSize:  readBufferSize,
		WriteBufferSize: writeBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	httpServer := ctx.GinHttpEngine

	if s.SocketRequestRouter != nil {
		httpServer.GET("/", func(context *gin.Context) {
			websocket, err := upgrader.Upgrade(context.Writer, context.Request, nil)
			if err != nil {
				log.Println(err)
				return
			}

			s.AddConnection(websocket, false)

			for {
				_, msg, err := websocket.ReadMessage()

				if err != nil {
					delete(s.Conns, websocket)
					break
				}

				// deserialize from bytes
				requestMsg := NewMessageFromBuffer(msg)

				s.SocketRequestRouter(ctx, requestMsg, websocket)
			}
		})
	}

	// subscribe to all table updates
	httpServer.GET("/subscribeAllTableUpdates", func(context *gin.Context) {
		websocket, err := upgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}

		s.AddConnection(websocket, true)

		for {
			_, _, err := websocket.ReadMessage()

			if err != nil {
				s.RemoveConnection(websocket)
				break
			}
		}
	})

	s.StartMessageBroadcastLoop()

	go func() {
		http.ListenAndServe(fmt.Sprintf("%s%d", ":", s.Port), ctx.GinHttpEngine)
	}()
}

func (ws *StreamServer) SetPlayerIdToConnection(playerId int, conn *websocket.Conn) {
	ws.PlayerIdToConnectionMutex.Lock()
	ws.PlayerIdToConnection[playerId] = conn
	ws.PlayerIdToConnectionMutex.Unlock()
}

func (ws *StreamServer) AddConnection(conn *websocket.Conn, subscribeToStateUpdates bool) {
	ws.ConnsMutex.Lock()
	connection := &ConnectionType{
		SubscribeAllStateUpdates: subscribeToStateUpdates,
	}
	ws.Conns[conn] = *connection
	ws.ConnsMutex.Unlock()
}

func (ws *StreamServer) RemoveConnection(conn *websocket.Conn) {
	ws.ConnsMutex.Lock()
	delete(ws.Conns, conn)
	ws.ConnsMutex.Unlock()
}

type WSMessage struct {
	EventType string `json:"type"`
	Payload   any    `json:"message"`
}

// Start message broadcast loop
func (ws *StreamServer) StartMessageBroadcastLoop() {
	ticker := time.NewTicker(time.Duration(ws.StreamInterval) * time.Millisecond)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case tableUpdates := <-ws.TableUpdatesCh:
				tableUpdateBytes, _ := state.EncodeTableUpdateArrayToBytes(tableUpdates)

				// broadcast all state updates to subscribers of state data
				for conn, connectionType := range ws.Conns {
					if connectionType.SubscribeAllStateUpdates {
						conn.WriteMessage(websocket.TextMessage, tableUpdateBytes)
					}
				}
			case packets := <-ws.ClientEventCh:
				if len(packets) == 0 {
					continue
				}

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

	// Example from event to NetworkMessage, should be modified for real use case
	ws.ClientEventCh <- clientEvents

	// push state updates to queue
	ws.TableUpdatesCh <- tableUpdates
}
