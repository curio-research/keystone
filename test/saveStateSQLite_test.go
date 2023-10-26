package test

import (
	"database/sql"
	"testing"

	"github.com/curio-research/keystone/db"
	gamedb "github.com/curio-research/keystone/db"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"github.com/curio-research/keystone/utils"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestSQLiteDB(t *testing.T, testGameID string, deleteTables bool, accessors map[interface{}]*state.TableBaseAccessor[any]) (*db.MySQLSaveStateHandler, *db.MySQLSaveTransactionHandler, *sql.DB) {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		require.Nil(t, err)
	}
	require.Nil(t, db.Ping())

	if deleteTables {
		// TODO: pass in DB file path
		deleteAllTablesSQLite(t)
	}

	gormDB, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	dialector := gormDB.Dialector

	mySQLSaveStateHandler, mySQLSaveTxHandler, err := gamedb.SQLHandlersFromDialector(dialector, testGameID, 0, accessors)
	require.Nil(t, err)

	return mySQLSaveStateHandler, mySQLSaveTxHandler, db
}

func TestSQLiteSaveStateHandler(t *testing.T) {
	mySQLStateHandler, _, db := setupTestSQLiteDB(t, testGameID1, true, testSchemaToAccessors)
	defer db.Close()

	var player1Entity, player2Entity, nt1Entity, nt2Entity int
	addVarsSystem := server.CreateGeneralSystem(func(ctx *server.TransactionCtx[any]) {
		w := ctx.W
		player1Entity = personTable.AddSpecific(w, 1, Person{
			Name:       testName1,
			MainWallet: testWallet1,
			Position:   testPos1,
		})
		player2Entity = personTable.AddSpecific(w, 2, Person{
			Name:       testName2,
			MainWallet: testWallet2,
		})
		nt1Entity = tokenTable.AddSpecific(w, 3, Token{
			OriginalOwnerId: 7,
			OwnerId:         player1Entity,
		})
		nt2Entity = tokenTable.AddSpecific(w, testEntity1, Token{
			OriginalOwnerId: 8,
			OwnerId:         player2Entity,
		})
		assert.Equal(t, testEntity1, nt2Entity)
	})

	gameEngine := initializeTestWorld(addVarsSystem)
	utils.TickWorldForward(gameEngine, 1)
	require.Nil(t, mySQLStateHandler.SaveState(gameEngine.PendingStateUpdatesToSave))

	newGameEngine := initializeTestWorld()
	newGw := newGameEngine.World
	require.Nil(t, mySQLStateHandler.RestoreState(newGameEngine, ""))

	p1Actual := personTable.Get(newGw, player1Entity)
	assert.Equal(t, testName1, p1Actual.Name)
	assert.Equal(t, testWallet1, p1Actual.MainWallet)
	assert.Equal(t, testPos1, p1Actual.Position)
	assert.Equal(t, player1Entity, p1Actual.Id)

	p2Actual := personTable.Get(newGw, player2Entity)
	assert.Equal(t, testName2, p2Actual.Name)
	assert.Equal(t, testWallet2, p2Actual.MainWallet)
	assert.Equal(t, player2Entity, p2Actual.Id)

	nt1Actual := tokenTable.Get(newGw, nt1Entity)
	assert.Equal(t, player1Entity, nt1Actual.OwnerId)
	assert.Equal(t, nt1Entity, nt1Actual.Id)

	nt2Actual := tokenTable.Get(newGw, nt2Entity)
	assert.Equal(t, player2Entity, nt2Actual.OwnerId)
	assert.Equal(t, testEntity1, nt2Actual.Id)
}
