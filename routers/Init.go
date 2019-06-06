package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zhanghe06/gin_project/middlewares"
)

var Router *gin.Engine
var Ver *gin.RouterGroup
var VerToken *gin.RouterGroup

func Init() *gin.Engine {
	// 创建路由
	Router = gin.New()

	// 测试模式禁用日志中间件
	if viper.GetString("mode") != "test" {
		Router.Use(
			middlewares.LoggingMiddleware(),
		)
	}

	Router.Use(
		//gin.Logger(),
		middlewares.RecoveryMiddleware(),
		middlewares.RequestIdMiddleware(),
	)

	ver := fmt.Sprintf("/%s", viper.GetString("ver"))
	accounts := gin.Accounts{
		viper.GetString("BasicAuth.Username"): viper.GetString("BasicAuth.Password"),
	}

	Ver = Router.Group(ver)

	VerToken = Router.Group(ver, gin.BasicAuth(accounts))
	VerToken.Use(middlewares.RequestIpMiddleware())

	// 注册路由
	RegisterIndex()
	RegisterAbout()
	RegisterDailySentence()
	RegisterStream()
	RegisterToken()
	RegisterInfo()
	RegisterDownload()

	return Router
}
