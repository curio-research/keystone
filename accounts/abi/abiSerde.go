package abi

import (
	"fmt"
	"math/big"
	"strings"
)

// ------------------------------
// function specific decoding
// ------------------------------

func DecodeComponentSpecArr(input []byte) ([]struct {
	Name      string   "json:\"name\""
	ValueType *big.Int "json:\"valueType\""
}, error) {
	tempRes, err := DecodeAbiFromBytes(input, registerComponentJson)
	if err != nil {
		return nil, err
	}

	decodedArr := (tempRes["method"].([]struct {
		Name      string   "json:\"name\""
		ValueType *big.Int "json:\"valueType\""
	}))

	return decodedArr, nil
}

// ------------------------------
// decode ECS data array
// ------------------------------

func DecodeEcsArr(input []byte) ([]struct {
	Name   string   "json:\"name\""
	Entity *big.Int "json:\"entity\""
	Value  []uint8  "json:\"value\""
}, error) {
	tempRes, err := DecodeAbiFromBytes(input, ecsDataJson)
	if err != nil {
		return nil, err
	}
	decodedArr := (tempRes["method"].([]struct {
		Name   string   "json:\"name\""
		Entity *big.Int "json:\"entity\""
		Value  []uint8  "json:\"value\""
	}))

	return decodedArr, nil
}

// ------------------------------
// decode ECS query
// ------------------------------

func DecodeEcsQuery(input []byte) ([]struct {
	Component string   "json:\"component\""
	Action    *big.Int "json:\"action\""
	Value     []uint8  "json:\"value\""
}, error) {
	tempRes, err := DecodeAbiFromBytes(input, ecsQueryJson)
	if err != nil {
		return nil, err
	}

	decodedQuery := tempRes["method"].([]struct {
		Component string   "json:\"component\""
		Action    *big.Int "json:\"action\""
		Value     []uint8  "json:\"value\""
	})

	return decodedQuery, nil
}

// ------------------------------
// decode get component value
// ------------------------------

// TODO: this is copy pasta'd from abiFragment instead of using type directly
func DecodeGetComponentValue(input []byte) (struct {
	Component string   "json:\"component\""
	Entity    *big.Int "json:\"entity\""
}, error) {
	tempRes, err := DecodeAbiFromBytes(input, getComponentValueAbiFragment)
	if err != nil {
		return struct {
			Component string   "json:\"component\""
			Entity    *big.Int "json:\"entity\""
		}{}, err
	}

	decodedArr := tempRes["method"].(struct {
		Component string   "json:\"component\""
		Entity    *big.Int "json:\"entity\""
	})

	return decodedArr, nil
}

// ------------------------------
// decode position
// ------------------------------

func AbiDecodePositionFromBytes(input []byte) (struct {
	X *big.Int "json:\"x\""
	Y *big.Int "json:\"y\""
}, error) {
	tempRes, err := DecodeAbiFromBytes(input, positionJsonFragment)
	if err != nil {
		return struct {
			X *big.Int "json:\"x\""
			Y *big.Int "json:\"y\""
		}{}, err
	}

	decodedPos := (tempRes["method"].(struct {
		X *big.Int "json:\"x\""
		Y *big.Int "json:\"y\""
	}))

	return decodedPos, nil
}

// ------------------------------------------------

// generalized method for all decoding
func DecodeAbiFromBytes(input []byte, fragment string) (map[string]interface{}, error) {

	def := fmt.Sprintf(`[{ "name" : "method", "type": "function", "outputs": %s}]`, fragment)
	queryConditionAbi, _ := JSON(strings.NewReader(def))

	tempRes := make(map[string]interface{})

	err := queryConditionAbi.UnpackIntoMap(tempRes, "method", input)
	if err != nil {
		return nil, err
	}

	return tempRes, nil

}

type BigIntPos struct {
	X *big.Int `json:"X"`
	Y *big.Int `json:"Y"`
}
