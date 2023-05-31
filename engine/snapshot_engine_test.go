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
	"testing"
	"time"
)

func TestSnapshotWorldInitialization(t *testing.T) {

	genesisWorld := InitializeSnapshottableWorld(nil)
	genesisWorld.RegisterComponents(mockComponentList)

	currentWorld := GenerateSnapshottableWorlds(genesisWorld, 5)

	// check for components
	for _, component := range mockComponentList {
		_, ok := currentWorld.Components[component.Name]
		if !ok {
			t.Fail()
		}
	}
}

func TestSnapshotWorldSetECSValue(t *testing.T) {

	genesisWorld := InitializeSnapshottableWorld(nil)
	genesisWorld.RegisterComponents(mockComponentList)

	count := 100

	for i := 0; i < count; i++ {
		genesisWorld.SetComponentValue(int64(i), Health, int64(i))
	}

	currentWorld := GenerateSnapshottableWorlds(genesisWorld, 5)

	// verify that each value has been committed
	for i := 0; i < count; i++ {
		val := currentWorld.GetComponentValue(Health, int64(i))
		if val != int64(i) {
			t.Fail()
		}
	}
}

func TestSnapshotWorldEntitiesToValue(t *testing.T) {

	genesisWorld := InitializeSnapshottableWorld(nil)
	genesisWorld.RegisterComponents(mockComponentList)

	count := 100

	for i := 0; i < count; i++ {
		genesisWorld.SetComponentValue(int64(i), Health, int64(i))
	}

	currentWorld := GenerateSnapshottableWorlds(genesisWorld, 5)

	// add additional to component
	entity := int64(420)
	componentVal := int64(20)

	currentWorld.SetComponentValue(entity, Health, componentVal)
	healthComponent, _ := currentWorld.GetComponent(Health)
	if healthComponent.ValueToEntities[componentVal].Size() != 2 {
		t.Fail()
	}

	// TODO: add more edge cases
}

// test recursively lookup speed
func TestSnapshotWorldRecursiveLookupSpeed(t *testing.T) {
	genesisWorld := InitializeSnapshottableWorld(nil)
	genesisWorld.RegisterComponents(mockComponentList)

	entity := int64(0)

	genesisWorld.SetComponentValue(entity, Health, int64(50))

	// 10-deep recursive look up
	currentWorld := GenerateSnapshottableWorlds(genesisWorld, 10)

	a := time.Now()
	currentWorld.GetComponentValue(Health, entity)
	fmt.Println("10 deep look up:", time.Since(a))

	// 100-deep recursive look up
	currentWorld = GenerateSnapshottableWorlds(genesisWorld, 100)

	a = time.Now()
	currentWorld.GetComponentValue(Health, entity)
	fmt.Println("100 deep look up:", time.Since(a))

	// 1000-deep recursive look up
	currentWorld = GenerateSnapshottableWorlds(genesisWorld, 1_000)

	a = time.Now()
	currentWorld.GetComponentValue(Health, entity)
	fmt.Println("1000 deep look up:", time.Since(a))

}

// ---------------------------------------------------------

// helper function. generates a chain of worlds
func GenerateSnapshottableWorlds(startingWorld *World, worldCount int) *World {
	prevWorldPointer := startingWorld

	for i := 0; i < worldCount; i++ {
		// construct to new world and link to parent
		newWorld := InitializeSnapshottableWorld(prevWorldPointer)

		// set prevWorld reference
		prevWorldPointer = newWorld
	}

	return prevWorldPointer
}
