package server

import (
	"fmt"
	"sync"
	"time"

	"github.com/curio-research/keystone/state"
	"github.com/fatih/color"
)

// ---------------------------------------
// context containing everything about the game server
// ---------------------------------------

type EngineCtx struct {
	// Unique game ID
	GameId string

	// Is the game live
	IsLive bool

	// Is the game state is being restored from db
	IsRestoringState bool

	// Game world containing table game state
	World *state.GameWorld

	// Game tick. The heartbeat of your game
	GameTick *GameTick

	// Stream server for broadcasting data such as table changes and errors to clients
	Stream *StreamServer

	// Transaction queue
	TransactionsToSaveLock sync.Mutex

	// Transactions to be stored in the data availability layer (aka a write ahead log basically)
	TransactionsToSave []TransactionSchema

	// Handles interactions for saving stae
	SaveStateHandler ISaveState

	SaveTransactionsHandler ISaveTransactions

	// Implementations on how to broadcast events and errors
	SystemErrorHandler     ISystemErrorHandler
	SystemBroadcastHandler ISystemBroadcastHandler

	// "dev", "prod"
	Mode GameMode

	// Whether game should record error in error log
	ShouldRecordError bool

	// Error log for printing when testing
	ErrorLog []ErrorLog

	StateUpdatesMutex sync.Mutex

	// State updates
	PendingStateUpdatesToSave []state.TableUpdate
}

// for debugging
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

// add a transactions that needs to be saved
func (ctx *EngineCtx) AddTransactionToSave(transaction TransactionSchema, tick int) error {
	ctx.TransactionsToSaveLock.Lock()
	defer ctx.TransactionsToSaveLock.Unlock()

	ctx.TransactionsToSave = append(ctx.TransactionsToSave, transaction)
	return nil
}

func (ctx *EngineCtx) AddTransactionsToSave() {
	tickNumber := ctx.GameTick.TickNumber

	transactionIds := GetTransactionsAtTickNumber(ctx.World, tickNumber)
	for _, transactionId := range transactionIds {
		transaction := TransactionTable.Get(ctx.World, transactionId)

		// only add external transactions aka ones from user requests
		if transaction.IsExternal {
			ctx.AddTransactionToSave(transaction, tickNumber)
		}
	}
}

// add to the list of state updates to save to database
func (ctx *EngineCtx) AddStateUpdatesToSave() {
	ctx.PendingStateUpdatesToSave = append(ctx.PendingStateUpdatesToSave, ctx.World.TableUpdates...)
}

func (ctx *EngineCtx) ClearStateUpdatesToSave() {
	ctx.PendingStateUpdatesToSave = []state.TableUpdate{}
}

// set whether game is live or not
func (ctx *EngineCtx) SetGameLiveliness(isLive bool) {
	ctx.IsLive = isLive
}

// clear transactions to save
func (ctx *EngineCtx) ClearTransactionsToSave() {
	ctx.TransactionsToSave = []TransactionSchema{}
}

func CopyTransactions(transactions []TransactionSchema) []TransactionSchema {
	newTransactions := make([]TransactionSchema, len(transactions))
	copy(newTransactions, transactions)
	return newTransactions
}

// Set game ID
func (ctx *EngineCtx) SetGameId(id string) {
	ctx.GameId = id
}

// Add tables to world
func (ctx *EngineCtx) AddTables(tables ...state.ITable) {
	for _, table := range tables {
		ctx.World.AddTable(table)
	}
}

// Set save state handler
func (ctx *EngineCtx) SetSaveStateHandler(saveStateHandler ISaveState, saveInterval time.Duration) {
	ctx.SaveStateHandler = saveStateHandler
	SetupSaveStateLoop(ctx, saveInterval)
}

// Set save transaction handler
func (ctx *EngineCtx) SetSaveTxHandler(saveTxHandler ISaveTransactions, saveInterval time.Duration) {
	ctx.SaveTransactionsHandler = saveTxHandler
	SetupSaveTxLoop(ctx, saveInterval)
}

// Interval: how frequently a system ticks (in milliseconds)
func (ctx *EngineCtx) AddSystem(IntervalMs int, tickFunction TickSystemFunction) {
	ctx.GameTick.Schedule.AddSystem(IntervalMs, tickFunction)
}

func (ctx *EngineCtx) Start() {
	color.HiYellow("")
	color.HiYellow("---- 🗝  Powered by Keystone 🗿 ----")
	fmt.Println()

	ctx.IsLive = true

}
