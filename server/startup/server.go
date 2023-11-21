package startup

import (
	"math/rand"
	"sync"
	"time"

	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/server/routes"
	"github.com/curio-research/keystone/state"
	"github.com/gin-gonic/gin"
)

// Initialize new Keystone game engine
func NewGameEngine() *server.EngineCtx {

	// Initialize http server
	gin.SetMode(gin.ReleaseMode)
	ginHttpServer := gin.Default()

	// TODO: make this modular in the future
	ginHttpServer.Use(server.CORSMiddleware())

	// Initialize new game tick (not started yet, no goroutines here)
	gameTick := server.NewGameTick(server.DefaultTickRate)

	// Initialize new table-based game world
	gameWorld := state.NewWorld()
	server.RegisterDefaultTables(gameWorld)

	// Create a random gameID
	gameId := randomString(8)

	// Create a stream server
	streamServer := server.NewStreamServer()

	// This is the master game context being passed around, containing pointers to everything
	ctx := &server.EngineCtx{
		GameId:                 gameId,
		IsLive:                 false,
		World:                  gameWorld,
		GameTick:               gameTick,
		TransactionsToSaveLock: sync.Mutex{},
		GinHttpEngine:          ginHttpServer,
		Stream:                 streamServer,
		HttpPort:               server.DefaultServerPort,
		ShouldSaveState:        false,
		ShouldSaveTransactions: false,
		TransactionChan:        make(chan server.TransactionSchema, server.DefaultChannelBuffer),
		StateUpdateChan:        make(chan []state.TableUpdate, server.DefaultChannelBuffer),
	}

	// Use protobuf based handlers as default
	ctx.SetEmitErrorHandler(&server.ProtoBasedErrorHandler{})
	ctx.SetEmitEventHandler(&server.ProtoBasedBroadcastHandler{})

	return ctx
}

func RegisterRewindEndpoint(ctx *server.EngineCtx, initWorld func(w *state.GameWorld)) {
	ctx.GinHttpEngine.POST("/rewindState", server.HandleRewindState(ctx, initWorld))
}

func RegisterGetStateEndpoint(ctx *server.EngineCtx) {
	ctx.GinHttpEngine.POST("/getState", routes.GetStateRouteHandler(ctx))
}

func RegisterGetEntityValueEndpoint(ctx *server.EngineCtx) {
	ctx.GinHttpEngine.POST("/entityValue", routes.GetEntityValueRouteHandler(ctx))
}

func RegisterGetStateRootHashEndpoint(ctx *server.EngineCtx) {
	ctx.GinHttpEngine.POST("/stateRoot", routes.StateRootRouteHandler(ctx))
}

func randomString(length int) string {
	// Define a character set from which to generate the random string
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Create a byte slice to store the random string
	result := make([]byte, length)

	// Generate random characters and append them to the result slice
	for i := 0; i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}

	// Convert the byte slice to a string and return it
	return string(result)
}
