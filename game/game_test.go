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

package game

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/engine"
	astar "github.com/ethereum/go-ethereum/pathfinder"
)

// ----------------------------------------------------------------
// world tick estimate stats

// 500 entities -> ~60ms with pathfinding
// pathfinding takes up majority of time

func TestEngineTick(t *testing.T) {

	world := SpawnTestWorld()

	// get all entities that can move
	query := []engine.QueryCondition{
		{QueryType: engine.Has, Component: PositionComp, Value: 0},
		{QueryType: engine.Has, Component: TargetPositionComp, Value: 0},
	}
	movableEntities := world.QueryAsArray(query)

	total := 0 // nanoseconds
	iterations := 5

	for i := 0; i < iterations; i++ {
		s := time.Now()

		TickWorld(world, movableEntities)

		total += int(time.Since(s).Nanoseconds())
	}

	fmt.Println("⏰ Average tick time: ", total/iterations, "ns", "(", total/iterations/1_000_000, "ms")
}

func SpawnTestWorld() *engine.World {

	w := engine.InitializeNewWorld()

	w.RegisterComponents(componentList)

	// ----------------------------------------------------------------

	// FIXME: this is for demo only
	// load hard-coded map as ECS state into the chain

	for rowId, row := range astar.GameMap {
		for colId, tile := range row {

			tileEntityId := w.AddEntity()

			w.SetComponentValue(tileEntityId, TagComp, "Tile")

			// FIXME: colID, rowID need to be uniform
			w.SetComponentValue(tileEntityId, PositionComp, engine.Pos{X: int64(colId), Y: int64(rowId)})

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

					w.SetComponentValue(entityId, PositionComp, engine.Pos{X: int64(colId), Y: int64(rowId)})
					w.SetComponentValue(entityId, TagComp, "Apple")
				}
			}
		}
	}

	movableEntityCount := 500

	// spawn many movable entities. This will simulate tick
	for i := 0; i < movableEntityCount; i++ {
		entityId := w.AddEntity()

		w.SetComponentValue(entityId, PositionComp, engine.Pos{X: int64(rand.Intn(100)), Y: int64(rand.Intn(100))})
		w.SetComponentValue(entityId, TargetPositionComp, engine.Pos{X: int64(rand.Intn(100)), Y: int64(rand.Intn(100))})
		w.SetComponentValue(entityId, TagComp, "Person")
	}

	// spawn many cities
	// cities generate health
	// cities have mock productions

	cityCount := 2_000

	for i := 0; i < cityCount; i++ {
		cityEntity := w.AddEntity()

		w.SetComponentValue(cityEntity, PositionComp, engine.Pos{X: int64(rand.Intn(100)), Y: int64(rand.Intn(100))})
		w.SetComponentValue(cityEntity, HealthComp, 0)
		w.SetComponentValue(cityEntity, MaxHealthComp, 100)
		w.SetComponentValue(cityEntity, TagComp, "City")

		w.SetComponentValue(cityEntity, OwnerComp, 0)

		productionEntity := w.AddEntity()

		w.SetComponentValue(productionEntity, TagComp, "Production")
		w.SetComponentValue(productionEntity, ProductionOwnerComp, cityEntity)
		w.SetComponentValue(productionEntity, StartTimeComp, productionEntity)
	}

	fmt.Println("Movable entities: ", movableEntityCount)

	return w
}
