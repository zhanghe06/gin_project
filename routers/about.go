package routers

import (
	"github.com/zhanghe06/gin_project/controllers"
)

func RegisterAbout() {
	Router.GET("/about", controllers.GetAboutHandler)
}
