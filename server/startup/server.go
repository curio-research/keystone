package startup

import (
	"sync"

	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/server/routes"
	"github.com/curio-research/keystone/state"
	"github.com/gin-gonic/gin"
	"github.com/tjarratt/babble"
)

// Initialize new Keystone game engine
func NewGameEngine() *server.EngineCtx {

	// Initialize http server
	gin.SetMode(gin.ReleaseMode)

	// TODO: set HTTP port
	ginHttpServer := gin.Default()

	// TODO: probably restore this
	ginHttpServer.Use(server.CORSMiddleware())

	// Initialize new game tick (not started yet, no goroutines here)
	gameTick := server.NewGameTick(server.DefaultTickRate)

	// Initialize new table-based game world
	gameWorld := state.NewWorld()
	server.RegisterDefaultTables(gameWorld)

	// Create a random gameID
	gameId := babble.NewBabbler().Babble()

	// Create a stream server
	streamServer := server.NewStreamServer()

	// This is the master game context being passed around, containing pointers to everything
	gameCtx := &server.EngineCtx{
		GameId:                 gameId,
		IsLive:                 false,
		World:                  gameWorld,
		GameTick:               gameTick,
		TransactionsToSaveLock: sync.Mutex{},
		GinHttpEngine:          ginHttpServer,
		Stream:                 streamServer,
	}

	return gameCtx
}

func RegisterRewindEndpoint(ctx *server.EngineCtx) {
	ctx.GinHttpEngine.POST("/rewindState", server.HandleRewindState(ctx))
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

// TODO: move this into init
// func RegisterWSRoutes(gameCtx *server.EngineCtx, g *gin.Engine, router server.ISocketRequestRouter, websocketPort int) error {

// 	// initialize a websocket streaming server for both incoming and outgoing requests
// 	streamServer, err := server.StartStreamServer(g, gameCtx, router, websocketPort)
// 	if err != nil {
// 		return err
// 	}

// 	gameCtx.Stream = streamServer
// 	return nil
// }
