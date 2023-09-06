package server

import (
	"encoding/json"
)

// bring your own implementation on how to submit and retrieve data to the DA layer!
type ITickTransactionHandler interface {
	UploadTickTransactions(tickTransactions TickTransactions) error

	// startTick: inclusive
	// endTick: inclusive
	DownloadTickTransactions(gameId string, startTick int, endTick int) (TickTransactions, error)
}

type TickTransactions []TickTransaction

// player requests (aka tick transactions) are objects that need to be made available such that
// anyone can recreate the state
type TickTransaction struct {
	GameId string

	// likely some sort of address
	Sender string

	// likely a cryptographic signature to verify sender
	Signature string

	// name of function (tick function for now) being triggered
	FunctionName string

	Tick int

	// payload data struct
	Payload any
}

// copy transactions and clear
func (ctx *EngineCtx) CopyClearTickTransactions() TickTransactions {
	ctx.TickTransactionLock.Lock()
	defer ctx.TickTransactionLock.Unlock()

	newRequests := TickTransactions{}

	for _, request := range ctx.TickTransactionsQueue {
		newRequests = append(newRequests, request)
	}

	ctx.TickTransactionsQueue = TickTransactions{}

	return newRequests
}

// serialize to bytes
func (req TickTransactions) Serialize() ([][]byte, error) {
	res := [][]byte{}

	for _, request := range req {
		jsonBytes, err := json.Marshal(request)
		if err != nil {
			return nil, err
		}

		res = append(res, jsonBytes)
	}

	return res, nil
}

// deserialize from bytes to array
func DeserializeTickTransactionsToBytes(rawRequests [][]byte) (TickTransactions, error) {
	res := TickTransactions{}

	for _, request := range rawRequests {
		var p = TickTransaction{}
		err := json.Unmarshal(request, &p)

		if err != nil {
			return nil, err
		}

		res = append(res, p)
	}

	return res, nil
}
