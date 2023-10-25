package test

import (
	"database/sql"
	"fmt"
	"github.com/curio-research/keystone/db"
	gamedb "github.com/curio-research/keystone/db"
	"os"
	"testing"

	"github.com/curio-research/keystone/utils"

	"github.com/DATA-DOG/go-txdb"
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"

	"github.com/joho/godotenv"
)

var sqlDSN string

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		panic(err)
	}
	sqlDSN = os.Getenv("SQL_DSN")

	txdb.Register("txdb", "mysql", sqlDSN)
}

func TestMySQLSaveStateHandler(t *testing.T) {
	mySQLStateHandler, _, db := setupTestDB(t, testGameID1, true, testSchemaToAccessors)
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

func TestMySQLSaveStateHandler_Removal(t *testing.T) {
	mySQLStateHandler, _, db := setupTestDB(t, testGameID1, true, testSchemaToAccessors)
	defer db.Close()

	var player1Entity int
	addFirst := true
	addVarsSystem := server.CreateGeneralSystem(func(ctx *server.TransactionCtx[any]) {
		w := ctx.W
		if addFirst == true {
			player1Entity = personTable.AddSpecific(w, 3, Person{
				Name:       testName1,
				MainWallet: testWallet1,
			})
			addFirst = false
		} else {
			personTable.RemoveEntity(w, player1Entity)
		}
	})

	gameEngine := initializeTestWorld(addVarsSystem)
	utils.TickWorldForward(gameEngine, 1)
	assert.Nil(t, mySQLStateHandler.SaveState(gameEngine.PendingStateUpdatesToSave))

	utils.TickWorldForward(gameEngine, 1)
	assert.Nil(t, mySQLStateHandler.SaveState(gameEngine.PendingStateUpdatesToSave))

	newGameEngine := initializeTestWorld()
	newGw := newGameEngine.World
	require.Nil(t, mySQLStateHandler.RestoreState(newGameEngine, ""))

	p1Actual := personTable.Get(newGw, player1Entity)
	assert.Equal(t, "", p1Actual.Name)
	assert.Equal(t, "", p1Actual.MainWallet)
}

func TestMySQLSaveStateHandler_NestedStructs(t *testing.T) {
	mySQLStateHandler, _, db := setupTestDB(t, testGameID1, true, testSchemaToAccessors)
	defer db.Close()

	addVarsSystem := server.CreateGeneralSystem(func(ctx *server.TransactionCtx[any]) {
		w := ctx.W
		embeddedStructTable.AddSpecific(w, testEntity1, EmbeddedStructSchema{
			Emb: NestedStruct{
				Name:  testName1,
				Age:   26,
				Happy: true,
				Pos:   testPos2,
			},
		})
	})

	gameEngine := initializeTestWorld(addVarsSystem)
	utils.TickWorldForward(gameEngine, 1)
	require.Nil(t, mySQLStateHandler.SaveState(gameEngine.PendingStateUpdatesToSave))

	newGameEngine := initializeTestWorld()
	newGw := newGameEngine.World
	require.Nil(t, mySQLStateHandler.RestoreState(newGameEngine, ""))

	esActual := embeddedStructTable.Get(newGw, testEntity1)
	assert.Equal(t, testEntity1, esActual.Id)

	embStruct := esActual.Emb
	assert.Equal(t, testPos2, embStruct.Pos)
	assert.Equal(t, true, embStruct.Happy)
	assert.Equal(t, 26, embStruct.Age)
	assert.Equal(t, testName1, embStruct.Name)
}

func TestMySQLRestoreStateFromTxs(t *testing.T) {
	_, mySQLTxHandler, db := setupTestDB(t, testGameID2, true, testSchemaToAccessors)
	defer db.Close()

	var p1Entity, p2Entity, p3Entity = testEntity1, testEntity2, testEntity3
	var p1Pos, p2Pos, p3Pos = testPos1, testPos2, testPos3

	// General system always gets called, so requests don't need to be saved
	initializePersonSystem := server.CreateGeneralSystem(func(ctx *server.TransactionCtx[any]) {
		if ctx.GameCtx.GameTick.TickNumber == 1 {
			p1Entity = personTable.AddSpecific(ctx.W, p1Entity, Person{
				Name:     testName1,
				Position: p1Pos,
				Id:       p1Entity,
			})

			p2Entity = personTable.AddSpecific(ctx.W, p2Entity, Person{
				Name:     testName2,
				Position: p2Pos,
				Id:       p2Entity,
			})

			p3Entity = personTable.AddSpecific(ctx.W, p3Entity, Person{
				Name:     testName3,
				Position: p3Pos,
				Id:       p3Entity,
			})
		}
	})

	type MovePersonRequest struct {
		TargetEntity int
		NewPosition  state.Pos
	}

	updatePersonSystem := server.CreateSystemFromRequestHandler(func(ctx *server.TransactionCtx[MovePersonRequest]) {
		req := ctx.Req

		person := personTable.Get(ctx.W, req.TargetEntity)
		person.Position = req.NewPosition

		personTable.Set(ctx.W, req.TargetEntity, person)
	})

	newGameEngine := func(t *testing.T) (*server.EngineCtx, *state.GameWorld) {
		gameEngine := initializeTestWorld(initializePersonSystem, updatePersonSystem)
		return gameEngine, gameEngine.World
	}

	initialGameEngine, initialGameWorld := newGameEngine(t)

	// transactions for tick 2
	p1Pos2 := testPos4
	p2Pos2 := testPos5

	server.QueueTxAtTime(initialGameWorld, 2, MovePersonRequest{
		TargetEntity: p1Entity,
		NewPosition:  p1Pos2,
	}, "", true)
	server.QueueTxAtTime(initialGameWorld, 2, MovePersonRequest{
		TargetEntity: p2Entity,
		NewPosition:  p2Pos2,
	}, "", true)

	// transactions for tick 3
	p1Pos3 := testPos6
	p3Pos2 := testPos7
	server.QueueTxAtTime(initialGameWorld, 3, MovePersonRequest{
		TargetEntity: p1Entity,
		NewPosition:  p1Pos3,
	}, "", false) // to see that internal requests are not being added to diffs
	server.QueueTxAtTime(initialGameWorld, 3, MovePersonRequest{
		TargetEntity: p3Entity,
		NewPosition:  p3Pos2,
	}, "", true)

	// apply transactions to the world
	utils.TickWorldForward(initialGameEngine, 3)
	require.Nil(t, mySQLTxHandler.SaveTransactions(initialGameEngine, initialGameEngine.TransactionsToSave))

	// reinitializing tick 1
	newCtx, newGw := newGameEngine(t)
	err := mySQLTxHandler.RestoreStateFromTxs(newCtx, 1, "")
	require.Nil(t, err)

	p1 := personTable.Get(newGw, p1Entity)
	assert.Equal(t, p1Pos, p1.Position)

	p2 := personTable.Get(newGw, p2Entity)
	assert.Equal(t, p2Pos, p2.Position)

	p3 := personTable.Get(newGw, p3Entity)
	assert.Equal(t, p3Pos, p3.Position)

	// reinitializing tick 2
	newCtx, newGw = newGameEngine(t)
	err = mySQLTxHandler.RestoreStateFromTxs(newCtx, 2, "")
	require.Nil(t, err)

	p1 = personTable.Get(newGw, p1Entity)
	assert.Equal(t, p1Pos2, p1.Position)

	p2 = personTable.Get(newGw, p2Entity)
	assert.Equal(t, p2Pos2, p2.Position)

	p3 = personTable.Get(newGw, p3Entity)
	assert.Equal(t, p3Pos, p3.Position)

	// reinitializing tick 3
	newCtx, newGw = newGameEngine(t)
	err = mySQLTxHandler.RestoreStateFromTxs(newCtx, 3, "")
	require.Nil(t, err)

	p1 = personTable.Get(newGw, p1Entity)
	assert.Equal(t, p1Pos2, p1.Position)

	p2 = personTable.Get(newGw, p2Entity)
	assert.Equal(t, p2Pos2, p2.Position)

	p3 = personTable.Get(newGw, p3Entity)
	assert.Equal(t, p3Pos2, p3.Position)
}

func TestMultipleGames_SaveState(t *testing.T) {
	game1System := server.CreateGeneralSystem(func(ctx *server.TransactionCtx[any]) {
		personTable.AddSpecific(ctx.W, 69, Person{
			Name: testName1,
		})
	})

	game2System := server.CreateGeneralSystem(func(ctx *server.TransactionCtx[any]) {
		personTable.AddSpecific(ctx.W, 69, Person{
			Name: testName2,
		})
	})

	newGameEngine := func(t *testing.T, system server.TickSystemFunction, gameID string) *server.EngineCtx {
		gameEngine := initializeTestWorld(system)
		gameEngine.GameId = gameID

		return gameEngine
	}

	game1 := newGameEngine(t, game1System, testGameID1)
	game2 := newGameEngine(t, game2System, testGameID2)

	saveStateHandler1, saveTxHandler1, db1 := setupTestDB(t, testGameID1, true, testSchemaToAccessors)
	defer db1.Close()

	game1.SaveStateHandler = saveStateHandler1
	game1.SaveTransactionsHandler = saveTxHandler1

	saveStateHandler2, saveTxHandler2, db2 := setupTestDB(t, testGameID2, false, testSchemaToAccessors)
	defer db2.Close()

	game2.SaveStateHandler = saveStateHandler2
	game2.SaveTransactionsHandler = saveTxHandler2

	utils.TickWorldForward(game1, 1)
	utils.TickWorldForward(game2, 1)

	game1.SaveStateHandler.SaveState(game1.PendingStateUpdatesToSave)
	game2.SaveStateHandler.SaveState(game2.PendingStateUpdatesToSave)

	newGameEngine1 := initializeTestWorld()
	newGw1 := newGameEngine1.World

	newGameEngine2 := initializeTestWorld()
	newGw2 := newGameEngine2.World

	require.Nil(t, saveStateHandler1.RestoreState(newGameEngine1, ""))
	require.Nil(t, saveStateHandler2.RestoreState(newGameEngine2, ""))

	player1 := personTable.Get(newGw1, 69)
	fmt.Println(player1)
	assert.Equal(t, testName1, player1.Name)

	player2 := personTable.Get(newGw2, 69)
	assert.Equal(t, testName2, player2.Name)
}

func TestMultipleGames_SaveTx(t *testing.T) {
	game1System := server.CreateGeneralSystem(func(ctx *server.TransactionCtx[any]) {
		tickNumber := ctx.GameCtx.GameTick.TickNumber
		w := ctx.W
		if tickNumber == 1 {
			personTable.AddSpecific(w, testEntity1, Person{
				Name:       testName1,
				MainWallet: testWallet1,
			})
		} else if tickNumber == 2 {
			player := personTable.Get(w, testEntity1)
			player.MainWallet = testWallet2
			personTable.Set(w, testEntity1, player)
		}
	})

	game2System := server.CreateGeneralSystem(func(ctx *server.TransactionCtx[any]) {
		tickNumber := ctx.GameCtx.GameTick.TickNumber
		w := ctx.W

		if tickNumber == 1 {
			personTable.AddSpecific(w, testEntity1, Person{
				Name:       testName2,
				MainWallet: testWallet1,
			})
		} else if tickNumber == 2 {
			personTable.AddSpecific(w, testEntity2, Person{
				Name:       testName1,
				MainWallet: testWallet2,
			})
		}
	})

	newGameEngine := func(t *testing.T, system server.TickSystemFunction, gameID string) *server.EngineCtx {
		gameEngine := initializeTestWorld()

		gameEngine.GameTick.Schedule.AddTickSystem(1, system)
		gameEngine.GameId = gameID

		return gameEngine
	}

	game1 := newGameEngine(t, game1System, testGameID1)
	game2 := newGameEngine(t, game2System, testGameID2)

	saveStateHandler1, saveTxHandler1, db1 := setupTestDB(t, testGameID1, true, testSchemaToAccessors)
	defer db1.Close()

	game1.SaveStateHandler = saveStateHandler1
	game1.SaveTransactionsHandler = saveTxHandler1

	saveStateHandler2, saveTxHandler2, db2 := setupTestDB(t, testGameID2, false, testSchemaToAccessors)
	defer db2.Close()

	game2.SaveStateHandler = saveStateHandler2
	game2.SaveTransactionsHandler = saveTxHandler2

	utils.TickWorldForward(game1, 2)
	utils.TickWorldForward(game2, 2)

	game1.SaveTransactionsHandler.SaveTransactions(game1, game1.TransactionsToSave)
	game2.SaveTransactionsHandler.SaveTransactions(game2, game2.TransactionsToSave)

	newGameEngine1 := newGameEngine(t, game1System, "")
	newGameEngine2 := newGameEngine(t, game2System, "")

	require.Nil(t, saveTxHandler1.RestoreStateFromTxs(newGameEngine1, 1, ""))
	require.Nil(t, saveTxHandler2.RestoreStateFromTxs(newGameEngine2, 1, ""))

	player1 := personTable.Get(newGameEngine1.World, testEntity1)
	assert.Equal(t, testName1, player1.Name)
	assert.Equal(t, testWallet1, player1.MainWallet)

	player2 := personTable.Get(newGameEngine2.World, testEntity1)
	assert.Equal(t, testName2, player2.Name)
	assert.Equal(t, testWallet1, player2.MainWallet)

	newGameEngine1 = newGameEngine(t, game1System, "")
	newGameEngine2 = newGameEngine(t, game2System, "")

	require.Nil(t, saveTxHandler1.RestoreStateFromTxs(newGameEngine1, 2, ""))
	require.Nil(t, saveTxHandler2.RestoreStateFromTxs(newGameEngine2, 2, ""))

	player1 = personTable.Get(newGameEngine1.World, testEntity1)
	assert.Equal(t, testName1, player1.Name)
	assert.Equal(t, testWallet2, player1.MainWallet)

	player2 = personTable.Get(newGameEngine2.World, testEntity1)
	assert.Equal(t, testName2, player2.Name)
	assert.Equal(t, testWallet1, player2.MainWallet)

	player3 := personTable.Get(newGameEngine2.World, testEntity2)
	assert.Equal(t, testName1, player3.Name)
	assert.Equal(t, testWallet2, player3.MainWallet)
}

func setupTestDB(t *testing.T, testGameID string, deleteTables bool, accessors map[interface{}]*state.TableBaseAccessor[any]) (*db.MySQLSaveStateHandler, *db.MySQLSaveTransactionHandler, *sql.DB) {
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
	mySQLSaveStateHandler, mySQLSaveTxHandler, err := gamedb.SQLHandlersFromDialector(sqlDialector, testGameID, 0, accessors)
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

	fmt.Println("-> Existing tables have been removed")
}
