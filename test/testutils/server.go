package testutils

import (
	"database/sql"
	"strconv"
	"testing"

	"github.com/curio-research/keystone/startup"

	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	pb_test "github.com/curio-research/keystone/test/proto/pb.test"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

// TODO refactor http to also be started inside here
func Server(t *testing.T, mode server.GameMode, websocketPort int, schemaToTableAccessors map[interface{}]*state.TableBaseAccessor[any]) (*gin.Engine, *server.EngineCtx, *sql.DB, error) {
	gin.SetMode(gin.ReleaseMode)
	s := gin.Default()
	s.Use(server.CORSMiddleware())

	tables := []state.ITable{}
	for _, accessor := range schemaToTableAccessors {
		tables = append(tables, accessor)
	}

	tickRate := 20 // 20 ms
	ctx := startup.NewGameEngine("test", tickRate, tables...)

	// initialize a websocket streaming server for both incoming and outgoing requests
	err := startup.RegisterWSRoutes(ctx, s, SocketRequestRouter, websocketPort)
	if err != nil {
		return nil, nil, nil, err
	}
	startup.RegisterErrorHandler(ctx, NewMockErrorHandler())

	var db *sql.DB
	if mode == server.Prod || mode == server.DevSQL {
		saveStateHandler, saveTxHandler, testDB := SetupTestDB(t, ctx.GameId, true, schemaToTableAccessors)
		db = testDB

		saveInterval := server.SaveStateInterval
		if mode == server.DevSQL {
			saveInterval = server.DevSQLSaveStateInterval
		}

		startup.RegisterSaveStateHandler(ctx, saveStateHandler, saveInterval)
		startup.RegisterSaveTxHandler(ctx, saveTxHandler, saveInterval)
		startup.RegisterRewindEndpoint(ctx, s)
	}

	// http api routes
	startup.RegisterGetEntityValueEndpoint(ctx, s)
	startup.RegisterGetStateEndpoint(ctx, s)
	startup.RegisterGetStateRootHashEndpoint(ctx, s)

	return s, ctx, db, nil
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
		queueTxIntoSystems[*pb_test.C2S_Test](ctx, requestMsg, server.KeystoneRequest[*pb_test.C2S_Test]{})
	}
}

// queue transactions for systems from the outside
func queueTxIntoSystems[T proto.Message](ctx *server.EngineCtx, requestMsg *server.NetworkMessage, req server.KeystoneRequest[T]) T {
	requestMsg.GetProtoMessage(req.Data)
	requestId := requestMsg.Param()

	server.QueueTxFromExternal(ctx, req, strconv.Itoa(int(requestId)))
	return req.Data
}
