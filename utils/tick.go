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
