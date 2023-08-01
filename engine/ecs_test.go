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
)

func SpawnTestWorld() *World {

	w := NewGameWorld()
	registerMockComponents(w)

	// 500 movable entities
	entityCount := 500

	// spawn many movable entities
	for i := 0; i < entityCount; i++ {
		entityId := w.AddEntity()

		PositionComp.Set(w, entityId, Pos{X: rand.Intn(100), Y: rand.Intn(100)})
		TargetPositionComp.Set(w, entityId, Pos{X: rand.Intn(100), Y: rand.Intn(100)})
		Tag.Set(w, entityId, "Person")

	}

	// spawn many cities
	// cities generate health
	// cities will have production

	cityCount := 2_000

	for i := 0; i < cityCount; i++ {
		entity := w.AddEntity()

		PositionComp.Set(w, entity, Pos{X: rand.Intn(100), Y: rand.Intn(100)})
		HealthComp.Set(w, entity, 0)
		MaxHealthComp.Set(w, entity, 100)
		Tag.Set(w, entity, "City")
		Owner.Set(w, entity, 0)

		// add productions
		productionEntity := w.AddEntity()

		Tag.Set(w, productionEntity, "Production")
		ProductionOwner.Set(w, productionEntity, entity)
		StartTime.Set(w, productionEntity, 0)

	}

	fmt.Println("Movable entities: ", entityCount)

	return w
}

func TestQueryPerformance(t *testing.T) {

	world := SpawnTestWorld()

	// get all entities that can move
	query := []QueryCondition{
		{QueryType: Has, Component: PositionComp.ComponentName, Value: 0},
		{QueryType: Has, Component: TargetPositionComp.ComponentName, Value: 0},
		{QueryType: HasExact, Component: Tag.ComponentName, Value: "Person"}}

	a := time.Now()
	movableEntities := world.Query(query)
	fmt.Println("Query time: ", time.Since(a).Microseconds(), "ps")
	fmt.Println("Result length: ", len(movableEntities))
}

// TODO: [WIP]
func TestApplyChildToParent(t *testing.T) {

}

func TestConcurrentSetComponentValue(t *testing.T) {

	world := NewGameWorld()
	registerMockComponents(world)

	entity := world.AddEntity()

	// spawn 1000 entities
	for i := 0; i < 1000; i++ {
		go HealthComp.Set(world, entity, 1)
		go PositionComp.Set(world, entity, Pos{X: 1, Y: 1})
	}
}

func TestRemoveEntityTest(t *testing.T) {

	world := NewGameWorld()
	registerMockComponents(world)

	entity := world.AddEntity()
	HealthComp.Set(world, entity, 100)
	PositionComp.Set(world, entity, Pos{X: 1, Y: 1})

	world.RemoveEntity(entity)

	allEntityData := world.GetAllEntityData()

	if len(allEntityData) != 0 {
		t.Errorf("wrong result got %v, want %v", len(allEntityData), 0)
	}
}

func registerMockComponents(w *World) {

	w.AddComponentNew(TargetPositionComp)
	w.AddComponentNew(PositionComp)
	w.AddComponentNew(Tag)
	w.AddComponentNew(Terrain)
	w.AddComponentNew(HealthComp)
	w.AddComponentNew(MaxHealthComp)
	w.AddComponentNew(ProductionOwner)
	w.AddComponentNew(StartTime)
	w.AddComponentNew(Owner)
	w.AddComponentNew(AddressComp)
}
