package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var Router *gin.Engine
var Ver *gin.RouterGroup
var VerToken *gin.RouterGroup

func Init() *gin.Engine {
	// 创建路由
	Router = gin.Default()

	ver := fmt.Sprintf("/%s", viper.GetString("ver"))
	accounts := gin.Accounts{
		viper.GetString("BasicAuth.Username"): viper.GetString("BasicAuth.Password"),
	}

	Ver = Router.Group(ver)
	VerToken = Router.Group(ver, gin.BasicAuth(accounts))

	// 注册路由
	RegisterIndex()
	RegisterAbout()
	RegisterDailySentence()
	RegisterToken()

	return Router
}
