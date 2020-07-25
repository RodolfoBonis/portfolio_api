package router

import (
	"github.com/gin-gonic/gin"
	"portfolio_api/src/handlers"
)

// Attach ...
func Attach(router *gin.Engine) {
	// declaração das rotas

	handlers.PostsHandlers(router)
}
