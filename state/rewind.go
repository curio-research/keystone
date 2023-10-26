package state

import (
	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RewindStateRequest struct {
	ElapsedSeconds int    `json:"elapsedSeconds"`
	GameId         string `json:"gameId"`
}

// TODO will only one person have the power to restore the state
// initialize the world before calling it
func HandleRewindState(ctx *server.EngineCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		if ctx.IsRestoringState {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "restoring state already in progress",
			})
			return
		}

		req, err := server.DecodeRequestBody[RewindStateRequest](c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid request body",
			})
			return
		}

		// TODO can we assume they already did this?
		ctx.IsRestoringState = true
		ctx.IsLive = false

		err = ctx.SaveTransactionsHandler.RestoreStateFromTxs(ctx, utils.CalcFutureTickFromS(ctx, req.ElapsedSeconds), req.GameId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.IsLive = true
		ctx.IsRestoringState = false
	}
}
