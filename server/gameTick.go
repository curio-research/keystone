package server

import (
	"encoding/json"
	"reflect"
	"sort"
	"time"

	"github.com/curio-research/keystone/keystone/ecs"
)

// game tick is the heartbeat of a game, processing data per game tick,
// such as moving troops, battling, regenerating health, and doing some math stuff

type GameTick struct {
	// the current game tick number
	TickNumber int

	// how often the game ticks in milliseconds
	TickRateMs int

	// schedule of ticks and corresponding functions
	Schedule *TickSchedule
}

type TickSystemFunction func(ctx *EngineCtx) error

type TickSchedule struct {
	// list of tick job systems that need to be triggered
	ScheduledTickSystems []TickSystem
}

type TickSystem struct {
	// interval for when the job should be triggered (milliseconds)
	TickInterval int

	// tick system function
	TickFunction TickSystemFunction
}

func NewTickSchedule() *TickSchedule {
	return &TickSchedule{
		ScheduledTickSystems: []TickSystem{},
	}
}

func (s *TickSchedule) AddTickSystem(tickInterval int, tickFunction TickSystemFunction) {
	s.ScheduledTickSystems = append(s.ScheduledTickSystems, TickSystem{TickInterval: tickInterval, TickFunction: tickFunction})
}

// system handler type
type SystemHandler[T any] func(ctx *EngineCtx, jobId int, w *ecs.GameWorld, req T, eventCtx *EventCtx) error

type SystemHandlerNew[T any] func(ctx *TransactionCtx[T]) error

// transaction ctx
type TransactionCtx[T any] struct {
	GameCtx *EngineCtx

	JobId int

	// access world variables through this
	W *ecs.GameWorld

	// request parameters
	Req T

	EventCtx *EventCtx
}

// emit event to client
func (ctx *TransactionCtx[T]) EmitEvent(eventType string, data any) {
	if ctx.EventCtx == nil {
		return
	}

	ctx.EventCtx.EmitEvent(eventType, data)
}

// creates a system from a handler
func CreateSystemFromRequestHandler[T any](handler SystemHandlerNew[T]) TickSystemFunction {

	return func(ctx *EngineCtx) error {
		jobIds := GetTickJobsOfType[T](ctx)

		for _, jobId := range sort.IntSlice(jobIds) {

			// create a world to temporarily record ecs changes
			w := ecs.StartRecordingStateChanges(ctx.World)

			req := DecodeJobData[T](ctx, jobId)

			// new client events array
			eventCtx := &EventCtx{}

			// create a transaction context for writing logic easier
			transactionCtx := &TransactionCtx[T]{
				GameCtx:  ctx,
				JobId:    jobId,
				W:        w,
				Req:      req,
				EventCtx: eventCtx,
			}

			err := handler(transactionCtx)

			// broadcast all ecs updates regardless whether it fails or not
			errStr := ""
			if err != nil {
				errStr = err.Error()
			}

			BroadcastMessage(ctx, w, jobId, errStr, eventCtx.ClientEvents)
		}

		return nil
	}
}

func CreateGeneralSystem(handler SystemHandlerNew[any]) TickSystemFunction {
	return func(ctx *EngineCtx) error {

		w := ecs.StartRecordingStateChanges(ctx.World)

		eventCtx := &EventCtx{}

		// create a transaction context for writing logic easier
		transactionCtx := &TransactionCtx[any]{
			GameCtx:  ctx,
			JobId:    -1,
			W:        w,
			Req:      nil,
			EventCtx: eventCtx,
		}

		err := handler(transactionCtx)

		errStr := ""
		if err != nil {
			errStr = err.Error()
		}

		// add state updates
		ctx.AddStateUpdatesToSave(w.TableUpdates)

		BroadcastMessage(ctx, w, -1, errStr, eventCtx.ClientEvents)

		return nil
	}

}

func NewGameTick(tickRate int) *GameTick {
	return &GameTick{
		TickNumber: 1,
		TickRateMs: tickRate,
	}
}

// set up a game tick
func (g *GameTick) Setup(ctx *EngineCtx, tickSchedule *TickSchedule) {
	tickerTime := time.Duration(g.TickRateMs) * time.Millisecond
	ticker := time.NewTicker(tickerTime)

	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:

				if ctx.IsLive {
					for _, tickSystem := range tickSchedule.ScheduledTickSystems {
						if ShouldTriggerTick(g.TickNumber, g.TickRateMs, tickSystem.TickInterval) {
							tickSystem.TickFunction(ctx)
						}
					}

					// get all jobs in this tick and delete them
					// TODO: test this performance
					DeleteAllJobsOfTick(ctx.World, g.TickNumber)

					g.TickNumber++
				}

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

// TODO: [WIP] State sync flow
// Load empty state from DA layer
// Download transactions from DA
// apply those transactions to the state
// this step will fail. save the previously correctly applied state
// snapshot state

// given a current state, apply the tick transactions to the state
// tick transactions are likely supplied within a certain range of ticks

func ShouldTriggerTick(tickNumber int, tickRate int, frequencyInMs int) bool {
	return tickNumber*tickRate%frequencyInMs < tickRate
}

// tick all systems
func ForceTickAllSystems(ctx *EngineCtx) {

	for _, tickSystem := range ctx.Ticker.Schedule.ScheduledTickSystems {
		tickSystem.TickFunction(ctx)
	}
}

func TickGameSystems(ctx *EngineCtx, tickSchedule *TickSchedule) {
	for _, tickSystem := range tickSchedule.ScheduledTickSystems {
		tickSystem.TickFunction(ctx)
	}
}

func SerializeRequestToString[T any](req T) (string, error) {

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func GetTickJobs(w *ecs.GameWorld, tickType string, tickNumber int) []int {

	query := JobSchema{TickNumber: tickNumber, JobType: tickType}
	queryFields := []string{"TickNumber", "TickType"}

	return JobTable.Filter(w, query, queryFields)
}

// get tick jobs of type at a current tick number
func GetTickJobsOfType[T any](ctx *EngineCtx) []int {
	var t T
	return GetTickJobs(ctx.World, reflect.TypeOf(t).String(), ctx.Ticker.TickNumber)

}

// queues tick job to be executed in the future
func QueueTickJobAtTime[T any](w *ecs.GameWorld, tickNumber int, jobData T, tickId string) error {
	var t T

	serializedJobString, err := SerializeRequestToString(jobData)
	if err != nil {
		return err
	}

	requestTypeString := reflect.TypeOf(t).String()

	AddTickJob(w, tickNumber, requestTypeString, serializedJobString, tickId, "")

	return nil
}

// queue tick job at the next tick. mainly used for testing
func QueueTickJob[T any](ctx *EngineCtx, jobData T) error {
	nextTickId := ctx.Ticker.TickNumber + 1
	return QueueTickJobAtTime[T](ctx.World, nextTickId, jobData, "")
}

// queue tick job at the next tick with a specific ID. mainly used for testing
func QueueTickJobWithId[T any](ctx *EngineCtx, jobData T, tickId string) error {
	nextTickId := ctx.Ticker.TickNumber + 1
	return QueueTickJobAtTime[T](ctx.World, nextTickId, jobData, tickId)
}

func GetAllJobsAtTick(w *ecs.GameWorld, tickNumber int) []int {
	query := JobSchema{TickNumber: tickNumber}
	queryFields := []string{"TickNumber"}

	return JobTable.Filter(w, query, queryFields)
}

func DeleteAllJobsOfTick(w *ecs.GameWorld, tickNumber int) {
	jobIds := GetAllJobsAtTick(w, tickNumber)

	for _, jobId := range jobIds {
		JobTable.RemoveEntity(w, jobId)
	}
}
