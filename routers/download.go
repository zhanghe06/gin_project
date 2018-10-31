package routers

import (
	"github.com/zhanghe06/gin_project/controllers"
)

func RegisterDownload() {
	Router.GET("/download", controllers.DownloadHandler)
}
