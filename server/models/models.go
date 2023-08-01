package models

type GameWorld struct {
	SnapshotId    string      `json:"snapshotId"`
	Entities      []int       `json:"entities"`
	EntitiesNonce int         `json:"entitiesNonce"`
	Components    []Component `json:"components"`
	Time          int         `json:"time"`
}

type Component struct {
	Name            string `json:"name"`
	DataType        int    `json:"dataType"`
	Entities        []int  `json:"entities"`
	EntitiesToValue any    `json:"entitiesToValue"`
	ValueToEntities any    `json:"valueToEntities"`
}

// this should be similarly structured as an Ethereum transaction
type TickTransactionModel struct {
	GameId       string `json:"gameId"`
	Sender       string `json:"sender"`    // likely some sort of address
	Signature    string `json:"signature"` // likely a cryptographic signature to verify sender
	Tick         int    `json:"tick"`
	FunctionName string `json:"functionName"`
	Payload      any    `json:"payload"` // struct
}

type Player struct {
	Address  string `json:"address"`
	PlayerId int    `json:"playerId"`
}
