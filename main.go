package main

import (
	"applemango/boorutan/backend/booru/moebooru"
	"applemango/boorutan/backend/utils/image"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	var booru = moebooru.CreateMoeBooru("https://konachan.com")
	//var booru = danbooru.CreateDanBooru("https://danbooru.donmai.us/")

	app := gin.Default()
	{
		app.Use(cors.Default())
		app.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong")
		})
	}
	{
		app.GET("/image", func(c *gin.Context) {
			url, in := c.GetQuery("url")
			if !in {
				c.JSON(http.StatusInternalServerError, "err")
			}
			uuid, err := image.GetImage(url)
			if err != nil {
				uuid = "e.png"
			}
			path := fmt.Sprintf("./static/images/%s", uuid)
			c.File(path)
		})
	}
	{
		app.GET("/post", func(c *gin.Context) {
			post, err := booru.GetPost(true)
			if err != nil {
				println(err)
				c.JSON(http.StatusInternalServerError, err)
			}
			c.JSON(http.StatusOK, post)
		})
	}
	app.Run(":8080")
}
