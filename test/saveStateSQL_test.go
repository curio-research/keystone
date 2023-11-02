package test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/curio-research/keystone/test/testutils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

var sqlDSN string

func TestMySQLSaveStateHandler(t *testing.T) {
	testutils.SkipTestIfShort(t)

	handler, _, db := testutils.SetupMySQLTestDB(t, testGameID1, true, testSchemaToAccessors)
	defer db.Close()

	coreTestSaveStateHandler(t, handler)
}

func TestMySQLSaveStateHandler_Removal(t *testing.T) {
	testutils.SkipTestIfShort(t)

	mySQLStateHandler, _, db := testutils.SetupMySQLTestDB(t, testGameID1, true, testSchemaToAccessors)
	defer db.Close()

	coreTestSaveStateRemovalHandler(t, mySQLStateHandler)
}

func TestMySQLSaveStateHandler_NestedStructs(t *testing.T) {
	testutils.SkipTestIfShort(t)

	mySQLStateHandler, _, db := testutils.SetupMySQLTestDB(t, testGameID1, true, testSchemaToAccessors)
	defer db.Close()

	coreTestSaveStateWithNestedStructsHandler(t, mySQLStateHandler)
}

func TestMySQLRestoreStateFromTxs(t *testing.T) {
	testutils.SkipTestIfShort(t)

	_, mySQLTxHandler, db := testutils.SetupMySQLTestDB(t, testGameID2, true, testSchemaToAccessors)
	defer db.Close()

	coreTestRestoreStateFromTransactionsHandler(t, mySQLTxHandler)

}

func TestMySQLMultipleGames_SaveState(t *testing.T) {
	testutils.SkipTestIfShort(t)

	saveStateHandler1, saveTxHandler1, db1 := testutils.SetupMySQLTestDB(t, testGameID1, true, testSchemaToAccessors)
	saveStateHandler2, saveTxHandler2, db2 := testutils.SetupMySQLTestDB(t, testGameID2, false, testSchemaToAccessors)

	coreTestMultipleGamesSaveState(t, saveStateHandler1, saveTxHandler1, db1, saveStateHandler2, saveTxHandler2, db2)
}

func TestMultipleGames_SaveTx(t *testing.T) {
	testutils.SkipTestIfShort(t)

	saveStateHandler1, saveTxHandler1, db1 := testutils.SetupMySQLTestDB(t, testGameID1, true, testSchemaToAccessors)
	saveStateHandler2, saveTxHandler2, db2 := testutils.SetupMySQLTestDB(t, testGameID2, false, testSchemaToAccessors)

	coreTestMultipleGamesSaveTransactions(t, saveStateHandler1, saveTxHandler1, db1, saveStateHandler2, saveTxHandler2, db2)
}

func deleteAllTables(t *testing.T, db *sql.DB) {
	rows, err := db.Query("SHOW TABLES")
	require.Nil(t, err)
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		require.Nil(t, rows.Scan(&table))
		tables = append(tables, table)
	}

	// Drop each table
	for _, table := range tables {
		_, err = db.Exec(fmt.Sprintf("DROP TABLE %s", table))
		if err != nil {
			fmt.Println("Failed to drop table", table, "err", err)
		}
	}
}
