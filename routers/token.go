package routers

import (
	"github.com/zhanghe06/gin_project/controllers"
)

func RegisterToken() {
	VerToken.GET("/token", controllers.GetTokenHandler)
}
