package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTagCategory(c *gin.Context) {
	b := GetBooru(c)
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
}
