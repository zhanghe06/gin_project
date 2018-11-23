package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zhanghe06/gin_project/logs"
	"time"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()

		c.Next()

		// 结束时间
		end := time.Now()
		// 执行时间
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		logFields := log.Fields{
			"project":     viper.GetString("project.name"),
			"latency":     fmt.Sprintf("%v", latency),
			"client_ip":   clientIP,
			"method":      method,
			"status_code": statusCode,
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
