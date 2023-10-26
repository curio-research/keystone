package db

import (
	"github.com/curio-research/keystone/core"
	"github.com/curio-research/keystone/server"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitializeSQLHandlers(ctx *server.EngineCtx, mySQLDSN string, accessors map[interface{}]*core.TableBaseAccessor[any]) error {
	dialector := mysql.Open(mySQLDSN)
	saveStateHandler, saveTransactionsHandler, err := SQLHandlersFromDialector(dialector, ctx.GameId, ctx.RandSeed, accessors)
	if err != nil {
		return err
	}

	ctx.SaveStateHandler = saveStateHandler
	ctx.SaveTransactionsHandler = saveTransactionsHandler
	return nil
}

func SQLHandlersFromDialector(dialector gorm.Dialector, gameId string, randSeed int, accessors map[interface{}]*core.TableBaseAccessor[any]) (*MySQLSaveStateHandler, *MySQLSaveTransactionHandler, error) {
	saveStateHandler, err := newSQLSaveStateHandler(dialector, gameId, accessors)
	if err != nil {
		return nil, nil, err
	}

	txHandler, err := newSQLSaveTransactionHandler(dialector, randSeed, gameId)
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