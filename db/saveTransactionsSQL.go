package db

import (
	"fmt"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/utils"
	"gorm.io/gorm"
)

const restoreTickRate = 100

type MySQLSaveTransactionHandler struct {
	transactionTable *SQLTransactionTable
	randSeed         int
	gameId           string
}

func newSQLSaveTransactionHandler(dialector gorm.Dialector, randSeed int, gameID string) (*MySQLSaveTransactionHandler, error) {
	db, err := gorm.Open(dialector, gormOpts(gameID))
	if err != nil {
		return nil, err
	}

	txTable, err := NewTransactionTable(db)
	if err != nil {
		return nil, err
	}

	handler := &MySQLSaveTransactionHandler{
		transactionTable: txTable,
		randSeed:         randSeed,
		gameId:           gameID,
	}
	return handler, nil
}

func (h *MySQLSaveTransactionHandler) SaveTransactions(ctx *server.EngineCtx, transactions []server.TransactionSchema) error {
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

func (h *MySQLSaveTransactionHandler) RestoreStateFromTxs(ctx *server.EngineCtx, tickNumber int, _ string) error {
	gw := ctx.World
	for _, table := range gw.Tables {
		if len(table.EntityToValue) != 0 {
			return fmt.Errorf("table %s is not empty", table.Name)
		}
	}

	entries, err := h.transactionTable.GetEntriesUntilTick(tickNumber)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		server.AddSystemTransaction(gw, entry.Tick, entry.Type, entry.Data, "", false)
	}
	utils.TickWorldForward(ctx, tickNumber)

	return nil
}
