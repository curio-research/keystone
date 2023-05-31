package state

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
)

type StorageDiff struct {
	Address  common.Address
	Key      common.Hash
	OldValue common.Hash
	NewValue common.Hash
}

func CompareDirtyStorage(statedbBefore *StateDB, statedbAfter *StateDB, addr common.Address) []StorageDiff {
	var diffs []StorageDiff

	dirtyStatedbBefore := statedbBefore.stateObjects[addr].dirtyStorage
	dirtyStatedbAfter := statedbAfter.stateObjects[addr].dirtyStorage

	for key := range dirtyStatedbAfter {
		valueBefore := dirtyStatedbBefore[key]
		valueAfter := dirtyStatedbAfter[key]
		valueBeforeBytes := valueBefore.Bytes()
		valueAfterBytes := valueAfter.Bytes()
		compareKeys := bytes.Compare(valueBeforeBytes, valueAfterBytes)

		if compareKeys != 0 {
			diffs = append(diffs, StorageDiff{
				Address:  addr,
				Key:      key,
				OldValue: valueBefore,
				NewValue: valueAfter,
			})
		}
	}

	return diffs
}
