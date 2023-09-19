package main

import (
	"applemango/boorutan/backend/booru"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	var booru = booru.Booru{
		Url: booru.BooruUrl{
			Base: "https://konachan.com",
			Post: "/post.json",
		},
		BooruType: booru.MoeBooru,
	}

	app := gin.Default()
	{
		app.Use(cors.Default())
		app.GET("/ping", func(c *gin.Context) {
			c.JSON(200, "pong")
		})
	}
	{
		app.GET("/post", func(c *gin.Context) {
			post, err := booru.GetPost()
			if err != nil {
				println(err)
				c.JSON(http.StatusInternalServerError, err)
			}
			c.JSON(200, post)
		})
	}
	app.Run(":8080")
}
