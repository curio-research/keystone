package test

import (
	"sync"

	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
)

var (
	// Person testing vars
	testName1 = "Alice"
	testName2 = "Bob"
	testName3 = "Francisco"

	testPos1 = state.Pos{1, 2}
	testPos2 = state.Pos{5, 6}

	testAddress1 = "123 Vitalik Drive"
	testAddress2 = "Home"

	testAge1 = 26
	testAge2 = 24

	// Book testing vars
	testBookTitle1  = "Cat in a Hat"
	testBookAuthor1 = "Dr. Seuss"

	testBookTitle2  = "Fault in Our Stars"
	testBookAuthor2 = "John Greene"

	testBookTitle3  = "The Order of the Phoenix"
	testBookAuthor3 = "J.K. Rowling"
)

var testErrMsg = "error in system"

type Person struct {
	Name     string
	Age      int
	Address  string
	Position state.Pos
	BookId   int // TODO: if we can automatically solve the linkage that'd be OP
	Id       int
}

type Book struct {
	Title   string
	Author  string
	OwnerID int
	Id      int
}

var personTable = state.NewTableAccessor[Person]()
var bookTable = state.NewTableAccessor[Book]()

func testRegisterTables(w *state.GameWorld) {
	w.AddTables(personTable, bookTable, server.TransactionTable)
}

type testPersonRequests struct {
	People   []testPersonRequest `json:"People"`
	PlayerID int64               `json:"playerID"`
}

type testPersonRequest struct {
	OP        state.TableOperationType `json:"OP"`
	Entity    int                      `json:"Entity"`
	Val       Person                   `json:"Val"`
	Id        int                      `json:"Id"`
	SendError bool                     `json:"SendError"`
}

type testIdentityPayload struct {
	jwtToken string
	playerID int64
}

func (t *testPersonRequests) GetIdentityPayload() testIdentityPayload {
	return testIdentityPayload{
		jwtToken: "",
		playerID: t.PlayerID,
	}
}

func initializeTestWorld(systems ...server.TickSystemFunction) (*server.EngineCtx, error) {
	gameTick := server.NewGameTick(100)

	// initiate an empty tick schedule
	tickSchedule := server.NewTickSchedule()
	for _, system := range systems {
		tickSchedule.AddTickSystem(0, system)
	}
	gameTick.Schedule = tickSchedule

	gameWorld := state.NewWorld()
	testRegisterTables(gameWorld)

	gameCtx := &server.EngineCtx{
		GameId:                 "prototype-game",
		IsLive:                 true,
		World:                  gameWorld,
		GameTick:               gameTick,
		TransactionsToSaveLock: sync.Mutex{},
		ShouldRecordError:      true,
		ErrorLog:               []server.ErrorLog{},
		Mode:                   "dev",
		SystemErrorHandler:     &testErrorHandler{},
		SystemBroadcastHandler: &testBroadcastHandler{},
	}

	return gameCtx, nil
}

type testErrorHandler struct {
}

func (t *testErrorHandler) FormatMessage(transactionUuidIdentifier int, errorMessage string) *server.NetworkMessage {
	return server.NewMessageFromBuffer([]byte(errorMessage))
}

type testBroadcastHandler struct {
}

func (t *testBroadcastHandler) BroadcastMessage(ctx *server.EngineCtx, clientEvents []server.ClientEvent) {
	for _, ev := range clientEvents {
		ctx.ErrorLog = append(ctx.ErrorLog, server.ErrorLog{
			Tick:    ctx.GameTick.TickNumber,
			Message: string(ev.NetworkMessage.ParseToBuffer()),
		})
	}
}
