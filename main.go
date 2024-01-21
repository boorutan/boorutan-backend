package main

import (
	"applemango/boorutan/backend/booru/moebooru"
	h "applemango/boorutan/backend/handler"
	"applemango/boorutan/backend/middleware"
	"github.com/gin-gonic/gin"
)

func __init__() {
	b := moebooru.CreateMoeBooru("https://konachan.net")
	_ = b.GetTags()
	_ = h.ReadTags()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	app.Use(middleware.Cors())
	app.Use(h.OptionMiddleware())
	app.GET("/ping", h.Ping)
	app.OPTIONS("/image", h.OptionMiddleware())
	app.GET("/image", h.DownloadImage)
	app.GET("/like", h.GetLikedPost)
	app.POST("/like/:booru/:id", h.LikePost)
	app.OPTIONS("/category", h.OptionMiddleware())
	app.GET("/category", h.GetTagCategory)
	app.OPTIONS("/tag/suggest", h.OptionMiddleware())
	app.GET("/tag/suggest", h.GetTagSuggest)
	app.OPTIONS("/tag", h.OptionMiddleware())
	app.GET("/tag", h.GetTag)
	app.OPTIONS("/post/:id", h.OptionMiddleware())
	app.GET("/post/:id", h.GetPost)
	app.OPTIONS("/post", h.OptionMiddleware())
	app.GET("/post", h.GetPosts)
	_ = app.Run(":8080")
}
