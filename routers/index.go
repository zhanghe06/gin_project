package routers

import (
	"github.com/zhanghe06/gin_project/controllers"
)

func RegisterIndex() {
	Router.GET("/", controllers.GetIndexHandler)
}
