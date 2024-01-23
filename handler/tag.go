package handler

import (
	"applemango/boorutan/backend/booru"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTagSuggest(c *gin.Context) {
	q, in := c.GetQuery("q")
	if !in {
		c.JSON(http.StatusInternalServerError, "err")
		return
	}
	tags := booru.SearchTags(q)
	c.JSON(http.StatusOK, tags)
}

func GetTagSuggestFast(c *gin.Context) {
	q, in := c.GetQuery("q")
	if !in {
		c.JSON(http.StatusInternalServerError, "err")
		return
	}
	tags := booru.SearchTagsFast(q)
	c.JSON(http.StatusOK, tags)
}

func GetTag(c *gin.Context) {
	b := GetBooru(c)
	tag, err := b.GetTag(booru.GetTagOption{
		Cache: true,
	})
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, tag)
}

func TagTranslate(c *gin.Context) {
	type Body struct {
		Tags []string `json:"tags"`
	}
	var b Body
	err := c.Bind(&b)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	tags := booru.TranslateTags(b.Tags)
	c.JSON(http.StatusOK, tags)
}
