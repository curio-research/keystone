package server

import (
	"math/rand"
	"time"

	"github.com/curio-research/go-backend/engine"
	"github.com/curio-research/go-backend/pathfinder"
	"github.com/curio-research/go-backend/server/components"
)

// get the tile at position using reverse lookup
// tiles shouldn't change so i think these should be cached
func GetSmallTileEntityAtPos(w *engine.World, x int, y int) int {
	res := w.Query([]engine.QueryCondition{
		{
			Component: components.PositionComponent.Name(),
			QueryType: engine.HasExact,
			Value:     engine.Pos{X: x, Y: y},
		},
		{
			Component: components.TagComponent.Name(),
			QueryType: engine.HasExact,
			Value:     components.SmallTileTag,
		},
	})

	if len(res) == 0 {
		return 0
	}

	return res[0]
}

func GetAllEntitiesOfTag(w *engine.World, tag string) []int {
	res := w.Query([]engine.QueryCondition{
		{
			Component: components.TagComponent.Name(),
			QueryType: engine.HasExact,
			Value:     tag,
		},
	})

	return res
}

// optimize this a lot
func ConvertLiveMapTo2DArray(w *engine.World) [][]string {
	world2dArr := [][]string{}

	// spawn empty map
	for i := 0; i < int(WorldHeight); i++ {
		row := []string{}
		for j := 0; j < int(WorldWidth); j++ {
			row = append(row, pathfinder.EmptySlotSymbol)
		}
		world2dArr = append(world2dArr, row)
	}

	// get all troop entities
	troopIds := ProbabilisticRemoval(GetAllEntitiesOfTag(w, components.InfantryTroop), 0.5)
	blockerIds := GetAllEntitiesOfTag(w, components.BlockerTag)

	allBlockerTiles := []int{}
	allBlockerTiles = append(allBlockerTiles, blockerIds...)
	allBlockerTiles = append(allBlockerTiles, troopIds...)

	// get all troops and building blockers and make them obstacles
	for _, obstacleEntity := range allBlockerTiles {
		pos, _ := components.PositionComponent.Get(w, obstacleEntity)

		world2dArr[pos.Y][pos.X] = pathfinder.ObstacleSymbol
	}

	return world2dArr

}

func ProbabilisticRemoval(arr []int, probability float64) []int {
	rand.Seed(time.Now().UnixNano())

	removedArr := make([]int, 0)
	for _, num := range arr {
		if rand.Float64() > probability {
			removedArr = append(removedArr, num)
		}
	}

	return removedArr
}

// three segment queries don't work
func GetTickJobs(w *engine.World, tickType string, tickNumber int) []int {
	res := w.Query([]engine.QueryCondition{
		{
			Component: components.TickNumberComponent.Name(),
			QueryType: engine.HasExact,
			Value:     tickNumber,
		},
		// {
		// 	Component: TagComp,
		// 	QueryType: engine.HasExact,
		// 	Value:     TickJobTag,
		// },
		{
			Component: components.TickJobTypeComponent.Name(),
			QueryType: engine.HasExact,
			Value:     tickType,
		},
	})

	return res
}

func GetInfantriesOfOwner(w *engine.World, ownerId int) []int {
	res := w.Query([]engine.QueryCondition{
		{
			Component: components.OwnerIdComponent.Name(),
			QueryType: engine.HasExact,
			Value:     ownerId,
		},
		{
			Component: components.TagComponent.Name(),
			QueryType: engine.HasExact,
			Value:     components.InfantryTroop,
		},
	})

	return res
}

func GetEntitiesAtPosition(w *engine.World, pos engine.Pos) []int {
	res := w.Query([]engine.QueryCondition{
		{
			Component: components.PositionComponent.Name(),
			QueryType: engine.HasExact,
			Value:     pos,
		},
	})

	return res
}

func GetEntitiesOfTickId(w *engine.World, tickId string) []int {
	res := w.Query([]engine.QueryCondition{
		{
			Component: components.TickIdComponent.Name(),
			QueryType: engine.HasExact,
			Value:     tickId,
		},
	})

	return res
}

func RemoveAllEntities(w *engine.World, entities []int) {
	for _, entity := range entities {
		w.RemoveEntity(entity)
	}
}
