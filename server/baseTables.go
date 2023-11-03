package server

import (
	"strconv"

	"github.com/curio-research/keystone/state"
)

type TransactionSchema struct {
	// Entity ID
	Id int `gorm:"primaryKey;autoIncrement:false"`

	// Transaction type
	Type string

	// A uuid that's sent from the client, which is usd to identify which quest has been satisfied
	Uuid string

	// Data payload serialized to string format
	Data string

	// Tick number when the transaction was to be processed
	TickNumber int

	// Timestamp of when the transaction was received by server
	UnixTimestamp int

	// whether transaction was submitted by player or other systems
	IsExternal bool
}

var (
	TransactionTable = state.NewTableAccessor[TransactionSchema]()
)

var BaseTableSchemasToAccessors = map[interface{}]*state.TableBaseAccessor[any]{
	&TransactionSchema{}: (*state.TableBaseAccessor[any])(TransactionTable),
}

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
