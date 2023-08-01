package server

// TODO: [WIP]
// syncs the game state progressively from the data availability layer
func SyncGameState(ctx *EngineCtx, startTick, endTick int) error {

	interval := 100

	lastSyncedTick := startTick
	for lastSyncedTick < endTick {
		// get all tick transactions from the data availability layer
		// tickTransactions, err := ctx.TickTransactionApi.DownloadTickTransactions(ctx.GameId, lastSyncedTick, lastSyncedTick+interval)
		// if err != nil {
		// 	return err
		// }

		// apply all tick transactions to the world

		// run all systems

		lastSyncedTick += interval
	}

	return nil
}
