package test

import (
	"net/http"
	"testing"
	"time"

	server "github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	pb_test "github.com/curio-research/keystone/test/proto/pb.test"
	"github.com/curio-research/keystone/test/testutils"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddBook(t *testing.T) {
	e, ws, s, _, _ := startTestServer(t, server.Dev)
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

	server.TickWorldForward(e, 2)

	b1 := bookTable.Get(w, specificBookEntity)
	assert.Equal(t, testBookTitle1, b1.Title)
	assert.Equal(t, testBookAuthor1, b1.Author)
	assert.Equal(t, playerID, b1.OwnerID)
	assert.Equal(t, specificBookEntity, b1.Id)

	b2 := bookTable.Filter(w, Book{
		Title:  testBookTitle2,
		Author: testBookAuthor2,
	}, []string{"Title", "Author"})
	require.Len(t, b2, 1)
	assert.Equal(t, playerID, bookTable.Get(w, b2[0]).OwnerID)

	b3 := bookTable.Filter(w, Book{
		Title:  testBookTitle3,
		Author: testBookAuthor3,
	}, []string{"Title", "Author"})
	require.Len(t, b3, 1)
	assert.Equal(t, player2ID, bookTable.Get(w, b3[0]).OwnerID)
}

func tearDown(ws *websocket.Conn, server *http.Server) {
	ws.Close()
	server.Close()
	time.Sleep(time.Millisecond * 50)
}

func TestUpdate(t *testing.T) {
	e, ws, s, mockErrorHandler, _ := startTestServer(t, server.Dev)
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

	server.TickWorldForward(e, 2)

	require.Equal(t, mockErrorHandler.ErrorCount(), 1)
	assert.Equal(t, "no book to update with entity 0", mockErrorHandler.LastError())

	b1 := bookTable.Get(w, b1Entity)
	assert.Equal(t, testBookTitle1, b1.Title)
	assert.Equal(t, testBookAuthor1, b1.Author)

	b2 := bookTable.Get(w, b2Entity)
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

	server.TickWorldForward(e, 2)

	b1 = bookTable.Get(w, b1Entity)
	assert.Equal(t, testBookTitle1, b1.Title)
	assert.Equal(t, testBookAuthor3, b1.Author)

	b2 = bookTable.Get(w, b2Entity)
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
			e, ws, s, errorHandler, _ := startTestServer(t, server.Dev)
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

			server.TickWorldForward(e, 3)

			m := make(map[int]interface{})
			for _, i := range testCase.remainingEntities {
				m[i] = nil
			}

			for i := 1; i <= 5; i++ {
				book := bookTable.Get(w, i)
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

func addBook(w state.IWorld, title, author string, ownerID int, entity int) int {
	return bookTable.AddSpecific(w, entity, Book{
		Title:   title,
		Author:  author,
		OwnerID: ownerID,
	})
}

func addBookSpecific(w state.IWorld, title, author string, ownerID, entity int) int {
	return bookTable.AddSpecific(w, entity, Book{
		Title:   title,
		Author:  author,
		OwnerID: ownerID,
	})
}

func sendWSMsg(ws *websocket.Conn, playerID int, bookInfos ...*pb_test.TestBookInfo) error {
	err := testutils.SendMessage(ws, testutils.C2S_Test_MessageType, server.NewKeystoneTx(&pb_test.C2S_Test{
		BookInfos:       bookInfos,
		IdentityPayload: testutils.CreateMockIdentityPayload(playerID),
	}, nil))
	if err != nil {
		return err
	}

	time.Sleep(time.Millisecond * 100)

	return nil
}
