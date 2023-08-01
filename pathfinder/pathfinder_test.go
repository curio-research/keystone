package pathfinder

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// Test parallel path finding with new path finding
// new a* path finder uses a cacheLayer which stores start and end locations
// this prevents the deepCopy step

// Speeds up aggregate ~4x
func TestPathfindingNew(t *testing.T) {

	// construct new map
	gameMap := ConstructMap2dArray(102, 102, 0.1)

	gameMapCopy := DeepCopy2DArr(gameMap)
	worldMap := ConstructWorldNew(gameMapCopy)

	wg := sync.WaitGroup{}
	a := time.Now()

	iterations := 500

	for i := 0; i < iterations; i++ {

		wg.Add(1)
		go func(id int) {

			fromPos := Pos{X: rand.Intn(100), Y: rand.Intn(100)}
			toPos := Pos{X: rand.Intn(100), Y: rand.Intn(100)}

			a := time.Now()

			path, _, _ := AstarPathfinder(fromPos, toPos, worldMap)

			fmt.Println("⏰ Path time: ", int(time.Since(a).Milliseconds()), "ms")
			fmt.Println("🏃 Path length: ", len(path))

			wg.Done()

		}(i)

	}

	wg.Wait()

	fmt.Println("⏰ Total time: ", int(time.Since(a).Milliseconds()), "ms")
	fmt.Println("⏰ Avg time: ", int(time.Since(a).Milliseconds()/int64(iterations)), "ms")
}
