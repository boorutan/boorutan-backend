package handler

import (
	"applemango/boorutan/backend/booru"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetPost(c *gin.Context) {
	b := GetBooru(c)
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
}

func GetPosts(c *gin.Context) {
	b := GetBooru(c)
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
	_, inBypassCache := c.GetQuery("bypasscache")
	post, err := b.GetPosts(booru.GetPostsOption{
		Cache: !inBypassCache,
		Page:  page,
		Tags:  tags,
	})
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, post)
}
