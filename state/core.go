package state

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/curio-research/keystone/utils"
	"github.com/golang-collections/collections/stack"
)

// ---------------------------------------
//	Table (aka component)
// ---------------------------------------

// Each table corresponds to a schema (ex: Player, Animal, Tile)
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
	// "name" => "Alice" => [1, 2, 3]
	// "age" => "20" => [3, 5, 11] // json representation of int
	Indexes map[string]map[string]*SparseSet // TODO create a struct that locks this

	mu *sync.Mutex
}

// Initialize a new table that points to a world
func NewTable(w *GameWorld, table ITable) Table {
	return Table{
		GameWorld:     w,
		Name:          table.Name(),
		Type:          table.Type(),
		Entities:      NewSparseSet(),
		EntityToValue: make(map[int]any),
		Indexes:       make(map[string]map[string]*SparseSet),
		mu:            &sync.Mutex{},
	}
}

// Core table set function
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
			fieldNameMapping = make(map[string]*SparseSet)
			t.Indexes[field.Name] = fieldNameMapping
		}

		// remove value => entity mapping for the previous value if the old value exists
		if prevVal != nil {
			prevFieldValue := _prevV.Field(i)
			fieldNameMapping[tableKey(prevFieldValue)].Remove(entity)
		}

		// create new set if not present

		key := tableKey(fieldValue)
		if fieldNameMapping[key] == nil {
			fieldNameMapping[key] = NewSparseSet()
		}

		// add value
		fieldNameMapping[key].Add(entity)
	}

	update := TableUpdate{Entity: entity, Table: t.Name, Value: value, OP: UpdateOP, Time: time.Now().Unix()}
	w.AddTableUpdate(update)

	return value
}

// Get from table with entity
func (t *Table) Get(entity int) (any, bool) {
	val, found := t.EntityToValue[entity]
	return val, found
}

// Get all entities numbers of a table
func (t *Table) All() []int {
	return t.Entities.GetAll()
}

// ------------------------------
// add entities to world
// ------------------------------

// Add entity to world. Fetches the next available entity
func (w *GameWorld) AddEntity() int {
	return w.entityManager.GetNextAvailableEntity()
}

// Add a specific entity to the world
// This is useful in debugging and assigning constant entities to constant numbers
func (w *GameWorld) AddSpecificEntity(entity int) {
	w.entityManager.Add(entity)
}

// Remove entity from world
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
		fieldMapping[tableKey(fieldValue)].Remove(entity)
	}

	update := TableUpdate{Entity: entity, Table: t.Name, Value: val, OP: RemovalOP, Time: time.Now().Unix()}
	w.AddTableUpdate(update)

	delete(t.EntityToValue, entity)
}

// Table interface
type ITable interface {
	Name() string
	Type() reflect.Type
}

// Table accessors are used to access table data with proper types of its schema
type TableBaseAccessor[T any] struct {
	TableName  string
	SchemaType reflect.Type
}

// Create type map registration
func CreateTypeRegistrationMapping() map[string]reflect.Type {
	return map[string]reflect.Type{}
}

// Initialize new table accessor
func NewTableAccessor[T any]() *TableBaseAccessor[T] {
	var s T
	t := reflect.TypeOf(s)

	base := &TableBaseAccessor[T]{
		SchemaType: t,
		TableName:  t.Name(),
	}

	// check if the schema has an id field
	field, exists := t.FieldByName("Id")
	if !exists {
		panic(t.Name() + " is missing an Id field (we use this when syncing game state)")
	}
	idTag := field.Tag.Get("gorm")
	if !strings.Contains(t.String(), "TransactionSchema") {
		if !strings.Contains(idTag, "primaryKey") {
			panic(fmt.Sprintf("Id field of %s needs `gorm:\"primaryKey\"` tag", t.Name()))
		}
	}

	jsonArrayName := reflect.TypeOf(utils.SerializableArray[any]{}).String()
	jsonArrayName = strings.Split(jsonArrayName, "[")[0]
	jsonArrayName = strings.Split(jsonArrayName, ".")[1]

	stk := stack.New()
	for i := 0; i < t.NumField(); i++ {
		stk.Push(t.Field(i))
	}

	for stk.Len() != 0 {
		f := stk.Pop().(reflect.StructField)
		kind := f.Type.Kind()
		gormTag := f.Tag.Get("gorm")
		if kind == reflect.Slice {
			if !strings.Contains(f.Type.Name(), jsonArrayName) {
				panic(fmt.Sprintf("Field %s of %s must be of SerializableArray type", f.Name, t.Name()))
			}

			if gormTag != "serializer:json" {
				panic(fmt.Sprintf("Field %s of %s needs `gorm:\"serializer:json\"` tag", f.Name, t.Name()))
			}
		} else if gormTag == "embedded" {
			for i := 0; i < f.Type.NumField(); i++ {
				stk.Push(f.Type.Field(i))
			}
		}
	}
	return base
}

// Get the table name
func (c *TableBaseAccessor[T]) Name() string {
	return c.TableName
}

// Receive the reflect.Type value of a table's schema
func (c *TableBaseAccessor[T]) Type() reflect.Type {
	return c.SchemaType
}

// Get object using entity from a world
func (c *TableBaseAccessor[T]) Get(w IWorld, entity int) T {
	v, _ := w.Get(entity, c.Name())
	val, _ := v.(T)
	return val
}

// Remove entity from world
func (c *TableBaseAccessor[T]) RemoveEntity(w IWorld, entity int) int {
	return w.Delete(entity, c.Name())
}

// Get all entities for a schema table
func (c *TableBaseAccessor[T]) Entities(w IWorld) []int {
	return w.Entities(c.Name())
}

// Set entity and its object value to world
func (c *TableBaseAccessor[T]) Set(w IWorld, entity int, value T) T {
	w.Set(entity, value, c.Name())
	return value
}

// Add the object to the world and returns the assigned entity number
func (c *TableBaseAccessor[T]) Add(w IWorld, obj T) int {
	return w.Add(obj, c.Name())
}

// Adds the object to the world with a specific entity ID
func (c *TableBaseAccessor[T]) AddSpecific(w IWorld, entity int, obj T) int {
	return w.AddSpecific(entity, obj, c.Name())
}

// Query by filtering the fields. returns list of entities
func (c *TableBaseAccessor[T]) Filter(w IWorld, filter T, fieldNames []string) []int {
	return w.Filter(filter, fieldNames, c.Name())
}

func checkStructMatchFieldValues(fullStruct any, structWithValsToMatch any, fieldNamesToMatch []string) bool {

	fullStructReflectVal := reflect.ValueOf(fullStruct)
	structWithValsToMatchRefecectVal := reflect.ValueOf(structWithValsToMatch)

	// check if the struct matches the field names
	for _, fieldName := range fieldNamesToMatch {
		fullStructFieldVal := fullStructReflectVal.FieldByName(fieldName).Interface()
		structWithValsToMatchFieldVal := structWithValsToMatchRefecectVal.FieldByName(fieldName).Interface()

		if fullStructFieldVal != structWithValsToMatchFieldVal {
			return false
		}
	}

	return true
}

// Core filter query function
func (t *Table) Filter(filter any, fieldNames []string) []int {

	// this needs to be an array
	// queryCtx := NewQueryContext()
	// var filteredEntities []int

	// // loop through all fields to check if the field has a tag.
	// // if there's a tag, fetch value and merge it with "res"

	// // if not, skip values
	// rt := reflect.TypeOf(filter)
	// foundFirstIndexCache := false
	// if rt.Kind() == reflect.Struct {
	// 	for i := 0; i < rt.NumField(); i++ {
	// 		field := rt.Field(i)
	// 		_, keyExists := field.Tag.Lookup("key")
	// 		if keyExists {

	// 			cachedEntities := t.Indexes[field.Name][tableKey(reflect.ValueOf(filter).Field(i))].GetAll()

	// 			// first time found, register idx
	// 			if !foundFirstIndexCache {
	// 				filteredEntities = append(filteredEntities, cachedEntities...)
	// 				foundFirstIndexCache = true

	// 			} else {

	// 				filteredEntities = ArrayIntersectionWithContext(queryCtx, filteredEntities, cachedEntities)
	// 			}
	// 		}
	// 	}
	// }

	var finalRes []int

	// for every entity, fetch its value
	// add to results if all fields match values in the Filter function
	for _, entity := range t.Entities.GetAll() {
		entityValue := t.EntityToValue[entity]
		isMatch := checkStructMatchFieldValues(entityValue, filter, fieldNames)
		if isMatch {
			finalRes = append(finalRes, entity)
		}
	}

	return finalRes

	// if len(fieldNames) == 0 {
	// 	return []int{}
	// }

	// t.mu.Lock()
	// defer t.mu.Unlock()

	// v := reflect.ValueOf(filter)

	// // if the query is only length 1, we can directly return
	// if len(fieldNames) == 1 {
	// 	fieldValue := v.FieldByName(fieldNames[0])

	// 	// search in the reverse mapping and register
	// 	fieldNameMapping := t.Indexes[fieldNames[0]]

	// 	// if the name doesn't exist (aka wrong field name), return nothing
	// 	if fieldNameMapping == nil {
	// 		return []int{}
	// 	}

	// 	return t.Indexes[fieldNames[0]][tableKey(fieldValue)].GetAll()
	// }

	// // we take the first one first
	// fieldValue := v.FieldByName(fieldNames[0])

	// var res []int = t.Indexes[fieldNames[0]][tableKey(fieldValue)].GetAll()

	// queryCtx := NewQueryContext()

	// // we start search from the 2nd element
	// for i := 1; i < len(fieldNames); i++ {
	// 	fieldName := fieldNames[i]

	// 	fieldVal := v.FieldByName(fieldName)

	// 	fieldMapping := t.Indexes[fieldName]

	// 	// field doesn't exist
	// 	if fieldMapping == nil {
	// 		return []int{}
	// 	} else {
	// 		// value => return entities that match the value of this struct's field
	// 		valueSet := fieldMapping[tableKey(fieldVal)]

	// 		res = ArrayIntersectionWithContext(queryCtx, res, valueSet.GetAll())
	// 	}
	// }

	// return res
}

func tableKey(val reflect.Value) string {
	v, _ := json.Marshal(val.Interface())
	return string(v)
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

type TableOperationType string

// Available op codes to indicate what type of update it is
var (
	// key indicating that the entity is being removed
	RemovalOP TableOperationType = "removal"

	// key indicating that the entity value is being set
	UpdateOP TableOperationType = "set"

	// key indicating that the entity is being added
	AddEntityOP TableOperationType = "add"
)

// Deep copy table table updates
func CopyTableUpdates(updates []TableUpdate) []TableUpdate {
	res := make([]TableUpdate, len(updates))
	copy(res, updates)
	return res
}

// Get and clear game world's table updates
func (w *GameWorld) GetAndClearTableUpdates() []TableUpdate {
	updates := CopyTableUpdates(w.TableUpdates)
	w.ClearTableUpdates()
	return updates
}

// Clear game world's table updates
func (w *GameWorld) ClearTableUpdates() {
	w.TableUpdates = TableUpdateArray{}
}

func assignIdFieldInSchemaWithEntity(obj interface{}, entity int) interface{} {
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Interface || val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	newStruct := reflect.New(val.Type()).Elem()
	newStruct.Set(val)

	newStruct.FieldByName("Id").SetInt(int64(entity))
	return newStruct.Interface()
}

// Universal position struct
type Pos struct {
	X int `json:"x"`
	Y int `json:"y"`
}
