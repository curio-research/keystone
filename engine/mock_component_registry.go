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

// ------------------------
// sample component registry
// (used for testing)
// ------------------------

var (
	TargetPositionComp = "TargetPosition"
	PositionComp       = "Position"
	TagComp            = "Tag"
	TerrainComp        = "Terrain"
	Health             = "Health"
	MaxHealth          = "MaxHealth"
	ProductionOwner    = "ProductionOwner"
	StartTime          = "StartTime"
	Owner              = "Owner"
	AddressComp        = "Address"
)

// create a list of component registrations based on the above components
var mockComponentList = []ComponentRegistration{
	{
		Name:                       PositionComp,
		Type:                       Position,
		ShouldStoreValueToEntities: true,
	},
	{
		Name:                       TargetPositionComp,
		Type:                       Position,
		ShouldStoreValueToEntities: true,
	},
	{
		Name:                       TagComp,
		Type:                       String,
		ShouldStoreValueToEntities: true,
	},
	{
		Name:                       TerrainComp,
		Type:                       Number,
		ShouldStoreValueToEntities: true,
	},
	{
		Name:                       Health,
		Type:                       Number,
		ShouldStoreValueToEntities: true,
	},
	{
		Name:                       MaxHealth,
		Type:                       Number,
		ShouldStoreValueToEntities: true,
	},
	{
		Name:                       ProductionOwner,
		Type:                       Number,
		ShouldStoreValueToEntities: true,
	},
	{
		Name:                       StartTime,
		Type:                       Number,
		ShouldStoreValueToEntities: true,
	},
	{
		Name:                       Owner,
		Type:                       Number,
		ShouldStoreValueToEntities: true,
	},
	{
		Name:                       AddressComp,
		Type:                       Address,
		ShouldStoreValueToEntities: true,
	},
}
