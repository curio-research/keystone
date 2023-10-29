package routes

import (
	"net/http"

	"github.com/curio-research/keystone/server"
	"github.com/curio-research/keystone/state"
	"github.com/gin-gonic/gin"
)

// request
type StateRootRequest struct {
}

// response
type StateRootResponse struct {
	GameId   string `json:"gameId"`
	RootHash string `json:"rootHash"`
}

// calculate state root hash for a game world
func StateRootRouteHandler(ctx *server.EngineCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := server.DecodeRequestBody[StateRootRequest](c)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid request body")
		}

		stateRoot := state.CalculateWorldStateRootHash(ctx.World)

		response := &StateRootResponse{
			GameId:   ctx.GameId,
			RootHash: stateRoot,
		}

		c.JSON(http.StatusOK, response)
	}
}
