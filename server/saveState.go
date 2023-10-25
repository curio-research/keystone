package server

import (
	"time"

	"github.com/curio-research/keystone/state"
)

// general interface to implement
type ISaveState interface {
	SaveState(tableUpdates []state.TableUpdate) error

	RestoreState(ctx *EngineCtx, gameId string) error
}

type ISaveTransactions interface {
	SaveTransactions(ctx *EngineCtx, updates []TransactionSchema) error

	// TODO: hook this up with a CLI for our SDK
	RestoreStateFromTxs(ctx *EngineCtx, tickNumber int, gameId string) error
}

// game loop that triggers the save world state
func SetupSaveStateLoop(ctx *EngineCtx, saveInterval int) {
	tickerTime := time.Second
	if saveInterval != 0 {
		tickerTime = time.Duration(saveInterval) * time.Second
	}
	ticker := time.NewTicker(tickerTime)

	go func() {
		for range ticker.C {
			if ctx.IsLive {
				// deep copy pending state updates and clear then
				updatesToPublish := state.CopyTableUpdates(ctx.PendingStateUpdatesToSave)
				ctx.ClearStateUpdatesToSave()
				ctx.SaveStateHandler.SaveState(updatesToPublish)

				transactionsToSave := CopyTransactions(ctx.TransactionsToSave)
				ctx.ClearTransactionsToSave()
				ctx.SaveTransactionsHandler.SaveTransactions(ctx, transactionsToSave)
			}
		}
	}()
}
