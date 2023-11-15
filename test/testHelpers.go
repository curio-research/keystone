package test

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/server/startup"
	"github.com/curio-research/keystone/state"
	pb_test "github.com/curio-research/keystone/test/proto/pb.test"
	"github.com/curio-research/keystone/test/testutils"
	"github.com/curio-research/keystone/utils"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
)

var p *testutils.PortManager

func init() {
	p = testutils.NewPortManager()
}

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
	Title   string `key:"true"`
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

	ctx := startup.NewGameEngine()
	ctx.SetGameId("test")
	ctx.SetTickRate(20)
	ctx.AddTables(testSchemaToAccessors)

	ctx.SetEmitErrorHandler(&testErrorHandler{})
	ctx.SetEmitEventHandler(&testBroadcastHandler{})

	for _, system := range systems {
		ctx.GameTick.Schedule.AddSystem(0, system)
	}

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

// Start test server
func startTestServer(t *testing.T, mode server.GameMode) (*server.EngineCtx, *websocket.Conn, *http.Server, *testutils.MockErrorHandler, *sql.DB) {
	httpPort, wsPort := p.GetPort(), p.GetPort()

	s, ctx, db, err := testutils.Server(t, mode, wsPort, testSchemaToAccessors)
	require.Nil(t, err)

	ctx.AddSystem(1, TestBookSystem)
	ctx.AddSystem(1, TestRemoveBookSystem)

	ctx.Stream.Start(ctx)

	// Serve HTTP server
	addr := ":" + strconv.Itoa(httpPort)
	httpServer := &http.Server{
		Addr:    addr,
		Handler: s,
	}

	// spin up the HTTP server
	go func() {
		err := httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			t.Errorf("http server closed with unexpected error %v", err)
			return
		}
	}()

	ws, err := testutils.EstablishWsConnection(t, wsPort)
	require.Nil(t, err)

	return ctx, ws, httpServer, ctx.SystemErrorHandler.(*testutils.MockErrorHandler), db
}

var TestBookSystem = server.CreateSystemFromRequestHandler(func(ctx *server.TransactionCtx[*pb_test.C2S_Test]) {
	req := ctx.Req.Data
	w := ctx.W

	playerID := int(req.GetIdentityPayload().GetPlayerId())

	for _, bookInfo := range req.BookInfos {
		switch bookInfo.Op {
		case pb_test.Operation_Add:
			bookTable.Add(w, Book{
				Title:   bookInfo.Title,
				Author:  bookInfo.Author,
				OwnerID: playerID,
			})
		case pb_test.Operation_AddSpecific:
			bookTable.AddSpecific(w, int(bookInfo.Entity), Book{
				Title:   bookInfo.Title,
				Author:  bookInfo.Author,
				OwnerID: playerID,
			})
		case pb_test.Operation_Remove:
			server.QueueTxFromInternal(w, ctx.GameCtx.GameTick.TickNumber+1, server.NewKeystoneTx(testRemoveRequest{
				Title:    bookInfo.Title,
				Author:   bookInfo.Author,
				PlayerID: playerID,
			}, nil), "")
		case pb_test.Operation_Update:
			book := bookTable.Get(w, int(bookInfo.Entity))
			if book.Title == "" {
				ctx.EmitError(fmt.Sprintf("no book to update with entity %v", bookInfo.Entity), []int{playerID})
				return
			}

			book.Title = bookInfo.Title
			book.Author = bookInfo.Author
			bookTable.Set(w, int(bookInfo.Entity), book)
		}
	}
})

type testRemoveRequest struct {
	Author   string
	Title    string
	PlayerID int
}

var TestRemoveBookSystem = server.CreateSystemFromRequestHandler(func(ctx *server.TransactionCtx[testRemoveRequest]) {
	req := ctx.Req.Data
	w := ctx.GameCtx.World

	bookFilter := Book{Author: req.Author, Title: req.Title, OwnerID: req.PlayerID}
	fieldNames := []string{"OwnerID"}
	if req.Author == "" && req.Title == "" {
		ctx.EmitError("author or title must be provided to remove a book", []int{req.PlayerID})
		return
	}

	if req.Author != "" {
		fieldNames = append(fieldNames, "Author")
	}
	if req.Title != "" {
		fieldNames = append(fieldNames, "Title")
	}

	bookEntities := bookTable.Filter(w, bookFilter, fieldNames)
	for _, e := range bookEntities {
		bookTable.RemoveEntity(w, e)
	}
})
