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
	SaveTransactions(updates []TransactionSchema) error

	// TODO: hook this up with a CLI for our SDK
	RestoreStateFromTxs(ctx *EngineCtx, tickNumber int, gameId string) error
}

// game loop that triggers the save world state
func SetupSaveStateLoop(ctx *EngineCtx, saveInterval time.Duration) {
	tickerTime := time.Second
	if saveInterval != 0 {
		tickerTime = saveInterval
	}

	ticker := time.NewTicker(tickerTime)
	updatesToPublish := []state.TableUpdate{}

	go func() {
		for {
			if ctx.IsLive && ctx.ShouldSaveState {
				select {
				case <-ticker.C:
					ctx.SaveStateHandler.SaveState(updatesToPublish)
					updatesToPublish = []state.TableUpdate{}
				case updates := <-ctx.StateUpdateChan:
					updatesToPublish = append(updatesToPublish, updates...)
				}
			}
		}
	}()
}

// Set up save transaction loop
func SetupSaveTxLoop(ctx *EngineCtx, saveInterval time.Duration) {
	tickerTime := time.Second
	if saveInterval != 0 {
		tickerTime = saveInterval
	}

	ticker := time.NewTicker(tickerTime)
	txToPublish := []TransactionSchema{}

	go func() {
		for {
			if ctx.IsLive && ctx.ShouldSaveTransactions {
				select {
				case <-ticker.C:
					ctx.SaveTransactionsHandler.SaveTransactions(txToPublish)
					txToPublish = []TransactionSchema{}
				case tx := <-ctx.TransactionChan:
					txToPublish = append(txToPublish, tx)
				}
			}
		}
	}()
}
