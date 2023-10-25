package core

import "encoding/json"

// "table update" represents 1 single table update to the state
type TableUpdateArray []TableUpdate

type TableUpdate struct {
	// "op codes" that represent what type of operation it is
	OP TableOperationType `json:"op"`

	// the entity
	Entity int `json:"entity"`

	// the name of the table
	Table string `json:"table"`

	// the value that's being updated
	Value interface{} `json:"value"`

	// unix timestamp
	Time int64 `json:"time"`
}

func EncodeTableUpdateArrayToBytes(tableUpdates []TableUpdate) ([]byte, error) {
	return json.Marshal(tableUpdates)
}
