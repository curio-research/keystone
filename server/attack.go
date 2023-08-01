package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// first version of attack: don't go to destination when done attacking
type SubmitAttackRequest struct {
	AttackerId int `json:"attackerId"`
	TargetId   int `json:"targetId"`
}

func SubmitAttackAction(ctx *EngineCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := SubmitAttackRequest{}
		DecodeRequestBody(c, &req)

		requestUuid := uuid.New().String()

		tickReq := SubmitAttackRequest{
			AttackerId: req.AttackerId,
			TargetId:   req.TargetId,
		}

		jsonBytes, _ := json.Marshal(tickReq)
		jsonString := string(jsonBytes)

		tick := ctx.Ticker.TickNumber + 1

		AddTickJob(ctx.World, tick, AttackTickID, jsonString, strconv.Itoa(req.TargetId))
		ctx.AddTickTransaction(MoveCalculationTickID, tick, jsonString)

		c.JSON(http.StatusOK, CreateBasicResponseObject(requestUuid))
	}
}

type AttackTickJob struct {
	AttackerId int `json:"attackerId"`
	TargetId   int `json:"targetId"`
}
