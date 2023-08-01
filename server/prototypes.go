package server

import (
	"time"

	"github.com/curio-research/go-backend/engine"
	"github.com/curio-research/go-backend/server/components"
)

func AddSmallTile(w *engine.World, x int, y int) int {
	entity := w.AddEntity()

	components.TagComponent.Set(w, entity, components.SmallTileTag)
	components.PositionComponent.Set(w, entity, engine.Pos{X: x, Y: y})

	return entity
}

func AddInfantry(w *engine.World, x int, y int, ownerId int) int {
	entity := w.AddEntity()

	components.TagComponent.Set(w, entity, components.InfantryTroop)
	components.PositionComponent.Set(w, entity, engine.Pos{X: x, Y: y})
	components.OwnerIdComponent.Set(w, entity, ownerId)
	components.HealthComponent.Set(w, entity, InfantryStartingHealth)

	return entity
}

func AddPlayer(w *engine.World, name string) int {
	entity := w.AddEntity()

	components.TagComponent.Set(w, entity, components.PlayerTag)
	components.NameComponent.Set(w, entity, name)
	components.LastActiveComponent.Set(w, entity, int(time.Now().Unix()))

	return entity
}

func AddBlocker(w *engine.World, x int, y int) int {
	entity := w.AddEntity()

	components.TagComponent.Set(w, entity, components.BlockerTag)
	components.PositionComponent.Set(w, entity, engine.Pos{X: x, Y: y})

	return entity
}

// input tree type
func AddTree(w *engine.World, x int, y int, treeType string) int {
	entity := w.AddEntity()

	components.TagComponent.Set(w, entity, treeType)
	components.PositionComponent.Set(w, entity, engine.Pos{X: x, Y: y})

	return entity
}

func AddBuilding(w *engine.World, x int, y int, ownerId int, buildingType string) int {
	entity := w.AddEntity()

	components.TagComponent.Set(w, entity, components.BuildingTag)
	components.NameComponent.Set(w, entity, buildingType)
	components.PositionComponent.Set(w, entity, engine.Pos{X: x, Y: y})
	components.OwnerIdComponent.Set(w, entity, ownerId)

	return entity
}

func AddTickJob(w *engine.World, tickNumber int, jobType string, jobData string, tickId string) int {
	entity := w.AddEntity()

	components.TagComponent.Set(w, entity, components.TickJobTag)
	components.TickNumberComponent.Set(w, entity, tickNumber)
	components.TickJobTypeComponent.Set(w, entity, jobType)
	components.TickDataStringComponent.Set(w, entity, jobData)
	components.TickIdComponent.Set(w, entity, tickId)

	return entity
}
