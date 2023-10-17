package state

import "sync"

type EntityManager struct {
	nonce int
	*SparseSet
	mu *sync.Mutex
}

func NewEntityMananger() *EntityManager {
	return &EntityManager{
		SparseSet: NewSparseSet(),
		mu:        &sync.Mutex{},
	}
}

func (e *EntityManager) GetEntity() int {
	e.mu.Lock()
	defer e.mu.Unlock()

	n := e.nonce
	for e.SparseSet.Contains(n) {
		e.nonce++
		n = e.nonce
	}

	e.nonce++
	return n
}
