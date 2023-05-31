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
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/engine"
)

type queryPrecompile struct{}

func (c *queryPrecompile) RequiredGas(input []byte) uint64 {
	return uint64(100)
}

func (c *queryPrecompile) Run(evm *EVM, input []byte) ([]byte, error) {

	rawQuery, err := abi.DecodeEcsQuery(input)
	if err != nil {
		return nil, err
	}

	// convert raw query to ECS query format
	ecsQuery := []engine.QueryCondition{}

	gameState := evm.StateDB.GetLatestWorldState()

	for _, val := range rawQuery {

		component, ok := gameState.GetComponent(val.Component)

		if !ok {
			return nil, errors.New("Component not found")
		}

		// decode from bytes based on component type
		decodedValue, err := engine.DecodeAbiBytesToDataType(component.DataType, val.Value)
		if err != nil {
			return nil, err
		}

		queryCondition := engine.QueryCondition{
			Component: val.Component,
			QueryType: val.Action.Int64(),
			Value:     decodedValue,
		}

		ecsQuery = append(ecsQuery, queryCondition)
	}

	queryResult := gameState.QueryAsArray(ecsQuery)

	abiEncodedQueryResult, err := abi.AbiEncodeInt64ArrayToBytes(queryResult)
	if err != nil {
		return nil, err
	}

	return abiEncodedQueryResult[:], nil
}
