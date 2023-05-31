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
	"fmt"
	"math/rand"
	"testing"
	"time"

	pathfinder "github.com/ethereum/go-ethereum/pathfinder"
)

func SpawnTestWorld() *World {

	w := InitializeNewWorld()

	w.RegisterComponents(mockComponentList)

	// ----------------------------------------------------------------

	// FIXME: this is for demo only
	// load hard-coded map as ECS state into the chain

	for rowId, row := range pathfinder.GameMap {
		for colId, tile := range row {

			tileEntityId := w.AddEntity()

			w.SetComponentValue(tileEntityId, TagComp, "Tile")

			// FIXME: colID, rowID need to be uniform
			w.SetComponentValue(tileEntityId, PositionComp, Pos{X: int64(colId), Y: int64(rowId)})

			// terrain is a number
			// 0 = passable
			// 1 = not passable
			terrain := 0
			if tile == "X" {
				terrain = 1
			}

			w.SetComponentValue(tileEntityId, TerrainComp, int64(terrain))

			// if empty land, spawn an apple with a 1/20th chance
			if terrain == 0 {

				// there's a 20% chance that we need to spawn an apple
				if rand.Intn(5) == 0 {
					entityId := w.AddEntity()

					w.SetComponentValue(entityId, PositionComp, Pos{X: int64(colId), Y: int64(rowId)})
					w.SetComponentValue(entityId, TagComp, "Apple")
				}
			}
		}
	}

	// 500 movable entities
	entityCount := 500

	// spawn many movable entities
	for i := 0; i < entityCount; i++ {
		entityId := w.AddEntity()

		w.SetComponentValue(entityId, PositionComp, Pos{X: int64(rand.Intn(100)), Y: int64(rand.Intn(100))})
		w.SetComponentValue(entityId, TargetPositionComp, Pos{X: int64(rand.Intn(100)), Y: int64(rand.Intn(100))})
		w.SetComponentValue(entityId, TagComp, "Person")
	}

	// spawn many cities
	// cities generate health
	// cities will have production

	cityCount := 2_000

	for i := 0; i < cityCount; i++ {
		cityEntity := w.AddEntity()

		w.SetComponentValue(cityEntity, PositionComp, Pos{X: int64(rand.Intn(100)), Y: int64(rand.Intn(100))})
		w.SetComponentValue(cityEntity, Health, 0)
		w.SetComponentValue(cityEntity, MaxHealth, 100)
		w.SetComponentValue(cityEntity, TagComp, "City")

		// assign no owner for now
		w.SetComponentValue(cityEntity, Owner, 0)

		// add productions
		productionEntity := w.AddEntity()

		w.SetComponentValue(productionEntity, TagComp, "Production")
		w.SetComponentValue(productionEntity, ProductionOwner, cityEntity)
		w.SetComponentValue(productionEntity, StartTime, productionEntity)
	}

	fmt.Println("Movable entities: ", entityCount)

	return w
}

func TestQueryPerformance(t *testing.T) {

	world := SpawnTestWorld()

	// get all entities that can move
	query := []QueryCondition{
		{QueryType: Has, Component: PositionComp, Value: 0},
		{QueryType: Has, Component: TargetPositionComp, Value: 0},
		{QueryType: HasExact, Component: TagComp, Value: "Person"}}

	a := time.Now()
	movableEntities := world.QueryAsArray(query)
	fmt.Println("Query time: ", time.Since(a).Microseconds(), "ps")
	fmt.Println("Result length: ", len(movableEntities))
}

func TestApplyChildToParent(t *testing.T) {
	parentWorld := InitializeSnapshottableWorld(nil)

	parentWorld.RegisterComponents(mockComponentList)

	// Case 1: Change child and sync diffs to parent

	childWorld := InitializeSnapshottableWorld(parentWorld)

	childWorld.SetComponentValue(1, TargetPositionComp, Pos{X: 100, Y: 100})
	childWorld.ApplyChildToParent()

	case1 := parentWorld.GetComponentValue(TargetPositionComp, 1).(Pos)

	if case1.X != 100 || case1.Y != 100 {
		t.Errorf("wrong result got %v, want %v", case1, Pos{X: 100, Y: 100})
	}

	// Case 2: Test different data types

	expectedAddress := "0xd9145CCE52D386f254917e481eB44e9943F39138"

	// Position
	childWorld.SetComponentValue(1, PositionComp, Pos{X: 200, Y: 200})
	// Number
	childWorld.SetComponentValue(1, Health, int64(1))
	// Address
	childWorld.SetComponentValue(1, AddressComp, expectedAddress)

	childWorld.ApplyChildToParent()

	case2A := parentWorld.GetComponentValue(PositionComp, 1).(Pos)
	case2B := parentWorld.GetComponentValue(Health, 1).(int64)
	case2C := parentWorld.GetComponentValue(AddressComp, 1).(string)

	if case2A.X != 200 || case2A.Y != 200 {
		t.Errorf("wrong result got %v, want %v", case2A, Pos{X: 200, Y: 200})
	}

	if case2B != int64(1) {
		t.Errorf("wrong result got %v, want %v", case2B, int64(1))
	}

	if case2C != expectedAddress {
		t.Errorf("wrong result got %v, want %v", case2C, expectedAddress)
	}

	// Case3: Test parent-child chain

	childWorldA := InitializeSnapshottableWorld(parentWorld)

	childWorldA.SetComponentValue(2, Health, int64(1)) // now child A entity 2 has health 1
	childWorldA.ApplyChildToParent()                   // sync to parent

	childWorldB := InitializeSnapshottableWorld(parentWorld)

	healthB := childWorldB.GetComponentValue(Health, 2).(int64) // child B get entity 2 health
	childWorldB.SetComponentValue(2, Health, healthB+1)         // entity 2 add health

	childWorldB.ApplyChildToParent() // sync to parent

	case3 := parentWorld.GetComponentValue(Health, 2).(int64)

	if case3 != 2 {
		t.Errorf("wrong result got %v, want %v", case3, int64(2))
	}
}

// ------------------------------------------------------

func printTestingBegin(input string) {
	fmt.Println("🔫", input)
}
