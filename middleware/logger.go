package middleware

import (
	"applemango/boorutan/backend/db/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

type RouterLogger struct {
	Latency   int64  `json:"duration,omitempty"`
	Ip        string `json:"ip,omitempty"`
	Method    string `json:"method,omitempty"`
	RequestId string `json:"requestId,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Message   string `json:"message,omitempty"`
}

func BaseLimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := uuid.New().String()
		start := time.Now()

		c.Set("requestId", requestId)

		c.Next()

		latency := time.Now().Sub(start).Milliseconds()

		logger.Ctx.SendEvent(RouterLogger{
			Latency:   latency,
			Ip:        c.ClientIP(),
			Method:    c.Request.Method,
			RequestId: requestId,
			Namespace: c.FullPath(),
			Message:   "New Request",
		})

	}
}
