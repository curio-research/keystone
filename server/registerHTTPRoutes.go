package server

import (
	"github.com/gin-gonic/gin"
)

func RegisterSQLRoutes(ctx *EngineCtx, g *gin.Engine) {
	g.POST("/rewindState", HandleRewindState(ctx))
}
