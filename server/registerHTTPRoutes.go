package server

import (
	"github.com/gin-gonic/gin"
)

func RegisterHTTPSQLRoutes(ctx *EngineCtx, g *gin.Engine) {
	g.POST("/rewindState", HandleRewindState(ctx))
}
