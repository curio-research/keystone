package ecs

import (
	"fmt"
	"testing"
)

type Person struct {
	Name     string
	Age      int
	Address  string
	Position Pos
	BookId   int // TODO: if we can automatically solve the linkage that'd be OP
}

type Book struct {
	Title  string
	Author string
}

func TestTable(t *testing.T) {
	w := NewWorld()

	// create new tables
	personTable := NewTable[Person]()
	bookTable := NewTable[Book]()

	w.AddTables(personTable, bookTable)

	person1 := Person{
		Name:     "Alice",
		Age:      0,
		Address:  "123 Main St",
		Position: Pos{X: 1, Y: 2},
		BookId:   0,
	}
	person2 := Person{
		Name:     "Bobby",
		Age:      0,
		Address:  "Metaverse St",
		Position: Pos{X: 1, Y: 2},
		BookId:   0,
	}

	book1 := Book{"The Great Gatsby", "F. Scott Fitzgerald"}

	person1Entity := personTable.Add(w, person1)
	personTable.Add(w, person2)
	bookTable.Add(w, book1)

	person1Retrieved := personTable.Get(w, person1Entity)
	person1Retrieved.Name = "VITALIK"
	personTable.Set(w, person1Entity, person1Retrieved)

	// --------------------
	//     query test
	// --------------------

	query := Person{Name: "Alice", Age: 0}
	queryFields := []string{"Name", "Age"}

	res := personTable.Filter(w, query, queryFields)
	fmt.Println("Filter res :", res)

}
