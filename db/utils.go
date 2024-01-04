package db

import (
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// initialize and set mySQL handlers to server context
func MySQLHandlers(ctx *server.EngineCtx, mySQLDSN string, accessors map[interface{}]*state.TableBaseAccessor[any]) (*SaveStateHandler, *SaveTransactionHandler, error) {
	dialector := mysql.Open(mySQLDSN)
	return SQLHandlersWithDialector(ctx, dialector, accessors)
}

// initialize SQLite handler through file path
func SQLiteHandlers(ctx *server.EngineCtx, sqliteDBFilePath string, accessors map[interface{}]*state.TableBaseAccessor[any]) (*SaveStateHandler, *SaveTransactionHandler, error) {
	dialector := sqlite.Open(sqliteDBFilePath)
	return SQLHandlersWithDialector(ctx, dialector, accessors)
}

// initialize SQLite handlers through dialector
func SQLHandlersWithDialector(ctx *server.EngineCtx, dialector gorm.Dialector, accessors map[interface{}]*state.TableBaseAccessor[any]) (*SaveStateHandler, *SaveTransactionHandler, error) {
	saveStateHandler, saveTransactionsHandler, err := SQLHandlersFromDialector(dialector, ctx.GameId, accessors)
	if err != nil {
		return nil, nil, err
	}

	return saveStateHandler, saveTransactionsHandler, nil
}

func SQLHandlersFromDialector(dialector gorm.Dialector, gameId string, accessors map[interface{}]*state.TableBaseAccessor[any]) (*SaveStateHandler, *SaveTransactionHandler, error) {
	saveStateHandler, err := SQLSaveStateHandler(dialector, gameId, accessors)
	if err != nil {
		return nil, nil, err
	}

	txHandler, err := SQLSaveTransactionHandler(dialector, gameId)
	if err != nil {
		return nil, nil, err
	}

	return saveStateHandler, txHandler, nil
}

func gormOpts(gameID string) *gorm.Config {
	return &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix(gameID),
			SingularTable: true, // this removes the "s" at the end of table names automatically added by GORM
			NoLowerCase:   true,
		},
	}
}

func tableNameWithPrefix(tableName, gameID string) string {
	return tablePrefix(gameID) + tableName
}

func tablePrefix(gameID string) string {
	return gameID + "_"
}
