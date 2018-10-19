package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取列表
// curl -u username:password -i -X GET http://0.0.0.0:8080/v1/token
// curl -u foo:bar -i -X GET http://0.0.0.0:8080/v1/token
func GetTokenHandler(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	c.JSON(http.StatusOK, gin.H{"user": user})
}
