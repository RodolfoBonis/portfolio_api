package utils

import (
	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if ValidateAcess(c) {
			c.Next()
		}
	}
}
