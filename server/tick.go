package server

// ticks the world forward for testing. time machine go brr
func TickWorldForward(ctx *EngineCtx, ticks int) int {
	if ctx.GameTick.Schedule == nil {
		panic("You need to set a ticker")
	}

	for i := 0; i < ticks; i++ {
		TickGameSystems(ctx)
		ctx.GameTick.TickNumber++
	}
	return ctx.GameTick.TickNumber
}

// time in milliseconds
func CalcFutureTickFromMs(ctx *EngineCtx, timeInMs int) int {
	return ctx.GameTick.TickNumber + (timeInMs / ctx.GameTick.TickRateMs)
}

func CalcFutureTickFromS(ctx *EngineCtx, timeInSeconds int) int {
	return CalcFutureTickFromMs(ctx, timeInSeconds*1000)
}
