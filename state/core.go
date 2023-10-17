package state

import (
	"reflect"
	"sync"
	"time"
)

// canonical position struct
type Pos struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// ---------------------------------------
//	Table (aka component)
// ---------------------------------------

// TODO we shouldn't be able to access these internals
// master table struct that holds table's entities
type Table struct {
	GameWorld *GameWorld

	// table name
	Name string

	// table type
	Type reflect.Type

	// entities
	Entities *SparseSet

	// entity int => value
	EntityToValue map[int]any

	// fieldName => fieldVal => entities sparse set
	//
	// ex: Person {name string, age int}
	// name => "Alice" => [1, 2, 3]
	// age => 20 => [3, 5, 11]
	Indexes map[string]map[any]*SparseSet // TODO create a struct that locks this

	mu *sync.Mutex
}

// initialize a new table that points to a world
func NewTable(w *GameWorld, table ITable) Table {
	return Table{
		GameWorld:     w,
		Name:          table.Name(),
		Type:          table.Type(),
		Entities:      NewSparseSet(),
		EntityToValue: make(map[int]any),
		Indexes:       make(map[string]map[any]*SparseSet),
		mu:            &sync.Mutex{},
	}
}

// core table set
func (t *Table) Set(w *GameWorld, entity int, value any) any {
	t.mu.Lock()
	defer t.mu.Unlock()

	w.entityManager.Add(entity)

	// this means that the entity is being set in the table for the first time
	if !t.Entities.Contains(entity) {
		t.Entities.Add(entity)

		// // add to list of updates
		update := TableUpdate{Entity: entity, Table: t.Name, Value: nil, OP: AddEntityOP, Time: time.Now().Unix()}
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

	update := TableUpdate{Entity: entity, Table: t.Name, Value: value, OP: UpdateOP, Time: time.Now().Unix()}
	w.AddTableUpdate(update)

	return value
}

func (t *Table) Get(entity int) (any, bool) {
	val, found := t.EntityToValue[entity]
	return val, found
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
	w.entityManager.Add(entity)
}

func (w *GameWorld) _addEntityNew() int {
	return w.entityManager.GetEntity()
}

// func remove entity from world
func (t *Table) RemoveEntity(w *GameWorld, entity int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	// remove from world association
	t.GameWorld.entityManager.Remove(entity)
	t.Entities.Remove(entity)

	// loop through each field in the struct and remove entities
	val, exists := t.EntityToValue[entity]
	if !exists {
		return
	}

	v := reflect.ValueOf(val)
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		fieldValue := v.Field(i)

		fieldMapping := t.Indexes[field.Name]
		fieldMapping[fieldValue.Interface()].Remove(entity)
	}

	update := TableUpdate{Entity: entity, Table: t.Name, Value: val, OP: RemovalOP, Time: time.Now().Unix()}
	w.AddTableUpdate(update)

	delete(t.EntityToValue, entity)
}

// the table interface
type ITable interface {
	Name() string
	Type() reflect.Type
}

// table accessor type
type TableBaseAccessor[T any] struct {
	TableName  string
	SchemaType reflect.Type
}

// create type map registration
func CreateTypeRegistrationMapping() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// new table accessor
func NewTableAccessor[T any]() *TableBaseAccessor[T] {
	var s T
	t := reflect.TypeOf(s)

	base := &TableBaseAccessor[T]{
		SchemaType: t,
		TableName:  t.Name(),
	}

	// check if the schema has an id field
	_, exists := t.FieldByName("Id")

	if !exists {
		panic("Every schema must have an Id field (we use this when syncing game state)")
	}

	return base
}

// get the table name
func (c *TableBaseAccessor[T]) Name() string {
	return c.TableName
}

// receive the reflect.Type value
func (c *TableBaseAccessor[T]) Type() reflect.Type {
	return c.SchemaType
}

// TODO: return a bool to show when the value is not found / set all table types to be pointers
func (c *TableBaseAccessor[T]) Get(w IWorld, entity int) T {
	v, _ := w.Get(entity, c.Name())
	val, _ := v.(T)
	return val
}

// needs to be used to communicate table updates
func (c *TableBaseAccessor[T]) RemoveEntity(w IWorld, entity int) int {
	return w.Delete(entity, c.Name())
}

// get all entities for this schema
func (c *TableBaseAccessor[T]) Entities(w IWorld) []int {
	return w.Entities(c.Name())
}

// needs to be used to communicate ecs updates
func (c *TableBaseAccessor[T]) Set(w IWorld, entity int, value T) T {
	w.Set(entity, value, c.Name())
	return value
}

// add the object to the world
func (c *TableBaseAccessor[T]) Add(w IWorld, obj T) int {
	return w.Add(obj, c.Name())
}

// adds the object to the world with a specific entity ID
func (c *TableBaseAccessor[T]) AddSpecific(w IWorld, entity int, obj T) int {
	return w.AddSpecific(entity, obj, c.Name())
}

// query by filtering the fields. returns list of entities
func (c *TableBaseAccessor[T]) Filter(w IWorld, filter T, fieldNames []string) []int {
	return w.Filter(filter, fieldNames, c.Name())
}

// core filter query function
func (t *Table) Filter(filter any, fieldNames []string) []int {
	if len(fieldNames) == 0 {
		return []int{}
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	v := reflect.ValueOf(filter)

	// if the query is only length 1, we can directly return
	if len(fieldNames) == 1 {
		fieldValue := v.FieldByName(fieldNames[0])

		// search in the reverse mapping and register
		fieldNameMapping := t.Indexes[fieldNames[0]]

		// if the name doesn't exist (aka wrong field name), return nothing
		if fieldNameMapping == nil {
			return []int{}
		}
		return t.Indexes[fieldNames[0]][fieldValue.Interface()].GetAll()
	}

	// we take the first one first
	fieldValue := v.FieldByName(fieldNames[0])

	var res []int = t.Indexes[fieldNames[0]][fieldValue.Interface()].GetAll()

	queryCtx := NewQueryContext()

	// we start search from the 2nd element
	for i := 1; i < len(fieldNames); i++ {
		fieldName := fieldNames[i]

		fieldVal := v.FieldByName(fieldName)

		fieldMapping := t.Indexes[fieldName]

		// field doesn't exist
		if fieldMapping == nil {
			return []int{}

		} else {

			// value => return entities that match the value of this struct's field
			valueSet := fieldMapping[fieldVal.Interface()]

			res = ArrayIntersectionWithContext(queryCtx, res, valueSet.GetAll())
		}
	}

	return res
}

type tableNameAndType struct {
	name string
	t    reflect.Type
}

func (t tableNameAndType) Name() string {
	return t.name
}

func (t tableNameAndType) Type() reflect.Type {
	return t.t
}

func (t *Table) TableInterface() ITable {
	return tableNameAndType{
		name: t.Name,
		t:    t.Type,
	}
}

func isEmptyValue(value reflect.Value) bool {
	zeroValue := reflect.Zero(value.Type())
	return reflect.DeepEqual(value.Interface(), zeroValue.Interface())
}

type TableOperationType string

// available op codes to indicate what type of update it is
var (
	// key indicating that the entity is being removed
	RemovalOP TableOperationType = "removal"

	// key indicating that the entity value is being set
	UpdateOP TableOperationType = "set"

	// key indicating that the entity is being added
	AddEntityOP TableOperationType = "add"
)

// deep copy table table updates
func CopyTableUpdates(updates []TableUpdate) []TableUpdate {
	res := make([]TableUpdate, len(updates))
	copy(res, updates)
	return res
}

// get and clear game world's table updates
func (w *GameWorld) GetAndClearTableUpdates() []TableUpdate {
	updates := CopyTableUpdates(w.TableUpdates)
	w.ClearTableUpdates()
	return updates
}

// clear game world's table updates
func (w *GameWorld) ClearTableUpdates() {
	w.TableUpdates = TableUpdateArray{}
}

func withEntity(obj interface{}, entity int) interface{} {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Interface || val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	newStruct := reflect.New(val.Type()).Elem()
	newStruct.Set(val)

	newStruct.FieldByName("Id").SetInt(int64(entity))
	return newStruct.Interface()
}
