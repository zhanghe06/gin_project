package tests

import (
	"github.com/magiconair/properties/assert"
	"github.com/zhanghe06/gin_project/routers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexGet(t *testing.T) {
	// 初始设置
	Setup()

	// 初始化路由
	router := routers.Init()

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	//assert.Equal(t, "pong", w.Body.String())

	// 退出设置
	TearDown()
}
