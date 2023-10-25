package http

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"

	"github.com/curio-research/keystone/server"
	"github.com/gin-gonic/gin"
)

type PauseGameRequest struct {
	Live bool `json:"live"`
}

// Toggles the liveliness of the game world.
// When game is paused, game tick will not run and transactions will be rejected
func AddPauseGameRoute(router *gin.Engine, gameCtx *server.EngineCtx) {

	// Setup any http requests here
	router.POST("/pause", func(ctx *gin.Context) {
		b, _ := io.ReadAll(ctx.Request.Body)

		var t PauseGameRequest
		err := json.Unmarshal(b, &t)
		if err != nil {
			ctx.Writer.Write([]byte(fmt.Sprintf("error unmarshalling request to type of %s: %s", reflect.TypeOf(t).String(), err.Error())))
			return
		}

		// set game state
		gameCtx.IsLive = t.Live
	})
}
