package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)


func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 若存在，则传递，便于链路追踪
		requestId := c.Request.FormValue("X-Request-Id")
		if requestId == "" {
			requestId = uuid.Must(uuid.NewV4()).String()
		}
		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
