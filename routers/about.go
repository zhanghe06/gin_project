package routers

import (
	"github.com/zhanghe06/gin_project/controllers"
	"os"
	"path/filepath"
)

func RegisterAbout() {
	projectPath := os.Getenv("PROJECT_PATH")
	templatePath := filepath.Join(projectPath, "templates/about/*")
	//Router.LoadHTMLFiles("templates/about/index.tmpl")
	Router.LoadHTMLGlob(templatePath)
	Router.GET("/about", controllers.GetAboutHandler)
}
