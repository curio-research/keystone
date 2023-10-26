package test

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/curio-research/keystone/core"
	server "github.com/curio-research/keystone/server"
	pb_test "github.com/curio-research/keystone/test/proto/pb.test"
	"github.com/curio-research/keystone/test/testutils"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var BookTable = core.NewTableAccessor[Book]()

var p *testutils.PortManager

func init() {
	p = testutils.NewPortManager()
}

func TestAddBook(t *testing.T) {
	e, ws, s, _, _ := startTestServer(t, core.Dev)
	defer tearDown(ws, s)

	w := e.World

	playerID := 7
	player2ID := 8

	specificBookEntity := 969
	err := sendWSMsg(ws, playerID, &pb_test.TestBookInfo{
		Op:     pb_test.Operation_AddSpecific,
		Title:  testBookTitle1,
		Author: testBookAuthor1,
		Entity: int64(specificBookEntity),
	})
	require.Nil(t, err)

	err = sendWSMsg(ws, playerID, &pb_test.TestBookInfo{
		Op:     pb_test.Operation_Add,
		Title:  testBookTitle2,
		Author: testBookAuthor2,
	})
	require.Nil(t, err)

	err = sendWSMsg(ws, player2ID, &pb_test.TestBookInfo{
		Op:     pb_test.Operation_Add,
		Title:  testBookTitle3,
		Author: testBookAuthor3,
	})
	require.Nil(t, err)

	b1 := BookTable.Get(w, specificBookEntity)
	assert.Equal(t, testBookTitle1, b1.Title)
	assert.Equal(t, testBookAuthor1, b1.Author)
	assert.Equal(t, playerID, b1.OwnerID)
	assert.Equal(t, specificBookEntity, b1.Id)

	b2 := BookTable.Filter(w, Book{
		Title:  testBookTitle2,
		Author: testBookAuthor2,
	}, []string{"Title", "Author"})
	require.Len(t, b2, 1)
	assert.Equal(t, playerID, BookTable.Get(w, b2[0]).OwnerID)

	b3 := BookTable.Filter(w, Book{
		Title:  testBookTitle3,
		Author: testBookAuthor3,
	}, []string{"Title", "Author"})
	require.Len(t, b3, 1)
	assert.Equal(t, player2ID, BookTable.Get(w, b3[0]).OwnerID)
}

func tearDown(ws *websocket.Conn, server *http.Server) {
	ws.Close()
	server.Close()
	time.Sleep(time.Millisecond * 50)
}

func TestUpdate(t *testing.T) {
	e, ws, s, mockErrorHandler, _ := startTestServer(t, core.Dev)
	defer tearDown(ws, s)

	w := e.World

	playerID := 7

	w.AddEntity()

	b1Entity := 1
	b2Entity := 2

	addBook(w, testBookTitle1, testBookAuthor1, playerID, b1Entity)
	addBook(w, testBookTitle2, testBookAuthor2, playerID, b2Entity)

	// error; first update missing entity => none of the books should be updated
	err := sendWSMsg(ws, playerID, &pb_test.TestBookInfo{
		Op:     pb_test.Operation_Update,
		Title:  testBookTitle1,
		Author: testBookAuthor3,
	}, &pb_test.TestBookInfo{
		Op:     pb_test.Operation_Update,
		Title:  testBookTitle3,
		Author: testBookAuthor2,
		Entity: int64(b2Entity),
	})
	require.Nil(t, err)

	require.Equal(t, mockErrorHandler.ErrorCount(), 1)
	assert.Equal(t, "no book to update with entity 0", mockErrorHandler.LastError())

	b1 := BookTable.Get(w, b1Entity)
	assert.Equal(t, testBookTitle1, b1.Title)
	assert.Equal(t, testBookAuthor1, b1.Author)

	b2 := BookTable.Get(w, b2Entity)
	assert.Equal(t, testBookTitle2, b2.Title)
	assert.Equal(t, testBookAuthor2, b2.Author)

	err = sendWSMsg(ws, playerID, &pb_test.TestBookInfo{
		Op:     pb_test.Operation_Update,
		Title:  testBookTitle1,
		Author: testBookAuthor3,
		Entity: int64(b1Entity),
	}, &pb_test.TestBookInfo{
		Op:     pb_test.Operation_Update,
		Title:  testBookTitle3,
		Author: testBookAuthor2,
		Entity: int64(b2Entity),
	})
	require.Nil(t, err)

	b1 = BookTable.Get(w, b1Entity)
	assert.Equal(t, testBookTitle1, b1.Title)
	assert.Equal(t, testBookAuthor3, b1.Author)

	b2 = BookTable.Get(w, b2Entity)
	assert.Equal(t, testBookTitle3, b2.Title)
	assert.Equal(t, testBookAuthor2, b2.Author)
}

// also testing sending a call to another system
func TestDeleteAndFilter(t *testing.T) {
	playerID := 7
	player2ID := 8

	testTable := []struct {
		name         string
		authorFilter string
		titleFilter  string
		playerID     int
		errorMsg     string

		remainingEntities []int
	}{
		{
			name:              "one author",
			authorFilter:      testBookAuthor1,
			playerID:          playerID,
			remainingEntities: []int{2, 4, 5},
		},
		{
			name:              "author and title",
			titleFilter:       testBookTitle1,
			authorFilter:      testBookAuthor2,
			playerID:          player2ID,
			remainingEntities: []int{1, 3, 4, 5},
		},
		{
			name:              "no matching queries - playerID",
			authorFilter:      testBookAuthor1,
			playerID:          player2ID,
			remainingEntities: []int{1, 2, 3, 4, 5},
		},
		{
			name:              "no matching entities - filters",
			titleFilter:       testBookTitle1,
			authorFilter:      testBookAuthor3,
			playerID:          playerID,
			remainingEntities: []int{1, 2, 3, 4, 5},
		},
		{
			name:              "no author or title - error",
			playerID:          playerID,
			remainingEntities: []int{1, 2, 3, 4, 5},
			errorMsg:          "author or title must be provided to remove a book",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			e, ws, s, errorHandler, _ := startTestServer(t, core.Dev)
			defer tearDown(ws, s)

			w := e.World

			addBookSpecific(w, testBookTitle1, testBookAuthor1, playerID, 1)  // entity = 1
			addBookSpecific(w, testBookTitle1, testBookAuthor2, player2ID, 2) // entity = 2
			addBookSpecific(w, testBookTitle2, testBookAuthor1, playerID, 3)  // entity = 3
			addBookSpecific(w, testBookTitle2, testBookAuthor2, player2ID, 4) // entity = 4
			addBookSpecific(w, testBookTitle3, testBookAuthor3, playerID, 5)  // entity = 5

			err := sendWSMsg(ws, testCase.playerID, &pb_test.TestBookInfo{
				Op:     pb_test.Operation_Remove,
				Title:  testCase.titleFilter,
				Author: testCase.authorFilter,
			})
			require.Nil(t, err)

			m := make(map[int]interface{})
			for _, i := range testCase.remainingEntities {
				m[i] = nil
			}

			for i := 1; i <= 5; i++ {
				book := BookTable.Get(w, i)
				if _, ok := m[i]; ok {
					assert.NotEqual(t, "", book.Title)
				} else {
					assert.Equal(t, "", book.Title)
				}
			}

			if testCase.errorMsg == "" {
				assert.Equal(t, errorHandler.ErrorCount(), 0)
			} else {
				require.Equal(t, errorHandler.ErrorCount(), 1)
				assert.Equal(t, testCase.errorMsg, errorHandler.LastError())
			}
		})
	}
}

func addBook(w core.IWorld, title, author string, ownerID int, entity int) int {
	return BookTable.AddSpecific(w, entity, Book{
		Title:   title,
		Author:  author,
		OwnerID: ownerID,
	})
}

func addBookSpecific(w core.IWorld, title, author string, ownerID, entity int) int {
	return BookTable.AddSpecific(w, entity, Book{
		Title:   title,
		Author:  author,
		OwnerID: ownerID,
	})
}

func sendWSMsg(ws *websocket.Conn, playerID int, bookInfos ...*pb_test.TestBookInfo) error {
	err := testutils.SendMessage(ws, testutils.C2S_Test_MessageType, &pb_test.C2S_Test{
		BookInfos:       bookInfos,
		IdentityPayload: testutils.CreateMockIdentityPayload(playerID),
	})
	if err != nil {
		return err
	}

	time.Sleep(time.Millisecond * 100)
	// time.Sleep(time.Second)

	return nil
}

func startTestServer(t *testing.T, mode core.GameMode) (*server.EngineCtx, *websocket.Conn, *http.Server, *testutils.MockErrorHandler, *sql.DB) {
	port, wsPort := p.GetPort(), p.GetPort()

	s, e, db, err := testutils.Server(t, mode, wsPort, 1, testSchemaToAccessors)
	require.Nil(t, err)

	mockErrorHandler := testutils.NewMockErrorHandler()
	e.SystemErrorHandler = mockErrorHandler

	addr := ":" + strconv.Itoa(port)
	httpServer := &http.Server{
		Addr:    addr,
		Handler: s,
	}

	go func() {
		fmt.Println("starting server at ", addr)
		err := httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			t.Errorf("http server closed with unexpected error %v", err)
			return
		}
	}()

	e.GameTick.Schedule.AddTickSystem(1, TestBookSystem)
	e.GameTick.Schedule.AddTickSystem(1, TestRemoveBookSystem)

	e.World.AddTable(BookTable)

	ws, err := testutils.SetupWS(t, wsPort)
	require.Nil(t, err)

	return e, ws, httpServer, mockErrorHandler, db
}

var TestBookSystem = server.CreateSystemFromRequestHandler(func(ctx *server.TransactionCtx[*pb_test.C2S_Test]) {
	req := ctx.Req
	w := ctx.W

	fmt.Println("tick", ctx.GameCtx.GameTick.TickNumber)
	playerID := int(req.GetIdentityPayload().GetPlayerId())

	for _, bookInfo := range req.BookInfos {
		switch bookInfo.Op {
		case pb_test.Operation_Add:
			BookTable.Add(w, Book{
				Title:   bookInfo.Title,
				Author:  bookInfo.Author,
				OwnerID: playerID,
			})
		case pb_test.Operation_AddSpecific:
			BookTable.AddSpecific(w, int(bookInfo.Entity), Book{
				Title:   bookInfo.Title,
				Author:  bookInfo.Author,
				OwnerID: playerID,
			})
		case pb_test.Operation_Remove:
			server.QueueTxFromInternal(w, ctx.GameCtx.GameTick.TickNumber+1, testRemoveRequest{
				Title:    bookInfo.Title,
				Author:   bookInfo.Author,
				PlayerID: playerID,
			}, "")
		case pb_test.Operation_Update:
			book := BookTable.Get(w, int(bookInfo.Entity))
			if book.Title == "" {
				ctx.EmitError(fmt.Sprintf("no book to update with entity %v", bookInfo.Entity), []int{playerID})
				return
			}

			book.Title = bookInfo.Title
			book.Author = bookInfo.Author
			BookTable.Set(w, int(bookInfo.Entity), book)
		}
	}
})

type testRemoveRequest struct {
	Author   string
	Title    string
	PlayerID int
}

var TestRemoveBookSystem = server.CreateSystemFromRequestHandler(func(ctx *server.TransactionCtx[testRemoveRequest]) {
	fmt.Println("tick2", ctx.GameCtx.GameTick.TickNumber)
	req := ctx.Req
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

	bookEntities := BookTable.Filter(w, bookFilter, fieldNames)
	for _, e := range bookEntities {
		BookTable.RemoveEntity(w, e)
	}
})
