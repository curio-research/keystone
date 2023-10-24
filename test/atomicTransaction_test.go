package test

import (
	"testing"

	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"github.com/curio-research/keystone/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorldUpdatedOnSuccess_RequestSystem(t *testing.T) {
	t.Error("Bad")
	ctx, err := initializeTestWorld(TestPersonRequestSystem)
	require.Nil(t, err)

	person1Req := testPersonRequest{OP: state.UpdateOP, Entity: 27, Val: Person{
		Name: testName1,
	}, SendError: false}
	person2Req := testPersonRequest{OP: state.UpdateOP, Entity: 28, Val: Person{
		Name: testName2,
	}, SendError: false}

	req := testPersonRequests{
		People:   []testPersonRequest{person1Req, person2Req},
		PlayerID: 6,
	}

	server.QueueTxAtTime(ctx.World, 1, req, "", true)
	utils.TickWorldForward(ctx, 1)

	person1 := personTable.Get(ctx.World, 27)
	assert.Equal(t, testName1, person1.Name)

	person2 := personTable.Get(ctx.World, 28)
	assert.Equal(t, testName2, person2.Name)
}

func TestWorldNotUpdatedOnFailure_RequestSystem(t *testing.T) {
	ctx, err := initializeTestWorld(TestPersonRequestSystem)
	require.Nil(t, err)

	person1Req := testPersonRequest{OP: state.UpdateOP, Entity: 27, Val: Person{
		Name: testName1,
	}, SendError: false}
	person2Req := testPersonRequest{OP: state.UpdateOP, Entity: 28, Val: Person{
		Name: testName2,
	}, SendError: true}

	req := testPersonRequests{
		People:   []testPersonRequest{person1Req, person2Req},
		PlayerID: 6,
	}

	server.QueueTxAtTime(ctx.World, 1, req, "", true)
	utils.TickWorldForward(ctx, 1)

	person1 := personTable.Get(ctx.World, 27)
	assert.NotEqual(t, testName1, person1.Name)

	person2 := personTable.Get(ctx.World, 28)
	assert.NotEqual(t, testName2, person2.Name)
}

func TestWorldUpdatedOnSuccess_GeneralSystem(t *testing.T) {
	ctx, err := initializeTestWorld(TestPersonSystem)
	require.Nil(t, err)

	utils.TickWorldForward(ctx, 1)

	person := personTable.Get(ctx.World, 0)
	assert.Equal(t, testName1, person.Name)

	book := bookTable.Get(ctx.World, 1)
	assert.Equal(t, testBookTitle1, book.Title)
}

func TestWorldNotUpdatedOnFailure_GeneralSystem(t *testing.T) {
	ctx, err := initializeTestWorld(TestPersonSystemWithError)
	require.Nil(t, err)

	utils.TickWorldForward(ctx, 1)

	person := personTable.Get(ctx.World, 0)
	assert.NotEqual(t, testName1, person.Name)

	book := bookTable.Get(ctx.World, 1)
	assert.NotEqual(t, testBookTitle1, book.Title)
}

var TestPersonRequestSystem = server.CreateSystemFromRequestHandler(func(ctx *server.TransactionCtx[testPersonRequests]) {
	w := ctx.W
	req := ctx.Req
	playerID := int(req.GetIdentityPayload().playerID)

	for _, person := range req.People {
		switch person.OP {
		case state.UpdateOP:
			personTable.Set(w, person.Entity, person.Val)
		case state.RemovalOP:
			personTable.RemoveEntity(w, person.Entity)
		case state.AddEntityOP:
			personTable.AddSpecific(w, person.Entity, person.Val)
		}

		if person.SendError {
			ctx.EmitError(testErrMsg, []int{playerID})
			return
		} else {
			ctx.EmitEvent(69, nil, []int{playerID}, true)
		}
	}
})

var TestPersonSystem = server.CreateGeneralSystem(func(ctx *server.TransactionCtx[any]) {
	if ctx.GameCtx.GameTick.TickNumber == 1 {
		w := ctx.W

		personTable.Add(w, Person{
			Name:    testName1,
			Age:     testAge1,
			Address: testAddress1,
		})
		bookTable.Add(w, Book{
			Title:  testBookTitle1,
			Author: testBookAuthor1,
		})

		ctx.EmitEvent(69, nil, nil, true)
	}
})

var TestPersonSystemWithError = server.CreateGeneralSystem(func(ctx *server.TransactionCtx[any]) {
	if ctx.GameCtx.GameTick.TickNumber == 1 {
		w := ctx.W

		personTable.Add(w, Person{
			Name:    testName1,
			Age:     testAge1,
			Address: testAddress1,
		})
		bookTable.Add(w, Book{
			Title:  testBookTitle1,
			Author: testBookAuthor1,
		})

		ctx.EmitError(testErrMsg, nil)
	}
})
