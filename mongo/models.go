package mongoHelper

// /////////////////////
//    mongoDB schemas
// /////////////////////

// table schema
type Table struct {
	// ID            primitive.ObjectID  `bson:"_id,omitempty"`
	GameId   string `bson:"gameid"`
	Name     string `bson:"name"`
	Entities []int  `bson:"entities"`

	// TODO: try using string
	EntityToValue map[int]string `bson:"entitytovalue"`
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
