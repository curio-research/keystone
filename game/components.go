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

import "github.com/ethereum/go-ethereum/engine"

// ------------------------
// game component registry
// ------------------------

// TODO: make them IDs instead of string for faster comparison (?)

var (
	PositionComp        = "A"
	TargetPositionComp  = "B"
	TagComp             = "C"
	TerrainComp         = "D"
	HealthComp          = "E"
	MaxHealthComp       = "F"
	ProductionOwnerComp = "G"
	StartTimeComp       = "H"
	OwnerComp           = "I"
	AddressComp         = "J"
)

var componentList = []engine.ComponentRegistration{
	{
		Name:                       TargetPositionComp,
		Type:                       engine.Position,
		ShouldStoreValueToEntities: true,
	},
	{
		Name:                       TargetPositionComp,
		Type:                       engine.Position,
		ShouldStoreValueToEntities: true,
	},
	{
		Name:                       TagComp,
		Type:                       engine.String,
		ShouldStoreValueToEntities: true,
	},
	{
		Name:                       TerrainComp,
		Type:                       engine.Number,
		ShouldStoreValueToEntities: true,
	},
	{
		Name:                       HealthComp,
		Type:                       engine.Number,
		ShouldStoreValueToEntities: true,
	},
	{
		Name:                       MaxHealthComp,
		Type:                       engine.Number,
		ShouldStoreValueToEntities: true,
	},
	{
		Name:                       ProductionOwnerComp,
		Type:                       engine.Number,
		ShouldStoreValueToEntities: true,
	},
	{
		Name:                       StartTimeComp,
		Type:                       engine.Number,
		ShouldStoreValueToEntities: true,
	},
	{
		Name:                       OwnerComp,
		Type:                       engine.Number,
		ShouldStoreValueToEntities: true,
	},
	{
		Name:                       AddressComp,
		Type:                       engine.Address,
		ShouldStoreValueToEntities: true,
	},
}
