package test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/curio-research/keystone/db"
	gamedb "github.com/curio-research/keystone/db"

	"github.com/curio-research/keystone/state"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
)

var sqlDSN string

// func init() {
// 	if err := godotenv.Load("../.env"); err != nil {
// 		fmt.Println("Failed to load .env file")
// 		return
// 	}
// 	sqlDSN = os.Getenv("SQL_DSN")

// 	txdb.Register("txdb", "mysql", sqlDSN)
// }

func TestMySQLSaveStateHandler(t *testing.T) {
	handler, _, db := setupMySQLTestDB(t, testGameID1, true, testSchemaToAccessors)
	defer db.Close()

	coreTestSaveStateHandler(t, handler)
}

func TestMySQLSaveStateHandler_Removal(t *testing.T) {
	mySQLStateHandler, _, db := setupMySQLTestDB(t, testGameID1, true, testSchemaToAccessors)
	defer db.Close()

	coreTestSaveStateRemovalHandler(t, mySQLStateHandler)
}

func TestMySQLSaveStateHandler_NestedStructs(t *testing.T) {
	mySQLStateHandler, _, db := setupMySQLTestDB(t, testGameID1, true, testSchemaToAccessors)
	defer db.Close()

	coreTestSaveStateWithNestedStructsHandler(t, mySQLStateHandler)
}

func TestMySQLRestoreStateFromTxs(t *testing.T) {
	_, mySQLTxHandler, db := setupMySQLTestDB(t, testGameID2, true, testSchemaToAccessors)
	defer db.Close()

	coreTestRestoreStateFromTransactionsHandler(t, mySQLTxHandler)

}

func TestMySQLMultipleGames_SaveState(t *testing.T) {
	saveStateHandler1, saveTxHandler1, db1 := setupMySQLTestDB(t, testGameID1, true, testSchemaToAccessors)
	saveStateHandler2, saveTxHandler2, db2 := setupMySQLTestDB(t, testGameID2, false, testSchemaToAccessors)

	coreTestMultipleGamesSaveState(t, saveStateHandler1, saveTxHandler1, db1, saveStateHandler2, saveTxHandler2, db2)
}

func TestMultipleGames_SaveTx(t *testing.T) {
	saveStateHandler1, saveTxHandler1, db1 := setupMySQLTestDB(t, testGameID1, true, testSchemaToAccessors)
	saveStateHandler2, saveTxHandler2, db2 := setupMySQLTestDB(t, testGameID2, false, testSchemaToAccessors)

	coreTestMultipleGamesSaveTransactions(t, saveStateHandler1, saveTxHandler1, db1, saveStateHandler2, saveTxHandler2, db2)
}

func setupMySQLTestDB(t *testing.T, testGameID string, deleteTables bool, accessors map[interface{}]*state.TableBaseAccessor[any]) (*db.MySQLSaveStateHandler, *db.MySQLSaveTransactionHandler, *sql.DB) {
	var db *sql.DB
	db, err := sql.Open("txdb", sqlDSN)
	if err != nil {
		require.Nil(t, err)
	}
	require.Nil(t, db.Ping())

	if deleteTables {
		deleteAllTables(t, db)
	}

	sqlDialector := mysql.New(mysql.Config{Conn: db})
	mySQLSaveStateHandler, mySQLSaveTxHandler, err := gamedb.SQLHandlersFromDialector(sqlDialector, testGameID, accessors)
	require.Nil(t, err)

	return mySQLSaveStateHandler, mySQLSaveTxHandler, db
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
