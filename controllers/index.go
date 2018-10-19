package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 首页页面
// curl -i -X GET http://0.0.0.0:8080
func GetIndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}
