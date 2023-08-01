package server

import (
	"math/rand"
	"time"

	"github.com/curio-research/go-backend/engine"
	"github.com/curio-research/go-backend/pathfinder"
	"github.com/curio-research/go-backend/server/components"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	NpcRespawnInterval int = 5_000 // milliseconds
	MaxPlayerCount     int = 20
	MaxNPCInWorld      int = 100

	InfantryStartingHealth int = 100
	MoveSpeed              int = 200
	MaxInfantryPerPlayer   int = 5
	AttackDamage           int = 30
	AttackDistance         int = 7

	MongoDatabaseName string = "test-db"

	WorldWidth  int = 100
	WorldHeight int = 100
)

// TODO: refactooor to somewhere else
var PlayerCollection *mongo.Collection
var WorldsCollection *mongo.Collection

// initialize world map and initial objects
func InitializeMap(w *engine.World, seed int) {

	RegisterAllComponents(w)

	farmBuildingCount := 5
	humanBarracksCount := 5
	humanTowerCount := 5
	workshopBuildingCount := 5

	buildingWidth := 9 // should correspond with pixels in Unity

	// add 100x100 small tiles
	for i := 0; i < int(WorldWidth); i++ {
		row := []string{}

		for j := 0; j < int(WorldHeight); j++ {
			row = append(row, pathfinder.EmptySlotSymbol)

			AddSmallTile(w, i, j)
		}
	}

	// spawn NPCs which are troops anyone can attack
	SpawnNPCsDeterministic(w, MaxNPCInWorld, HashNumbers(seed, w.EntitiesNonce))

	farm := "farm"
	humanbarracks := "humanbarracks"
	humantower := "humantower"
	workshop := "workshop"

	tree1Density := 0.01 // percentage
	tree2Density := 0.01

	// spawn trees on random zero points on map
	SpawnTreeWithPercentage(w, components.Tree1, tree1Density, HashNumbers(seed, w.EntitiesNonce))
	SpawnTreeWithPercentage(w, components.Tree2, tree2Density, HashNumbers(seed, w.EntitiesNonce))

	for i := 0; i < farmBuildingCount; i++ {
		SpawnBuilding(w, 0, farm, buildingWidth, HashNumbers(seed, w.EntitiesNonce))
	}

	for i := 0; i < humanBarracksCount; i++ {
		SpawnBuilding(w, 0, humanbarracks, buildingWidth, HashNumbers(seed, w.EntitiesNonce))
	}

	for i := 0; i < humanTowerCount; i++ {
		SpawnBuilding(w, 0, humantower, buildingWidth, HashNumbers(seed, w.EntitiesNonce))
	}

	for i := 0; i < workshopBuildingCount; i++ {
		SpawnBuilding(w, 0, workshop, buildingWidth, HashNumbers(seed, w.EntitiesNonce))
	}
}

func SpawnTreeWithPercentage(w *engine.World, treeType string, percentage float64, seed int) {

	totalTreeNumber := float64(WorldWidth) * float64(WorldHeight) * percentage
	plantedTree := 0

	source := rand.NewSource(int64(seed))
	random := rand.New(source)

	for plantedTree < int(totalTreeNumber) {
		x := random.Intn(WorldWidth)
		y := random.Intn(WorldHeight)

		if IsPositionInMap(x, y, WorldWidth, WorldHeight) {
			if len(GetEntitiesAtPosition(w, engine.Pos{X: x, Y: y})) == 1 {
				AddTree(w, x, y, treeType)
				plantedTree++
			}
		}
	}
}

// this should be deterministic
func GetEmptyPosition(w *engine.World) engine.Pos {
	for {
		x := rand.Intn(int(WorldWidth))
		y := rand.Intn(int(WorldHeight))

		if len(GetEntitiesAtPosition(w, engine.Pos{X: x, Y: y})) == 1 {
			return engine.Pos{X: x, Y: y}
		}
	}
}

func GetEmptyPositionDeterministic(w *engine.World, seed int) engine.Pos {
	source := rand.NewSource(int64(seed))
	random := rand.New(source)

	for {
		x := random.Intn(int(WorldWidth))
		y := random.Intn(int(WorldHeight))

		if len(GetEntitiesAtPosition(w, engine.Pos{X: x, Y: y})) == 1 {
			return engine.Pos{X: x, Y: y}
		}
	}
}

func SpawnBuilding(w *engine.World, owner int, buildingType string, buildingWidth int, seed int) {
	source := rand.NewSource(int64(seed))
	random := rand.New(source)

	// randomize position and initialize
	x := random.Intn(WorldWidth)
	y := random.Intn(WorldHeight)

	AddBuilding(w, x, y, owner, buildingType)

	regionPositions := GetPositionsInRegion(x, y, buildingWidth)

	for _, pos := range regionPositions {
		if IsPositionInMap(pos.X, pos.Y, WorldWidth, WorldHeight) {
			AddBlocker(w, pos.X, pos.Y)
		}
	}
}

func GetPositionsInRegion(x int, y int, width int) []engine.Pos {
	positions := []engine.Pos{}

	halfWidth := (width - 1) / 2

	for i := x - halfWidth; i <= x+halfWidth; i++ {
		for j := y - halfWidth; j <= y+halfWidth; j++ {
			if IsPositionInMap(j, i, WorldWidth, WorldHeight) {
				positions = append(positions, engine.Pos{X: i, Y: j})
			}

		}
	}

	return positions
}

func IsPositionInMap(x int, y int, worldWidth int, worldHeight int) bool {
	return x >= 0 && x < worldWidth && y >= 0 && y < worldHeight
}

func GetRandomArrayElement(arr []int) int {
	randomIndex := rand.Intn(len(arr))
	return arr[randomIndex]
}

func GetRandomPositionArrayElement(arr []engine.Pos) engine.Pos {
	randomIndex := rand.Intn(len(arr))
	return arr[randomIndex]
}

// helpers

func InitializePokerCardsStringArray() []string {
	pokerSuits := []string{"club", "spade", "square", "heart"}
	numbers := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	pokerCardsStringArray := make([]string, len(pokerSuits)*len(numbers))

	// write a double for loop. initialize the 4 suits and the 13 cards in each suit
	for i, suit := range pokerSuits {
		for j, cardNumber := range numbers {
			pokerCardsStringArray[i*13+j] = suit + "-" + cardNumber
		}
	}

	return pokerCardsStringArray
}

func PosInPosArray(posArray []engine.Pos, pos engine.Pos) bool {
	for _, posInArray := range posArray {
		if posInArray.X == pos.X && posInArray.Y == pos.Y {
			return true
		}
	}
	return false
}

func ShuffleArray(arr []int) []int {
	// Create a new slice with the same length as the original array
	shuffled := make([]int, len(arr))
	copy(shuffled, arr)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}

func ShufflePositionArray(arr []engine.Pos) []engine.Pos {
	// Create a new slice with the same length as the original array
	shuffled := make([]engine.Pos, len(arr))
	copy(shuffled, arr)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled
}

// spawn an PC at a random location on the world map
func SpawnNPCs(w *engine.World, npcCount int) {

	for i := 0; i < npcCount; i++ {
		pos := GetEmptyPosition(w)

		AddInfantry(w, pos.X, pos.Y, 0)
	}
}

func SpawnNPCsDeterministic(w *engine.World, npcCount int, seed int) {
	for i := 0; i < npcCount; i++ {
		pos := GetEmptyPositionDeterministic(w, seed)

		AddInfantry(w, pos.X, pos.Y, 0)
	}
}

// register all components for game world
func RegisterAllComponents(w *engine.World) {
	w.AddComponentNew(components.TagComponent)
	w.AddComponentNew(components.PositionComponent)
	w.AddComponentNew(components.OwnerIdComponent)
	w.AddComponentNew(components.LevelComponent)
	w.AddComponentNew(components.HealthComponent)
	w.AddComponentNew(components.NameComponent)
	w.AddComponentNew(components.LastActiveComponent)
	w.AddComponentNew(components.TickJobTypeComponent)
	w.AddComponentNew(components.TickIdComponent)
	w.AddComponentNew(components.TickDataStringComponent)
	w.AddComponentNew(components.TickNumberComponent)
}
