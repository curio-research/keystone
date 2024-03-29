package test

import (
	"github.com/curio-research/keystone/utils"
	"testing"

	"github.com/curio-research/keystone/state"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTable(t *testing.T) {
	w := state.NewWorld()
	testRegisterTables(w)

	alice1 := Person{
		Name:     "Alice",
		Age:      0,
		Address:  "123 Main St",
		Position: state.Pos{X: 1, Y: 2},
		BookId:   0,
	}
	alice2 := Person{
		Name:     "Alice",
		Age:      1,
		Address:  "123 Main St",
		Position: state.Pos{X: 1, Y: 2},
		BookId:   0,
	}
	bobby := Person{
		Name:     "Bobby",
		Age:      0,
		Address:  "Metaverse St",
		Position: state.Pos{X: 1, Y: 2},
		BookId:   0,
	}

	book1 := Book{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"}

	alice1Entity := personTable.Add(w, alice1)
	alice2Entity := personTable.Add(w, alice2)
	bobbyEntity := personTable.Add(w, bobby)

	bookEntity := bookTable.Add(w, book1)

	person1Retrieved := personTable.Get(w, alice1Entity)
	assert.Equal(t, alice1.Name, person1Retrieved.Name)

	person1Retrieved.Name = "VITALIK"
	personTable.Set(w, alice1Entity, person1Retrieved)

	// --------------------
	//     query test
	// --------------------

	personQuery := Person{Name: "Alice", Age: 0}
	queryFields := []string{"Name", "Age"}
	res := personTable.Filter(w, personQuery, queryFields)
	assert.Len(t, res, 0)

	personQuery = Person{Name: "Alice"}
	queryFields = []string{"Name"}
	res = personTable.Filter(w, personQuery, queryFields)
	require.Len(t, res, 1)
	assert.Equal(t, alice2Entity, res[0])

	personQuery = Person{Name: "VITALIK", Age: 0}
	queryFields = []string{"Name", "Age"}
	res = personTable.Filter(w, personQuery, queryFields)
	require.Len(t, res, 1)
	assert.Equal(t, alice1Entity, res[0])

	personQuery = Person{Address: "Metaverse St"}
	queryFields = []string{"Address"}
	res = personTable.Filter(w, personQuery, queryFields)
	require.Len(t, res, 1)
	assert.Equal(t, bobbyEntity, res[0])

	bookQuery := Book{Title: "The Great Gatsby", Author: "Dr. Seuss"}
	queryFields = []string{"Title", "Author"}
	res = bookTable.Filter(w, bookQuery, queryFields)
	assert.Len(t, res, 0)

	bookQuery = Book{Title: "The Great Gatsb", Author: "Dr. Seuss"}
	queryFields = []string{"Title", "Author"}
	res = bookTable.Filter(w, bookQuery, queryFields)
	assert.Len(t, res, 0)

	bookQuery = Book{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"}
	queryFields = []string{"Title", "Author"}
	res = bookTable.Filter(w, bookQuery, queryFields)
	require.Len(t, res, 1)
	assert.Equal(t, bookEntity, res[0])
}

func TestFilter(t *testing.T) {
	type TransactionSchema struct {
		Type       string
		Uuid       string // a uuid that's sent from the client, which is usd to identify which quest has been satisfied
		Data       string
		TickNumber int
		Id         int `gorm:"primaryKey"`
		IsExternal bool
	}

	w := state.NewWorld()

	// create new tables
	TransactionTable := state.NewTableAccessor[TransactionSchema]()

	w.AddTables(TransactionTable)

	// create transactions
	count := 10
	for i := 0; i < count; i++ {
		tx := TransactionSchema{
			Type:       "aaa",
			TickNumber: i,
		}
		TransactionTable.Add(w, tx)
	}

	// query
	query := TransactionSchema{TickNumber: 1, Type: "move"}
	queryFields := []string{"TickNumber", "Type"}
	res := TransactionTable.Filter(w, query, queryFields)
	if len(res) != 0 {
		t.Error("Expected empty result because we query by a wrong field name")
	}
}

func TestProcessTxs(t *testing.T) {
	w := state.NewWorld()

	type TransactionSchema struct {
		Type       string
		Data       string
		TickNumber int
		TickUuid   string
		Id         int `gorm:"primaryKey"`
	}

	var (
		TransactionTable = state.NewTableAccessor[TransactionSchema]()
	)

	w.AddTables(TransactionTable)

	txId := TransactionSchema{
		Type:       "bbb",
		TickNumber: 1,
	}
	TransactionTable.Add(w, txId)

	// query
	query := TransactionSchema{TickNumber: 1, Type: "move"}
	queryFields := []string{"TickNumber", "Type"}
	res := TransactionTable.Filter(w, query, queryFields)
	if len(res) != 0 {
		t.Error("Expected empty result because we query by a wrong field name")
	}
}

func Test_PanicWithNoID(t *testing.T) {
	type structWithoutID struct {
		Name string
	}
	assert.Panicsf(t, func() {
		state.NewTableAccessor[structWithoutID]()
	}, "Every schema must have an Id field (we use this when syncing game state)")

	type structWithoutTag struct {
		Name string
		Id   int
	}
	assert.Panicsf(t, func() {
		state.NewTableAccessor[structWithoutTag]()
	}, "Id field needs `gorm:\"primaryKey\"` tag")

	type structWithID struct {
		Name string
		Id   int `gorm:"primaryKey"`
	}
	assert.NotPanics(t, func() {
		state.NewTableAccessor[structWithID]()
	})
}

func Test_PanicWithWrongArrayType(t *testing.T) {
	type structWithWrongArray struct {
		Arr []string
		Id  int `gorm:"primaryKey"`
	}
	assert.Panicsf(t, func() {
		state.NewTableAccessor[structWithWrongArray]()
	}, "Every array in the top level of a schema must be of SerializableArray type")

	type structWithWrongArrayTag struct {
		Arr utils.SerializableArray[string] `gorm:"type:json"`
		Id  int                             `gorm:"primaryKey"`
	}
	assert.Panicsf(t, func() {
		state.NewTableAccessor[structWithWrongArrayTag]()
	}, "Array field in top level of a struct needs `gorm:\"serializer:json\"` tag")

	type correctStruct struct {
		Arr utils.SerializableArray[string] `gorm:"serializer:json"`
		Id  int                             `gorm:"primaryKey"`
	}
	assert.NotPanics(t, func() {
		state.NewTableAccessor[correctStruct]()
	})
}

func Test_PanicWithWrongArrayType_EmbeddedArray(t *testing.T) {
	type embeddedStruct struct {
		Arr []string
	}

	type structWithWrongArray struct {
		Emb embeddedStruct `gorm:"embedded"`
		Id  int            `gorm:"primaryKey"`
	}

	assert.Panicsf(t, func() {
		state.NewTableAccessor[structWithWrongArray]()
	}, "Every array in its own column must be of SerializableArray type")

	type embeddedStruct2 struct {
		Arr utils.SerializableArray[string]
	}

	type structWithWrongArrayTag struct {
		Emb embeddedStruct2 `gorm:"embedded"`
		Id  int             `gorm:"primaryKey"`
	}

	assert.Panicsf(t, func() {
		state.NewTableAccessor[structWithWrongArrayTag]()
	}, "Array field in top level of a struct needs `gorm:\"serializer:json\"` tag")

	type embeddedStruct3 struct {
		Arr utils.SerializableArray[string] `gorm:"serializer:json"`
	}

	type correctStruct struct {
		Emb embeddedStruct3 `gorm:"embedded"`
		Id  int             `gorm:"primaryKey"`
	}
	assert.NotPanics(t, func() {
		state.NewTableAccessor[correctStruct]()
	})
}
