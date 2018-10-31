package controllers

import (
	"github.com/gin-gonic/gin"
)

// 关于页面
// curl -i -X GET http://0.0.0.0:8080/download
func DownloadHandler(c *gin.Context) {
	c.Header("content-disposition", "attachment; filename=download.txt")
	c.Data(200, "application/octet-stream", []byte("test"))
}
