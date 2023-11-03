package routes

import (
	"net/http"

	"github.com/curio-research/keystone/server"
	"github.com/gin-gonic/gin"
)

// request
type DownloadStateRequest struct {
	Tables []string `json:"tables"`
}

// response
type GameStateResponse struct {
	Tick   int         `json:"tick"`
	Tables []TableData `json:"tables"`
}

// fetch the world state
func GetStateRouteHandler(ctx *server.EngineCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		req, err := server.DecodeRequestBody[DownloadStateRequest](c)
		if err != nil {
			c.JSON(http.StatusBadRequest, "Invalid request body")
		}

		gameStateResponse := &GameStateResponse{
			Tick:   ctx.GameTick.TickNumber,
			Tables: make([]TableData, 0),
		}

		// loop through all world tables
		for tableName, table := range ctx.World.Tables {

			// if table array is empty, return all tables
			if len(req.Tables) == 0 || ContainsString(req.Tables, tableName) {

				tableData := TableData{
					Name:   tableName,
					Values: make([]Value, 0),
				}

				for entity, value := range table.EntityToValue {
					tableData.Values = append(tableData.Values, Value{
						Entity: entity,
						Value:  value,
					})
				}
				gameStateResponse.Tables = append(gameStateResponse.Tables, tableData)

			}
		}

		c.JSON(http.StatusOK, gameStateResponse)
	}
}

type TableData struct {
	Name   string  `json:"name"`
	Values []Value `json:"values"`
}

type Value struct {
	Entity int `json:"entity"`
	Value  any `json:"value"`
}

func ContainsString(arr []string, target string) bool {
	for _, str := range arr {
		if str == target {
			return true
		}
	}
	return false
}
