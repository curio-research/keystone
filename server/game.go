package server

import (
	"fmt"
	"sync"

	"github.com/curio-research/keystone/keystone/ecs"
)

// contains entire server's state
type EngineCtx struct {

	// unique game ID
	GameId string

	// pause game world
	// 0: paused
	// 1: live
	IsLive bool

	// is the game state is being restored from db
	IsRestoringState bool

	// game world containing ECS data
	World *ecs.GameWorld

	// game tick aka the heart of the game
	Ticker *GameTick

	// stream server responsible for broadcasting data such as ecs changes and errors to clients
	Stream *StreamServer

	RandSeed int

	TickTransactionLock sync.Mutex

	// player tick requests queue that needs to be published to DA
	TickTransactionsQueue TickTransactions

	// TODO: [WIP] Data availability API
	TickTransactionApi TickTransactionApi

	// interface that handles all interactions for saving transactions
	SaveDatabase ISaveDatabase

	// "dev", "debug", "prod"
	// TODO: refactor this to constants
	Mode string

	// error log for printing when testing
	ErrorLog []ErrorLog

	// TODO: add
	StateUpdatesMutex sync.Mutex

	// state updates
	PendingStateUpdatesToSave []ecs.ECSUpdate
}

type ErrorLog struct {
	Tick    int
	Message string
}

// used in debug mode, print all the errors the game has collected
// should be the same set of errors that are broadcasted to clients
func PrintErrorLog(ctx *EngineCtx) {
	if len(ctx.ErrorLog) == 0 {
		return
	}

	fmt.Println()

	for _, errorLog := range ctx.ErrorLog {
		fmt.Println(errorLog.Message)
		fmt.Println()
	}

	fmt.Println()
}

// add a player request to the player world state
func (ctx *EngineCtx) AddTickTransaction(functionName string, tick int, payload any) {
	ctx.TickTransactionLock.Lock()
	defer ctx.TickTransactionLock.Unlock()

	newRequest := TickTransaction{
		GameId:       ctx.GameId,
		FunctionName: functionName,
		Tick:         tick,
		Payload:      payload,
	}

	ctx.TickTransactionsQueue = append(ctx.TickTransactionsQueue, newRequest)
}

// add to the list of state updates to save to database
func (ctx *EngineCtx) AddStateUpdatesToSave(updates []ecs.ECSUpdate) {
	// TODO: deep copy updates?
	ctx.PendingStateUpdatesToSave = append(ctx.PendingStateUpdatesToSave, updates...)
}
