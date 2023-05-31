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

// ECS based world struct
type World struct {
	SnapshotMode bool
	Parent       *World

	Entities        *SparseSet
	EntitiesNonce   int64
	Components      map[string]Component
	ComponentsNonce int64
}

// ECS component
type Component struct {
	DataType                   DataType
	shouldStoreValueToEntities bool

	Entities        *SparseSet
	EntitiesToValue map[int64][]byte
	ValueToEntities map[any]*SparseSet // mapping: value (bytes) -> Entity set
}

type DataType int64

const (
	Number   DataType = 0
	String   DataType = 1
	Position DataType = 2
	Address  DataType = 3
	Bool     DataType = 4
)

type Pos struct {
	X int64 `json:"X"`
	Y int64 `json:"Y"`
}

type EcsValueUnit struct {
	Component string
	Value     []byte
	Entity    int64
}

type EntityStep struct {
	EntityId int64
	Pos      Pos
}

// query types
var (
	Has              int64 = 0
	Not              int64 = 1
	HasExact         int64 = 2
	HasExactMultiple int64 = 3
	HasExactNot      int64 = 4
)

type QueryCondition struct {
	QueryType int64
	Component string
	Value     any
}

// add entity
func (w *World) AddEntity() int64 {
	w.EntitiesNonce++

	w.Entities.Add(w.EntitiesNonce)

	return w.EntitiesNonce
}

func InitializeNewWorld() *World {
	w := &World{}

	w.Entities = NewSparseSet()
	w.Components = make(map[string]Component)

	return w
}

// Snapshottable world
func InitializeSnapshottableWorld(parent *World) *World {
	w := InitializeNewWorld()
	w.SnapshotMode = true
	w.Parent = parent

	// if there's a parent, we first copy all the components and entities from their parents
	// the assumption is that components will mostly be the same
	if parent != nil {

		for componentName, component := range parent.Components {
			w.AddComponent(componentName, component.DataType, component.shouldStoreValueToEntities)

			// copy entities over
			parentComponent, ok := parent.Components[componentName]
			if ok {
				component.Entities = parentComponent.Entities.DeepCopy()
			}
		}
	} else {
		w.RegisterComponents(mockComponentList)
	}

	return w
}

type ComponentRegistration struct {
	Name                       string
	Type                       DataType
	ShouldStoreValueToEntities bool
}

func (w *World) RegisterComponents(componentRegistrations []ComponentRegistration) {
	for _, componentRegistration := range componentRegistrations {
		w.AddComponent(componentRegistration.Name, componentRegistration.Type, componentRegistration.ShouldStoreValueToEntities)
	}
}

// add component
func (w *World) AddComponent(name string, dataType DataType, shouldStoreValueToEntities bool) (int64, bool) {
	w.ComponentsNonce++

	// check if component already exists
	_, ok := w.Components[name]
	if ok {
		return 0, false
	}

	w.Components[name] = Component{DataType: dataType, EntitiesToValue: make(map[int64][]byte), Entities: NewSparseSet(), ValueToEntities: make(map[interface{}]*SparseSet), shouldStoreValueToEntities: shouldStoreValueToEntities}

	return w.ComponentsNonce, true
}

func (w *World) GetComponent(name string) (Component, bool) {
	component, ok := w.Components[name]
	return component, ok
}

// Get componnet from self or its parent
func (w *World) GetComponentFull(name string) (Component, bool) {
	component, ok := w.Components[name]

	if ok {
		return component, ok
	}

	if w.Parent == nil {
		return component, false
	}

	parentComponents, parentOk := w.Parent.Components[name]

	return parentComponents, parentOk
}

func (w *World) SetComponentValue(entity int64, componentName string, value any) ECSUpdate {

	component, ok := w.GetComponent(componentName)

	// this component isn't present
	if !ok {
		return ECSUpdate{}
	}

	w.Entities.Add(entity)

	component.Entities.Add(entity)

	switch dataType := component.DataType; dataType {
	case Number:

		// TODO: add proper error handling
		// TODO: only works if it's casted to int64
		val, ok := value.(int64)

		if ok {
			w.Components[componentName].EntitiesToValue[entity] = EncodeToBytesBasedOnDataType(dataType, val)
		}

	case String:
		val, ok := value.(string)

		if ok {
			w.Components[componentName].EntitiesToValue[entity] = EncodeToBytesBasedOnDataType(dataType, val)
		}

	case Position:
		val, ok := value.(Pos)

		if ok {
			w.Components[componentName].EntitiesToValue[entity] = EncodeToBytesBasedOnDataType(dataType, val)
		}

	case Address:
		val, ok := value.(string)

		if ok {
			w.Components[componentName].EntitiesToValue[entity] = EncodeToBytesBasedOnDataType(dataType, val)
		}

	default:

	}

	// store in reverse mapping
	if component.shouldStoreValueToEntities {

		if w.SnapshotMode {
			// recursively find the last sparseSet
			lastFoundSparsSet := w.FindLastSparseSet(componentName, value)

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
			// normal mode
			if component.ValueToEntities[value] == nil {
				component.ValueToEntities[value] = NewSparseSet()
			}

			component.ValueToEntities[value].Add(entity)

		}
	}

	return ECSUpdate{Entity: entity, Component: componentName, Value: value}
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

// TODO: add proper typing here
// Get the value associated with such entity. Returns a generic type that should be casted
func (w *World) GetComponentValue(componentName string, entity int64) interface{} {

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

		if comp.EntitiesToValue[entity] == nil {

			if w.Parent == nil {
				return nil
			}

			parentComp, parentCompOk := w.Parent.Components[componentName]

			if !parentCompOk {
				return nil
			}

			// We just find one layer
			return DecodeBytesBasedOnDataType(parentComp.DataType, parentComp.EntitiesToValue[entity])
		} else {
			// value found!
			return DecodeBytesBasedOnDataType(comp.DataType, comp.EntitiesToValue[entity])
		}

	} else {
		// normal mode
		comp, ok := w.Components[componentName]

		if !ok {
			return nil
		}

		return DecodeBytesBasedOnDataType(comp.DataType, comp.EntitiesToValue[entity])
	}

}

// query
// TODO: in a query, should we return a empty set or nil ?
func (w *World) Query(query []QueryCondition) *SparseSet {

	if w == nil {
		return NewSparseSet()
	}

	// handle cases where query has only one element
	if len(query) == 1 {
		queryCondition := query[0]
		if queryCondition.QueryType == Has {
			component, ok := w.GetComponent(queryCondition.Component)

			if !ok {
				return NewSparseSet()
			}

			return component.Entities
		} else if queryCondition.QueryType == HasExact {
			component, ok := w.GetComponent(queryCondition.Component)

			if !ok {
				return NewSparseSet()
			}

			return component.ValueToEntities[queryCondition.Value]
		}
	}

	// TODO: deep copied only once per operation
	// can we eliminate this?
	allEntities := w.Entities.DeepCopy()

	allWorldEntitySet := w.Entities

	for i := 0; i < len(query); i++ {
		currentQuery := query[i]

		// get the component
		component, ok := w.Components[currentQuery.Component]
		if ok {

			if currentQuery.QueryType == Has {

				allEntities = SetIntersection(allEntities, component.Entities)

			} else if currentQuery.QueryType == Not {
				// Does not have component

				allEntities = SetIntersection(allEntities, SetDifference(allWorldEntitySet, component.Entities))

			} else if currentQuery.QueryType == HasExact {

				allEntities = SetIntersection(allEntities, component.ValueToEntities[currentQuery.Value])
			} else if currentQuery.QueryType == HasExactMultiple {
				// Has multiple

				// TODO: implement this
			} else if currentQuery.QueryType == HasExactNot {
				// has exactly not this

				allEntities = SetIntersection(allEntities, component.GetNot(currentQuery.Value))
			}
		}
	}

	return allEntities
}

func (w *World) QueryAsArray(query []QueryCondition) []int64 {
	return w.Query(query).elements
}

func (c *Component) GetNot(val interface{}) *SparseSet {
	return SetDifference(c.Entities, c.ValueToEntities[val])
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
		parentComponent.shouldStoreValueToEntities = childComponent.shouldStoreValueToEntities

		if parentComponent.EntitiesToValue == nil {
			parentComponent.EntitiesToValue = make(map[int64][]byte)
		}

		for entity := range childComponent.EntitiesToValue {
			parentComponent.EntitiesToValue[entity] = childComponent.EntitiesToValue[entity]
		}

		if childComponent.shouldStoreValueToEntities {
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

// ------------------------------------------------
// "ECS Update" represents 1 single ecs state update

type ECSUpdateArray []ECSUpdate

type ECSUpdate struct {
	Entity    int64
	Component string
	Value     interface{}
}

func InitializeEcsUpdateArray() ECSUpdateArray {
	return make([]ECSUpdate, 0)
}

func (arr *ECSUpdateArray) AddUpdate(entity int64, component string, value interface{}) {
	*arr = append(*arr, ECSUpdate{Entity: entity, Component: component, Value: value})
}

func (c *Component) DeepCopy() Component {

	newComponent := Component{}
	newComponent.EntitiesToValue = make(map[int64][]byte)
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

func (world *World) DeepCopy() World {
	newWorld := World{}

	newWorld.Parent = world.Parent
	newWorld.SnapshotMode = world.SnapshotMode

	newWorld.Entities = world.Entities.DeepCopy()
	newWorld.EntitiesNonce = world.EntitiesNonce

	newWorld.Components = make(map[string]Component)

	// TODO: should be a way to deep copy in parallel
	for componentName, component := range world.Components {
		newWorld.Components[componentName] = component.DeepCopy()
	}

	newWorld.ComponentsNonce = world.ComponentsNonce

	return newWorld
}
