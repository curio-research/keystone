package server

import (
	"github.com/gin-gonic/gin"
	"io"
)

type RewindState struct {
}

func registerHTTPRoutes(g *gin.Engine) {
	g.POST("/rewindState", func(context *gin.Context) {
		request := context.Request
		writer := context.Writer

		b, err := io.ReadAll(request.Body)
		if err != nil {
			writer.Write([]byte(err.Error()))
			return
		}

	})
}
