package middlewares


import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)


func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.Must(uuid.NewV4()).String())
		c.Next()
	}
}
