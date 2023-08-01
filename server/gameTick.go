package server

import (
	"math"
	"time"
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

const (
	TickRate = 100 // milliseconds
)

func NewGameTick() *GameTick {
	return &GameTick{
		TickNumber: 1,
		TickRateMs: TickRate,
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

				if ctx.Paused != true {
					for _, tickSystem := range tickSchedule.ScheduledTickSystems {
						if ShouldTriggerTick(g.TickNumber, g.TickRateMs, tickSystem.TickInterval) {
							tickSystem.TickFunction(ctx)
						}
					}
				}

				g.TickNumber++
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

// [WIP] State sync flow
// Load empty state from DA layer
// Download transactions from DA
// apply those transactions to the state
// this step will fail. save the previously correctly applied state
// snapshot state

// given a current state, apply the tick transactions to the state
// tick transactions are likely supplied within a certain range of ticks
func TransitionStateFromTickTransactions(ctx *EngineCtx, tickTransactions TickTransactions, startTick int, endTick int) {

	// TODO: should we deep copy an entire state?

	// inject all tick transactions in ecs world state
	for _, tickTransaction := range tickTransactions {
		ctx.AddTickTransaction(tickTransaction.FunctionName, tickTransaction.Tick, tickTransaction.Payload)
	}

	for i := startTick; i < endTick; i++ {
		// run systems on each function
		for _, tickSystem := range ctx.Ticker.Schedule.ScheduledTickSystems {
			if ShouldTriggerTick(ctx.Ticker.TickNumber, ctx.Ticker.TickRateMs, tickSystem.TickInterval) {
				tickSystem.TickFunction(ctx)
			}
		}
	}

	// snapshot entire state
}

func CalculateTickBasedOnTimeMilliseconds(ms int, tickRateMs int) int {
	return ms * 1_000_000 / int((time.Duration(tickRateMs) * time.Millisecond).Nanoseconds())
}

func ShouldTriggerTick(tickNumber int, tickRate int, frequencyInMs int) bool {
	return math.Floor(float64(tickNumber*tickRate%frequencyInMs)) == 0
}
