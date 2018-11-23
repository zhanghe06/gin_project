package middlewares

import (
	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 处理异常
		defer func() {
			for _, err := range c.Errors {
				c.AbortWithStatusJSON(c.Writer.Status(), err)
			}
		}()
		c.Next()
	}
}
