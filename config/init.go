package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func Init() (err error) {
	projectPath := os.Getenv("PROJECT_PATH")
	confPath := filepath.Join(projectPath, "config")
	confMode := os.Getenv("CONF_MODE")

	viper.AddConfigPath(confPath)		// 设置配置文件路径
	viper.SetConfigName(confMode)		// 设置配置文件名称
	viper.SetConfigType("yaml")		// 设置配置文件类型
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return
}

func Watch() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed: ", e.Name)
	})
}
