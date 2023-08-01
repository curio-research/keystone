package server

import (
	"encoding/json"
	"testing"
)

type FakePayload struct {
	Id string `json:"id"`
}

func TestTickTransactions(t *testing.T) {

	ctx := &EngineCtx{
		TickTransactionsQueue: TickTransactions{},
	}

	fakePayload := FakePayload{
		Id: "1",
	}

	ctx.AddTickTransaction("move", 1, fakePayload)

	data := ctx.CopyClearTickTransactions()
	res, err := data.Serialize()

	if err != nil {
		t.Error(err)
		return
	}

	// deserialize
	serializedResult, err := DeserializeTickTransactionsToBytes(res)
	if err != nil {
		t.Error(err)
	}

	for _, tickTransaction := range serializedResult {

		// serialize to bytes first
		tickTxAsBytes, _ := json.Marshal(tickTransaction.Payload)

		var fakePayload FakePayload
		err := json.Unmarshal(tickTxAsBytes, &fakePayload)

		if err != nil {
			t.Error(err)
		}

	}
}
