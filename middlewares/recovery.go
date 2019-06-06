package middlewares

import (
	"fmt"
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

		contentType := c.ContentType()
		if contentType != "" {
			c.Writer.Header().Set("Content-Type", fmt.Sprintf("%s; charset=utf-8", c.ContentType()))
		}

		//c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		c.Next()
	}
}
