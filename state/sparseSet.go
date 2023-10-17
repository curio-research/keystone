package state

import (
	"math/rand"
	"sync"
)

type SparseSet struct {
	// array of all elements
	elements []int

	// general purpose mutex
	mutex sync.Mutex

	// element => index mapping in elements variable
	elementIndexMap map[int]int
}

func NewSparseSet() *SparseSet {
	s := &SparseSet{}

	s.mutex = sync.Mutex{}
	s.elementIndexMap = make(map[int]int)

	return s
}

func (s *SparseSet) Add(value int) {
	if s.Contains(value) {
		return
	}

	s.mutex.Lock()

	s.elements = append(s.elements, value)
	s.elementIndexMap[value] = s.Size() - 1

	s.mutex.Unlock()
}

func (s *SparseSet) Remove(value int) {
	if !s.Contains(value) {
		return
	}

	s.mutex.Lock()

	indexToRemove := s.elementIndexMap[value]
	lastElement := s.elements[s.Size()-1]
	s.elements[indexToRemove] = lastElement
	s.elementIndexMap[lastElement] = indexToRemove

	delete(s.elementIndexMap, value)

	s.elements = s.elements[:len(s.elements)-1]

	s.mutex.Unlock()
}

func (s *SparseSet) Size() int {
	if s == nil {
		return 0
	}
	return len(s.elements)
}

func (s *SparseSet) GetAll() []int {
	if s == nil {
		return []int{}
	}
	return s.elements
}

func (s *SparseSet) Contains(value int) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s == nil {
		return false
	}

	if len(s.elements) == 0 {
		return false
	}

	return s.elementIndexMap[value] != 0 || s.elements[0] == value
}

func (s *SparseSet) DeepCopy() *SparseSet {
	newSet := NewSparseSet()

	for _, element := range s.elements {
		newSet.Add(element)
	}

	return newSet

}

func (s *SparseSet) GetRandomElement() int {
	return s.elements[rand.Intn(s.Size()-1)]
}

func ArrayToSparseSet(array []int) *SparseSet {
	set := NewSparseSet()

	for _, element := range array {
		set.Add(element)
	}

	return set
}
