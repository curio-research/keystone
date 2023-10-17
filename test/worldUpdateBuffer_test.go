package test

import (
	"github.com/curio-research/keystone/keystone/state"
	"github.com/stretchr/testify/assert"
	"testing"
)

// These tests are structured in this general way:
// 	1. Create a new GameWorld and a wrapping WorldUpdateBuffer (both considered IWorlds), both with registered Person/Book tables
//	2. Apply updates to either one or both IWorlds
//	3. Assert the updates for the buffer world
//	4. Assert that the current world does not have the updates made to the buffer
//	5. Assert that once the buffer updates are applied, the current world has the same updates as the buffer

func TestWorldUpdateBuffer_Add(t *testing.T) {
	currWorld := state.NewWorld()
	testRegisterTables(currWorld)
	bufferWorld := state.NewWorldUpdateBuffer(currWorld)

	p1 := Person{
		Name:     testName1,
		Position: testPos1,
		Address:  testAddress1,
	}
	person1Entity := personTable.Add(currWorld, p1)

	p2 := Person{
		Name:     testName2,
		Position: testPos2,
		Address:  testAddress1,
	}
	person2Entity := personTable.Add(bufferWorld, p2)

	p1.Id = 0
	p2.Id = 1

	assertUpdates := func(t *testing.T, w state.IWorld) {
		assert.Equal(t, p1, personTable.Get(w, person1Entity))
		assert.Equal(t, p2, personTable.Get(w, person2Entity))
	}

	assertUpdates(t, bufferWorld)

	assert.Equal(t, p1, personTable.Get(currWorld, person1Entity))
	assertEmpty(t, personTable.Get(currWorld, person2Entity))

	bufferWorld.ApplyUpdates()
	assertUpdates(t, currWorld)
}

func TestWorldUpdateBuffer_AddSpecific(t *testing.T) {
	currWorld := state.NewWorld()
	testRegisterTables(currWorld)
	bufferWorld := state.NewWorldUpdateBuffer(currWorld)

	p1 := Person{
		Name: testName1,
		Age:  testAge1,
	}
	person1Entity := personTable.AddSpecific(currWorld, 0, p1)

	b1 := Book{
		Title:  testBookTitle1,
		Author: testBookAuthor1,
	}
	book1Entity := bookTable.AddSpecific(bufferWorld, 1, b1)

	b2 := Book{
		Title:  testBookTitle2,
		Author: testBookAuthor2,
	}
	book2Entity := bookTable.Add(currWorld, b2)

	p2 := Person{
		Name:     testName2,
		Position: testPos2,
	}
	person2Entity := personTable.Add(bufferWorld, p2)

	p1.Id = 0
	b1.Id = 1
	b2.Id = 2
	p2.Id = 3

	assertUpdates := func(t *testing.T, w state.IWorld) {
		assert.Equal(t, p1, personTable.Get(w, person1Entity))
		assert.Equal(t, b1, bookTable.Get(w, book1Entity))
		assert.Equal(t, b2, bookTable.Get(w, book2Entity))
		assert.Equal(t, p2, personTable.Get(w, person2Entity))
	}

	assertUpdates(t, bufferWorld)

	assertEmpty(t, bookTable.Get(currWorld, book1Entity))
	assertEmpty(t, personTable.Get(currWorld, person2Entity))
	assert.Equal(t, p1, personTable.Get(currWorld, person1Entity))
	assert.Equal(t, b2, bookTable.Get(currWorld, book2Entity))

	bufferWorld.ApplyUpdates()
	assertUpdates(t, currWorld)
}

func TestWorldUpdateBuffer_Set(t *testing.T) {
	currWorld := state.NewWorld()
	testRegisterTables(currWorld)
	bufferWorld := state.NewWorldUpdateBuffer(currWorld)

	person := Person{
		Name:     testName1,
		Position: testPos1,
	}
	personEntity := personTable.Add(currWorld, person)

	book := Book{
		Title:  testBookTitle1,
		Author: testBookAuthor1,
	}
	bookEntity := bookTable.Add(bufferWorld, book)

	updatedPerson := person
	updatedPerson.Position = testPos2
	updatedPerson.Id = 0
	personTable.Set(bufferWorld, personEntity, updatedPerson)

	updatedBook := book
	updatedBook.Author = testBookAuthor2
	updatedBook.Id = 1
	bookTable.Set(bufferWorld, bookEntity, updatedBook)

	assertUpdates := func(t *testing.T, w state.IWorld) {
		assert.Equal(t, updatedPerson, personTable.Get(w, personEntity))
		assert.Equal(t, updatedBook, bookTable.Get(w, bookEntity))
	}

	assertUpdates(t, bufferWorld)

	assert.Equal(t, person, personTable.Get(currWorld, personEntity))
	assertEmpty(t, bookTable.Get(currWorld, bookEntity))

	bufferWorld.ApplyUpdates()
	assertUpdates(t, currWorld)
}

func TestWorldUpdateBuffer_Remove(t *testing.T) {
	currWorld := state.NewWorld()
	testRegisterTables(currWorld)
	bufferWorld := state.NewWorldUpdateBuffer(currWorld)

	book1 := Book{
		Title:  testBookTitle1,
		Author: testBookAuthor1,
	}
	book2 := Book{
		Title:  testBookTitle2,
		Author: testBookAuthor2,
	}
	book3 := book1
	book4 := Book{
		Title:  testBookTitle3,
		Author: testBookAuthor3,
	}

	book1Entity := bookTable.Add(bufferWorld, book1)
	book2Entity := bookTable.AddSpecific(bufferWorld, 34, book2)
	book3Entity := bookTable.Add(currWorld, book3)
	book4Entity := bookTable.AddSpecific(currWorld, 45, book4)

	book1.Id = 0
	book2.Id = 34
	book3.Id = 1
	book4.Id = 45

	bookTable.RemoveEntity(bufferWorld, book2Entity)
	bookTable.RemoveEntity(bufferWorld, book3Entity)

	assertUpdates := func(t *testing.T, w state.IWorld) {
		assert.Equal(t, book1, bookTable.Get(w, book1Entity))
		assertEmpty(t, bookTable.Get(w, book2Entity))
		assertEmpty(t, bookTable.Get(w, book3Entity))
		assert.Equal(t, book4, bookTable.Get(w, book4Entity))
	}

	assertUpdates(t, bufferWorld)

	assertEmpty(t, bookTable.Get(currWorld, book1Entity))
	assertEmpty(t, bookTable.Get(currWorld, book2Entity))
	assert.Equal(t, book3, bookTable.Get(currWorld, book3Entity))
	assert.Equal(t, book4, bookTable.Get(currWorld, book4Entity))

	bufferWorld.ApplyUpdates()
	assertUpdates(t, currWorld)
}

func TestWorldUpdateBuffer_Filter(t *testing.T) {
	currWorld := state.NewWorld()
	testRegisterTables(currWorld)
	bufferWorld := state.NewWorldUpdateBuffer(currWorld)

	person1Entity := personTable.Add(currWorld, Person{
		Name:    testName1,
		Age:     testAge1,
		Address: testAddress1,
	})

	person2Entity := personTable.Add(bufferWorld, Person{
		Name:    testName2,
		Age:     testAge1,
		Address: testAddress2,
	})

	person3Entity := personTable.Add(bufferWorld, Person{
		Name:    testName3,
		Age:     testAge2,
		Address: testAddress1,
	})

	person4Entity := personTable.Add(currWorld, Person{
		Name:    testName1,
		Age:     testAge2,
		Address: testAddress2,
	})

	person5Entity := personTable.Add(bufferWorld, Person{
		Name:    testName2,
		Age:     testAge1,
		Address: testAddress1,
	})

	// to test entity is marked for deletion
	personTable.RemoveEntity(bufferWorld, person4Entity)

	assertChanges := func(t *testing.T, w state.IWorld) {
		filter1 := personTable.Filter(w, Person{Age: testAge1}, []string{"Age"})
		assert.Len(t, filter1, 3)
		assert.Contains(t, filter1, person1Entity)
		assert.Contains(t, filter1, person2Entity)
		assert.Contains(t, filter1, person5Entity)

		filter2 := personTable.Filter(w, Person{Age: testAge2}, []string{"Age"})
		assert.Len(t, filter2, 1)
		assert.Contains(t, filter2, person3Entity)

		filter3 := personTable.Filter(w, Person{Address: testAddress1}, []string{"Address"})
		assert.Len(t, filter3, 3)
		assert.Contains(t, filter3, person1Entity)
		assert.Contains(t, filter3, person3Entity)
		assert.Contains(t, filter3, person5Entity)

		filter4 := personTable.Filter(w, Person{Address: testAddress2}, []string{"Address"})
		assert.Len(t, filter4, 1)
		assert.Contains(t, filter4, person2Entity)

		filter5 := personTable.Filter(w, Person{Age: testAge1, Address: testAddress1}, []string{"Age", "Address"})
		assert.Len(t, filter5, 2)
		assert.Contains(t, filter3, person1Entity)
		assert.Contains(t, filter3, person5Entity)
	}

	assertChanges(t, bufferWorld)

	filter1 := personTable.Filter(currWorld, Person{Age: testAge1}, []string{"Age"})
	assert.Len(t, filter1, 1)
	assert.Contains(t, filter1, person1Entity)

	filter2 := personTable.Filter(currWorld, Person{Address: testAddress2}, []string{"Address"})
	assert.Len(t, filter2, 1)
	assert.Contains(t, filter2, person4Entity)

	bufferWorld.ApplyUpdates()
	assertChanges(t, currWorld)
}

func TestWorldUpdateBuffer_Entities_Add(t *testing.T) {
	currWorld := state.NewWorld()
	testRegisterTables(currWorld)
	bufferWorld := state.NewWorldUpdateBuffer(currWorld)

	personTable.Add(currWorld, Person{})
	personTable.Add(currWorld, Person{})
	personTable.Add(bufferWorld, Person{})
	personTable.Add(bufferWorld, Person{})

	personTable.AddSpecific(currWorld, 45, Person{})
	personTable.AddSpecific(bufferWorld, 56, Person{})
	personTable.AddSpecific(currWorld, 69, Person{})
	personTable.AddSpecific(currWorld, 78, Person{})

	assertUpdates := func(t *testing.T, w state.IWorld) {
		entities := personTable.Entities(w)
		assert.Len(t, entities, 8)
		for _, i := range []int{0, 1, 2, 3, 45, 56, 69, 78} {
			assert.Contains(t, entities, i)
		}
	}

	assertUpdates(t, bufferWorld)

	entities := personTable.Entities(currWorld)
	assert.Len(t, entities, 5)
	for _, i := range []int{0, 1, 45, 69, 78} {
		assert.Contains(t, entities, i)
	}

	bufferWorld.ApplyUpdates()
	assertUpdates(t, currWorld)
}

func TestWorldUpdateBuffer_Entities_Remove(t *testing.T) {
	currWorld := state.NewWorld()
	testRegisterTables(currWorld)
	bufferWorld := state.NewWorldUpdateBuffer(currWorld)

	personTable.Add(currWorld, Person{})
	personTable.Add(currWorld, Person{})
	personTable.Add(bufferWorld, Person{})
	personTable.Add(bufferWorld, Person{})

	personTable.AddSpecific(currWorld, 45, Person{})
	personTable.AddSpecific(bufferWorld, 56, Person{})
	personTable.AddSpecific(currWorld, 69, Person{})
	personTable.AddSpecific(currWorld, 78, Person{})

	personTable.RemoveEntity(bufferWorld, 1)
	personTable.RemoveEntity(bufferWorld, 3)
	personTable.RemoveEntity(bufferWorld, 45)
	personTable.RemoveEntity(bufferWorld, 78)

	assertUpdates := func(t *testing.T, w state.IWorld) {
		entities := personTable.Entities(w)
		assert.Len(t, entities, 4)
		for _, i := range []int{0, 2, 56, 69} {
			assert.Contains(t, entities, i)
		}
	}

	assertUpdates(t, bufferWorld)

	entities := personTable.Entities(currWorld)
	assert.Len(t, entities, 5)
	for _, i := range []int{0, 1, 45, 69, 78} {
		assert.Contains(t, entities, i)
	}

	bufferWorld.ApplyUpdates()
	assertUpdates(t, currWorld)
}

func assertEmpty[T any](t *testing.T, obj T) {
	var empty T
	assert.Equal(t, empty, obj)
}
