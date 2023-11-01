package server

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/curio-research/keystone/state"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
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

	// Gin HTTP server
	GinHttpEngine *gin.Engine

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
func (ctx *EngineCtx) AddTables(tables map[interface{}]*state.TableBaseAccessor[any]) {
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

// Set broadcast event handler
func (ctx *EngineCtx) SetEmitEventHandler(broadcastHandler ISystemBroadcastHandler) {
	ctx.SystemBroadcastHandler = broadcastHandler
}

// Set broadcast error handler
func (ctx *EngineCtx) SetEmitErrorHandler(errorHandler ISystemErrorHandler) {
	ctx.SystemErrorHandler = errorHandler
}

// Set tick rate (milliseconds)
func (ctx *EngineCtx) SetTickRate(tickRateMs int) {
	ctx.GameTick.TickRateMs = tickRateMs
}

// Start Keystone game server
func (ctx *EngineCtx) Start() {
	color.HiYellow("")
	color.HiYellow("---- 🗝  Powered by Keystone 🗿 ----")
	fmt.Println()

	color.HiWhite("Tick rate:         " + strconv.Itoa(ctx.GameTick.TickRateMs) + "ms")

	// TODO: change to log library

	if ctx.SystemErrorHandler == nil {
		fmt.Println("system error handler not provided")
	}

	if ctx.SystemBroadcastHandler == nil {
		fmt.Println("system broadcast handler not provided")
	}

	if ctx.SaveTransactionsHandler == nil {
		fmt.Println("save transactions handler not provided")
	}

	if ctx.SaveStateHandler == nil {
		fmt.Println("save state handler not provided")
	}

	if ctx.Stream == nil {
		fmt.Println("websocket routes not registered")
	}

	if len(ctx.World.Tables) == 0 {
		fmt.Println("no tables registered")
	}

	if len(ctx.GameTick.Schedule.ScheduledTickSystems) == 0 {
		fmt.Println("no tables registered")
	}

	ctx.IsLive = true

	// TODO: start HTTP server
	// ctx.GinHttpEngine.Run(":" + "8080")

	// addr := ":" + strconv.Itoa(8080)

	// httpServer := &http.Server{
	// 	Addr:    addr,
	// 	Handler: ctx.GinHttpEngine,
	// }

	// go func() {
	// 	err := httpServer.ListenAndServe()
	// 	if err != nil && !errors.Is(err, http.ErrServerClosed) {
	// 		t.Errorf("http server closed with unexpected error %v", err)
	// 		return
	// 	}
	// }()

	// TODO: start stream server
	// TODO: start tick system
}
