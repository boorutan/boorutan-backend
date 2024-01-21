package handler

import (
	"applemango/boorutan/backend/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var b user.User
	err := c.Bind(&b)
	if err != nil || len(b.Id) == 0 {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, user.GetUser(b.Id))
}
