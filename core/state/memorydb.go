package state

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

type FuncionSlotItem struct {
	lock     sync.Locker
	slotList []common.Hash
}

type FuncionSlotMap = map[string]FuncionSlotItem

type FunctionItem struct {
	address      common.Address
	functionName string
}

// A in-memory db caching function slot relationship
type MemoryDB struct {
	addressFunctionSlotMap map[common.Address]FuncionSlotMap
	slotStoreMutex         sync.Locker
	serialFuntions         []FunctionItem
	parallelFunctions      []FunctionItem
}

func NewMemoryDB() *MemoryDB {
	memorydb := &MemoryDB{}

	memorydb.addressFunctionSlotMap = make(map[common.Address]map[string]FuncionSlotItem)
	memorydb.slotStoreMutex = &sync.RWMutex{}

	return memorydb
}

func (m *MemoryDB) SetSlotStore(addr common.Address, functionName string, statedbBefore *StateDB, statedbAfter *StateDB) {
	functionSlotMap := m.addressFunctionSlotMap[addr]

	if functionSlotMap == nil {
		functionSlotMap = make(map[string]FuncionSlotItem)
	}

	slotItem := functionSlotMap[functionName]

	if slotItem.slotList != nil {
		return
	}

	slotItem = FuncionSlotItem{
		lock:     &sync.RWMutex{},
		slotList: []common.Hash{},
	}
	m.AddLock()

	// compare statedb before and after
	diffs := CompareDirtyStorage(statedbBefore, statedbAfter, addr)

	for _, diff := range diffs {
		slotItem.slotList = append(slotItem.slotList, diff.Key)
	}

	functionSlotMap[functionName] = slotItem
	m.addressFunctionSlotMap[addr] = functionSlotMap

	m.RemoveLock()
}

func (m *MemoryDB) UpdateFunctionItem(addr common.Address, functionName string) {
	functionSlotMap := m.addressFunctionSlotMap[addr]

	if functionSlotMap == nil {
		m.parallelFunctions = append(m.parallelFunctions, FunctionItem{addr, functionName})
		return
	}

	selfSlotItem := functionSlotMap[functionName]
	selfSlotList := selfSlotItem.slotList

	slotCounter := 0

	for _, slotItem := range functionSlotMap {
		slotList := slotItem.slotList

		if hasOverlay(selfSlotList, slotList) {
			slotCounter += 1
		}
	}

	if slotCounter > 1 {
		m.serialFuntions = append(m.serialFuntions, FunctionItem{addr, functionName})
	} else {
		m.parallelFunctions = append(m.parallelFunctions, FunctionItem{addr, functionName})
	}
}

func (m *MemoryDB) GetSerialFuntions() []FunctionItem {
	return m.serialFuntions
}

func (m *MemoryDB) GetParallelFuntions() []FunctionItem {
	return m.parallelFunctions
}

func (item *FunctionItem) GetAddress() common.Address {
	return item.address
}

func (item *FunctionItem) GetFunctionName() string {
	return item.functionName
}

func (item *FunctionItem) SetAddress(address common.Address) {
	item.address = address
}

func (item *FunctionItem) SetFunctionName(functionName string) {
	item.functionName = functionName
}

func hasOverlay(s1, s2 []common.Hash) bool {
	hash := make(map[common.Hash]bool)
	inter := []common.Hash{}

	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		// If elements present in the hashmap then append intersection list.
		if hash[e] {
			inter = append(inter, e)
		}
	}

	return len(inter) > 0
}

func (m *MemoryDB) AddLock() {
	m.slotStoreMutex.Lock()
}

func (m *MemoryDB) RemoveLock() {
	m.slotStoreMutex.Unlock()
}

// Get Contract Address, just for test
func GetAddressList(statedb *StateDB) []common.Address {
	if statedb.stateObjects == nil {
		return []common.Address{}
	}

	addressList := []common.Address{}

	for key := range statedb.stateObjects {
		if len(statedb.stateObjects[key].originStorage) > 0 {
			addressList = append(addressList, key)
		}
	}

	return addressList
}
