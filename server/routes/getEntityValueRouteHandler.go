package routes

import (
	"net/http"

	"github.com/curio-research/keystone/server"
	"github.com/gin-gonic/gin"
)

// request
type GetEntityRequest struct {
	Entity int `json:"entity"`
}

// response
type GetEntityResponse struct {
	Table string `json:"table"`
	Value any    `json:"value"`
}

// calculate state root hash for a game world
func GetEntityValueRouteHandler(ctx *server.EngineCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := server.DecodeRequestBody[GetEntityRequest](c)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid request body")
		}

		for tableName, table := range ctx.World.Tables {
			if table.Entities.Contains(req.Entity) {
				value, _ := table.Get(req.Entity)

				c.JSON(http.StatusOK, GetEntityResponse{
					Table: tableName,
					Value: value,
				})
				return
			}
		}

		c.JSON(http.StatusOK, GetEntityResponse{
			Table: "",
			Value: nil,
		})

		return
	}
}
