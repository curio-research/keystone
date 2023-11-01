package test

import (
	"encoding/json"
	"fmt"
	"io"
	"testing"

	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/server/routes"
	"github.com/curio-research/keystone/state"
)

func TestGetEntityValueAPI(t *testing.T) {
	ctx, _, s, _, _ := startTestServer(t, server.Dev)

	book1Entity := 10000000
	book1 := Book{
		Title:  "a",
		Author: "b",
		Id:     book1Entity,
	}

	bookTable.AddSpecific(ctx.World, book1Entity, book1)

	res := sendPostRequest(t, s, "entityValue", routes.GetEntityRequest{
		Entity: book1Entity,
	})

	fmt.Println(3)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	type EntityBookResponse struct {
		Table string `json:"table"`
		Value Book   `json:"value"`
	}

	var response EntityBookResponse
	json.Unmarshal(body, &response)

	if response.Table != bookTable.Name() {
		t.Errorf("Expected table to be Book, got %s", response.Table)
	}

	if response.Value != book1 {
		t.Errorf("Expected value to be %v, got %v", book1, response.Value)
	}

	// test fetching an entity that doesn't exist
	emptyRawResponse := sendPostRequest(t, s, "entityValue", routes.GetEntityRequest{
		Entity: 1000000000,
	})

	body, err = io.ReadAll(emptyRawResponse.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var emptyResponse routes.GetEntityResponse
	json.Unmarshal(body, &response)

	if emptyResponse.Table != "" {
		t.Errorf("Expected table to be empty, got %s", response.Table)
	}
	if emptyResponse.Value != nil {
		t.Errorf("Expected value to be nil, got %v", response.Value)
	}

}

func TestCalculateStateRootAPI(t *testing.T) {
	ctx, _, s, _, _ := startTestServer(t, server.Dev)

	book1Entity := 10
	book2Entity := 11
	book1 := Book{
		Title:  "a",
		Author: "b",
	}
	book2 := Book{
		Title: "c",
	}

	bookTable.AddSpecific(ctx.World, book1Entity, book1)
	bookTable.AddSpecific(ctx.World, book2Entity, book2)

	// calculate state root
	root := state.CalculateWorldStateRootHash(ctx.World)

	res := sendPostRequest(t, s, "stateRoot", routes.StateRootRequest{})

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var response routes.StateRootResponse
	json.Unmarshal(body, &response)

	if response.RootHash != root {
		t.Fatalf("Expected root hash to be %s, got %s", root, response.RootHash)
	}
}
