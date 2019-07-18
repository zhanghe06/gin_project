package controllers

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 流式请求（利用 H5 的新特性 SSE）
// curl -i -X GET http://0.0.0.0:8080/v1/stream/sse
// https://www.html5rocks.com/en/tutorials/eventsource/basics/
func StreamSSEHandler(c *gin.Context) {
	chanStream := make(chan int, 10)
	go func() {
		defer close(chanStream)
		for i := 0; i < 5; i++ {
			chanStream <- i
			time.Sleep(time.Second * 1)
		}
	}()
	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-chanStream; ok {
			c.SSEvent("message", msg)
			return true
		}
		return false
	})
}

// 流式请求
// curl -i -X GET http://0.0.0.0:8080/v1/stream/crd
func StreamCRDHandler(c *gin.Context) {
	chanStream := make(chan int, 10)
	go func() {
		defer close(chanStream)
		for i := 0; i < 5; i++ {
			chanStream <- i
			time.Sleep(time.Second * 1)
		}
	}()
	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-chanStream; ok {
			data := fmt.Sprintf("%d\n", msg)
			c.Data(http.StatusOK, "text/event-stream", []byte(data))
			return true
		}
		return false
	})
}
