package controllers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/zhanghe06/gin_project/dbs"
	"github.com/zhanghe06/gin_project/logs"
	"github.com/zhanghe06/gin_project/models"
	"github.com/zhanghe06/gin_project/requests"
	"github.com/zhanghe06/gin_project/utils"
	"net/http"
)

// 获取列表
// curl -i -X GET http://0.0.0.0:8080/v1/daily_sentences
func ListsDailySentenceHandler(c *gin.Context) {
	var dailySentences []models.DailySentence

	if err := dbs.DbClient.Find(&dailySentences).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		//fmt.Println(err)
		logs.Logger.Error(err)
		return
	}
	c.JSON(http.StatusOK, dailySentences)
	//c.String(200, "test list daily sentences")
}

// 获取详情
// curl -i -X GET http://0.0.0.0:8080/v1/daily_sentence/1
func GetDailySentenceHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	var dailySentence models.DailySentence

	if err := dbs.DbClient.Where("id = ?", id).First(&dailySentence).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		//fmt.Println(err)
		logs.Logger.Error(err)
		return
	}
	c.JSON(http.StatusOK, dailySentence)
}

// 创建记录
// curl -i -X POST http://0.0.0.0:8080/v1/daily_sentence -d '{"Author": "Tom", "Title": "this is a test", "Classification": "news"}'
func CreateDailySentenceHandler(c *gin.Context) {
	var dailySentence models.DailySentence
	c.BindJSON(&dailySentence)
	if err := dbs.DbClient.Create(&dailySentence).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		//fmt.Println(err)
		logs.Logger.Error(err)
		return
	}
	c.JSON(http.StatusOK, dailySentence)
}

// 更新记录
// curl -i -X PUT http://0.0.0.0:8080/v1/daily_sentence/1 -d '{"Author": "Tom"}'
func UpdateDailySentenceHandler(c *gin.Context) {
	var dailySentence models.DailySentence
	id := c.Params.ByName("id")
	if err := dbs.DbClient.Where("id = ?", id).First(&dailySentence).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		//fmt.Println(err)
		logs.Logger.Error(err)
		return
	}
	c.BindJSON(&dailySentence)
	if err := dbs.DbClient.Save(&dailySentence).Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		//fmt.Println(err)
		logs.Logger.Error(err)
		return
	}
	c.JSON(http.StatusOK, dailySentence)
}

// 删除记录
// curl -i -X DELETE http://0.0.0.0:8080/v1/daily_sentence/2
func DeleteDailySentenceHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	var dailySentence models.DailySentence
	if err := dbs.DbClient.Where("id = ?", id).Delete(&dailySentence).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		//fmt.Println(err)
		logs.Logger.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"id #" + id: " has deleted"})
}

// 打分
// curl -i -X PATCH http://0.0.0.0:8080/v1/daily_sentence/1/score -d '{"score": 0}'
func ScoreDailySentenceHandler(c *gin.Context) {
	// 参数校验
	var scoreDailySentenceRequests requests.ScoreDailySentenceRequests
	err := c.ShouldBindJSON(&scoreDailySentenceRequests)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		logs.Logger.Error(err)
		return
	}

	// 判断记录是否存在
	var dailySentence models.DailySentence
	id := c.Params.ByName("id")
	//if err := dbs.DbClient.Where("id = ?", id).First(&dailySentence).Error; err != nil {
	if err := dbs.DbClient.First(&dailySentence, id).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		//fmt.Println(err)
		logs.Logger.Error(err)
		return
	}

	// 更新数据
	data := utils.Struct2Map(scoreDailySentenceRequests)
	res := dbs.DbClient.Model(&dailySentence).Updates(data)
	if err = res.Error; err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		//fmt.Println(err)
		logs.Logger.Error(err)
		return
	}
	// 注意: 这里的res.Value和dailySentence一样, 更新的字段仅为updates传入的参数, 其余数据库自动修改的字段不会更新到这里
	c.JSON(http.StatusOK, dailySentence)
}
