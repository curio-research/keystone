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

import "math/rand"

var exists = struct{}{}

type SparseSet struct {
	// array of all elements
	elements []int64

	// mapping between element to index position
	elementIndexMap []int64
}

func NewSparseSet() *SparseSet {
	s := &SparseSet{}

	// TODO: we assume a 100k upper ceiling here for number of entities
	s.elementIndexMap = make([]int64, 100_000)

	return s
}

func (s *SparseSet) Add(value int64) {
	if s.Contains(value) {
		return
	}

	s.elements = append(s.elements, value)
	s.elementIndexMap[value] = int64(s.Size() - 1)
}

func (s *SparseSet) Remove(value int64) {
	if !s.Contains(value) {
		return
	}

	indexToRemove := s.elementIndexMap[value]
	lastElement := s.elements[s.Size()-1]
	s.elements[indexToRemove] = lastElement
	s.elementIndexMap[lastElement] = indexToRemove

	s.elementIndexMap[value] = 0

	s.elements = s.elements[:len(s.elements)-1]

}

func (s *SparseSet) Size() int {
	if s == nil {
		return 0
	}
	return len(s.elements)
}

func (s *SparseSet) GetAll() []int64 {
	return s.elements
}

func (s *SparseSet) Contains(value int64) bool {
	// null pointer check
	if s == nil {
		return false
	}

	if len(s.elements) == 0 {
		return false
	}

	return s.elementIndexMap[value] != 0 || s.elements[0] == value
}

func SetIntersection(set1 *SparseSet, set2 *SparseSet) *SparseSet {
	// sanity checkers
	if set1 == nil {
		set1 = NewSparseSet()
	}
	if set2 == nil {
		set2 = NewSparseSet()
	}

	intersection := NewSparseSet()

	for _, element := range set1.elements {
		if set2.Contains(element) {
			intersection.Add(element)
		}
	}

	return intersection
}

func SetDifference(set1 *SparseSet, set2 *SparseSet) *SparseSet {
	difference := NewSparseSet()

	if set1.Size() > set2.Size() {
		// set2 is smaller
		// loop through set2
		for _, element := range set2.elements {
			if !set1.Contains(element) {
				difference.Add(element)
			}
		}
		return difference

	} else {
		// set1 is smaller
		// loop through set1
		for _, element := range set1.elements {
			if !set2.Contains(element) {
				difference.Add(element)
			}
		}
		return difference

	}

}

func SetUnion(set1 *SparseSet, set2 *SparseSet) *SparseSet {
	union := NewSparseSet()

	for _, element := range set1.elements {
		union.Add(element)
	}

	for _, element := range set2.elements {
		union.Add(element)
	}

	return union
}

func (s *SparseSet) DeepCopy() *SparseSet {
	newSet := NewSparseSet()

	for _, element := range s.elements {
		newSet.Add(element)
	}

	return newSet

}

func CreateAndPopulateSparseSet(count int) SparseSet {
	set := NewSparseSet()

	for i := 0; i < count; i++ {
		set.Add(int64(i))
	}

	return *set
}

func (s *SparseSet) GetRandomElement() int64 {
	return s.elements[rand.Intn(s.Size()-1)]
}
