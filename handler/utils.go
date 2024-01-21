package handler

import (
	"applemango/boorutan/backend/booru"
	"applemango/boorutan/backend/booru/danbooru"
	"applemango/boorutan/backend/booru/moebooru"
	"applemango/boorutan/backend/db/redis"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func GetBooruFromString(booru string) booru.Booru {
	switch booru {
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

func GetBooru(c *gin.Context) booru.Booru {
	b, in := c.GetQuery("booru")
	if !in {
		return danbooru.CreateDanBooru("https://danbooru.donmai.us/")
	}
	return GetBooruFromString(b)
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

func PushTag(tag *booru.DanbooruTag, json string) error {
	err := redis.Push(fmt.Sprintf("cache:tag:%v", tag.Name), json)
	return err
}

func ReadTags() error {
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
		err = PushTag(tag, string(line))
		if err != nil {
			println(err.Error())
			continue
		}
	}
	return nil
}
