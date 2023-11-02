package testutils

import (
	"database/sql"
	"strconv"
	"testing"

	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/server/startup"
	"github.com/curio-research/keystone/state"
	pb_test "github.com/curio-research/keystone/test/proto/pb.test"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

// TODO: refactor http to also be started inside here
func Server(t *testing.T, mode server.GameMode, websocketPort int, schemaToTableAccessors map[interface{}]*state.TableBaseAccessor[any]) (*gin.Engine, *server.EngineCtx, *sql.DB, error) {

	ctx := startup.NewGameEngine()

	ctx.SetGameId("test")
	ctx.SetTickRate(20)
	ctx.AddTables(schemaToTableAccessors)

	ctx.SetEmitErrorHandler(NewMockErrorHandler())

	// TODO: add set http server

	// Websocket handler
	ctx.SetSocketRequestRouter(SocketRequestRouter)
	ctx.SetWebsocketPort(websocketPort)

	var db *sql.DB
	if mode == server.Prod || mode == server.DevSQL {
		saveStateHandler, saveTxHandler, testDB := SetupTestDB(t, ctx.GameId, true, schemaToTableAccessors)
		db = testDB

		saveInterval := server.SaveStateInterval
		if mode == server.DevSQL {
			saveInterval = server.DevSQLSaveStateInterval
		}

		ctx.SetSaveStateHandler(saveStateHandler, saveInterval)
		ctx.SetSaveTxHandler(saveTxHandler, saveInterval)
		startup.RegisterRewindEndpoint(ctx)
	}

	// http api routes
	startup.RegisterGetEntityValueEndpoint(ctx)
	startup.RegisterGetStateEndpoint(ctx)
	startup.RegisterGetStateRootHashEndpoint(ctx)

	return ctx.GinHttpEngine, ctx, db, nil
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
