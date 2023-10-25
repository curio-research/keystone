package testutils

import (
	"strconv"
	"sync"

	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	pb_test "github.com/curio-research/keystone/test/proto/pb.test"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

func StartMainServer(mode string, websocketPort int, mySQLdsn string, randSeedNumber int) (*gin.Engine, *server.EngineCtx, error) {
	gin.SetMode(gin.ReleaseMode)
	s := gin.Default()
	s.Use(server.CORSMiddleware())

	gameWorld := state.NewWorld()

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

	// initialize a websocket streaming server for both incoming and outgoing requests
	streamServer, err := server.NewStreamServer(s, gameCtx, SocketRequestRouter, websocketPort)
	if err != nil {
		return nil, nil, err
	}
	gameCtx.Stream = streamServer
	gameTick.Setup(gameCtx, gameTick.Schedule)

	// setup server routes
	// TODO: Restore HTTP routes in future integration tests

	return s, gameCtx, nil
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
