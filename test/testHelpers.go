package test

import (
	"sync"

	"github.com/curio-research/keystone/core"
	"github.com/curio-research/keystone/server"
)

var (
	// Person testing vars
	testName1 = "Alice"
	testName2 = "Bob"
	testName3 = "Francisco"

	testPos1 = core.Pos{1, 2}
	testPos2 = core.Pos{5, 6}
	testPos3 = core.Pos{1, 3}
	testPos4 = core.Pos{2, 1}
	testPos5 = core.Pos{2, 2}
	testPos6 = core.Pos{2, 3}
	testPos7 = core.Pos{3, 1}

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
	Position   core.Pos `gorm:"embedded"`
	BookId     int      // TODO: if we can automatically solve the linkage that'd be OP
	Id         int      `gorm:"primaryKey"`
}

type Book struct {
	Title   string
	Author  string
	OwnerID int
	Id      int `gorm:"primaryKey"`
}

type Token struct {
	OriginalOwnerId int
	OwnerId         int
	Id              int `gorm:"primaryKey"`
}

type NestedStruct struct {
	Name  string
	Age   int
	Happy bool
	Pos   core.Pos `gorm:"embedded"`
}

type EmbeddedStructSchema struct {
	Emb NestedStruct `gorm:"embedded"`
	Id  int          `gorm:"primaryKey"`
}

var personTable = core.NewTableAccessor[Person]()
var bookTable = core.NewTableAccessor[Book]()
var tokenTable = core.NewTableAccessor[Token]()
var embeddedStructTable = core.NewTableAccessor[EmbeddedStructSchema]()

var testSchemaToAccessors = map[interface{}]*core.TableBaseAccessor[any]{
	&Person{}:                   (*core.TableBaseAccessor[any])(personTable),
	&Book{}:                     (*core.TableBaseAccessor[any])(bookTable),
	&Token{}:                    (*core.TableBaseAccessor[any])(tokenTable),
	&server.TransactionSchema{}: (*core.TableBaseAccessor[any])(server.TransactionTable),
	&EmbeddedStructSchema{}:     (*core.TableBaseAccessor[any])(embeddedStructTable),
}

func testRegisterTables(w *core.GameWorld) {
	for _, accessor := range testSchemaToAccessors {
		w.AddTable(accessor)
	}
}

type testPersonRequests struct {
	People   []testPersonRequest `json:"People"`
	PlayerID int64               `json:"playerID"`
}

type testPersonRequest struct {
	OP        core.TableOperationType `json:"OP"`
	Entity    int                     `json:"Entity"`
	Val       Person                  `json:"Val"`
	Id        int                     `json:"Id"`
	SendError bool                    `json:"SendError"`
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
	gameTick := server.NewGameTick(100)

	// initiate an empty tick schedule
	tickSchedule := server.NewTickSchedule()
	for _, system := range systems {
		tickSchedule.AddTickSystem(0, system)
	}
	gameTick.Schedule = tickSchedule

	gameWorld := core.NewWorld()
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
