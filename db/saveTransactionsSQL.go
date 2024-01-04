package db

import (
	"fmt"

	"github.com/curio-research/keystone/server"
	"gorm.io/gorm"
)

type SaveTransactionHandler struct {
	transactionTable *SQLTransactionTable
	gameId           string
}

func SQLSaveTransactionHandler(dialector gorm.Dialector, gameID string) (*SaveTransactionHandler, error) {
	db, err := gorm.Open(dialector, gormOpts(gameID))
	if err != nil {
		return nil, err
	}

	txTable, err := NewTransactionTable(db)
	if err != nil {
		return nil, err
	}

	handler := &SaveTransactionHandler{
		transactionTable: txTable,
		gameId:           gameID,
	}
	return handler, nil
}

func (h *SaveTransactionHandler) SaveTransactions(transactions []server.TransactionSchema) error {
	updatesForSql := []TransactionSQLFormat{}
	for _, transaction := range transactions {
		updatesForSql = append(updatesForSql, TransactionSQLFormat{
			GameId:        h.gameId,
			UnixTimestamp: transaction.UnixTimestamp,
			Tick:          transaction.TickNumber,
			Data:          transaction.Data,
			Type:          transaction.Type,
		})
	}

	return h.transactionTable.AddEntries(updatesForSql...)
}

// initialize the world to the initial state before calling
func (h *SaveTransactionHandler) RestoreStateFromTxs(ctx *server.EngineCtx, tickNumber int, _ string) error {
	if ctx.GameTick.TickNumber != 1 {
		return fmt.Errorf("game tick was not reset to 1")
	}

	gw := ctx.World
	entries, err := h.transactionTable.GetEntriesUntilTick(tickNumber)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		server.AddSystemTransaction(gw, entry.Tick, entry.Type, entry.Data, "", false)
	}
	server.TickWorldForward(ctx, tickNumber)

	return nil
}
