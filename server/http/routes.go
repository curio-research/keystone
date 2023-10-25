package http

import (
	"github.com/curio-research/keystone/server"
	"github.com/gin-gonic/gin"
)

// Adds the default set of recommended routes
func AddDefaultRoutes(router *gin.Engine, gameCtx *server.EngineCtx) {
	AddPauseGameRoute(router, gameCtx)

	// TODO: add restore state from db route

	// TODO: add restore state from transactions route (DA layer)
}
