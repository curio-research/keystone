package utils

import "github.com/curio-research/keystone/server"

// ticks the world forward for testing. time machine go brr
func TickWorldForward(ctx *server.EngineCtx, ticks int) int {
	if ctx.GameTick.Schedule == nil {
		panic("You need to set a ticker")
	}

	for i := 0; i < ticks; i++ {
		server.TickGameSystems(ctx, ctx.GameTick.Schedule)
		ctx.GameTick.TickNumber++
	}
	return ctx.GameTick.TickNumber
}

// time in milliseconds
func CalcFutureTickFromMs(ctx *server.EngineCtx, timeInMs int) int {
	return ctx.GameTick.TickNumber + (timeInMs / ctx.GameTick.TickRateMs)
}

func CalcFutureTickFromS(ctx *server.EngineCtx, timeInSeconds int) int {
	return CalcFutureTickFromMs(ctx, timeInSeconds*1000)
}
