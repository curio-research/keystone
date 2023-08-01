// SPDX-License-Identifier: BUSL-1.1

// Copyright (C) 2023, Curiosity Research. All rights reserved.
// Use of this software is covered by the Business Source License included
// in the LICENSE file in the license folder of this repository and at www.mariadb.com/bsl11.

// Any use of the Licensed Work in violation of this License will automatically
// terminate your rights under this License for the current and all other
// versions of the Licensed Work.

// This License does not grant you any right in any trademark or logo of
// Licensor or its affiliates (provided that you may use a trademark or logo of
// Licensor as expressly required by this License).

// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN "AS IS" BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package engine

import (
	"errors"
	"reflect"
	"sync"
)

// ECS based world struct
type World struct {
	// snapshot mode doesn't make changes to the underlying world. Instead it creates another
	// ecs world on top to record changes, allowing the base layer to be read only
	SnapshotMode bool

	// the parent world in snapshot mode
	Parent *World

	Entities      *SparseSet
	EntitiesNonce int

	Components map[string]Component

	// ecs changes will be recorded in an array which can be exported
	ShouldRecordEcsUpdates bool

	// record ecs changes. used in snapshot mode
	// This allows you to dump all ecs changes
	EcsChanges []ECSData

	ShellMode bool
}

// ECS component
type Component struct {
	DataType                   DataType
	ShouldStoreValueToEntities bool

	ComponentMutex  *sync.Mutex
	Entities        *SparseSet
	EntitiesToValue map[int]string
	ValueToEntities map[any]*SparseSet // mapping: value (bytes) -> Entity set
}

type ExportECSWorld map[string]ExportComponent

type ExportComponent struct {
	DataType DataType `json:"dataType"`
	Entities []int    `json:"entities"`

	EntitiesToValue map[int]interface{} `json:"entitiesToValue"`
	ValueToEntities map[string][]int    `json:"valueToEntities"`
}

type DataType int

const (
	Number   DataType = 0
	String   DataType = 1
	Position DataType = 2
	Address  DataType = 3
	Bool     DataType = 4
)

type Pos struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type EcsValueUnit struct {
	Component string
	Value     []byte
	Entity    int
}

type EntityStep struct {
	EntityId int
	Pos      Pos
}

// query types
var (
	Has      int = 0
	HasExact int = 2
)

type QueryCondition struct {
	QueryType int
	Component string
	Value     any
}

// add entity
func (w *World) AddEntity() int {

	world := w
	if w.ShellMode {
		world = w.Parent
	}

	world.EntitiesNonce++

	world.Entities.Add(world.EntitiesNonce)

	if world.ShouldRecordEcsUpdates {

		w.AddEcsChange(ECSData{
			Entity:    w.EntitiesNonce,
			Component: "+",
		})
	}

	return world.EntitiesNonce
}

func (w *World) AddSpecificEntity(entityId int) {
	w.Entities.Add(entityId)
}

func NewGameWorld() *World {
	w := &World{}

	w.Entities = NewSparseSet()
	w.Components = make(map[string]Component)

	return w
}

// Snapshottable world
func NewSnapshottableWorld(parent *World) *World {
	w := NewGameWorld()

	w.SnapshotMode = true
	w.ShouldRecordEcsUpdates = true

	// if there's a parent, we first copy all the components and entities from their parents
	// the assumption is that components will mostly be the same
	if parent != nil {

		w.EntitiesNonce = parent.EntitiesNonce
		w.Parent = parent
		w.Entities = parent.Entities.DeepCopy()

		for componentName, component := range parent.Components {
			w.AddComponent(componentName, component.DataType, component.ShouldStoreValueToEntities)

			// copy entities of a component over
			parentComponent, ok := parent.Components[componentName]
			if ok {
				component.Entities = parentComponent.Entities.DeepCopy()
			}
		}
	}

	return w
}

// shell mode's parent must not be null and is responsible for recording ecs changes
func StartRecordingStateChanges(parentWorld *World) *World {
	w := &World{}

	w.ShellMode = true
	w.ShouldRecordEcsUpdates = true
	w.Parent = parentWorld

	return w
}

func (w *World) ExportEcsChanges() []ECSData {
	return w.EcsChanges
}

func (w *World) ClearEcsChanges() {
	w.EcsChanges = []ECSData{}
}

func (w *World) AddEcsChange(ecsUpdate ECSData) {
	w.EcsChanges = append(w.EcsChanges, ecsUpdate)
}

// apply multiple ecs updates to a world
// TODO: add error detecting and reversion
func ApplyEcsUpdatesToWorld(ecsUpdates []ECSData, w *World) {
	for _, ecsUpdate := range ecsUpdates {
		ApplyEcsUpdateToWorld(ecsUpdate, w)
	}
}

// apply single ecs update to a world
func ApplyEcsUpdateToWorld(ecsUpdate ECSData, w *World) {

	if ecsUpdate.Component == "-" {
		w.RemoveEntity(ecsUpdate.Entity)
	} else if ecsUpdate.Component == "+" {
		// add entity
		w.AddSpecificEntity(ecsUpdate.Entity)
	} else {
		// normal mode. apply ECS update
		w.SetComponentValue(ecsUpdate.Entity, ecsUpdate.Component, ecsUpdate.Value)
	}
}

// target world is the one that's being changed
func ApplyWorldChangesToOtherWorld(source *World, target *World) {
	target.EntitiesNonce = source.EntitiesNonce

	ApplyEcsUpdatesToWorld(source.ExportEcsChanges(), target)
}

// add component
func (w *World) AddComponent(name string, dataType DataType, ShouldStoreValueToEntities bool) error {

	// check if component already exists
	_, ok := w.Components[name]
	if ok {
		return errors.New("component already exists")
	}

	w.Components[name] = Component{
		DataType:                   dataType,
		EntitiesToValue:            make(map[int]string),
		ComponentMutex:             &sync.Mutex{},
		Entities:                   NewSparseSet(),
		ValueToEntities:            make(map[interface{}]*SparseSet),
		ShouldStoreValueToEntities: ShouldStoreValueToEntities,
	}

	return nil
}

func (w *World) AddComponentNew(component NewComponentInterface) error {

	_, ok := w.Components[component.Name()]
	if ok {
		return errors.New("component already exists")
	}

	w.Components[component.Name()] = Component{
		DataType:                   ReflectTypeToDataType(component.Type()),
		EntitiesToValue:            make(map[int]string),
		ComponentMutex:             &sync.Mutex{},
		Entities:                   NewSparseSet(),
		ValueToEntities:            make(map[interface{}]*SparseSet),
		ShouldStoreValueToEntities: component.ShouldReverseLookup(),
	}

	return nil
}

// TODO: get rid of this intermediary type completely
func ReflectTypeToDataType(dataType reflect.Type) DataType {

	switch dataType.String() {
	case reflect.String.String():
		return String
	case reflect.Int.String():
		return Number
	case reflect.Bool.String():
		return Bool
	case reflect.TypeOf(Pos{}).String():
		return Position

		// TODO: add case for address
	default:
		return String
	}
}

func (w *World) GetComponent(name string) (Component, bool) {
	component, ok := w.Components[name]
	return component, ok
}

func (w *World) SetComponentValue(entity int, componentName string, value any) ECSData {

	world := w
	if w.ShellMode {
		world = w.Parent
	}

	component, ok := world.GetComponent(componentName)

	component.ComponentMutex.Lock()

	if !ok {
		return ECSData{}
	}

	oldValueString := component.EntitiesToValue[entity]
	decodedOldValue := DecodeBytesBasedOnDataType(component.DataType, oldValueString)

	if decodedOldValue != nil {
		// remove from valueToEntities mapping
		elements, _ := component.ValueToEntities[decodedOldValue]
		elements.Remove(entity)
	}

	world.Entities.Add(entity)

	component.Entities.Add(entity)

	switch dataType := component.DataType; dataType {
	case Number:
		switch val := value.(type) {
		case int:
			world.Components[componentName].EntitiesToValue[entity] = EncodeToStringBasedOnDataType(dataType, val)

		case int64:
			world.Components[componentName].EntitiesToValue[entity] = EncodeToStringBasedOnDataType(dataType, int(val))
		}

	case String:
		val, ok := value.(string)

		if ok {
			world.Components[componentName].EntitiesToValue[entity] = EncodeToStringBasedOnDataType(dataType, val)
		}

	case Position:
		val, ok := value.(Pos)

		if ok {
			world.Components[componentName].EntitiesToValue[entity] = EncodeToStringBasedOnDataType(dataType, val)
		}

	case Address:
		val, ok := value.(string)

		if ok {
			world.Components[componentName].EntitiesToValue[entity] = EncodeToStringBasedOnDataType(dataType, val)
		}

	default:

	}

	// store in reverse mapping
	if component.ShouldStoreValueToEntities {

		if world.SnapshotMode {
			// recursively find the last sparseSet
			lastFoundSparsSet := world.FindLastSparseSet(componentName, value)

			if lastFoundSparsSet == nil {

				// if it doesn't exist, create a new one
				component.ValueToEntities[value] = NewSparseSet()

			} else {

				// if it exists, deep copy and add element
				component.ValueToEntities[value] = lastFoundSparsSet.DeepCopy()

			}

			// in both cases, add element
			component.ValueToEntities[value].Add(entity)

		} else {

			// take care of special case if it's an integer.
			var encodingVal any = value

			switch typedVal := value.(type) {
			case int:
				encodingVal = typedVal
			case int64:
				encodingVal = int(typedVal)
			}

			if component.ValueToEntities[encodingVal] == nil {
				component.ValueToEntities[encodingVal] = NewSparseSet()
			}

			component.ValueToEntities[encodingVal].Add(entity)

		}
	}

	ecsUpdate := ECSData{Entity: entity, Component: componentName, Value: value}

	if w.ShouldRecordEcsUpdates {
		w.AddEcsChange(ecsUpdate)
	}

	component.ComponentMutex.Unlock()

	return ecsUpdate
}

func (w *World) GetAllEntityData() []ECSData {

	ecsData := []ECSData{}

	for componentName, component := range w.Components {

		for _, entity := range component.Entities.GetAll() {
			entityValue := w.GetComponentValue(componentName, entity)

			ecsData = append(ecsData, ECSData{
				Component: componentName,
				Entity:    entity,
				Value:     entityValue,
			})
		}

	}

	return ecsData
}

func (w *World) FindLastSparseSet(componentName string, value any) *SparseSet {

	// TODO: ideally add error handling
	// we here assume that component exist
	component, _ := w.GetComponent(componentName)

	if component.ValueToEntities[value] != nil {
		return component.ValueToEntities[value]
	}

	if w.Parent == nil {
		return nil
	}

	parentComponent, _ := w.Parent.GetComponent(componentName)

	if parentComponent.ValueToEntities[value] != nil {
		return parentComponent.ValueToEntities[value]
	}

	return nil
}

// TODO: add proper error handling here
// Get the value associated with such entity. Returns a generic type that should be casted
func (w *World) GetComponentValue(componentName string, entity int) interface{} {

	if w.SnapshotMode {
		comp, ok := w.Components[componentName]

		if !ok {
			if w.Parent != nil {
				_, parentOk := w.Parent.Components[componentName]

				if !parentOk {
					return nil
				}
			}
		}

		// component exists
		// if it's nil, find parent
		// if the parent is nil, return  nil

		if comp.EntitiesToValue[entity] == "" {

			if w.Parent == nil {
				return nil
			}

			parentComp, ok := w.Parent.Components[componentName]

			if !ok {
				return nil
			}

			// We just find one layer
			return DecodeBytesBasedOnDataType(parentComp.DataType, parentComp.EntitiesToValue[entity])
		} else {
			// value found!
			return DecodeBytesBasedOnDataType(comp.DataType, comp.EntitiesToValue[entity])
		}

	} else {

		world := w
		if w.ShellMode {
			world = w.Parent
		}

		// normal mode
		comp, ok := world.Components[componentName]
		if !ok {
			return nil
		}

		entityValue, _ := comp.EntitiesToValue[entity]

		// cast to a type and return basically

		return DecodeBytesBasedOnDataType(comp.DataType, entityValue)
	}
}

func (w *World) RemoveEntity(entity int) ECSData {

	world := w
	if w.ShellMode {
		world = w.Parent
	}

	// remove from global entity set
	world.Entities.Remove(entity)

	// loop through all components
	for componentName, component := range world.Components {
		// get entity value
		component.ComponentMutex.Lock()

		value := world.GetComponentValue(componentName, entity)

		component.Entities.Remove(entity)

		// delete entity value
		// TODO: add mutex here
		delete(component.EntitiesToValue, entity)

		// remove from valueToEntities mapping
		if component.ShouldStoreValueToEntities {
			component.ValueToEntities[value].Remove(entity)
		}

		component.ComponentMutex.Unlock()
	}

	ecsUpdate := ECSData{Entity: entity, Component: "-"}

	w.AddEcsChange(ecsUpdate)

	return ecsUpdate
}

// query for entities
func (world *World) Query(query []QueryCondition) []int {

	w := world
	if w.ShellMode {
		w = world.Parent
	}

	if w == nil || len(query) == 0 {
		return []int{}
	}

	// handle cases where query has only one element
	if len(query) == 1 {
		queryCondition := query[0]

		if queryCondition.QueryType == Has {
			component, ok := w.GetComponent(queryCondition.Component)

			if !ok {
				return []int{}
			}

			return component.Entities.GetAll()
		} else if queryCondition.QueryType == HasExact {
			component, ok := w.GetComponent(queryCondition.Component)

			if !ok {
				return []int{}
			}

			// child values
			childValues := component.ValueToEntities[queryCondition.Value]

			// if child values are empty, look at the parent values
			if w.SnapshotMode {
				// if child values are empty, look at the parent values
				if childValues == nil {
					// check for parents value
					parentComponent, _ := w.Parent.GetComponent(queryCondition.Component)
					parentValues := parentComponent.ValueToEntities[queryCondition.Value]

					return parentValues.GetAll()
				}
			}

			return childValues.GetAll()
		}
	}

	var resultEntities []int

	queryCtx := NewQueryContext()

	for i := 0; i < len(query); i++ {
		currentQuery := query[i]

		component, ok := w.Components[currentQuery.Component]

		if ok {
			if currentQuery.QueryType == Has {

				// if it's the first query
				if i == 0 {
					resultEntities = component.Entities.GetAll()

				} else {
					// not the first query, so resultEntities should have something
					resultEntities = ArrayIntersectionWithContext(queryCtx, resultEntities, component.Entities.GetAll())
				}

			} else if currentQuery.QueryType == HasExact {

				// TODO: add back query for snapshot mode

				if i == 0 {
					// if it's the first query
					resultEntities = component.ValueToEntities[currentQuery.Value].GetAll()

				} else {

					// not the first query
					resultEntities = ArrayIntersectionWithContext(queryCtx, resultEntities, component.ValueToEntities[currentQuery.Value].GetAll())

				}

			}
		}
	}

	queryCtx.Terminate()

	return resultEntities
}

// Apply child changed components and values to its parent
func (childWorld *World) ApplyChildToParent() *World {
	if !childWorld.SnapshotMode {
		return &World{}
	}

	parentWorld := *childWorld.Parent

	for componentName, component := range childWorld.Components {
		_, ok := parentWorld.GetComponent(componentName)

		if !ok {
			parentWorld.AddComponent(componentName, component.DataType, true)
		}

		parentComponent := parentWorld.Components[componentName]
		childComponent := childWorld.Components[componentName]

		parentComponent.Entities = childComponent.Entities.DeepCopy()
		parentComponent.DataType = childComponent.DataType
		parentComponent.ShouldStoreValueToEntities = childComponent.ShouldStoreValueToEntities

		if parentComponent.EntitiesToValue == nil {
			parentComponent.EntitiesToValue = make(map[int]string)
		}

		for entity := range childComponent.EntitiesToValue {
			parentComponent.EntitiesToValue[entity] = childComponent.EntitiesToValue[entity]
		}

		if childComponent.ShouldStoreValueToEntities {
			if parentComponent.ValueToEntities == nil {
				parentComponent.ValueToEntities = map[any]*SparseSet{}
			}

			for value := range childComponent.ValueToEntities {
				parentComponent.ValueToEntities[value] = childComponent.ValueToEntities[value].DeepCopy()
			}
		}
	}

	return &parentWorld
}

// "ecs update" represents 1 single ecs state update
type ECSUpdateArray []ECSData

type ECSData struct {
	Entity    int         `json:"entity"`
	Component string      `json:"component"`
	Value     interface{} `json:"value"`
}

func (arr *ECSUpdateArray) AddUpdate(entity int, component string, value interface{}) {
	*arr = append(*arr, ECSData{Entity: entity, Component: component, Value: value})
}

func (c *Component) DeepCopy() Component {

	newComponent := Component{}
	newComponent.EntitiesToValue = make(map[int]string)
	newComponent.ValueToEntities = make(map[interface{}]*SparseSet)

	newComponent.DataType = c.DataType
	newComponent.Entities = c.Entities.DeepCopy()

	for entity, value := range c.EntitiesToValue {
		newComponent.EntitiesToValue[entity] = value
	}

	for value, entitySet := range c.ValueToEntities {
		newComponent.ValueToEntities[value] = entitySet.DeepCopy()
	}

	return newComponent
}

type NewComponentInterface interface {
	// name of component
	Name() string

	// should the component store in a reverse mapping for value based reverse lookup
	ShouldReverseLookup() bool

	// type of component
	Type() reflect.Type
}

type ComponentBaseNew[T any] struct {
	ComponentName              string
	ShouldStoreValueToEntities bool
}

func NewComponentBase[T any](s T) *ComponentBaseNew[T] {
	componentType := &ComponentBaseNew[T]{}
	return componentType
}

func (c *ComponentBaseNew[T]) Get(w *World, entity int) (comp T, err error) {
	world := w
	if w.ShellMode {
		world = w.Parent
	}

	component, ok := world.Components[c.ComponentName]
	if !ok {
		return comp, nil
	}

	entityValue := component.EntitiesToValue[entity]
	if err != nil {
		return comp, err
	}

	res := DecodeBytesBasedOnDataType(component.DataType, entityValue)

	val, ok := res.(T)
	if !ok {
		return comp, errors.New("Component type insertion failed")
	}

	return val, nil
}

func (c *ComponentBaseNew[T]) Set(w *World, entity int, value T) {
	w.SetComponentValue(entity, c.ComponentName, value)
}

func (c *ComponentBaseNew[T]) Name() string {
	return c.ComponentName
}

func (c *ComponentBaseNew[T]) Type() reflect.Type {
	var _type T
	return reflect.TypeOf(_type)
}

func (c *ComponentBaseNew[T]) ShouldReverseLookup() bool {
	return c.ShouldStoreValueToEntities
}

func NewComponent[T any](name string, shouldStoreValueToEntities bool) *ComponentBaseNew[T] {
	var _type T
	comp := NewComponentBase(_type)
	comp.ComponentName = name
	comp.ShouldStoreValueToEntities = shouldStoreValueToEntities
	return comp
}
