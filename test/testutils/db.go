package testutils

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	gamedb "github.com/curio-research/keystone/db"
	"github.com/curio-research/keystone/state"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
)

var testSQLDSN string

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}
	testSQLDSN = os.Getenv("SQL_DSN")

	txdb.Register("txdb", "mysql", testSQLDSN)
}

func SetupTestDB(t *testing.T, testGameID string, deleteTables bool, accessors map[interface{}]*state.TableBaseAccessor[any]) (*gamedb.MySQLSaveStateHandler, *gamedb.MySQLSaveTransactionHandler, *sql.DB) {
	var db *sql.DB
	db, err := sql.Open("txdb", testSQLDSN)
	if err != nil {
		require.Nil(t, err)
	}
	require.Nil(t, db.Ping())

	if deleteTables {
		DeleteAllTables(t, db)
	}

	sqlDialector := mysql.New(mysql.Config{Conn: db})
	mySQLSaveStateHandler, mySQLSaveTxHandler, err := gamedb.SQLHandlersFromDialector(sqlDialector, testGameID, accessors)
	require.Nil(t, err)

	return mySQLSaveStateHandler, mySQLSaveTxHandler, db
}

func DeleteAllTables(t *testing.T, db *sql.DB) {
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
