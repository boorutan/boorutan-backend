package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	{
		app.Use(cors.Default())
		app.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "pong")
		})
	}
	app.Run(":8080")
}
