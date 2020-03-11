package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/zhanghe06/gin_project/logs"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()

		// 获取请求
		bodyCopy := new(bytes.Buffer)
		// Read the whole body
		_, err := io.Copy(bodyCopy, c.Request.Body)
		if err != nil {
			return // 退出中间件
		}
		bodyData := bodyCopy.Bytes()
		// Replace the body with a reader that reads from the buffer
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyData))

		var requestParams string
		requestParams = string(bodyData) // 默认原始

		decoder := json.NewDecoder(bodyCopy)
		var paramsMap map[string]interface{}
		if err = decoder.Decode(&paramsMap); err == nil {
			// 数据脱敏
			// delete(paramsMap, "password")
			if _, ok := paramsMap["password"]; ok {
				paramsMap["password"] = "******"
			}

			paramsByte, e := json.Marshal(paramsMap)
			if e == nil {
				requestParams = string(paramsByte) // 反向解析
			}
		}

		c.Next()

		// 结束时间
		end := time.Now()
		// 执行时间
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		logFields := log.Fields{
			"project":        viper.GetString("project.name"),
			"latency":        fmt.Sprintf("%v", latency),
			"client_ip":      clientIP,
			"host":           c.Request.Host,
			"method":         method,
			"request_id":     c.Writer.Header().Get("X-Request-Id"),
			"request_uri":    c.Request.RequestURI,
			"request_params": requestParams,
			"status_code":    statusCode,
		}
		logs.Log = logs.Logger.WithFields(logFields)

		// 捕获异常
		if c.Errors != nil {
			for _, err := range c.Errors {
				logs.Log.Error(err)
			}
		} else {
			logs.Log.Info()
		}
	}
}
