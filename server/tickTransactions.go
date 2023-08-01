package server

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// bring your own implementation on how to submit and retrieve data to the DA layer!
type TickTransactionApi interface {
	UploadTickTransactions(tickTransactions TickTransactions) error

	// startTick: inclusive
	// endTick: inclusive
	DownloadTickTransactions(gameId string, startTick int, endTick int) (TickTransactions, error)
}

type TickTransactions []TickTransaction

// player requests (aka tick transactions) are objects that need to be made available such that
// anyone can recreate the state
type TickTransaction struct {
	GameId string `json:"gameId"`

	// likely some sort of address
	Sender string `json:"sender"`

	// likely a cryptographic signature to verify sender
	Signature string `json:"signature"`

	// name of function (tick function for now) being triggered
	FunctionName string `json:"functionName"`

	Tick int `json:"tick"`

	// payload data struct
	Payload any `json:"payload"`
}

// add a player request to the player world state
func (ctx *EngineCtx) AddTickTransaction(functionName string, tick int, payload any) {
	ctx.TickTransactionLock.Lock()
	defer ctx.TickTransactionLock.Unlock()

	newRequest := TickTransaction{
		GameId:       ctx.GameId,
		FunctionName: functionName,
		Tick:         tick,
		Payload:      payload,
	}

	ctx.TickTransactionsQueue = append(ctx.TickTransactionsQueue, newRequest)
}

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

func PublishTickTransactions(ctx *EngineCtx) gin.HandlerFunc {
	return func(c *gin.Context) {

		transactions := ctx.CopyClearTickTransactions()

		ctx.TickTransactionApi.UploadTickTransactions(transactions)

		c.JSON(http.StatusOK, CreateBasicResponseObjectWithData("", transactions))

	}
}

type GetTickTransactionsRequest struct {
	GameId    string
	StartTick int
	EndTick   int
}

func GetTickTransactions(ctx *EngineCtx) gin.HandlerFunc {
	return func(c *gin.Context) {

		req := GetTickTransactionsRequest{}
		DecodeRequestBody(c, &req)

		transactions, err := ctx.TickTransactionApi.DownloadTickTransactions(req.GameId, req.StartTick, req.EndTick)

		if err != nil {
			c.JSON(http.StatusOK, CreateBasicResponseObjectWithData("", transactions))
		}

		c.JSON(http.StatusOK, CreateBasicResponseObjectWithData("", transactions))
	}
}
