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
	testPos3 = state.Pos{1, 3}
	testPos4 = state.Pos{2, 1}
	testPos5 = state.Pos{2, 2}
	testPos6 = state.Pos{2, 3}
	testPos7 = state.Pos{3, 1}

	testEntity1 = 69
	testEntity2 = 70
	testEntity3 = 71

	testWallet1 = "wallet1"
	testWallet2 = "wallet2"

	testAddress1 = "123 Vitalik Drive"
	testAddress2 = "Home"

	testGameID1 = "game1"
	testGameID2 = "game2"

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
	Name       string
	MainWallet string
	Age        int
	Address    string
	Position   state.Pos `gorm:"embedded"`
	BookId     int
	Id         int `gorm:"primaryKey;autoIncrement:false"`
}

type Book struct {
	Title   string
	Author  string
	OwnerID int
	Id      int `gorm:"primaryKey;autoIncrement:false"`
}

type Token struct {
	OriginalOwnerId int
	OwnerId         int
	Id              int `gorm:"primaryKey;autoIncrement:false"`
}

type NestedStruct struct {
	Name  string
	Age   int
	Happy bool
	Pos   state.Pos `gorm:"embedded"`
}

type EmbeddedStructSchema struct {
	Emb NestedStruct `gorm:"embedded"`
	Id  int          `gorm:"primaryKey;autoIncrement:false"`
}

var personTable = state.NewTableAccessor[Person]()
var bookTable = state.NewTableAccessor[Book]()
var tokenTable = state.NewTableAccessor[Token]()
var embeddedStructTable = state.NewTableAccessor[EmbeddedStructSchema]()

var testSchemaToAccessors = map[interface{}]*state.TableBaseAccessor[any]{
	&Person{}:                   (*state.TableBaseAccessor[any])(personTable),
	&Book{}:                     (*state.TableBaseAccessor[any])(bookTable),
	&Token{}:                    (*state.TableBaseAccessor[any])(tokenTable),
	&server.TransactionSchema{}: (*state.TableBaseAccessor[any])(server.TransactionTable),
	&EmbeddedStructSchema{}:     (*state.TableBaseAccessor[any])(embeddedStructTable),
}

func testRegisterTables(w *state.GameWorld) {
	for _, accessor := range testSchemaToAccessors {
		w.AddTable(accessor)
	}
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

func initializeTestWorld(systems ...server.TickSystemFunction) *server.EngineCtx {
	// initiate an empty tick schedule
	tickSchedule := server.NewTickSchedule()
	for _, system := range systems {
		tickSchedule.AddTickSystem(0, system)
	}

	gameTick := server.NewGameTick(server.TickRate)
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

	return gameCtx
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
