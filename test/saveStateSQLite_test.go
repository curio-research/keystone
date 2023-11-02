package test

import (
	"github.com/curio-research/keystone/test/testutils"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestSQLiteSaveStateHandler(t *testing.T) {
	testutils.ResetSQLiteTestDB()
	mySQLStateHandler, _, db := testutils.SetupSQLiteTestDB(t, testGameID1, true, testSchemaToAccessors)
	defer db.Close()

	coreTestSaveStateHandler(t, mySQLStateHandler)
}

func TestSQLiteSaveStateHandler_Removal(t *testing.T) {
	testutils.ResetSQLiteTestDB()
	mySQLStateHandler, _, db := testutils.SetupSQLiteTestDB(t, testGameID1, true, testSchemaToAccessors)
	defer db.Close()

	coreTestSaveStateRemovalHandler(t, mySQLStateHandler)
}

func TestSQLiteSaveStateHandler_NestedStructs(t *testing.T) {
	testutils.ResetSQLiteTestDB()
	mySQLStateHandler, _, db := testutils.SetupSQLiteTestDB(t, testGameID1, true, testSchemaToAccessors)
	defer db.Close()

	coreTestSaveStateWithNestedStructsHandler(t, mySQLStateHandler)
}

func TestSQLitesRestoreStateFromTxs(t *testing.T) {
	testutils.ResetSQLiteTestDB()
	_, mySQLTxHandler, db := testutils.SetupSQLiteTestDB(t, testGameID2, true, testSchemaToAccessors)
	defer db.Close()

	coreTestRestoreStateFromTransactionsHandler(t, mySQLTxHandler)
}

func TestSQLiteMultipleGames_SaveState(t *testing.T) {
	testutils.ResetSQLiteTestDB()
	saveStateHandler1, saveTxHandler1, db1 := testutils.SetupSQLiteTestDB(t, testGameID1, true, testSchemaToAccessors)
	saveStateHandler2, saveTxHandler2, db2 := testutils.SetupSQLiteTestDB(t, testGameID2, false, testSchemaToAccessors)

	coreTestMultipleGamesSaveState(t, saveStateHandler1, saveTxHandler1, db1, saveStateHandler2, saveTxHandler2, db2)
}

func TestSQLiteMultipleGames_SaveTx(t *testing.T) {
	testutils.ResetSQLiteTestDB()
	saveStateHandler1, saveTxHandler1, db1 := testutils.SetupSQLiteTestDB(t, testGameID1, true, testSchemaToAccessors)
	saveStateHandler2, saveTxHandler2, db2 := testutils.SetupSQLiteTestDB(t, testGameID2, false, testSchemaToAccessors)

	coreTestMultipleGamesSaveTransactions(t, saveStateHandler1, saveTxHandler1, db1, saveStateHandler2, saveTxHandler2, db2)
}

// local sqlite db for testing
