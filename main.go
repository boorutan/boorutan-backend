package main

import (
	"applemango/boorutan/backend/booru"
	"applemango/boorutan/backend/booru/danbooru"
	"applemango/boorutan/backend/utils/image"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//var b = moebooru.CreateMoeBooru("https://lolibooru.moe")
	//var b = moebooru.CreateMoeBooru("https://konachan.com")
	var b = danbooru.CreateDanBooru("https://danbooru.donmai.us/")

	app := gin.Default()
	{
		app.Use(cors.New(cors.Config{
			AllowOrigins: []string{
				"http://127.0.0.1:3001",
			},
			AllowMethods: []string{
				"POST",
				"GET",
				"OPTIONS",
			},
			AllowCredentials: true,
			AllowHeaders: []string{
				"Access-Control-Allow-Credentials",
				"Access-Control-Allow-Headers",
				"Content-Type",
				"Content-Length",
				"Accept-Encoding",
				"Authorization",
			},
			MaxAge: 0,
		}))
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
		app.GET("/tags", func(c *gin.Context) {
			tag, err := b.GetTag(booru.GetTagOption{
				Cache: true,
			})
			if err != nil {
				println(err.Error())
				c.JSON(http.StatusInternalServerError, err)
			}
			c.JSON(http.StatusOK, tag)
		})
		app.GET("/post", func(c *gin.Context) {
			pageStr, in := c.GetQuery("page")
			var page int
			if !in {
				page = 1
			} else {
				p, err := strconv.Atoi(pageStr)
				if err != nil {
					println(err.Error())
					c.JSON(http.StatusInternalServerError, err)
				}
				page = p
			}
			println(page)
			post, err := b.GetPost(booru.GetPostOption{
				Cache: true,
				Page:  page,
			})
			if err != nil {
				println(err.Error())
				c.JSON(http.StatusInternalServerError, err)
			}
			c.JSON(http.StatusOK, post)
		})
	}
	app.Run(":8080")
}
