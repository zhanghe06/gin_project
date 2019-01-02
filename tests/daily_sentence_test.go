package tests

import (
	"bytes"
	"encoding/json"
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

func DailySentenceCreateSuccessTest(t *testing.T) {
	// 初始化路由
	router := routers.Init()

	response := httptest.NewRecorder()
	url := fmt.Sprintf("/%s/daily_sentence", viper.GetString("ver"))
	bodyMap := map[string]string{
		"author": "Test",
		"title": "this is a test",
		"classification": "news",
	}
	bodyByte, _ := json.Marshal(bodyMap)

	// 新建记录
	request, _ := http.NewRequest("POST", url, bytes.NewReader(bodyByte))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	//assert.Equal(t, "", response.Body.String())
}

func DailySentenceCreateFailureTest(t *testing.T) {
	// 初始化路由
	router := routers.Init()

	response := httptest.NewRecorder()
	url := fmt.Sprintf("/%s/daily_sentence", viper.GetString("ver"))
	bodyMap := map[string]string{
		"author": "Test",
		"title": "我是一句话",
		"classification": "news",
	}
	bodyByte, _ := json.Marshal(bodyMap)

	// 重复记录
	requestRepeat, _ := http.NewRequest("POST", url, bytes.NewReader(bodyByte))
	router.ServeHTTP(response, requestRepeat)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	//assert.Equal(t, "", response.Body.String())
}

func DailySentenceDeleteTest(t *testing.T) {
	// 初始化路由
	router := routers.Init()

	response := httptest.NewRecorder()
	url := fmt.Sprintf("/%s/daily_sentence/1", viper.GetString("ver"))
	request, _ := http.NewRequest("DELETE", url, nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusNoContent, response.Code)
	//assert.Equal(t, "", response.Body.String())
}

func TestDailySentence(t *testing.T) {
	// 初始设置
	Setup()
	// This Run will not return until the parallel tests finish.
	t.Run("group", func(t *testing.T) {
		t.Run("DailySentenceListTest", DailySentenceListTest)
		t.Run("DailySentenceGetTest", DailySentenceGetTest)
		t.Run("DailySentenceCreateSuccessTest", DailySentenceCreateSuccessTest)
		t.Run("DailySentenceCreateFailureTest", DailySentenceCreateFailureTest)
		t.Run("DailySentenceDeleteTest", DailySentenceDeleteTest)
	})
	// 退出设置
	TearDown()
}
