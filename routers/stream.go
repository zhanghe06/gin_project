package routers

import (
	"github.com/zhanghe06/gin_project/controllers"
)

func RegisterStream() {
	Ver.GET("/stream/sse", controllers.StreamSSEHandler)
	Ver.GET("/stream/crd", controllers.StreamCRDHandler)
}
