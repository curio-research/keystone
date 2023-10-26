package testutils

import (
	"database/sql"
	"strconv"
	"sync"
	"testing"

	"github.com/curio-research/keystone/core"
	"github.com/curio-research/keystone/server"
	pb_test "github.com/curio-research/keystone/test/proto/pb.test"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

// TODO refactor http to also be started inside here
func Server(t *testing.T, mode core.GameMode, websocketPort int, randSeedNumber int, schemaToTableAccessors map[interface{}]*core.TableBaseAccessor[any]) (*gin.Engine, *server.EngineCtx, *sql.DB, error) {
	gin.SetMode(gin.ReleaseMode)
	s := gin.Default()
	s.Use(server.CORSMiddleware())

	gameWorld := core.NewWorld()

	gameTick := server.NewGameTick(20)
	gameTick.Schedule = server.NewTickSchedule()

	// TODO: Kevin: create handler for this. less footgun plz!
	server.RegisterDefaultTables(gameWorld)

	// this is the master game context being passed around, containing pointers to everything
	gameCtx := &server.EngineCtx{ // TODO create a constructor for this
		GameId:                 "test",
		IsLive:                 true,
		World:                  gameWorld,
		GameTick:               gameTick,
		TransactionsToSaveLock: sync.Mutex{},
		Mode:                   mode,
		// SystemErrorHandler:     &network.ProtoBasedErrorHandler{},
		// SystemBroadcastHandler: &network.ProtoBasedBroadcastHandler{},
		RandSeed: randSeedNumber,
	}

	var db *sql.DB
	if mode == core.Prod || mode == core.DevSQL {
		saveStateHandler, saveTxHandler, testDB := SetupTestDB(t, gameCtx.GameId, true, schemaToTableAccessors)
		gameCtx.SaveStateHandler = saveStateHandler
		gameCtx.SaveTransactionsHandler = saveTxHandler
		db = testDB

		server.RegisterHTTPSQLRoutes(gameCtx, s)
		saveInterval := core.SaveStateInterval
		if mode == core.DevSQL {
			saveInterval = core.DevSQLSaveStateInterval
		}
		server.SetupSaveStateLoop(gameCtx, saveInterval)
	}

	// initialize a websocket streaming server for both incoming and outgoing requests
	streamServer, err := server.NewStreamServer(s, gameCtx, SocketRequestRouter, websocketPort)
	if err != nil {
		return nil, nil, nil, err
	}
	gameCtx.Stream = streamServer
	//gameTick.Setup(gameCtx, gameTick.Schedule)

	return s, gameCtx, db, nil
}

// message types
const (
	C2S_Test_MessageType = 9000
)

// the websocket router routes incoming requests based on protobuf types
func SocketRequestRouter(ctx *server.EngineCtx, requestMsg *server.NetworkMessage, socketConnection *websocket.Conn) {

	// data received through websocket from game clients
	requestType := requestMsg.GetCommand()

	// route incoming data based on command routes
	switch requestType {
	case C2S_Test_MessageType: // No-op, only used in integration tests
		queueTxIntoSystems[*pb_test.C2S_Test](ctx, requestMsg, &pb_test.C2S_Test{})
	}
}

// queue transactions for systems from the outside
func queueTxIntoSystems[T proto.Message](ctx *server.EngineCtx, requestMsg *server.NetworkMessage, req T) T {
	requestMsg.GetProtoMessage(req)
	requestId := requestMsg.Param()

	server.QueueTxFromExternal(ctx, req, strconv.Itoa(int(requestId)))
	return req
}
