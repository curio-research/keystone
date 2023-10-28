package test

import (
	"math/rand"
	"testing"

	"github.com/curio-research/keystone/state"
)

func TestStateHash(t *testing.T) {
	world1 := createMockWorld(456456)
	world2 := createMockWorld(123123)

	world1stateRoot := state.CalculateWorldStateRootHash(world1)
	world2stateRoot := state.CalculateWorldStateRootHash(world2)

	if world1stateRoot != world2stateRoot {
		t.Errorf("World state roots should be equal")
	}

}

type tableASchema struct {
	Name string
	Id   int
}

type tableBSchema struct {
	Gold   int
	IsLive bool
	Id     int
}

var tableAAccessor = state.NewTableAccessor[tableASchema]()
var tableBAccessor = state.NewTableAccessor[tableBSchema]()

type pairToInsert struct {
	TableName string
	Entity    int
	Data      interface{}
}

func createMockWorld(seed int64) *state.GameWorld {
	world := state.NewWorld()

	world.AddTable(tableAAccessor)
	world.AddTable(tableBAccessor)

	// get elements that need to be initialized

	pairsToInsert := []pairToInsert{
		{
			TableName: tableAAccessor.Name(),
			Entity:    1,
			Data: tableASchema{
				Name: "test",
				Id:   1,
			},
		},
		{
			TableName: tableAAccessor.Name(),
			Entity:    2,
			Data: tableASchema{
				Name: "test2",
				Id:   2,
			},
		},
		{
			TableName: tableAAccessor.Name(),
			Entity:    3,
			Data: tableASchema{
				Name: "",
				Id:   3,
			},
		},
		{
			TableName: tableBAccessor.Name(),
			Entity:    100,
			Data: tableBSchema{
				Gold:   100,
				IsLive: true,
			},
		},
		{
			TableName: tableBAccessor.Name(),
			Entity:    101,
			Data: tableBSchema{
				Gold:   -100,
				IsLive: false,
			},
		},
		{
			TableName: tableBAccessor.Name(),
			Entity:    102,
			Data: tableBSchema{
				Gold:   2012,
				IsLive: true,
			},
		},
	}

	// randomize the order of the pairs
	arr := shufflePairs(pairsToInsert, seed)

	for _, pair := range arr {
		world.AddSpecific(pair.Entity, pair.Data, pair.TableName)
	}

	return world

}

func copyPairsToInsertArray(pairs []pairToInsert) []pairToInsert {
	newPairs := make([]pairToInsert, len(pairs))
	copy(newPairs, pairs)

	return newPairs
}

func shufflePairs(pairs []pairToInsert, seed int64) []pairToInsert {
	newArr := copyPairsToInsertArray(pairs)
	// Initialize the random number generator with a seed
	r := rand.New(rand.NewSource(seed))

	// Shuffle the slice using the Fisher-Yates shuffle algorithm
	for i := len(newArr) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		newArr[i], newArr[j] = newArr[j], newArr[i]
	}

	return newArr
}
