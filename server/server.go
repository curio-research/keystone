package server

import (
	"fmt"
	"sync"

	// "github.com/curio-research/keystone/server/routes"
	"github.com/curio-research/keystone/server/routes"
	"github.com/curio-research/keystone/state"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
	"github.com/tjarratt/babble"
)

// Initialize new Keystone game engine
func NewGameEngine() *EngineCtx {

	// Initialize http server
	gin.SetMode(gin.ReleaseMode)
	s := gin.Default()
	s.Use(CORSMiddleware())

	// Initialize new game tick (not started yet, no goroutines here)
	gameTick := NewGameTick(DefaultTickRate)

	// Initialize new table-based game world
	gameWorld := state.NewWorld()
	RegisterDefaultTables(gameWorld)

	// Create a random gameID
	gameId := babble.NewBabbler().Babble()

	// This is the master game context being passed around, containing pointers to everything
	gameCtx := &EngineCtx{
		GameId:                 gameId,
		IsLive:                 false,
		World:                  gameWorld,
		GameTick:               gameTick,
		TransactionsToSaveLock: sync.Mutex{},
	}

	return gameCtx
}

func RegisterRewindEndpoint(ctx *EngineCtx, g *gin.Engine) {
	g.POST("/rewindState", HandleRewindState(ctx))
}

func RegisterGetStateEndpoint(ctx *EngineCtx, g *gin.Engine) {
	g.POST("/getState", routes.GetStateRouteHandler(ctx))
}

func RegisterGetEntityValueEndpoint(ctx *EngineCtx, g *gin.Engine) {
	g.POST("/entityValue", routes.GetEntityValueRouteHandler(ctx))
}

func RegisterGetStateRootHashEndpoint(ctx *EngineCtx, g *gin.Engine) {
	g.POST("/stateRoot", routes.StateRootRouteHandler(ctx))
}

// TODO: move this into init
func RegisterWSRoutes(gameCtx *EngineCtx, g *gin.Engine, router ISocketRequestRouter, websocketPort int) error {
	// initialize a websocket streaming server for both incoming and outgoing requests
	streamServer, err := NewStreamServer(g, gameCtx, router, websocketPort)
	if err != nil {
		return err
	}

	gameCtx.Stream = streamServer
	return nil
}

func RegisterBroadcastHandler(gameCtx *EngineCtx, broadcastHandler ISystemBroadcastHandler) {
	gameCtx.SystemBroadcastHandler = broadcastHandler
}

func RegisterErrorHandler(gameCtx *EngineCtx, errorHandler ISystemErrorHandler) {
	gameCtx.SystemErrorHandler = errorHandler
}

func Start(gameCtx *EngineCtx) error {
	if gameCtx.SystemErrorHandler == nil {
		log.Info("system error handler not provided")
	}

	if gameCtx.SystemBroadcastHandler == nil {
		log.Info("system broadcast handler not provided")
	}

	if gameCtx.SaveTransactionsHandler == nil {
		log.Info("save transactions handler not provided")
	}

	if gameCtx.SaveStateHandler == nil {
		log.Info("save state handler not provided")
	}

	if gameCtx.Stream == nil {
		log.Info("websocket routes not registered")
	}

	if len(gameCtx.World.Tables) == 0 {
		return fmt.Errorf("no tables registered")
	}

	if len(gameCtx.GameTick.Schedule.ScheduledTickSystems) == 0 {
		return fmt.Errorf("no systems registered")
	}

	gameCtx.GameTick.Start(gameCtx)
	return nil
}
