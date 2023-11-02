package server

import (
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// creates a corresponding request from a Gin request to user
func CreateRequestForTick[T any](ctx *EngineCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := DecodeRequestBody[T](c)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid request body")
		}

		// if the state is being restored, return an error
		if ctx.IsRestoringState {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "cannot submit request while restoring state",
			})
			return
		}

		if !ctx.IsLive {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "cannot submit request while game is paused",
			})
			return
		}

		reqStr, _ := SerializeRequestToString(req)

		tick := ctx.GameTick.TickNumber + 1
		AddSystemTransaction(ctx.World, tick, reflect.TypeOf(req).String(), reqStr, "", false)

		requestUuid := uuid.New().String()
		c.JSON(http.StatusOK, CreateBasicResponseObject(requestUuid))
	}
}

// decode a string
func DecodeTxData[T any](ctx *EngineCtx, transactionId int) KeystoneTx[T] {
	transaction := TransactionTable.Get(ctx.World, transactionId)

	var req KeystoneTx[T]
	json.Unmarshal([]byte(transaction.Data), &req)

	return req
}
