package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RewindStateRequest struct {
	ElapsedSeconds int    `json:"elapsedSeconds"`
	GameId         string `json:"gameId"`
}

// TODO will only one person have the power to restore the state
func HandleRewindState(ctx *EngineCtx, initWorld func(ctx *EngineCtx)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if ctx.IsRestoringState {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "restoring state already in progress",
			})
			return
		}

		req, err := DecodeRequestBody[RewindStateRequest](c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid request body",
			})
			return
		}

		// TODO can we assume they already did this?
		ctx.IsRestoringState = true

		resetWorldAndTick(ctx)
		initWorld(ctx)

		futureTick := CalcFutureTickFromS(ctx, req.ElapsedSeconds)
		err = ctx.SaveTransactionsHandler.RestoreStateFromTxs(ctx, futureTick, req.GameId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.IsRestoringState = false
	}
}

func resetWorldAndTick(ctx *EngineCtx) {
	w := ctx.World
	for _, table := range w.Tables {
		entities := table.Entities.GetAll()
		for _, entity := range entities {
			table.RemoveEntity(w, entity)
		}
	}

	ctx.GameTick.TickNumber = 1
}
