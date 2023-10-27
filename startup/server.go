package startup

import (
	"database/sql"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"github.com/gin-gonic/gin"
	"sync"
)

func NewGameEngine(tickRate int, setupDB bool) (*server.EngineCtx, error) {
	gin.SetMode(gin.ReleaseMode)
	s := gin.Default()
	s.Use(server.CORSMiddleware())

	gameWorld := state.NewWorld()

	gameTick := server.NewGameTick(tickRate)
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
		SystemErrorHandler:     systemErrorHandler,
		SystemBroadcastHandler: systemBroadcastHandler,
		RandSeed:               randSeedNumber,
	}

	var db *sql.DB
	if setupDB {
		saveStateHandler, saveTxHandler, testDB := SetupTestDB(t, gameCtx.GameId, true, schemaToTableAccessors)
		gameCtx.SaveStateHandler = saveStateHandler
		gameCtx.SaveTransactionsHandler = saveTxHandler
		db = testDB

		RegisterHTTPSQLRoutes(gameCtx, s)
		saveInterval := server.SaveStateInterval
		if mode == server.DevSQL {
			saveInterval = server.DevSQLSaveStateInterval
		}
		server.SetupSaveStateLoop(gameCtx, saveInterval)
	}

	// initialize a websocket streaming server for both incoming and outgoing requests
	streamServer, err := server.NewStreamServer(s, gameCtx, SocketRequestRouter, websocketPort)
	if err != nil {
		return nil, nil, nil, err
	}
	gameCtx.Stream = streamServer

	return s, gameCtx, db, nil
}

func SetErrorHandler() {

}
