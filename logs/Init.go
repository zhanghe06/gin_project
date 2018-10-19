package logs

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

var LogFields log.Fields

func Init() (err error) {
	f, err := os.OpenFile("logs/app.log", os.O_WRONLY | os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	debug := viper.GetBool("debug")
	// 生产环境写入文件
	if debug == false{
		log.SetOutput(f)
		gin.SetMode(gin.ReleaseMode)
	}
	// 开发环境标准输出（默认）DEBUG模式

	LogFields = log.Fields {
		"project": "gin_project",
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
	log.WithFields(LogFields)

	return
}
