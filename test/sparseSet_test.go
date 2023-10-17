package test

import (
	"fmt"
	"github.com/curio-research/keystone/keystone/state"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestSparseSetGeneralOperations(t *testing.T) {
	fmt.Println("Start test")

	set1 := state.NewSparseSet()

	count := 10_000

	elements := make([]int, 10_000)

	a := time.Now()
	for i := 0; i < count; i++ {
		set1.Add(i)

		// add random element
		set1.Add(rand.Intn(1_000))
		elements[i] = i
	}

	setTime := time.Since(a)
	fmt.Println("Insertion time:", setTime)

	// test for inclusion
	for _, element := range elements {
		if !set1.Contains(element) {
			t.Error("Sparse set does not contain element")
		}
	}

	// deletion test
	for _, element := range elements {
		set1.Remove(element)
	}

	fmt.Println("Size after deletion: ", set1.Size())
}

func TestSparseSetDeepCopy(t *testing.T) {
	sparseSet := state.NewSparseSet()

	itemCount := 100

	for i := 0; i < itemCount; i++ {
		sparseSet.Add(rand.Intn(1_000))
	}

	sparseSetCopy := sparseSet.DeepCopy()

	for i := 0; i < itemCount; i++ {
		if !sparseSetCopy.Contains(sparseSet.GetRandomElement()) {
			t.Error("Sparse set copy does not contain element")
		}
	}
}

// Result: parallel is significantly faster in larger number of things to copy
func TestParallelizeDeepcopySparseSet(t *testing.T) {

	count := 100_000
	sets_to_copy := 100

	// add to mapping from ID to set
	sets := make(map[int]state.SparseSet)

	for i := 0; i < sets_to_copy; i++ {
		tempSet := CreateAndPopulateSparseSet(count)
		sets[i] = tempSet
	}

	// serial time
	startTime := time.Now()

	for i := 0; i < sets_to_copy; i++ {
		set := sets[i]
		set.DeepCopy()
	}

	fmt.Println("Serial time: ", time.Since(startTime).Microseconds(), "ps")

	// parallel time
	startTime = time.Now()

	wg := sync.WaitGroup{}

	for i := 1; i < len(sets); i++ {
		set := sets[i]

		wg.Add(1)

		go func(setId int) {
			set.DeepCopy()

			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("Parallel time: ", time.Since(startTime).Microseconds(), "us")
	fmt.Println("Total elements: ", sets_to_copy*count)
}

func CreateAndPopulateSparseSet(count int) state.SparseSet {
	set := state.NewSparseSet()

	for i := 0; i < count; i++ {
		set.Add(i)
	}

	return *set
}
