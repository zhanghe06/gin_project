package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 首页页面
// curl -i -X GET http://0.0.0.0:8080
func GetIndexHandler(c *gin.Context) {
	// 传递requestId
	requestId := c.Writer.Header().Get("X-Request-Id")

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
		"requestId": requestId,
	})
}
