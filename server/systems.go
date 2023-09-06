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
		var req T
		DecodeRequestBody(c, &req)

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

		tick := ctx.Ticker.TickNumber

		requestUuid := uuid.New().String()

		AddTickJob(ctx.World, tick, reflect.TypeOf(req).String(), reqStr, "", requestUuid)

		// TODO: Re-enable this at some point
		// ctx.AddTickTransaction(reflect.TypeOf(req).String(), tick, reqStr)

		c.JSON(http.StatusOK, CreateBasicResponseObject(requestUuid))
	}
}

// decode a string
func DecodeJobData[T any](ctx *EngineCtx, jobId int) T {
	job := JobTable.Get(ctx.World, jobId)

	var req T
	json.Unmarshal([]byte(job.TickDataString), &req)

	return req
}

type TickJobWithId interface {
	ID() string
}
