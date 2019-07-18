package routers

import (
	"github.com/zhanghe06/gin_project/controllers"
	"github.com/zhanghe06/gin_project/middlewares"
)

func RegisterInfo() {
	RouterGroupVer.Use(middlewares.ApiTokenAuthMiddleware())
	RouterGroupVer.GET("/info", controllers.GetInfoHandler)
}
