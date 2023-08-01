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
	"strconv"
	"strings"
)

// TODO: add error handling
func DecodeBytesBasedOnDataType(dataType DataType, data string) any {
	switch dataType {
	case Number:
		return DecodeStringToInt(data)
	case String:
		return DecodeStringFromString(data)
	case Position:
		return DecodePositionFromString(data)
	case Address:
		return DecodeStringFromString(data)
	default:
		return nil
	}
}

func EncodeToStringBasedOnDataType(dataType DataType, data any) string {
	switch dataType {
	case Number:
		switch val := data.(type) {
		case int:
			return EncodeIntAsString(val)

		case int64:
			return EncodeIntAsString(int(val))
		}

		return ""
	case String:
		return EncodeStringToString(data.(string))
	case Position:
		return EncodePositionAsString(data.(Pos))
	case Address:
		return EncodeStringToString(data.(string))
	default:
		return ""
	}
}

func EncodeIntAsString(num int) string {
	return strconv.Itoa(num)
}

func DecodeStringToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func EncodeStringToString(str string) string {
	return str
}

func DecodeStringFromString(str string) string {
	return str
}

func EncodePositionAsString(pos Pos) string {
	x := strconv.Itoa(pos.X)
	y := strconv.Itoa(pos.Y)
	return x + "," + y
}

// TODO: add proper error handling
func DecodePositionFromString(str string) Pos {

	parts := strings.Split(str, ",")
	if len(parts) != 2 {
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return Pos{}
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return Pos{}
	}

	return Pos{X: x, Y: y}
}
