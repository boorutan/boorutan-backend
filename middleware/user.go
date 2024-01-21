package middleware

import (
	"applemango/boorutan/backend/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		account := c.GetHeader("Account")
		if len(account) == 0 {
			c.JSON(http.StatusBadRequest, "error")
			c.Abort()
			return
		}
		u := user.GetUser(account)
		c.Set("account", u)
		c.Next()
	}
}
