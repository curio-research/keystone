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
	"math/rand"
	"sync"

	"github.com/ethereum/go-ethereum/engine"
	pathfinder "github.com/ethereum/go-ethereum/pathfinder"
)

// ------------------------------------
// game logic systems
// ------------------------------------

func TickWorld(world *engine.World, entitiesToMove []int64) engine.ECSUpdateArray {

	ecsUpdates := TickMovement(world, entitiesToMove)

	TickCityHealth(world)

	TickProduction(world)

	return ecsUpdates
}

func TickMovement(world *engine.World, entitiesToMove []int64) engine.ECSUpdateArray {

	// TODO: can we avoid deep copy on here?
	gameMapCopy := pathfinder.DeepCopy2DArr(pathfinder.GameMap)

	// set all starting positions of all entities as obstacles to make sure path finding works
	for i := 0; i < len(entitiesToMove); i++ {
		entity := entitiesToMove[i]
		entityPosition := engine.DecodePositionFromBytes(world.Components[PositionComp].EntitiesToValue[entity])

		gameMapCopy[entityPosition.Y][entityPosition.X] = "X"
	}

	// construct a fixed, sample map that's loaded in from memory
	worldMap := pathfinder.ConstructWorldNew(gameMapCopy)

	// parallel path finding ------------------------------------

	entityFirstSteps := []engine.EntityStep{}
	newTargetPositions := engine.ECSUpdateArray{}
	allOriginStartingPositions := []engine.Pos{}

	wg := sync.WaitGroup{}

	for i := 0; i < len(entitiesToMove); i++ {
		entityId := entitiesToMove[i]

		wg.Add(1)

		go func(entityId int64) {

			startPos := engine.DecodePositionFromBytes(world.Components[PositionComp].EntitiesToValue[entityId])
			targetPos := engine.DecodePositionFromBytes(world.Components[TargetPositionComp].EntitiesToValue[entityId])

			allOriginStartingPositions = append(allOriginStartingPositions, startPos)

			path, _, _ := pathfinder.AstarPathfinder(pathfinder.Pos{X: startPos.X, Y: startPos.Y}, pathfinder.Pos{X: targetPos.X, Y: targetPos.Y}, worldMap)

			// if the path is more than 0 steps, get the first step and apply
			if len(path) > 0 {
				// apply the first path step to result and set state
				firstPathStep := (path[len(path)-1]).(*pathfinder.Tile)

				entityFirstSteps = append(entityFirstSteps, engine.EntityStep{EntityId: entityId, Pos: engine.Pos{X: int64(firstPathStep.X), Y: int64(firstPathStep.Y)}})
			} else if len(path) == 0 {
				// if the path is zero, it has arrived at the destination
				// then, randomly generate a new location on the map

				newTargetPos := engine.Pos{X: int64(rand.Intn(len(pathfinder.GameMap[0]))), Y: int64(rand.Intn(len(pathfinder.GameMap)))}
				newTargetPositions.AddUpdate(entityId, TargetPositionComp, newTargetPos)
			}

			wg.Done()
		}(entityId)
	}
	wg.Wait()

	// // list of all ecs changes that need to be pushed
	ecsChanges := engine.ECSUpdateArray{}

	// get all apples and their locations
	appleQuery := []engine.QueryCondition{}
	appleQuery = append(appleQuery, engine.QueryCondition{QueryType: engine.HasExact, Component: TagComp, Value: "Apple"})
	appleEntities := world.QueryAsArray(appleQuery)

	// executed once per game loop
	// iterate through all the apple entities. get their locations, and cache that location in a mapping with locations as key and the entityID as the value
	appleLocations := make(map[string]int64)
	for _, appleEntity := range appleEntities {
		appleLocation := engine.DecodePositionFromBytes(world.Components[PositionComp].EntitiesToValue[appleEntity])
		appleLocations[SerializePosAsStr(appleLocation)] = appleEntity
	}

	for _, singleStep := range RemoveDuplicatesFromEntityStepArray(entityFirstSteps) {
		// for all the next steps, check if there's an apple. If there is, then eat the apple!
		appleEntityId, exists := appleLocations[SerializePosAsStr(singleStep.Pos)]
		if exists {
			update := world.SetComponentValue(appleEntityId, TagComp, "EatenApple")

			// if apple exists, eat apple and remove it from world
			ecsChanges = append(ecsChanges, update)

		}

		// apply next step always
		update := world.SetComponentValue(singleStep.EntityId, PositionComp, singleStep.Pos)
		ecsChanges = append(ecsChanges, update)
	}

	for _, update := range newTargetPositions {
		world.SetComponentValue(update.Entity, update.Component, update.Value)
		ecsChanges = append(ecsChanges, update)
	}

	return ecsChanges

}

func TickCityHealth(world *engine.World) engine.ECSUpdateArray {

	ecsUpdates := engine.ECSUpdateArray{}

	query := []engine.QueryCondition{}
	query = append(query, engine.QueryCondition{QueryType: engine.Has, Component: TagComp, Value: "City"})
	query = append(query, engine.QueryCondition{QueryType: engine.Has, Component: HealthComp, Value: 0})
	cityEntities := world.QueryAsArray(query)

	for _, cityEntity := range cityEntities {
		health := engine.DecodeInt64AsBytes(world.Components[HealthComp].EntitiesToValue[cityEntity])
		maxHealth := engine.DecodeInt64AsBytes(world.Components[MaxHealthComp].EntitiesToValue[cityEntity])

		if health < maxHealth {
			// regenerate health based on math curve
			update := world.SetComponentValue(cityEntity, HealthComp, health+1)
			ecsUpdates = append(ecsUpdates, update)
		}
	}

	return ecsUpdates
}

func TickProduction(world *engine.World) []engine.ECSUpdate {

	ecsUpdates := engine.ECSUpdateArray{}

	query := []engine.QueryCondition{}
	query = append(query, engine.QueryCondition{QueryType: engine.Has, Component: ProductionOwnerComp, Value: 0})
	productions := world.QueryAsArray(query)

	for _, productionEntity := range productions {
		startTime := engine.DecodeInt64AsBytes(world.Components[StartTimeComp].EntitiesToValue[productionEntity])

		// modify this.
		if startTime > 0 {
			// spawn a troop
			entity := world.AddEntity()

			world.SetComponentValue(entity, TagComp, "Troop")
			world.SetComponentValue(entity, HealthComp, 100)
			world.SetComponentValue(entity, OwnerComp, 0)
		}

		// if there's an end time, do something
	}

	return ecsUpdates

}

// function that removes all duplicate items from entityStep array
func RemoveDuplicatesFromEntityStepArray(entitySteps []engine.EntityStep) []engine.EntityStep {
	// map to keep track of seen values
	seen := make(map[string]bool)
	result := []engine.EntityStep{}

	for _, val := range entitySteps {
		if _, ok := seen[SerializePosAsStr(val.Pos)]; ok {
			// do not add duplicate
		} else {
			seen[SerializePosAsStr(val.Pos)] = true
			result = append(result, val)
		}
	}

	return result
}

// TODO: add this back
func RemovePositionsFromEntitySteps(positions []engine.Pos, steps []engine.EntityStep) []engine.EntityStep {
	// map to keep track of seen values
	seen := make(map[string]bool)
	result := []engine.EntityStep{}

	// serialize positions first
	for _, val := range positions {
		seen[SerializePosAsStr(val)] = true
	}

	for _, val := range steps {
		if _, ok := seen[SerializePosAsStr(val.Pos)]; ok {
			// do not add duplicate
		} else {
			result = append(result, val)
		}
	}

	return result

}
