package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostsHandlers(r *gin.Engine) {
	applications := r.Group("/posts")
	{
		applications.GET("/", getAllPosts)
	}
}

func getAllPosts(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{"data": "teste"})
}