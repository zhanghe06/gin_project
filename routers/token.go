package routers

import (
	"github.com/zhanghe06/gin_project/controllers"
)

func RegisterToken() {
	RouterGroupVerToken.GET("/token", controllers.GetTokenHandler)
}
