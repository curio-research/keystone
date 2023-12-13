package test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"github.com/curio-research/keystone/test/testutils"
)

func TestEmptyQuerySpeed(t *testing.T) {
	testutils.SkipTestIfShort(t)

	testing.Short()

	ctx, _, _, _, _ := startTestServer(t, server.Dev)

	startTime := time.Now()
	times := 100

	for i := 0; i < times; i++ {
		server.GetTickTransactionsOfType(ctx.World, "abc", i)
	}

	elapsed := time.Since(startTime)
	fmt.Println("Time elapsed: ", elapsed)
	fmt.Println("Each empty query took: ", divideTimeDuration(elapsed, times))

}

func TestQuerySpeed(t *testing.T) {
	testutils.SkipTestIfShort(t)

	// testing query through different methods
	counts := []int{10, 100, 1000, 10000, 100000, 200000, 300000}

	for _, count := range counts {
		ctx, _, _, _, _ := startTestServer(t, server.Dev)

		// populate world with troops
		for i := 0; i < count; i++ {
			bookTable.Add(ctx.World, Book{
				Title:   "a",
				Author:  "b",
				OwnerID: i,
			})
		}

		query := Book{Title: "a", Author: "b"}

		// 1) default filter query method
		startTime := time.Now()
		bookTable.Filter(ctx.World, query, []string{"Title", "Author"})
		elapsed1 := time.Since(startTime)

		// 2) loop through directly. in larger sets it's slower
		startTime = time.Now()

		res1 := []int{}
		table := ctx.World.Tables[bookTable.Name()]
		for entityId, rawBook := range table.EntityToValue {
			book := rawBook.(Book)
			if book.Title == "a" && book.Author == "b" {
				res1 = append(res1, entityId)
			}
		}
		elapsed2 := time.Since(startTime)

		fmt.Println(elapsed1, " vs ", elapsed2)
	}
}

func TestBufferTableCreationSpeed(t *testing.T) {
	testutils.SkipTestIfShort(t)

	counts := []int{10, 100, 1000, 10000, 100000, 200000, 300000}

	ctx, _, _, _, _ := startTestServer(t, server.Dev)

	for _, count := range counts {
		startTime := time.Now()
		for i := 0; i < count; i++ {
			state.NewWorldUpdateBuffer(ctx.World)
		}

		elapsed := time.Since(startTime)

		fmt.Println("Count: ", count, " | Time Elapsed: ", elapsed, " | Each took: ", divideTimeDuration(elapsed, count))
	}
}

func randomInRange(a, b int) int {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Ensure a <= b
	if a > b {
		a, b = b, a
	}

	// Generate a random number in the range [a, b]
	return rand.Intn(b-a+1) + a
}

// performance test for how fast we can apply a state
func TestApplyTx(t *testing.T) {
	testutils.SkipTestIfShort(t)

	ctx, _, _, _, _ := startTestServer(t, server.Dev)
	go func() {
		for {
			select {
			case <-ctx.StateUpdateCh:
			case <-ctx.TransactionCh:
			}
		}
	}()

	ctx.GameTick.Schedule.AddSystem(0, perfTestSystem)

	startTime := time.Now()

	txCount := 15325

	server.TickWorldForward(ctx, txCount)

	elapsed := time.Since(startTime)

	fmt.Println()
	fmt.Println("Applying " + strconv.Itoa(txCount) + " tx took " + elapsed.String() + "✅")
	fmt.Println("Time per tx: ", divideTimeDuration(elapsed, txCount))
	fmt.Println()
}

var perfTestSystem = server.CreateGeneralSystem(func(ctx *server.TransactionCtx[any]) {
	w := ctx.W

	// this should be similar to a real transaction size
	for i := 0; i < 5; i++ {
		troopId := bookTable.Add(w, Book{
			OwnerID: i,
			Title:   "a",
			Author:  "b",
		})

		bookTable.RemoveEntity(w, troopId)
	}
})

func divideTimeDuration(duration time.Duration, divider int) time.Duration {
	nanoseconds := duration.Nanoseconds()
	resultNanoseconds := nanoseconds / int64(divider)
	return time.Duration(resultNanoseconds)
}
