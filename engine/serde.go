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
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"math/big"
	"unsafe"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func EncodeAbiBytesBasedOnDataType(dataType DataType, input interface{}) ([]byte, error) {

	switch dataType {
	case Number:
		val, ok := input.(int64)
		if !ok {
			return nil, errors.New("Invalid input type for number")
		}

		res, err := abi.PackUint256(val)
		if err != nil {
			return nil, err
		}

		return res, nil

	case String:
		val, ok := input.(string)
		if !ok {
			return nil, errors.New("Invalid input type for string")
		}

		res, err := abi.PackString(val)
		if err != nil {
			return nil, err
		}

		return res, nil

	case Address:
		val, ok := input.(string)
		if !ok {
			return nil, errors.New("Invalid input type for address")
		}

		res, err := abi.PackAddress(common.HexToAddress(val))
		if err != nil {
			return nil, err
		}

		return res, nil

	case Bool:
		val, ok := input.(bool)
		if !ok {
			return nil, errors.New("Invalid input type for boolean")
		}

		res, err := abi.PackBool(val)
		if err != nil {
			return nil, err
		}

		return res, nil

	case Position:
		val, ok := input.(Pos)
		if !ok {
			return nil, errors.New("Invalid input type for position")
		}

		// first convert to bigInt for ABI conversion
		bigIntifyPosition := abi.BigIntPos{X: big.NewInt(val.X), Y: big.NewInt(val.Y)}

		res, err := abi.PackPosition(bigIntifyPosition)
		if err != nil {
			return nil, err
		}

		return res, nil

	default:
		return nil, nil
	}
}

// TODO: move this to appropriate folder
// used while interfacing with abi encoded objects usually submitted through smart contracts
func DecodeAbiBytesToDataType(dataType DataType, data []byte) (any, error) {

	switch dataType {

	case Number:
		res, err := abi.UnpackUint256(data)
		if err != nil {
			return nil, err
		}

		return res, nil

	case String:
		res, err := abi.UnpackString(data)
		if err != nil {
			return nil, err
		}

		return res, nil

	case Position:
		res, err := abi.AbiDecodePositionFromBytes(data)
		if err != nil {
			return nil, err
		}

		return Pos{X: (res.X).Int64(), Y: (res.Y).Int64()}, nil

	case Address:
		res, err := abi.UnpackAddress(data)
		if err != nil {
			return nil, err
		}

		return res, nil

	case Bool:
		res, err := abi.UnpackBool(data)
		if err != nil {
			return nil, err
		}

		return res, nil

	default:
		return nil, nil
	}
}

// TODO: add error handling
func DecodeBytesBasedOnDataType(dataType DataType, data []byte) any {
	switch dataType {
	case Number:
		return DecodeInt64AsBytes(data)
	case String:
		return DecodeStringFromBytes(data)
	case Position:
		return DecodePositionFromBytes(data)
	case Address:
		return DecodeStringFromBytes(data)
	default:
		return nil
	}
}

func EncodeToBytesBasedOnDataType(dataType DataType, data any) []byte {
	switch dataType {
	case Number:
		return EncodeInt64AsBytes(data.(int64))
	case String:
		return EncodeStringToBytes(data.(string))
	case Position:
		return EncodePositionAsBytes(data.(Pos))
	case Address:
		return EncodeStringToBytes(data.(string))
	default:
		return nil
	}
}

func EncodeInt64AsBytes(num int64) []byte {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutVarint(buf, num)
	b := buf[:n]

	return b
}

func DecodeInt64AsBytes(bytes []byte) int64 {
	num, _ := binary.Varint(bytes)
	return num
}

func EncodeStringToBytes(str string) []byte {
	return []byte(str)
}

func DecodeStringFromBytes(bytes []byte) string {
	return string(bytes)
}

func EncodePositionAsBytes(pos Pos) []byte {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(pos)
	return reqBodyBytes.Bytes() // this is the []byte
}

func DecodePositionFromBytes(bytes []byte) Pos {
	var pos Pos
	json.Unmarshal(bytes, &pos)
	return pos
}

func ByteArrayToInt(arr []byte) int64 {
	val := int64(0)
	size := len(arr)
	for i := 0; i < size; i++ {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}
	return val
}
