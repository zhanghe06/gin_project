package tests

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"github.com/zhanghe06/gin_project/routers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func DailySentenceListTest(t *testing.T) {
	// 初始化路由
	router := routers.Init()

	response := httptest.NewRecorder()
	url := fmt.Sprintf("/%s/daily_sentences", viper.GetString("ver"))
	request, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	//assert.Equal(t, "", response.Body.String())
}

func DailySentenceGetTest(t *testing.T) {
	// 初始化路由
	router := routers.Init()

	response := httptest.NewRecorder()
	url := fmt.Sprintf("/%s/daily_sentence/1", viper.GetString("ver"))
	request, _ := http.NewRequest("GET", url, nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	//assert.Equal(t, "", response.Body.String())
}

func TestDailySentence(t *testing.T) {
	// 初始设置
	Setup()
	// This Run will not return until the parallel tests finish.
	t.Run("group", func(t *testing.T) {
		t.Run("DailySentenceListTest", DailySentenceListTest)
		t.Run("DailySentenceGetTest", DailySentenceGetTest)
	})
	// 退出设置
	TearDown()
}
