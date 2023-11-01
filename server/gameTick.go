package server

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"time"

	"github.com/curio-research/keystone/state"
	"google.golang.org/protobuf/proto"
)

// --------------------------- ---------------------------
// game tick is the heartbeat of a game, processing data per game tick,
// such as moving troops, battling, regenerating health, and doing some math stuff
// --------------------------- ---------------------------

type GameTick struct {
	// the current game tick number
	TickNumber int

	// per how many milliseconds the game ticks
	TickRateMs int

	// schedule of ticks and corresponding functions
	Schedule *TickSchedule
}

type TickSystemFunction func(ctx *EngineCtx) error

type TickSchedule struct {
	// list of systems that need to be triggered
	ScheduledTickSystems []TickSystem
}

type TickSystem struct {
	// how fast the game ticks
	TickInterval int

	// tick system function
	TickFunction TickSystemFunction
}

// Initialize an empty tick schedule. Systems are added to the tick schedule
func NewTickSchedule() *TickSchedule {
	return &TickSchedule{
		ScheduledTickSystems: []TickSystem{},
	}
}

// tick interval in milliseconds
func (s *TickSchedule) AddTickSystem(tickInterval int, tickFunction TickSystemFunction) {
	s.ScheduledTickSystems = append(s.ScheduledTickSystems, TickSystem{TickInterval: tickInterval, TickFunction: tickFunction})
}

type ISystemHandler[T any] func(ctx *TransactionCtx[T])

// transaction ctx
type TransactionCtx[T any] struct {
	GameCtx *EngineCtx

	// the transaction entity ID
	TxId int

	// access world variables through this
	W state.IWorld

	// transaction request parameters
	Req KeystoneRequest[T]

	EventCtx *EventCtx

	ErrorReturned bool

	Meta map[string]any
}

type CMD uint32

func (ctx *TransactionCtx[T]) EmitEvent(cmd CMD, data proto.Message, playerIds []int, isBlocking bool) {
	command := uint32(cmd)
	if ctx.EventCtx == nil {
		return
	}

	var transactionUuidIdentifier uint32
	if isBlocking {
		transactionUuidIdentifier = uint32(GetTransactionUuid(ctx.W, ctx.TxId))
	} else {
		transactionUuidIdentifier = 0
	}

	msg, err := NewMessage(0, command, transactionUuidIdentifier, data)

	if err != nil {
		return
	}

	ctx.EventCtx.AddEvent(msg, playerIds)
}

// error handling interface
type ISystemErrorHandler interface {
	FormatMessage(transactionUuidIdentifier int, errorMessage string) *NetworkMessage
}

// error broadcasting interface
type ISystemBroadcastHandler interface {
	BroadcastMessage(ctx *EngineCtx, clientEvents []ClientEvent)
}

// emit error
func (ctx *TransactionCtx[T]) EmitError(errorMessage string, playerIds []int) {
	if ctx.EventCtx == nil {
		return
	}

	jobParamIdentifier := GetTransactionUuid(ctx.W, ctx.TxId)

	msg := ctx.GameCtx.SystemErrorHandler.FormatMessage(jobParamIdentifier, errorMessage)

	ctx.EventCtx.AddEvent(msg, playerIds)

	ctx.ErrorReturned = true
}

func CreateSystemFromRequestHandler[T any](handler ISystemHandler[T], middlewareFunctions ...IMiddleware[T]) TickSystemFunction {
	return func(ctx *EngineCtx) error {
		transactionIds := GetSystemTransactionsOfType[KeystoneRequest[T]](ctx)

		for _, transactionId := range sort.IntSlice(transactionIds) {
			// create a world to temporarily record ecs changes
			ctx.World.ClearTableUpdates()

			// new client events array
			eventCtx := &EventCtx{}

			req := DecodeTxData[T](ctx, transactionId)
			worldUpdateBuffer := state.NewWorldUpdateBuffer(ctx.World)

			// create a transaction context for writing logic easier
			transactionCtx := &TransactionCtx[T]{
				GameCtx:  ctx,
				TxId:     transactionId,
				W:        worldUpdateBuffer,
				Req:      req,
				EventCtx: eventCtx,
				Meta:     map[string]any{},
			}

			callHandler := true
			for _, middlewareFunc := range middlewareFunctions {
				prevErrorCount := len(transactionCtx.GameCtx.ErrorLog)
				ok := middlewareFunc(transactionCtx)
				if !ok {
					callHandler = false
					if len(eventCtx.ClientEvents) == prevErrorCount {
						transactionCtx.EmitError(fmt.Sprintf("req %v did not pass verification", req), nil)
					}

					break
				}
			}

			if callHandler {
				handler(transactionCtx)
				if !transactionCtx.ErrorReturned {
					worldUpdateBuffer.ApplyUpdates() // updates the current world inside with the updates
				}
			}

			BroadcastMessage(ctx, eventCtx.ClientEvents)
		}

		return nil
	}
}

// general are not triggered by user inputs
func CreateGeneralSystem(handler ISystemHandler[any]) TickSystemFunction {
	return func(ctx *EngineCtx) error {
		ctx.World.ClearTableUpdates()

		eventCtx := &EventCtx{}

		worldUpdateBuffer := state.NewWorldUpdateBuffer(ctx.World)

		// create a transaction context for writing logic easier
		transactionCtx := &TransactionCtx[any]{
			GameCtx:  ctx,
			TxId:     -1,
			W:        worldUpdateBuffer,
			EventCtx: eventCtx,
		}

		handler(transactionCtx)
		if !transactionCtx.ErrorReturned {
			worldUpdateBuffer.ApplyUpdates() // updates the current world inside with the updates
		}

		// add state updates
		BroadcastMessage(ctx, eventCtx.ClientEvents)

		return nil
	}
}

func NewGameTick(tickRateMs int) *GameTick {
	return &GameTick{
		TickNumber: 1,
		TickRateMs: tickRateMs,
	}
}

// set up a game tick
func (g *GameTick) Setup(ctx *EngineCtx) {
	tickerTime := time.Duration(g.TickRateMs) * time.Millisecond
	ticker := time.NewTicker(tickerTime)

	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				if ctx.IsLive {
					ctx.AddTransactionsToSave()

					for _, tickSystem := range g.Schedule.ScheduledTickSystems {
						if ShouldTriggerTick(g.TickNumber, g.TickRateMs, tickSystem.TickInterval) {
							tickSystem.TickFunction(ctx)
						}
					}

					ctx.AddStateUpdatesToSave()

					DeleteAllTicksAtTickNumber(ctx.World, g.TickNumber)
					g.TickNumber++
				}

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

// if the tick happened between the previous and the current tick, we can still trigger it
func ShouldTriggerTick(tickNumber int, tickRate int, frequencyInMs int) bool {
	if frequencyInMs == 0 {
		return false
	}
	return tickNumber*tickRate%frequencyInMs < tickRate
}

// used for tests
func TickGameSystems(ctx *EngineCtx) {
	ctx.AddTransactionsToSave()
	gameTick := ctx.GameTick
	for _, tickSystem := range gameTick.Schedule.ScheduledTickSystems {
		tickSystem.TickFunction(ctx)
	}
	DeleteAllTicksAtTickNumber(ctx.World, gameTick.TickNumber)

	ctx.AddStateUpdatesToSave()
}

func SerializeRequestToString[T any](req T) (string, error) {
	jsonBytes, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

// get tick transactions
func GetTickTransactionsOfType(w *state.GameWorld, transactionType string, tickNumber int) []int {
	query := TransactionSchema{TickNumber: tickNumber, Type: transactionType}
	queryFields := []string{"TickNumber", "Type"}

	return TransactionTable.Filter(w, query, queryFields)
}

// get tick transactions of type at a current tick number
func GetSystemTransactionsOfType[T any](ctx *EngineCtx) []int {
	var t T
	return GetTickTransactionsOfType(ctx.World, reflect.TypeOf(t).String(), ctx.GameTick.TickNumber)
}

func GetTransactionsAtTickNumber(w *state.GameWorld, tickNumber int) []int {
	query := TransactionSchema{TickNumber: tickNumber}
	queryFields := []string{"TickNumber"}

	return TransactionTable.Filter(w, query, queryFields)
}

// queue transactions that are internal (ex: move planning)
func QueueTxFromInternal[T any](w state.IWorld, tickNumber int, data T, tickId string) error {
	return QueueTxAtTime(w, tickNumber, data, tickId, false)
}

// queue transactions that are user-initiated aka external
func QueueTxFromExternal[T any](ctx *EngineCtx, data T, tickId string) error {
	nextTickId := ctx.GameTick.TickNumber + 1
	return QueueTxAtTime(ctx.World, nextTickId, data, tickId, true)
}

// queues tick transactions to be executed in the future
func QueueTxAtTime(w state.IWorld, tickNumber int, data interface{}, uuid string, isExternal bool) error {
	serializedStringData, err := SerializeRequestToString(data)
	if err != nil {
		return err
	}

	requestTypeString := reflect.TypeOf(data).String()

	TransactionTable.Add(w, TransactionSchema{
		Type:          requestTypeString,
		Uuid:          uuid,
		Data:          serializedStringData,
		TickNumber:    tickNumber,
		IsExternal:    isExternal,
		UnixTimestamp: int(time.Now().UnixNano()),
	})

	return nil
}

// queue system tx at the next tick
func QueueTransaction(ctx *EngineCtx, data interface{}, isExternal bool) error {
	nextTickId := ctx.GameTick.TickNumber + 1
	return QueueTxAtTime(ctx.World, nextTickId, data, "", isExternal)
}

func DeleteAllTicksAtTickNumber(w *state.GameWorld, tickNumber int) {
	transactionIds := GetTransactionsAtTickNumber(w, tickNumber)

	for _, transactionId := range transactionIds {
		TransactionTable.RemoveEntity(w, transactionId)
	}
}
