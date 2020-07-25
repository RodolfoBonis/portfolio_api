package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"portfolio_api/src/router"
)

func main() {

	gin.ForceConsoleColor()
	server := gin.Default()

	server.Use(cors.Default())

	router.Attach(server)

	server.Run(":3000")

}
