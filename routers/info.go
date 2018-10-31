package routers

import (
	"github.com/zhanghe06/gin_project/controllers"
	"github.com/zhanghe06/gin_project/middlewares"
)

func RegisterInfo() {
	Ver.Use(middlewares.ApiTokenAuthMiddleware())
	Ver.GET("/info", controllers.GetInfoHandler)
}
