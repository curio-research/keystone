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
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/engine"
)

type getComponentValuePrecompile struct{}

func (c *getComponentValuePrecompile) RequiredGas(input []byte) uint64 {
	return uint64(100)
}

func (c *getComponentValuePrecompile) Run(evm *EVM, input []byte) ([]byte, error) {

	getComponentValueRequest, err := abi.DecodeGetComponentValue(input)
	if err != nil {
		return nil, err
	}

	gameState := evm.StateDB.GetLatestWorldState()
	component, ok := gameState.GetComponent(getComponentValueRequest.Component)

	if !ok {
		return nil, errors.New("Component not found")
	}

	val := gameState.GetComponentValue(getComponentValueRequest.Component, getComponentValueRequest.Entity.Int64())

	fmt.Println(val)

	// encode value based on type
	valInBytes, err := engine.EncodeAbiBytesBasedOnDataType(component.DataType, val)
	if err != nil {
		return nil, err
	}

	return valInBytes, nil
}
