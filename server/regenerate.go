package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/curio-research/go-backend/engine"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RegenerateTroopsRequest struct {
	PlayerId int        `json:"playerId"`
	Position engine.Pos `json:"position"`
}

func RegenerateTroops(ctx *EngineCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := RegenerateTroopsRequest{}
		DecodeRequestBody(c, &req)

		requestUuid := uuid.New().String()

		tickReq := RegenerateTroopsRequest{
			PlayerId: req.PlayerId,
			Position: req.Position,
		}

		jsonBytes, _ := json.Marshal(tickReq)
		jsonString := string(jsonBytes)

		tick := ctx.Ticker.TickNumber + 1

		AddTickJob(ctx.World, tick, RegenerateTickId, jsonString, strconv.Itoa(req.PlayerId))
		ctx.AddTickTransaction(MoveCalculationTickID, tick, jsonString)

		c.JSON(http.StatusOK, CreateBasicResponseObject(requestUuid))
	}
}
