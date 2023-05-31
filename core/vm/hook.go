package vm

import (
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/params"
)

func ApplyHooks(statedb *state.StateDB, functionList []state.FunctionItem) {
	// Set up EVM
	sender := AccountRef(common.BytesToAddress([]byte{0})) // fake and send as origin address

	vmctx := BlockContext{
		CanTransfer: func(StateDB, common.Address, *big.Int) bool { return true },
		Transfer:    func(StateDB, common.Address, common.Address, *big.Int) {},
	}

	txContent := TxContext{Origin: sender.Address(), GasPrice: new(big.Int)}

	chainConfig := params.ChainConfig{
		IstanbulBlock: big.NewInt(0),
	}

	// use the latest instruction set
	vmConfig := Config{}

	memorydb := state.NewMemoryDB()

	wg := sync.WaitGroup{}

	// Stage 1: Pre process, collect FunctionName and Slot affected relation map
	predb := statedb.Copy()

	for _, function := range functionList {
		addr := function.GetAddress()
		functionName := function.GetFunctionName()

		inputData, _ := hookEncodedFunction(functionName)

		wg.Add(1)

		go func(address common.Address, functionName string) {
			statedbBefore := predb.Copy()
			statedbAfter := predb.Copy()

			preEVM := NewEVM(vmctx, txContent, statedbAfter, &chainConfig, vmConfig)
			preEVM.Call(sender, address, inputData, 1000000000, new(big.Int))

			memorydb.SetSlotStore(address, functionName, statedbBefore, statedbAfter)

			wg.Done()
		}(addr, functionName)
	}

	wg.Wait()

	for _, function := range functionList {
		addr := function.GetAddress()
		functionName := function.GetFunctionName()

		memorydb.UpdateFunctionItem(addr, functionName)
	}

	// Stage 2: Iterate all address, do the real statedb effect with locks
	serialFuntions := memorydb.GetSerialFuntions()
	parallelFuntions := memorydb.GetParallelFuntions()

	// Stage 2.1: Execute serial functions firstly
	vm := NewEVM(vmctx, txContent, statedb, &chainConfig, vmConfig)

	for _, functionItem := range serialFuntions {
		address := functionItem.GetAddress()
		functionName := functionItem.GetFunctionName()

		inputData, _ := hookEncodedFunction(functionName)

		vm.Call(sender, address, inputData, 1000000000, new(big.Int))
	}

	// Stage 2.2: Execute parallel functions
	diffs := []state.StorageDiff{}
	var appnedLock sync.Mutex

	for _, functionItem := range parallelFuntions {
		address := functionItem.GetAddress()
		functionName := functionItem.GetFunctionName()

		inputData, _ := hookEncodedFunction(functionName)

		wg.Add(1)

		go func(address common.Address, functionName string) {
			statedbCopy := statedb.Copy()
			vm := NewEVM(vmctx, txContent, statedbCopy, &chainConfig, vmConfig)
			vm.Call(sender, address, inputData, 1000000000, new(big.Int))

			diff := state.CompareDirtyStorage(statedb, statedbCopy, address)

			appnedLock.Lock()
			diffs = append(diffs, diff...)
			appnedLock.Unlock()

			wg.Done()
		}(address, functionName)
	}

	wg.Wait()

	// apply diffs to statedb
	for _, diff := range diffs {
		statedb.SetState(diff.Address, diff.Key, diff.NewValue)
	}
}

func hookEncodedFunction(functionName string) ([]byte, error) {
	// Create a new instance of the ABI for the contract
	myABI, err := abi.JSON(strings.NewReader(ContractABI))

	if err != nil {
		return nil, err
	}

	// Pack the function arguments into a byte array
	data, err := myABI.Pack(functionName)
	if err != nil {
		return nil, err
	}

	// Combine the function signature with the packed arguments to create the final calldata
	return append(myABI.Methods[functionName].ID, data...), nil
}

// TODO: replace this with the proper treaty hook ABI when necessary. This is a temporary placeholder.
var ContractABI = `
[
	{
		"inputs": [],
		"name": "AnotherSetA",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "SetA",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "SetAB",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "SetB",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "SetC",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "a",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "b",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "c",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
]
`
