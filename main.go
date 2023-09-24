package main

import (
	"applemango/boorutan/backend/booru"
	"applemango/boorutan/backend/booru/danbooru"
	"applemango/boorutan/backend/booru/moebooru"
	"applemango/boorutan/backend/db/redis"
	"applemango/boorutan/backend/utils/image"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getBooru(c *gin.Context) booru.Booru {
	b, in := c.GetQuery("booru")
	if !in {
		return danbooru.CreateDanBooru("https://danbooru.donmai.us/")
	}
	switch b {
	case "konachan":
		return moebooru.CreateMoeBooru("https://konachan.com")
	case "safekonachan":
		return moebooru.CreateMoeBooru("https://konachan.net")
	case "yandere":
		return moebooru.CreateMoeBooru("https://yande.re")
	case "lolibooru":
		return moebooru.CreateMoeBooru("https://lolibooru.moe")
	case "danbooru":
		return danbooru.CreateDanBooru("https://danbooru.donmai.us/")
	}
	return danbooru.CreateDanBooru("https://danbooru.donmai.us/")
}

func OptionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method == "OPTION" || method == "OPTIONS" {
			c.JSON(http.StatusOK, "pong")
			c.Abort()
		}
	}
}

func pushTag(tag *booru.DanbooruTag, json string) error {
	err := redis.Push(fmt.Sprintf("cache:tag:%v", tag.Name), json)
	return err
}

func readTags() error {
	fp, err := os.Open("./tags.json")
	if err != nil {
		return err
	}
	defer fp.Close()

	reader := bufio.NewReader(fp)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			println(err.Error())
			continue
		}
		var tag *booru.DanbooruTag
		err = json.Unmarshal(line, &tag)
		if err != nil {
			println(err.Error())
			continue
		}
		err = pushTag(tag, string(line))
		if err != nil {
			println(err.Error())
			continue
		}
	}
	return nil
}

func __init__() {
	b := moebooru.CreateMoeBooru("https://konachan.net")
	_ = b.GetTags()

	_ = readTags()
}

func main() {

	gin.SetMode(gin.ReleaseMode)
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
				"Access-Control-Allow-Origin",
				"Content-Type",
				"Content-Length",
				"Accept-Encoding",
				"Authorization",
			},
			MaxAge: 0,
		}))
		app.Use(OptionMiddleware())
		app.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "pong")
		})
	}
	{
		app.OPTIONS("/image", OptionMiddleware())
		app.GET("/image", func(c *gin.Context) {
			url, in := c.GetQuery("url")
			if !in {
				c.JSON(http.StatusInternalServerError, "err")
				return
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
		app.POST("")
	}
	{
		app.OPTIONS("/category", OptionMiddleware())
		app.GET("/category", func(c *gin.Context) {
			b := getBooru(c)
			tags, in := c.GetQuery("tag")
			if !in {
				c.JSON(http.StatusInternalServerError, "err")
				return
			}
			category, err := b.GetTagsCategoryFromString(tags)
			if err != nil {
				println(err.Error())
				c.JSON(http.StatusInternalServerError, err)
				return
			}
			c.JSON(http.StatusOK, category)
		})
		app.OPTIONS("/tag/suggest", OptionMiddleware())
		app.GET("/tag/suggest", func(c *gin.Context) {
			q, in := c.GetQuery("q")
			if !in {
				c.JSON(http.StatusInternalServerError, "err")
				return
			}
			tags := booru.SearchTags(q)
			c.JSON(http.StatusOK, tags)
		})
		app.OPTIONS("/tag", OptionMiddleware())
		app.GET("/tag", func(c *gin.Context) {
			b := getBooru(c)
			tag, err := b.GetTag(booru.GetTagOption{
				Cache: true,
			})
			if err != nil {
				println(err.Error())
				c.JSON(http.StatusInternalServerError, err)
				return
			}
			c.JSON(http.StatusOK, tag)
		})
		app.OPTIONS("/post/:id", OptionMiddleware())
		app.GET("/post/:id", func(c *gin.Context) {
			b := getBooru(c)
			idStr := c.Param("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				println(err.Error())
				c.JSON(http.StatusInternalServerError, err)
				return
			}
			post, err := b.GetPost(booru.GetPostOption{
				ID: id,
			})
			if err != nil {
				println(err.Error())
				c.JSON(http.StatusInternalServerError, err)
				return
			}
			c.JSON(http.StatusOK, post)
		})
		app.OPTIONS("/post", OptionMiddleware())
		app.GET("/post", func(c *gin.Context) {
			b := getBooru(c)
			pageStr, in := c.GetQuery("page")
			var tags any
			tags, inTags := c.GetQuery("tags")
			var page int
			if !inTags || tags == "" {
				tags = nil
			}
			if !in {
				page = 1
			} else {
				p, err := strconv.Atoi(pageStr)
				if err != nil {
					println(err.Error())
					c.JSON(http.StatusInternalServerError, err)
					return
				}
				page = p
			}
			println(page)
			post, err := b.GetPosts(booru.GetPostsOption{
				Cache: true,
				Page:  page,
				Tags:  tags,
			})
			if err != nil {
				println(err.Error())
				c.JSON(http.StatusInternalServerError, err)
				return
			}
			c.JSON(http.StatusOK, post)
		})
	}
	_ = app.Run(":8080")
}
