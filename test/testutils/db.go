package testutils

import (
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
var testSQLiteDBPath = "test.db"

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println("Failed to load .env file")
	}
	testSQLDSN = os.Getenv("SQL_DSN")

	txdb.Register("txdb", "mysql", testSQLDSN)
}

func SetupMySQLTestDB(t *testing.T, testGameID string, deleteTables bool, accessors map[interface{}]*state.TableBaseAccessor[any]) (*gamedb.SaveStateHandler, *gamedb.SaveTransactionHandler, *sql.DB) {
	var db *sql.DB
	db, err := sql.Open("txdb", testSQLDSN)
	if err != nil {
		require.Nil(t, err)
	}
	require.Nil(t, db.Ping())

	if deleteTables {
		deleteAllTablesMySQL(t, db)
	}

	sqlDialector := mysql.New(mysql.Config{Conn: db})
	mySQLSaveStateHandler, mySQLSaveTxHandler, err := gamedb.SQLHandlersFromDialector(sqlDialector, testGameID, accessors)
	require.Nil(t, err)

	return mySQLSaveStateHandler, mySQLSaveTxHandler, db
}

func deleteAllTablesMySQL(t *testing.T, db *sql.DB) {
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

// setup local sqlite test db
func SetupSQLiteTestDB(t *testing.T, testGameID string, deleteTables bool, accessors map[interface{}]*state.TableBaseAccessor[any]) (*gamedb.SaveStateHandler, *gamedb.SaveTransactionHandler, *sql.DB) {
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

func ResetSQLiteTestDB() error {
	dbFileName := testSQLiteDBPath

	// Check if the file exists
	if _, err := os.Stat(dbFileName); err == nil {
		// File exists, so delete it
		err := os.Remove(dbFileName)
		if err != nil {
			return err
		}
	}

	// Create an empty file
	file, err := os.Create(dbFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
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
