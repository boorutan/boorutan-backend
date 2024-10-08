package main

import (
	"applemango/boorutan/backend/booru/feed"
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
	//__init__()
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()

	app.Use(middleware.Cors())
	app.Use(h.OptionMiddleware())
	feed.RegisterBooruStreamer(app)
	app.Use(middleware.BaseLimeMiddleware())

	account := app.Group("/account")
	{
		account.Use(middleware.LoginMiddleware())
		account.GET("/like", h.GetLikedPost)
		account.POST("/like/:booru/:id", h.LikePost)
		// u, _ := c.Get("account")
	}
	app.GET("/ping", h.Ping)
	app.GET("/image", h.DownloadImage)
	//app.GET("/like", h.GetLikedPost)
	//app.POST("/like/:booru/:id", h.LikePost)
	app.GET("/category", h.GetTagCategory)
	app.GET("/tag/suggest", h.GetTagSuggest)
	app.GET("/tag/suggest/fast", h.GetTagSuggestFast)
	app.GET("/tag", h.GetTag)

	app.POST("/tag/translate", h.TagTranslate)
	app.GET("/post/:id", h.GetPost)
	app.GET("/post", h.GetPosts)
	_ = app.Run(":8080")
}
