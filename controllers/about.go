package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 关于页面
// curl -i -X GET http://0.0.0.0:8080/about
func GetAboutHandler(c *gin.Context) {
	// 传递requestId
	requestId := c.Writer.Header().Get("X-Request-Id")

	c.HTML(http.StatusOK, "about.tmpl", gin.H{
		"title": "About website",
		"requestId": requestId,
	})
}
