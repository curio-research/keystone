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

// mock list of components
var (
	TargetPositionComp = NewComponent[Pos]("TargetPosition", true)
	PositionComp       = NewComponent[Pos]("Position", true)
	Tag                = NewComponent[string]("Tag", true)
	Terrain            = NewComponent[int]("Terrain", true)
	HealthComp         = NewComponent[int]("HealthComp", true)
	MaxHealthComp      = NewComponent[int]("MaxHealthComp", true)
	ProductionOwner    = NewComponent[int]("ProductionOwner", true)
	StartTime          = NewComponent[int]("StartTime", true)
	Owner              = NewComponent[int]("Owner", true)
	AddressComp        = NewComponent[string]("Address", true)
)
