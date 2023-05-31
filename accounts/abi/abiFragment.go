package abi

import "math/big"

// ----------------------- register component fragment -----------------------

var registerComponentJson = `[
		{
			"components": [
				{
					"internalType": "string",
					"name": "name",
					"type": "string"
				},
				{
					"internalType": "uint256",
					"name": "valueType",
					"type": "uint256"
				}
			],
			"name": "method",
			"type": "tuple[]"
		}
	]`

type registerComponentDecodeStruct struct {
	Name      string   "json:\"name\""
	ValueType *big.Int "json:\"valueType\""
}

// ----------------------- bulk upload ecs fragment -----------------------

// struct EcsData {
// 	string name;
// 	uint256 entity;
// 	bytes value;
// }

var ecsDataJson = `[
		{
			"components": [
				{
					"internalType": "string",
					"name": "name",
					"type": "string"
				},
				{
					"internalType": "uint256",
					"name": "entity",
					"type": "uint256"
				},
				{
					"internalType": "bytes",
					"name": "value",
					"type": "bytes"
				}
			],
			"name": "method",
			"type": "tuple[]"
		}
	]`

type ecsDataDecodeStruct struct {
	Name   string   "json:\"name\""
	Entity *big.Int "json:\"entity\""
	Value  []uint8  "json:\"value\""
}

// ------------------------ decode position ---------------

var positionJsonFragment = `[
		{
			"components": [
				{
					"internalType": "uint256",
					"name": "x",
					"type": "uint256"
				},
				{
					"internalType": "uint256",
					"name": "y",
					"type": "uint256"
				}
			],
			"name": "method",
			"type": "tuple"
		}
	]`

// ----------------------- query ecs -----------------------

var ecsQueryJson = `[
	{
		"components": [
			{
				"internalType": "string",
				"name": "component",
				"type": "string"
			},
			{
				"internalType": "uint256",
				"name": "action",
				"type": "uint256"
			},
			{
				"internalType": "bytes",
				"name": "value",
				"type": "bytes"
			}
		],
		"name": "method",
		"type": "tuple[]"
	}
]`

type ecsQueryJsonStruct struct {
	Component string   "json:\"component\""
	Action    *big.Int "json:\"action\""
	Value     []uint8  "json:\"value\""
}

// ----------------------- get component value -----------------------

var getComponentValueAbiFragment = `[
		{
			"components": [
				{
					"internalType": "string",
					"name": "component",
					"type": "string"
				},
				{
					"internalType": "uint256",
					"name": "entity",
					"type": "uint256"
				}
			],
			"name": "method",
			"type": "tuple"
		}
	]`

type getComponentValueStruct struct {
	Component string   "json:\"component\""
	Entity    *big.Int "json:\"entity\""
}
