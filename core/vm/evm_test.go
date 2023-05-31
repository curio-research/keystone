package vm

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/params"
)

func TestEVM(t *testing.T) {
	println("⏰ Testing EVM")

	// initialize a new EVM

	contractAddress := common.BytesToAddress([]byte("contract"))

	statedb, _ := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)

	sender := AccountRef(common.BytesToAddress([]byte{0})) // fake and send as origin address

	// TODO: initialize a smart contract here
	simpleContractBytecode := "0x608060405234801561001057600080fd5b506004361061007d5760003560e01c80634df7e3d01161005b5780634df7e3d0146100b4578063c12fe9bd146100d2578063c3da42b8146100dc578063e92be978146100fa5761007d565b80630dbe671f146100825780632baf11a4146100a05780633dd5fd4a146100aa575b600080fd5b61008a610104565b604051610097919061019c565b60405180910390f35b6100a861010a565b005b6100b2610126565b005b6100bc610141565b6040516100c9919061019c565b60405180910390f35b6100da610147565b005b6100e4610162565b6040516100f1919061019c565b60405180910390f35b610102610168565b005b60005481565b60016002600082825461011d91906101e6565b92505081905550565b600160008082825461013891906101e6565b92505081905550565b60015481565b600160008082825461015991906101e6565b92505081905550565b60025481565b600180600082825461017a91906101e6565b92505081905550565b6000819050919050565b61019681610183565b82525050565b60006020820190506101b1600083018461018d565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006101f182610183565b91506101fc83610183565b9250828201905080821115610214576102136101b7565b5b9291505056fea2646970667358221220cb972c4444fcb9e2a48960aa917bbe093ab63fa758bbb75099486fea1323f69764736f6c63430008120033"

	statedb.CreateAccount(contractAddress)
	statedb.CreateAccount(sender.Address())
	// get account

	statedb.SetCode(contractAddress, hexutil.MustDecode(simpleContractBytecode))

	statedb.SetState(contractAddress, common.Hash{}, common.BytesToHash([]byte{1}))
	statedb.Finalise(true) // Push the state into the "original" slot

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

	evm := NewEVM(vmctx, txContent, statedb, &chainConfig, vmConfig)

	// empty input data
	inputData, _ := encodedFunction("hello")

	ret, _, callErr := evm.Call(sender, contractAddress, inputData, 1000000000, new(big.Int))

	if ret != nil {
		fmt.Println("Return value: ")
		fmt.Println(ret)
	}

	if callErr != nil {
		fmt.Println("Error: ")
		fmt.Println(callErr)
	}
}

func TestHooks(t *testing.T) {
	// Setup Statedb
	contractAddress := common.BytesToAddress([]byte("contract"))

	statedb, _ := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)

	sender := AccountRef(common.BytesToAddress([]byte{0})) // fake and send as origin address
	simpleContractBytecode := "0x608060405234801561001057600080fd5b50600436106100885760003560e01c806398439a4a1161005b57806398439a4a146100dd578063c12fe9bd146100e7578063c3da42b8146100f1578063e92be9781461010f57610088565b80630dbe671f1461008d5780632baf11a4146100ab5780633dd5fd4a146100b55780634df7e3d0146100bf575b600080fd5b610095610119565b6040516100a291906101e5565b60405180910390f35b6100b361011f565b005b6100bd61013b565b005b6100c7610156565b6040516100d491906101e5565b60405180910390f35b6100e561015c565b005b6100ef610190565b005b6100f96101ab565b60405161010691906101e5565b60405180910390f35b6101176101b1565b005b60005481565b600160026000828254610132919061022f565b92505081905550565b600160008082825461014d919061022f565b92505081905550565b60015481565b600160008082825461016e919061022f565b925050819055506001806000828254610187919061022f565b92505081905550565b60016000808282546101a2919061022f565b92505081905550565b60025481565b60018060008282546101c3919061022f565b92505081905550565b6000819050919050565b6101df816101cc565b82525050565b60006020820190506101fa60008301846101d6565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061023a826101cc565b9150610245836101cc565b925082820190508082111561025d5761025c610200565b5b9291505056fea2646970667358221220e1e7f2ad1d7debefcfa925a3ea6a1667fb0664bc2cc9c07abaed0f131885c9c264736f6c63430008120033"

	statedb.CreateAccount(contractAddress)
	statedb.CreateAccount(sender.Address())
	// get account

	statedb.SetCode(contractAddress, hexutil.MustDecode(simpleContractBytecode))
	statedb.Finalise(true)

	// Case 1: SetA, SetB, SetC
	case1StateDB := statedb.Copy()
	case1FunctionNames := []string{"SetA", "SetB", "SetC"}
	case1FunctionList := getFunctionList(contractAddress, case1FunctionNames)

	ApplyHooks(case1StateDB, case1FunctionList)

	case1A := case1StateDB.GetState(contractAddress, common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"))
	case1B := case1StateDB.GetState(contractAddress, common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"))
	case1C := case1StateDB.GetState(contractAddress, common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002"))

	expectedCase1A := common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001")
	expectedCase1B := common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001")
	expectedCase1C := common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001")

	if case1A != expectedCase1A {
		t.Errorf("wrong result got %v, want %v", case1A, expectedCase1A)
	}
	if case1B != expectedCase1B {
		t.Errorf("wrong result got %v, want %v", case1B, expectedCase1B)
	}
	if case1C != expectedCase1C {
		t.Errorf("wrong result got %v, want %v", case1C, expectedCase1C)
	}

	// Case 2: SetA, AnotherSetA, SetB, SetC
	case2StateDB := statedb.Copy()
	case2FunctionNames := []string{"SetA", "AnotherSetA", "SetB", "SetC"}
	case2FunctionList := getFunctionList(contractAddress, case2FunctionNames)

	ApplyHooks(case2StateDB, case2FunctionList)

	case2A := case2StateDB.GetState(contractAddress, common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"))
	case2B := case2StateDB.GetState(contractAddress, common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"))
	case2C := case2StateDB.GetState(contractAddress, common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002"))

	expectedCase2A := common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002")
	expectedCase2B := common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001")
	expectedCase2C := common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001")

	if case2A != expectedCase2A {
		t.Errorf("wrong result got %v, want %v", case2A, expectedCase2A)
	}
	if case2B != expectedCase2B {
		t.Errorf("wrong result got %v, want %v", case2B, expectedCase2B)
	}
	if case2C != expectedCase2C {
		t.Errorf("wrong result got %v, want %v", case2C, expectedCase2C)
	}

	// Case 3: SetA, AnotherSetA, SetB, SetC, SetAB
	case3StateDB := statedb.Copy()
	case3FunctionNames := []string{"SetA", "AnotherSetA", "SetB", "SetC", "SetAB"}
	case3FunctionList := getFunctionList(contractAddress, case3FunctionNames)

	ApplyHooks(case3StateDB, case3FunctionList)

	case3A := case3StateDB.GetState(contractAddress, common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"))
	case3B := case3StateDB.GetState(contractAddress, common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"))
	case3C := case3StateDB.GetState(contractAddress, common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002"))

	expectedCase3A := common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000003")
	expectedCase3B := common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002")
	expectedCase3C := common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001")

	if case3A != expectedCase3A {
		t.Errorf("wrong result got %v, want %v", case3A, expectedCase3A)
	}
	if case3B != expectedCase3B {
		t.Errorf("wrong result got %v, want %v", case3B, expectedCase3B)
	}
	if case3C != expectedCase3C {
		t.Errorf("wrong result got %v, want %v", case3C, expectedCase3C)
	}
}

// helpers

func getFunctionList(address common.Address, functionNames []string) []state.FunctionItem {
	functionList := []state.FunctionItem{}

	for _, functionName := range functionNames {
		functionItem := state.FunctionItem{}
		functionItem.SetAddress(address)
		functionItem.SetFunctionName(functionName)

		functionList = append(functionList, functionItem)
	}

	return functionList
}

func encodedFunction(functionName string) ([]byte, error) {
	// Create a new instance of the ABI for the contract
	myABI, err := abi.JSON(strings.NewReader(testContractABI))

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

var testContractABI = `
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
