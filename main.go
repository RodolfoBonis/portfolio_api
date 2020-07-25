package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"portfolio_api/src/database"
	"portfolio_api/src/router"
)

func main() {

	err := database.OpenConnection()

	if err != nil {
		panic(err)
	}

	//database.RunMigration()
	gin.ForceConsoleColor()
	server := gin.Default()

	server.Use(cors.Default())

	router.Attach(server)

	server.Run(":3000")

}
