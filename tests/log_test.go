package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zhanghe06/gin_project/models"
	"github.com/zhanghe06/gin_project/utils"
	"os"
	"testing"
	"time"
)

// 你可以创建很多instance
//var logF = logrus.New()
//var logJ = logrus.New()

func getLogSourceStart()  {
	return
}

func logJson(log logrus.Logger) {

}


func logSourceStartJson() string {
	// TODO 日志记录触发条件

	f, err := os.OpenFile("../logs/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(f)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 资源日志 - 开始
	var logSourceStart models.LogSourceStart
	// 固定参数
	logSourceStart.AuditType = "resource"
	logSourceStart.LogType = "LOG_START"
	logSourceStart.Action = "TestAction"
	// 动态参数
	logSourceStart.RequestId = "RequestId" // TODO header

	var ctx *gin.Context

	logSourceStart.Action = ""       // 操作action
	logSourceStart.RequestId = ctx.Request.Header.Get("X-Request-Id")    // 操作请求ID
	logSourceStart.Version = ""      // 版本号
	logSourceStart.VisitorId = ""    // 操作者ID
	logSourceStart.CloudId = ""      // 云环境ID
	logSourceStart.DataCenterId = "" // 数据中心ID
	logSourceStart.OwnerId = ""      // 项目ID
	logSourceStart.OwnerName = ""    // 项目名称




	log := logrus.WithFields(utils.Struct2Map(logSourceStart))
	log.Info("")

	//b, err := json.Marshal(logSourceStart)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(b))
	//
	return time.Now().Format("2006-01-06 15:04:05")
}

func TestLogSourceStartJson(t *testing.T) {
	logSourceStartJson()
}
