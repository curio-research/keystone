package server

import (
	"sync"

	"github.com/curio-research/go-backend/engine"
	"github.com/gin-gonic/gin"
)

type EngineCtx struct {

	// unique game ID
	GameId string

	// pause game world
	Paused bool

	// game world containing ECS data
	World *engine.World

	// game tick aka the heart of the game
	Ticker *GameTick

	// stream server responsible for broadcasting data to clients
	Stream *StreamServer

	RandSeed int

	TickTransactionLock sync.Mutex

	// player tick requests queue that needs to be published to DA
	TickTransactionsQueue TickTransactions

	TickTransactionApi TickTransactionApi
}

// server routes
func SetupRoutes(s *gin.Engine, ctx *EngineCtx) {

	// main user interactions

	// move many troops at once
	s.POST("/moveMany", SubmitMoveMultipleRequest(ctx))

	// one troop attacking another
	s.POST("/attack", SubmitAttackAction(ctx))

	// player login
	s.POST("/login", Login(ctx))

	// regenerates troops for a player
	s.POST("/regenerate", RegenerateTroops(ctx))

	// helper routes

	// fetches entire game state
	s.POST("/gameState", DownloadWorld(ctx))

	s.POST("/saveWorld", SaveWorld(ctx))

	s.POST("/loadWorld", FetchLoadWorld(ctx))

	// save to the data availability layer
	s.POST("/publishTransactions", PublishTickTransactions(ctx))

	// will be used by validators to reconstruct state
	s.GET("/getTransactions", GetTickTransactions(ctx))

	s.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
