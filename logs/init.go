package logs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var LogFields log.Fields
var Logger *log.Logger
//var F *os.File				// 普通文件
var Fl *rotatelogs.RotateLogs   // 日志切割

func Init() (err error) {
	// 普通文件
	//F, err := os.OpenFile("logs/app.log", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0755)
	//if err != nil {
	//	return err
	//}

	// 日志切割
	Fl, err := rotatelogs.New(
		fmt.Sprintf("%s.%%Y%%m%%d", viper.GetString("log.file_path")),
		rotatelogs.WithLinkName(viper.GetString("log.file_path")),
		rotatelogs.WithClock(rotatelogs.UTC),      // default: rotatelogs.Local
		rotatelogs.WithRotationTime(24*time.Hour), // default: 86400 sec
		rotatelogs.WithMaxAge(7*24*time.Hour),     // default: 7 days
	)
	if err != nil {
		return err
	}
	debug := viper.GetBool("debug")
	// 生产环境写入文件
	if debug == false {
		//log.SetOutput(F)	// 普通文件
		log.SetOutput(Fl) // 日志切割
		gin.SetMode(gin.ReleaseMode)
	}
	// 开发环境标准输出（默认）DEBUG模式

	LogFields = log.Fields{
		"project": viper.GetString("project.name"),
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
	log.WithFields(LogFields)

	return
}

func Close() {
	//F.Close()
	Fl.Close()
}
