package middlewares

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ApiTokenAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.Request.FormValue("api_token")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "API token required"})
			return
		}

		if token != os.Getenv("API_TOKEN") {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"msg": "Invalid API token"})
			return
		}
		c.Next()
	}
}
