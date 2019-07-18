package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zhanghe06/gin_project/middlewares"
	"os"
	"path/filepath"
)

var Router *gin.Engine

var RouterGroupVer *gin.RouterGroup
var RouterGroupVerToken *gin.RouterGroup


func Init() *gin.Engine {
	// 创建路由 - REST
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

	RouterGroupVer = Router.Group(ver)

	RouterGroupVerToken = Router.Group(ver, gin.BasicAuth(accounts))
	RouterGroupVerToken.Use(middlewares.RequestIpMiddleware())

	// 模板加载
	projectPath := os.Getenv("PROJECT_PATH")
	templatePath := filepath.Join(projectPath, "templates/website/*")
	Router.LoadHTMLGlob(templatePath)

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
