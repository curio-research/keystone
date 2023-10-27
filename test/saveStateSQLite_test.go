package test

import (
	"database/sql"
	"testing"

	gamedb "github.com/curio-research/keystone/db"
	"github.com/curio-research/keystone/state"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestSQLiteSaveStateHandler(t *testing.T) {
	mySQLStateHandler, _, db := setupSQLiteTestDB(t, testGameID1, true, testSchemaToAccessors)
	defer db.Close()

	coreTestSaveStateHandler(t, mySQLStateHandler)
}

func TestSQLiteSaveStateHandler_Removal(t *testing.T) {
	mySQLStateHandler, _, db := setupSQLiteTestDB(t, testGameID1, true, testSchemaToAccessors)
	defer db.Close()

	coreTestSaveStateRemovalHandler(t, mySQLStateHandler)
}

func TestSQLiteSaveStateHandler_NestedStructs(t *testing.T) {
	mySQLStateHandler, _, db := setupSQLiteTestDB(t, testGameID1, true, testSchemaToAccessors)
	defer db.Close()

	coreTestSaveStateWithNestedStructsHandler(t, mySQLStateHandler)
}

func TestSQLitesRestoreStateFromTxs(t *testing.T) {
	_, mySQLTxHandler, db := setupSQLiteTestDB(t, testGameID2, true, testSchemaToAccessors)
	defer db.Close()

	coreTestRestoreStateFromTransactionsHandler(t, mySQLTxHandler)
}

func TestSQLiteMultipleGames_SaveState(t *testing.T) {
	saveStateHandler1, saveTxHandler1, db1 := setupSQLiteTestDB(t, testGameID1, true, testSchemaToAccessors)
	saveStateHandler2, saveTxHandler2, db2 := setupSQLiteTestDB(t, testGameID2, false, testSchemaToAccessors)

	coreTestMultipleGamesSaveState(t, saveStateHandler1, saveTxHandler1, db1, saveStateHandler2, saveTxHandler2, db2)
}

func TestSQLiteMultipleGames_SaveTx(t *testing.T) {
	saveStateHandler1, saveTxHandler1, db1 := setupSQLiteTestDB(t, testGameID1, true, testSchemaToAccessors)
	saveStateHandler2, saveTxHandler2, db2 := setupSQLiteTestDB(t, testGameID2, false, testSchemaToAccessors)

	coreTestMultipleGamesSaveTransactions(t, saveStateHandler1, saveTxHandler1, db1, saveStateHandler2, saveTxHandler2, db2)
}

// local sqlite db for testing
var testSQLiteDBPath = "test.db"

// setup local sqlite test db
func setupSQLiteTestDB(t *testing.T, testGameID string, deleteTables bool, accessors map[interface{}]*state.TableBaseAccessor[any]) (*gamedb.MySQLSaveStateHandler, *gamedb.MySQLSaveTransactionHandler, *sql.DB) {
	db, err := sql.Open("sqlite3", testSQLiteDBPath)
	if err != nil {
		require.Nil(t, err)
	}
	require.Nil(t, db.Ping())

	if deleteTables {
		deleteAllTablesSQLite(t)
	}

	gormDB, err := gorm.Open(sqlite.Open(testSQLiteDBPath))

	mySQLSaveStateHandler, mySQLSaveTxHandler, err := gamedb.SQLHandlersFromDialector(gormDB.Dialector, testGameID, accessors)
	require.Nil(t, err)

	return mySQLSaveStateHandler, mySQLSaveTxHandler, db
}

// delete all tables in a sqlite db
func deleteAllTablesSQLite(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(testSQLiteDBPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// get list of table names
	tableNames := getSQLiteTableNames(db)

	// Iterate through the table names and drop each table
	for _, tableName := range tableNames {
		if err := db.Exec("DROP TABLE " + tableName + ";").Error; err != nil {
			panic("Failed to drop table " + tableName + ": " + err.Error())
		}
	}

	// verify that table names array is empty
	updatedTableNames := getSQLiteTableNames(db)
	assert.Equal(t, 0, len(updatedTableNames))

}

func getSQLiteTableNames(db *gorm.DB) []string {
	var tableNames []string
	if err := db.Raw("SELECT name FROM sqlite_master WHERE type='table';").Scan(&tableNames).Error; err != nil {
		panic("Failed to fetch table names: " + err.Error())
	}
	return tableNames
}
