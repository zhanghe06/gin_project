package middlewares

import "github.com/gin-gonic/gin"

func RequestIpMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Set header variable
		c.Writer.Header().Set("Request-Ip", c.ClientIP())

		// before request

		c.Next()

		// after request

	}
}
