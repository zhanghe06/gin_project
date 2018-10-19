package dbs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
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
	DbClient.SingularTable(true)
	return
}


func Close() (err error)  {
	if DbClient != nil{
		err = DbClient.Close()
		return err
	}
	return
}
