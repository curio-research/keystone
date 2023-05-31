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

package vm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/engine"
)

type bulkUploadEcs struct{}

func (c *bulkUploadEcs) RequiredGas(input []byte) uint64 {
	return uint64(0)
}

func (c *bulkUploadEcs) Run(evm *EVM, input []byte) ([]byte, error) {

	ecsDataArr, err := abi.DecodeEcsArr(input)
	if err != nil {
		return nil, err
	}

	currentWorld := evm.StateDB.GetLatestWorldState()

	for _, val := range ecsDataArr {
		component, exists := currentWorld.GetComponentFull(val.Name)

		if exists {
			// decode bytes based on value type
			decodedVal, err := engine.DecodeAbiBytesToDataType(component.DataType, val.Value)
			if err != nil {
				return nil, err
			}

			fmt.Println("💿 ECS Data Upload: ", val.Name, decodedVal, val.Entity.Int64())
			currentWorld.SetComponentValue(val.Entity.Int64(), val.Name, decodedVal)
		}
	}

	evm.StateDB.SetWorldState(currentWorld)

	return nil, nil
}
