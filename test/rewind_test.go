package test

import (
	"bytes"
	"encoding/json"
	"github.com/curio-research/keystone/server/startup"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/curio-research/keystone/server"
	pb_test "github.com/curio-research/keystone/test/proto/pb.test"
	"github.com/curio-research/keystone/test/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMySQLRewind(t *testing.T) {
	testutils.SkipTestIfShort(t)

	ctx, _, s, _, db := startTestServer(t, server.DevMySQL)
	defer db.Close()

	coreRewindTest(t, ctx, s)
}

func TestSQLiteRewind(t *testing.T) {
	ctx, _, s, _, db := startTestServer(t, server.DevSQLite)
	defer db.Close()

	coreRewindTest(t, ctx, s)
	testutils.ResetSQLiteTestDB()
}

func coreRewindTest(t *testing.T, ctx *server.EngineCtx, s *http.Server) {
	testutils.SkipTestIfShort(t)

	ctx.SetGameLiveliness(true)

	player1Entity := testEntity1
	book1Entity, book2Entity, book3Entity := testEntity2, testEntity3, testEntity4

	initWorld := func(ctx *server.EngineCtx) {
		bookTable.AddSpecific(ctx.World, testEntity4, Book{Title: testBookTitle3, Author: testBookAuthor3})
	}
	startup.RegisterRewindEndpoint(ctx, initWorld)

	// second 1
	server.QueueTxFromExternal(ctx, server.NewKeystoneTx(&pb_test.C2S_Test{ // at tick 2
		BookInfos: []*pb_test.TestBookInfo{
			{
				Op:     pb_test.Operation_AddSpecific,
				Entity: int64(book1Entity),
				Author: testBookAuthor1,
				Title:  testBookTitle1,
			},
			{
				Op:     pb_test.Operation_AddSpecific,
				Entity: int64(book2Entity),
				Author: testBookAuthor2,
				Title:  testBookTitle2,
			},
		},
		IdentityPayload: testutils.CreateMockIdentityPayload(player1Entity),
	}, nil), "")
	server.TickWorldForward(ctx, 50) // 50 * 20 ms => 1s

	// second 2
	server.QueueTxFromExternal(ctx, server.NewKeystoneTx(&pb_test.C2S_Test{
		BookInfos: []*pb_test.TestBookInfo{
			{
				Op:     pb_test.Operation_Update,
				Entity: int64(book1Entity),
				Title:  testBookTitle2,
				Author: testBookAuthor1,
			},
			{
				Op:     pb_test.Operation_Update,
				Entity: int64(book2Entity),
				Title:  testBookTitle1,
				Author: testBookAuthor2,
			},
		},
		IdentityPayload: testutils.CreateMockIdentityPayload(player1Entity),
	}, nil), "")
	server.TickWorldForward(ctx, 15)

	server.QueueTxFromExternal(ctx, server.NewKeystoneTx(&pb_test.C2S_Test{
		BookInfos: []*pb_test.TestBookInfo{
			{
				Op:     pb_test.Operation_Remove,
				Author: testBookAuthor1,
				Title:  testBookTitle2,
			},
		},
		IdentityPayload: testutils.CreateMockIdentityPayload(player1Entity),
	}, nil), "")
	server.TickWorldForward(ctx, 35)

	time.Sleep(time.Second * 2)

	sendPostRequest(t, s, "rewindState", server.NewKeystoneTx(server.RewindStateRequest{
		ElapsedSeconds: 1,
		GameId:         testGameID1,
	}, nil))

	book1 := bookTable.Get(ctx.World, book1Entity)
	assert.Equal(t, testBookTitle1, book1.Title)
	assert.Equal(t, testBookAuthor1, book1.Author)

	book2 := bookTable.Get(ctx.World, book2Entity)
	assert.Equal(t, testBookTitle2, book2.Title)
	assert.Equal(t, testBookAuthor2, book2.Author)

	book3 := bookTable.Get(ctx.World, book3Entity)
	assert.Equal(t, testBookTitle3, book3.Title)
	assert.Equal(t, testBookAuthor3, book3.Author)

	sendPostRequest(t, s, "rewindState", server.NewKeystoneTx(server.RewindStateRequest{
		ElapsedSeconds: 10,
		GameId:         ctx.GameId,
	}, nil))

	book1 = bookTable.Get(ctx.World, book1Entity)
	assert.Equal(t, 0, book1.Id)

	book2 = bookTable.Get(ctx.World, book2Entity)
	assert.Equal(t, testBookTitle1, book2.Title)
	assert.Equal(t, testBookAuthor2, book2.Author)

	book3 = bookTable.Get(ctx.World, book3Entity)
	assert.Equal(t, testBookTitle3, book3.Title)
	assert.Equal(t, testBookAuthor3, book3.Author)
}

func sendPostRequest[T any](t *testing.T, s *http.Server, route string, data server.KeystoneTx[T]) *http.Response {
	httpServer := httptest.NewServer(s.Handler)

	b, err := json.Marshal(data)
	require.Nil(t, err)

	req, err := http.NewRequest("POST", "http://"+s.Addr+"/"+route, bytes.NewBuffer(b))
	require.Nil(t, err)

	resp, err := httpServer.Client().Do(req)
	require.Nil(t, err)

	return resp
}
