package routers

import (
	"github.com/zhanghe06/gin_project/controllers"
	"os"
	"path/filepath"
)

func RegisterIndex() {
	projectPath := os.Getenv("PROJECT_PATH")
	templatePath := filepath.Join(projectPath, "templates/index.tmpl")
	Router.LoadHTMLFiles(templatePath)
	Router.GET("/", controllers.GetIndexHandler)
}
