package routers

import (
	"github.com/zhanghe06/gin_project/controllers"
)

func RegisterStream() {
	RouterGroupVer.GET("/stream/sse", controllers.StreamSSEHandler)
	RouterGroupVer.GET("/stream/crd", controllers.StreamCRDHandler)
}
