package state

import "reflect"

type IWorld interface {
	Add(obj interface{}, name string) int
	AddSpecific(entity int, obj interface{}, name string) int
	Get(entity int, tableName string) (any, bool)
	Set(entity int, obj interface{}, name string) interface{}
	Delete(entity int, tableName string) int
	Filter(filter interface{}, fieldNames []string, name string) []int
	Entities(string) []int
	GetTableUpdates() TableUpdateArray
}

// master game world holding entities and tables
// all data belong to a table (aka component in ECS) that have different schemas (ex: Cat{}, Dog{})
type GameWorld struct {
	// current entity nonce and sparse set of entities
	entityManager *EntityManager

	// all tables that stores data
	Tables map[string]Table

	// array of all table updates
	TableUpdates TableUpdateArray
}

// initialize new game world
func NewWorld() *GameWorld {
	return &GameWorld{
		Tables:        make(map[string]Table),
		entityManager: NewEntityMananger(),
	}
}

func (w *GameWorld) Add(obj interface{}, tableName string) int {
	return AddToWorld(w, obj)
}

func (w *GameWorld) AddSpecific(entity int, obj interface{}, tableName string) int {
	return AddToWorldSpecific(w, entity, obj)
}

func (w *GameWorld) Set(entity int, obj interface{}, tableName string) interface{} {
	return Set(w, entity, obj)
}

func (w *GameWorld) Get(entity int, tableName string) (any, bool) {
	table := w.Tables[tableName]
	return table.Get(entity)
}

func (w *GameWorld) Delete(entity int, tableName string) int {
	table := w.Tables[tableName]
	table.RemoveEntity(w, entity)
	return entity
}

func (w *GameWorld) Filter(filter interface{}, fieldNames []string, tableName string) []int {
	return Filter(w, filter, fieldNames)
}

func (w *GameWorld) Entities(tableName string) []int {
	table := w.Tables[tableName]
	return table.All()
}

func (w *GameWorld) GetTableUpdates() TableUpdateArray {
	return w.TableUpdates
}

// add single table to game world
func (w *GameWorld) AddTable(table ITable) {
	w.Tables[table.Name()] = NewTable(w, table)
}

// add multiple tables to the game world
func (w *GameWorld) AddTables(tables ...ITable) {
	for _, table := range tables {
		w.AddTable(table)
	}
}

// add table updates to game world
func (w *GameWorld) AddTableUpdate(tableUpdate TableUpdate) {
	w.TableUpdates = append(w.TableUpdates, tableUpdate)
}

// adds a filled schema to the world. Creates the proper entity, etc.
func AddToWorld(w *GameWorld, obj any) int {
	entity := w.AddEntityNew()
	obj = withEntity(obj, entity)

	tableName := reflect.TypeOf(obj).Name()
	table := w.Tables[tableName]
	table.Set(w, entity, obj)

	return entity

}

// core add a struct to a world on a specific entity
func AddToWorldSpecific(w *GameWorld, entity int, obj any) int {
	w.AddSpecificEntityNew(entity)
	obj = withEntity(obj, entity)

	tableName := reflect.TypeOf(obj).Name()
	table := w.Tables[tableName]
	table.Set(w, entity, obj)

	return entity
}

func Set[T any](w *GameWorld, entity int, obj T) int {
	tableName := reflect.TypeOf(obj).Name()

	table := w.Tables[tableName]
	table.Set(w, entity, obj)

	return entity
}

func Filter[T any](w *GameWorld, filter T, fieldNames []string) []int {
	tableName := reflect.TypeOf(filter).Name()

	table := w.Tables[tableName]
	return table.Filter(filter, fieldNames)
}
