package server

import (
	"github.com/curio-research/keystone/state"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPSQLRoutes(ctx *EngineCtx, g *gin.Engine) {
	g.POST("/rewindState", state.HandleRewindState(ctx))
}
