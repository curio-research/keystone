// SPDX-License-Identifier: BUSL-1.1

// Copyright (C) 2023, Curiosity Research. All rights reserved.
// Use of this software is covered by the Business Source License included
// in the LICENSE file in the license folder of this repository and at www.mariadb.com/bsl11.

// Any use of the Licensed Work in violation of this License will automatically
// terminate your rights under this License for the current and all other
// versions of the Licensed Work.

// This License does not grant you any right in any trademark or logo of
// Licensor or its affiliates (provided that you may use a trademark or logo of
// Licensor as expressly required by this License).

// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN "AS IS" BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package engine

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestSparseSetGeneralOperations(t *testing.T) {
	fmt.Println("Start test")

	set1 := NewSparseSet()

	count := 10_000

	elements := make([]int64, 10_000)

	a := time.Now()
	for i := 0; i < count; i++ {
		set1.Add(int64(i))

		// add random element
		set1.Add(int64(rand.Intn(1_000)))
		elements[i] = int64(i)
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
	sparseSet := NewSparseSet()

	itemCount := 100

	for i := 0; i < itemCount; i++ {
		sparseSet.Add(int64(rand.Intn(1_000)))
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
	sets := make(map[int]SparseSet)

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
