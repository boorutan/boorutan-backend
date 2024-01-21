package handler

import (
	"applemango/boorutan/backend/utils/image"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DownloadImage(c *gin.Context) {
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
}
