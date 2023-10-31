package server

// ------------------------------------
// broadcasts events to client
// 1) table state updates
// 2) user defined events
// ------------------------------------

// calls the game-defined implementation of the broadcast interface

func BroadcastMessage(ctx *EngineCtx, clientEvents []ClientEvent) {
	if ctx.SystemBroadcastHandler == nil {
		return
	}
	ctx.SystemBroadcastHandler.BroadcastMessage(ctx, clientEvents)
}
