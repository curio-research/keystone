package startup

import (
	"fmt"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

func NewGameEngine(tickRate, randSeed int, tables ...state.ITable) *server.EngineCtx {
	gin.SetMode(gin.ReleaseMode)
	s := gin.Default()
	s.Use(server.CORSMiddleware())

	gameWorld := state.NewWorld()
	gameWorld.AddTables(tables...)

	gameTick := server.NewGameTick(tickRate)
	gameTick.Schedule = server.NewTickSchedule()

	server.RegisterDefaultTables(gameWorld)

	// this is the master game context being passed around, containing pointers to everything
	gameCtx := &server.EngineCtx{
		GameId:                 "test",
		IsLive:                 true,
		World:                  gameWorld,
		GameTick:               gameTick,
		TransactionsToSaveLock: sync.Mutex{},
		RandSeed:               randSeed,
	}

	return gameCtx
}

func RegisterSaveStateHandler(gameCtx *server.EngineCtx, saveStateHandler server.ISaveState, saveInterval time.Duration) {
	gameCtx.SaveStateHandler = saveStateHandler
	server.SetupSaveStateLoop(gameCtx, saveInterval)
}

func RegisterSaveTxHandler(gameCtx *server.EngineCtx, saveTxHandler server.ISaveTransactions, saveInterval time.Duration) {
	gameCtx.SaveTransactionsHandler = saveTxHandler
	server.SetupSaveTxLoop(gameCtx, saveInterval)
}

func RegisterRewindEndpoint(ctx *server.EngineCtx, g *gin.Engine) {
	g.POST("/rewindState", server.HandleRewindState(ctx))
}

func RegisterWSRoutes(gameCtx *server.EngineCtx, g *gin.Engine, router server.ISocketRequestRouter, websocketPort int) error {
	// initialize a websocket streaming server for both incoming and outgoing requests
	streamServer, err := server.NewStreamServer(g, gameCtx, router, websocketPort)
	if err != nil {
		return err
	}

	gameCtx.Stream = streamServer
	return nil
}

func RegisterBroadcastHandler(gameCtx *server.EngineCtx, broadcastHandler server.ISystemBroadcastHandler) {
	gameCtx.SystemBroadcastHandler = broadcastHandler
}

func RegisterErrorHandler(gameCtx *server.EngineCtx, errorHandler server.ISystemErrorHandler) {
	gameCtx.SystemErrorHandler = errorHandler
}

func Start(gameCtx *server.EngineCtx) error {
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

	gameCtx.GameTick.Setup(gameCtx)
	return nil
}
