package components

import (
	"github.com/curio-research/go-backend/engine"
)

// component list
var (
	TagComponent        = engine.NewComponent[string]("Tag", true)
	PositionComponent   = engine.NewComponent[engine.Pos]("Position", true)
	OwnerIdComponent    = engine.NewComponent[int]("OwnerId", true)
	LevelComponent      = engine.NewComponent[int]("Level", true)
	HealthComponent     = engine.NewComponent[int]("HealthComp", true)
	NameComponent       = engine.NewComponent[string]("NameComp", true)
	LastActiveComponent = engine.NewComponent[int]("LastActive", true)

	TickJobTypeComponent    = engine.NewComponent[string]("TickJobType", true)
	TickIdComponent         = engine.NewComponent[string]("TickId", true)
	TickDataStringComponent = engine.NewComponent[string]("TickData", false)
	TickNumberComponent     = engine.NewComponent[int]("TickNumber", true)
)

// list of all tags (aka archetypes)
const (
	SmallTileTag = "SmallTileTag"
	SoldierTag   = "Soldier"
	CapitalTag   = "Capital"
	PlayerTag    = "PlayerTag"
	BlockerTag   = "BlockerTag"
	BuildingTag  = "BuildingTag"
	TickJobTag   = "TickJobTag"
	Tree1        = "Tree1"
	Tree2        = "Tree2"

	// inventory type names
	InventoryTypeGold = "InventoryTypeGold"
	NationalToken     = "NationalToken"

	// troop type names
	InfantryTroop = "InfantryTroop"
)
