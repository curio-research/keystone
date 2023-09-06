package ecs

import (
	"reflect"
)

// canonical position struct
type Pos struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// master game world holding entities and such
type GameWorld struct {
	Entities *SparseSet

	EntitiesNonce int

	Tables map[string]Table

	TableUpdates ECSUpdateArray
}

// add single table to game world
func (w *GameWorld) AddTable(table ITable) {
	w.Tables[table.Name()] = Table{
		GameWorld:     w,
		Name:          table.Name(),
		Entities:      NewSparseSet(),
		EntityToValue: make(map[int]any),
		Indexes:       make(map[string]map[any]*SparseSet),
	}
}

// add multiple tables to the game world
func (w *GameWorld) AddTables(tables ...ITable) {
	for _, table := range tables {
		w.AddTable(table)
	}
}

// adds a filled schema to the world. Creates the proper entity, etc.
func AddToWorld[T any](w *GameWorld, obj T) int {
	entity := w.AddEntityNew()

	tableName := reflect.TypeOf(obj).Name()

	table := w.Tables[tableName]

	table.Set(w, entity, obj)

	return entity

}

// core add a struct to a world on a specific entity
func AddToWorldSpecific[T any](w *GameWorld, entity int, obj T) int {
	tableName := reflect.TypeOf(obj).Name()

	table := w.Tables[tableName]

	w.AddSpecificEntityNew(entity)

	table.Set(w, entity, obj)

	return entity

}

// initialize new game world
func NewWorld() *GameWorld {
	return &GameWorld{
		Entities: NewSparseSet(),
		Tables:   make(map[string]Table),
	}
}

// add table updates to game world
func (w *GameWorld) AddTableUpdate(tableUpdate ECSUpdate) {
	w.TableUpdates = append(w.TableUpdates, tableUpdate)
}

// ---------------------------------------
//	table struct (aka component)
// ---------------------------------------

// master table struct that holds table's entities
type Table struct {
	GameWorld *GameWorld

	// table name
	Name string

	// entities
	Entities *SparseSet

	// entity int => value
	EntityToValue map[int]any

	// fieldName => fieldVal => entities sprase set
	//
	// ex: Person {name string, age int}
	// name => "Alice" => [1, 2, 3]
	// age => 20 => [3, 5, 11]
	Indexes map[string]map[any]*SparseSet
}

// core table set
func (t *Table) Set(w *GameWorld, entity int, value any) any {
	w.Entities.Add(entity)

	// this means that the entity is being set in the table for the first time
	if !t.Entities.Contains(entity) {
		t.Entities.Add(entity)

		// TODO: add this back if needed
		// // add to list of updates
		update := ECSUpdate{Entity: w.EntitiesNonce, Table: t.Name, Value: nil, OP: AddEntityOP}
		w.AddTableUpdate(update)
	}

	// get the previous value, and remove that association from the fields
	prevVal := t.EntityToValue[entity]
	_prevV := reflect.ValueOf(prevVal)

	t.EntityToValue[entity] = value

	reflectVal := reflect.ValueOf(value)
	reflectType := reflect.ValueOf(value).Type()

	// register reverse mappings
	for i := 0; i < reflectVal.NumField(); i++ {
		field := reflectType.Field(i)
		fieldValue := reflectVal.Field(i)

		fieldNameMapping := t.Indexes[field.Name]

		if fieldNameMapping == nil {
			fieldNameMapping = make(map[any]*SparseSet)
			t.Indexes[field.Name] = fieldNameMapping
		}

		// remove value => entity mapping for the previous value if the old value exists
		if prevVal != nil {
			prevFieldValue := _prevV.Field(i)
			fieldNameMapping[prevFieldValue.Interface()].Remove(entity)
		}

		// create new set if not present
		if fieldNameMapping[fieldValue.Interface()] == nil {
			fieldNameMapping[fieldValue.Interface()] = NewSparseSet()
		}

		// add value
		fieldNameMapping[fieldValue.Interface()].Add(entity)

	}

	update := ECSUpdate{Entity: entity, Table: t.Name, Value: value, OP: UpdateOP}
	w.AddTableUpdate(update)

	return value
}

// TODO: add type safety?
func (t *Table) Get(entity int) any {
	return t.EntityToValue[entity]
}

// get all entities
func (t *Table) All() []int {
	return t.Entities.GetAll()
}

// ------------------------------
// add entities to world
// ------------------------------

// add entity to world
func (w *GameWorld) AddEntityNew() int {
	entity := w._addEntityNew()
	return entity
}

func (w *GameWorld) AddSpecificEntityNew(entity int) {
	w.Entities.Add(entity)
}

func (w *GameWorld) _addEntityNew() int {
	w.EntitiesNonce++

	// if the world already allocated this entity, then increase it
	if w.Entities.Contains(w.EntitiesNonce) {

		return w._addEntityNew()

	} else {
		w.Entities.Add(w.EntitiesNonce)

	}

	return w.EntitiesNonce
}

// func remove entity from world
func (t *Table) RemoveEntity(w *GameWorld, entity int) {
	// remove from world association
	t.GameWorld.Entities.Remove(entity)

	// entity's value
	val := t.EntityToValue[entity]

	t.Entities.Remove(entity)

	// loop through each field in the struct and remove entities
	v := reflect.ValueOf(val)
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		fieldValue := v.Field(i)

		fieldMapping := t.Indexes[field.Name]
		fieldMapping[fieldValue.Interface()].Remove(entity)
	}

	update := ECSUpdate{Entity: entity, Table: t.Name, Value: nil, OP: RemovalOP}
	w.AddTableUpdate(update)

	delete(t.EntityToValue, entity)
}

// TODO: New
type ITable interface {
	Name() string
}

type TableBaseAccessor[T any] struct {
	TableName string
}

// create type map registration
func CreateTypeRegistrationMapping() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

var RegisteredTypes = map[string]reflect.Type{}

// new table accessor
func NewTable[T any]() *TableBaseAccessor[T] {
	var s T
	base := &TableBaseAccessor[T]{
		TableName: reflect.TypeOf(s).Name(),
	}

	// TODO: experimental
	RegisteredTypes[base.TableName] = reflect.TypeOf(s)

	return base
}

func (c *TableBaseAccessor[T]) Name() string {
	return c.TableName
}

func (c *TableBaseAccessor[T]) Get(w *GameWorld, entity int) T {
	table := w.Tables[c.TableName]

	value := table.Get(entity)

	// TODO: add error handling
	val, _ := value.(T)

	return val
}

func (c *TableBaseAccessor[T]) RemoveEntity(w *GameWorld, entity int) {
	table := w.Tables[c.TableName]

	table.RemoveEntity(w, entity)
}

func (c *TableBaseAccessor[T]) Entities(w *GameWorld) []int {
	table := w.Tables[c.TableName]
	return table.All()
}

func (c *TableBaseAccessor[T]) Set(w *GameWorld, entity int, value T) T {
	table := w.Tables[c.TableName]

	table.Set(w, entity, value)

	return value
}

// add the object to the world
func (c *TableBaseAccessor[T]) Add(w *GameWorld, obj T) int {
	return AddToWorld(w, obj)
}

// adds the object to the world with a specific entity ID
func (c *TableBaseAccessor[T]) AddSpecific(w *GameWorld, entity int, obj T) int {
	return AddToWorldSpecific(w, entity, obj)
}

// query by filtering the fields. returns list of entities
func (c *TableBaseAccessor[T]) Filter(w *GameWorld, filter T, fieldNames []string) []int {
	table := w.Tables[c.TableName]

	return table.Filter(filter, fieldNames)
}

// core filter query function
func (t *Table) Filter(filter any, fieldNames []string) []int {
	v := reflect.ValueOf(filter)

	// TODO: is deep copy needed?
	// not sure if this is needed
	var res []int = t.Entities.GetAll()

	queryCtx := NewQueryContext()

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		fieldValue := v.Field(i)

		// if it's included in the list of field Names, then filter
		if IncludesString(fieldNames, field.Name) || (IncludesString(fieldNames, field.Name) && isEmptyValue(fieldValue)) {
			fieldMapping := t.Indexes[field.Name]

			// value => return entities that match the value of this struct's field
			valueSet := fieldMapping[fieldValue.Interface()]

			res = ArrayIntersectionWithContext(queryCtx, res, valueSet.GetAll())

		}
	}

	return res
}

func isEmptyValue(value reflect.Value) bool {
	zeroValue := reflect.Zero(value.Type())
	return reflect.DeepEqual(value.Interface(), zeroValue.Interface())
}

// "ecs update" represents 1 single ecs state update
type ECSUpdateArray []ECSUpdate

type ECSUpdate struct {
	// "op codes" that represent what type of operation it is
	OP     string      `json:"op"`
	Entity int         `json:"entity"`
	Table  string      `json:"table"`
	Value  interface{} `json:"value"`
}

var (
	// key indicating that the entity is being removed
	RemovalOP = "removal"

	// key indicating that the entity value is being set
	UpdateOP = "set"

	// key indicating that the entity is being added
	AddEntityOP = "add"
)

func StartRecordingStateChanges(w *GameWorld) *GameWorld {

	// clear table updates
	w.TableUpdates = ECSUpdateArray{}

	return w
}
