package server

import (
	"strconv"

	"github.com/curio-research/keystone/state"
)

type TransactionSchema struct {
	Type string

	// a uuid that's sent from the client, which is usd to identify which quest has been satisfieds
	Uuid string

	Data string

	TickNumber int

	Id int

	UnixTimestamp int

	IsExternal bool
}

var (
	TransactionTable = state.NewTableAccessor[TransactionSchema]()
)

func AddSystemTransaction(w *state.GameWorld, tickNumber int, transactionType string, data string, uuid string, isExternal bool) int {
	entity := TransactionTable.Add(w, TransactionSchema{
		Type:       transactionType,
		Uuid:       uuid,
		Data:       data,
		TickNumber: tickNumber,
		IsExternal: isExternal,
	})

	return entity
}

// registers default tables keystone must operates on such as tick related
func RegisterDefaultTables(w *state.GameWorld) {
	w.AddTables(TransactionTable)
}

func GetTransactionUuid(w state.IWorld, transactionId int) int {
	transaction := TransactionTable.Get(w, transactionId)
	i, _ := strconv.ParseInt(transaction.Uuid, 10, 32)

	return int(i)
}
