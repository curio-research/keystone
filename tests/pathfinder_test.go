package tests

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	pathfinder "github.com/ethereum/go-ethereum/pathfinder"
)

// Test parallel path finding with new path finding
// new a* path finder uses a cacheLayer which stores start and end locations
// this prevents the deepCopy step
//
//	Speeds up aggregate ~4x
func TestPathfindingNew(t *testing.T) {

	gameMapCopy := pathfinder.DeepCopy2DArr(pathfinder.GameMap)
	worldMap := pathfinder.ConstructWorldNew(gameMapCopy)

	wg := sync.WaitGroup{}
	a := time.Now()

	iterations := 500

	for i := 0; i < iterations; i++ {

		wg.Add(1)
		go func(id int) {

			fromPos := pathfinder.Pos{X: int64(rand.Intn(100)), Y: int64(rand.Intn(100))}
			toPos := pathfinder.Pos{X: int64(rand.Intn(100)), Y: int64(rand.Intn(100))}

			a := time.Now()

			path, _, _ := pathfinder.AstarPathfinder(fromPos, toPos, worldMap)

			fmt.Println("⏰ Path time: ", int(time.Since(a).Milliseconds()), "ms")
			fmt.Println("🏃 Path length: ", len(path))

			wg.Done()

		}(i)

	}

	wg.Wait()

	fmt.Println("⏰ Total time: ", int(time.Since(a).Milliseconds()), "ms")
	fmt.Println("⏰ Avg time: ", int(time.Since(a).Milliseconds()/int64(iterations)), "ms")
}
