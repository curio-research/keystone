package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/curio-research/go-backend/engine"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SubmitMoveMultiple struct {
	TroopIds []int
	TileId   int
}

func SubmitMoveMultipleRequest(ctx *EngineCtx) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := SubmitMoveMultiple{}
		DecodeRequestBody(c, &req)

		id := uuid.New().String()

		tickReq := SubmitMoveMultiple{
			TroopIds: req.TroopIds,
			TileId:   req.TileId,
		}

		jsonBytes, _ := json.Marshal(tickReq)
		jsonString := string(jsonBytes)

		tick := ctx.Ticker.TickNumber + 1
		AddTickJob(ctx.World, tick, MoveCalculationTickID, jsonString, strconv.Itoa(req.TileId))
		ctx.AddTickTransaction(MoveCalculationTickID, tick, jsonString)

		c.JSON(http.StatusOK, CreateBasicResponseObject(id))
	}
}

type MovementTickJob struct {
	TroopId int `json:"troopId"`
	TileId  int `json:"tileId"`
}

func FindValidPositionForMove(world *engine.World, centerPosition engine.Pos, filledPositions []engine.Pos) engine.Pos {

	for _, pos := range SpiralSearchGrid {
		searchPos := engine.Pos{X: centerPosition.X + pos.X, Y: centerPosition.Y + pos.Y}

		if !engine.ContainsPositions(filledPositions, searchPos) {

			// if it's inside the world
			if IsPositionInMap(searchPos.X, searchPos.Y, WorldWidth, WorldHeight) {

				// see if it's filled
				entitiesAtPosition := GetEntitiesAtPosition(world, searchPos)

				// if there's only a tile, return this as a valid position
				if len(entitiesAtPosition) == 1 {
					return searchPos
				}
			}
		}
	}

	return centerPosition
}

// inner to outward spiral search grid used to help spawn objects
var SpiralSearchGrid = []engine.Pos{
	// 3x3
	{X: 0, Y: 0},
	{X: -1, Y: -1},
	{X: -1, Y: 0},
	{X: -1, Y: 1},
	{X: 0, Y: -1},
	{X: 0, Y: 1},
	{X: 1, Y: -1},
	{X: 1, Y: 0},
	{X: 1, Y: 1},

	// 4x4
	{X: -2, Y: -2},
	{X: -2, Y: -1},
	{X: -2, Y: 0},
	{X: -2, Y: 1},
	{X: -2, Y: 2},
	{X: -1, Y: -2},
	{X: -1, Y: 2},
	{X: 0, Y: -2},
	{X: 0, Y: 2},
	{X: 1, Y: -2},
	{X: 1, Y: 2},
	{X: 2, Y: -2},
	{X: 2, Y: -1},
	{X: 2, Y: 0},
	{X: 2, Y: 1},
	{X: 2, Y: 2},

	// 5x5
	{X: -3, Y: -3},
	{X: -3, Y: -2},
	{X: -3, Y: -1},
	{X: -3, Y: 0},
	{X: -3, Y: 1},
	{X: -3, Y: 2},
	{X: -3, Y: 3},
	{X: -2, Y: -3},
	{X: -2, Y: 3},
	{X: -1, Y: -3},
	{X: -1, Y: 3},
	{X: 0, Y: -3},
	{X: 0, Y: 3},
	{X: 1, Y: -3},
	{X: 1, Y: 3},
	{X: 2, Y: -3},
	{X: 2, Y: 3},
	{X: 3, Y: -3},
	{X: 3, Y: -2},
	{X: 3, Y: -1},
	{X: 3, Y: 0},
	{X: 3, Y: 1},
	{X: 3, Y: 2},
	{X: 3, Y: 3},

	// 6x6
	{X: -4, Y: -4},
	{X: -4, Y: -3},
	{X: -4, Y: -2},
	{X: -4, Y: -1},
	{X: -4, Y: 0},
	{X: -4, Y: 1},
	{X: -4, Y: 2},
	{X: -4, Y: 3},
	{X: -4, Y: 4},
	{X: -3, Y: -4},
	{X: -3, Y: 4},
	{X: -2, Y: -4},
	{X: -2, Y: 4},
	{X: -1, Y: -4},
	{X: -1, Y: 4},
	{X: 0, Y: -4},
	{X: 0, Y: 4},
	{X: 1, Y: -4},
	{X: 1, Y: 4},
	{X: 2, Y: -4},
	{X: 2, Y: 4},
	{X: 3, Y: -4},
	{X: 3, Y: 4},
	{X: 4, Y: -4},
	{X: 4, Y: -3},
	{X: 4, Y: -2},
	{X: 4, Y: -1},
	{X: 4, Y: 0},
	{X: 4, Y: 1},
	{X: 4, Y: 2},
	{X: 4, Y: 3},
	{X: 4, Y: 4},
}
