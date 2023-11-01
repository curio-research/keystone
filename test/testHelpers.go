package test

import (
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/startup"
	"github.com/curio-research/keystone/state"
	"github.com/curio-research/keystone/utils"
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

type PetKind int

const (
	Dog PetKind = iota
	Cat
)

type Pet struct {
	Kind PetKind `json:"kind"`
	Name string  `json:"name"`
}

type Owner struct {
	Name  string    `json:"name"`
	Age   int       `json:"age"`
	Happy bool      `json:"happy"`
	Pets  []Pet     `json:"pets"`
	Pos   state.Pos `json:"pos" gorm:"embedded"`
}

type Committee struct {
	CommitteeOwners utils.SerializableArray[Owner] `gorm:"serializer:json"`
}

type PetCommunity struct {
	Owners    utils.SerializableArray[Owner] `gorm:"serializer:json"`
	Committee Committee                      `gorm:"embedded"`
	Id        int                            `gorm:"primaryKey;autoIncrement:false"`
}

var personTable = state.NewTableAccessor[Person]()
var bookTable = state.NewTableAccessor[Book]()
var tokenTable = state.NewTableAccessor[Token]()
var petCommunityTable = state.NewTableAccessor[PetCommunity]()

var testSchemaToAccessors = map[interface{}]*state.TableBaseAccessor[any]{
	&Person{}:                   (*state.TableBaseAccessor[any])(personTable),
	&Book{}:                     (*state.TableBaseAccessor[any])(bookTable),
	&Token{}:                    (*state.TableBaseAccessor[any])(tokenTable),
	&server.TransactionSchema{}: (*state.TableBaseAccessor[any])(server.TransactionTable),
	&PetCommunity{}:             (*state.TableBaseAccessor[any])(petCommunityTable),
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
	var tables []state.ITable
	for _, accessor := range testSchemaToAccessors {
		tables = append(tables, accessor)
	}

	ctx := startup.NewGameEngine("test", server.TickRate, tables...)
	for _, system := range systems {
		ctx.GameTick.Schedule.AddSystem(0, system)
	}

	startup.RegisterErrorHandler(ctx, &testErrorHandler{})
	startup.RegisterBroadcastHandler(ctx, &testBroadcastHandler{})

	return ctx
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
