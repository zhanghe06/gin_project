package dbs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"time"
)


var DbClient *gorm.DB

func Init() (err error) {
	if DbClient != nil{
		return
	}

	dbStr := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.ip"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.name"),
		viper.GetString("mysql.charset"),
		true,
		"Local",
	)

	DbClient, err = gorm.Open("mysql", dbStr)
	if err != nil {
		return err
	}
	DbClient.SingularTable(true)				// 不考虑表名单复数变化
	DbClient.LogMode(viper.GetBool("debug"))	// 是否显示sql语句

	// 连接池配置
	DbClient.DB().SetMaxOpenConns(5)
	DbClient.DB().SetMaxIdleConns(10)
	DbClient.DB().SetConnMaxLifetime(500 * time.Second)
	// maxBadConnRetries 默认重试2次

	// 因为连接惰性创建, 这里预先创建
	err = DbClient.DB().Ping()
	if err != nil{
		return err
	}

	return
}


func Close() (err error)  {
	if DbClient != nil{
		err = DbClient.Close()
		return err
	}
	return
}
