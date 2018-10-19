package routers

import (
	"github.com/zhanghe06/gin_project/controllers"
	"github.com/zhanghe06/gin_project/middlewares"
)

func RegisterToken() {
	VerToken.Use(middlewares.RequestIpMiddleware())
	VerToken.GET("/token", controllers.GetTokenHandler)
}
