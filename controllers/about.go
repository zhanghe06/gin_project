package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 关于页面
// curl -i -X GET http://0.0.0.0:8080/about
func GetAboutHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "About website",
	})
}
