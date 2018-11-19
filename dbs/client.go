package dbs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"strings"
	"time"
)

var DbClient *gorm.DB

func Init() (err error) {
	if DbClient != nil {
		return
	}

	dbStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.ip"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.name"),
		viper.GetString("mysql.charset"),
		true,
		"Local",
	)
	// 超时配置
	timeout := viper.GetString("mysql.timeout")
	if timeout != "" {
		dbStr = strings.Join([]string{dbStr, fmt.Sprintf("timeout=%s", timeout)}, "&")
	}
	readTimeout := viper.GetString("mysql.timeout_read")
	if readTimeout != "" {
		dbStr = strings.Join([]string{dbStr, fmt.Sprintf("readTimeout=%s", readTimeout)}, "&")
	}
	writeTimeout := viper.GetString("mysql.timeout_write")
	if writeTimeout != "" {
		dbStr = strings.Join([]string{dbStr, fmt.Sprintf("writeTimeout=%s", writeTimeout)}, "&")
	}

	DbClient, err = gorm.Open("mysql", dbStr)
	if err != nil {
		return err
	}
	DbClient.SingularTable(true)             // 不考虑表名单复数变化
	DbClient.LogMode(viper.GetBool("debug")) // 是否显示sql语句

	// 连接池配置
	DbClient.DB().SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))                                    // 默认值0，无限制
	DbClient.DB().SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))                                    // 默认值2
	DbClient.DB().SetConnMaxLifetime(time.Duration(viper.GetInt("mysql.conn_max_lifetime")) * time.Second) // 默认值0，永不过期
	// maxBadConnRetries 默认重试2次

	// 因为连接惰性创建, 这里预先创建1个连接
	err = DbClient.DB().Ping()
	if err != nil {
		return err
	}

	return
}

func Close() (err error) {
	if DbClient != nil {
		err = DbClient.Close()
		return err
	}
	return
}
