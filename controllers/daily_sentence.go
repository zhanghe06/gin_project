package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/zhanghe06/gin_project/apis"
	"github.com/zhanghe06/gin_project/dbs"
	"github.com/zhanghe06/gin_project/models"
	"github.com/zhanghe06/gin_project/requests"
	"github.com/zhanghe06/gin_project/utils"
)

// 获取列表
// curl -i -X GET http://0.0.0.0:8080/v1/daily_sentences
func ListsDailySentenceHandler(c *gin.Context) {
	// 意外异常
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			err := fmt.Errorf("%v", rec)
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
	}(c)

	var dailySentences []models.DailySentence

	if err := dbs.DbClient.Find(&dailySentences).Error; err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, dailySentences)
	//c.String(200, "test list daily sentences")
}

// 获取详情
// curl -i -X GET http://0.0.0.0:8080/v1/daily_sentence/1
func GetDailySentenceHandler(c *gin.Context) {
	// 意外异常
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			err := fmt.Errorf("%v", rec)
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
	}(c)

	//id := c.Params.ByName("id")
	id := c.Param("id")
	var dailySentence models.DailySentence

	if err := dbs.DbClient.Where("id = ?", id).First(&dailySentence).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			// 记录不存在
			_ = c.AbortWithError(http.StatusNotFound, err)
		} else {
			// 数据库异常
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}
	c.JSON(http.StatusOK, dailySentence)
}

// 创建记录
// curl -i -X POST http://0.0.0.0:8080/v1/daily_sentence -d '{"Author": "Tom", "Title": "this is a test", "Classification": "news"}'
func CreateDailySentenceHandler(c *gin.Context) {
	// 意外异常
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			err := fmt.Errorf("%v", rec)
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
	}(c)

	var dailySentence models.DailySentence

	// 参数校验
	if err := c.ShouldBindJSON(&dailySentence); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 记录时间
	currentTime := time.Now()
	dailySentence.CreateTime = currentTime
	dailySentence.UpdateTime = currentTime

	if err := dbs.DbClient.Create(&dailySentence).Error; err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok { // Now the error number is accessible directly
			if driverErr.Number == 1062 {
				// 记录重复
				_ = c.AbortWithError(http.StatusBadRequest, err)
			}
		} else {
			// 数据库异常
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}
	c.JSON(http.StatusOK, dailySentence)
}

// 更新记录
// curl -i -X PUT http://0.0.0.0:8080/v1/daily_sentence/1 -d '{"Author": "Tom"}'
func UpdateDailySentenceHandler(c *gin.Context) {
	// 意外异常
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			err := fmt.Errorf("%v", rec)
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
	}(c)

	var dailySentence models.DailySentence

	//id := c.Params.ByName("id")
	id := c.Param("id")
	if err := dbs.DbClient.Where("id = ?", id).First(&dailySentence).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			// 记录不存在
			_ = c.AbortWithError(http.StatusNotFound, err)
		} else {
			// 数据库异常
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	// 参数校验, 注意顺序, 放在获取数据之后
	if err := c.ShouldBindJSON(&dailySentence); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// 保存所有字段
	if err := dbs.DbClient.Save(&dailySentence).Error; err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok { // Now the error number is accessible directly
			if driverErr.Number == 1062 {
				// 记录重复
				_ = c.AbortWithError(http.StatusBadRequest, err)
			}
		} else {
			// 数据库异常
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}
	c.JSON(http.StatusOK, dailySentence)
}

// 修改记录标题
// curl -i -X PUT http://0.0.0.0:8080/v1/daily_sentence/1/title -d '{"title": "测试更换标题"}'
func ReTitleDailySentenceHandler(c *gin.Context) {
	// 意外异常
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			err := fmt.Errorf("%v", rec)
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
	}(c)

	var jsonRequests requests.ReTitleJsonRequests

	// 参数校验
	if err := c.ShouldBindJSON(&jsonRequests); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//id := c.Params.ByName("id")
	id := c.Param("id")

	// 开启事务
	tx := dbs.DbClient.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			// 事务异常传递
			panic(r)
		}
	}()

	if err := tx.Error; err != nil {
		// 事务异常
		panic(err)
	}

	// 判断记录是否存在
	var dailySentence models.DailySentence
	if err := tx.Where("id = ?", id).Set("gorm:query_option", "FOR UPDATE").First(&dailySentence).Error; err != nil {
		tx.Rollback()
		if gorm.IsRecordNotFoundError(err) {
			// 记录不存在
			_ = c.AbortWithError(http.StatusNotFound, err)
		} else {
			// 数据库异常
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	updateData := models.DailySentence{
		Title:      jsonRequests.Title,
		UpdateTime: time.Now(),
	}
	// 更新指定字段
	if err := tx.Model(&dailySentence).UpdateColumns(updateData).Error; err != nil {
		tx.Rollback()
		if driverErr, ok := err.(*mysql.MySQLError); ok { // Now the error number is accessible directly
			if driverErr.Number == 1062 {
				// 记录重复
				_ = c.AbortWithError(http.StatusBadRequest, err)
			}
		} else {
			// 数据库异常
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		// 事务异常
		panic(err)
	}

	c.JSON(http.StatusOK, dailySentence)
}

// 删除记录
// curl -i -X DELETE http://0.0.0.0:8080/v1/daily_sentence/2
func DeleteDailySentenceHandler(c *gin.Context) {
	// 意外异常
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			err := fmt.Errorf("%v", rec)
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
	}(c)

	var uriRequests requests.DeleteDailySentenceUriRequests

	// 参数校验
	if err := c.ShouldBindUri(&uriRequests); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//id := c.Params.ByName("id")
	//id := c.Param("id")
	id := uriRequests.ID

	var dailySentence models.DailySentence
	if err := dbs.DbClient.Where("id = ?", id).Delete(&dailySentence).Error; err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// 打分
// curl -i -X PATCH http://0.0.0.0:8080/v1/daily_sentence/1/score -d '{"score": 0}'
func ScoreDailySentenceHandler(c *gin.Context) {
	// 意外异常
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			err := fmt.Errorf("%v", rec)
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
	}(c)

	var uriRequests requests.ScoreDailySentenceUriRequests

	// 参数校验
	if err := c.ShouldBindUri(&uriRequests); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var jsonRequests requests.ScoreDailySentenceJsonRequests

	// 参数校验
	if err := c.ShouldBindJSON(&jsonRequests); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//id := c.Params.ByName("id")
	//id := c.Param("id")
	id := uriRequests.ID

	// 判断记录是否存在
	var dailySentence models.DailySentence
	//if err := dbs.DbClient.Where("id = ?", id).First(&dailySentence).Error; err != nil {
	if err := dbs.DbClient.First(&dailySentence, id).Error; err != nil { // 此种写法仅仅支持整形主键
		if gorm.IsRecordNotFoundError(err) {
			// 记录不存在
			_ = c.AbortWithError(http.StatusNotFound, err)
		} else {
			// 数据库异常
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
		return
	}

	// 更新数据
	data := utils.Struct2Map(jsonRequests)
	if err := dbs.DbClient.Model(&dailySentence).Updates(data).Error; err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	// 注意: 这里的res.Value和dailySentence一样, 更新的字段仅为updates传入的参数, 其余数据库自动修改的字段不会更新到这里
	c.JSON(http.StatusOK, dailySentence)
}



// 更新记录
//curl -X POST \
// http://0.0.0.0:8080/v1/daily_sentence/transaction \
// -H 'Content-Type: application/json' \
// -d '{
//   "id": "1",
//   "Title": "this is a test",
//   "Classification": "news"
//}'
func UpdateDailySentenceTransactionHandler(c *gin.Context) {
	// 意外异常
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {
			err := fmt.Errorf("%v", rec)
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
	}(c)

	var dailySentenceRequest requests.UpdateDailySentenceTransactionRequests

	// 参数校验
	if err := c.ShouldBindJSON(&dailySentenceRequest); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	condition := map[string]interface{}{"id": dailySentenceRequest.ID}
	updateData := map[string]interface{}{}
	if dailySentenceRequest.Title != "" {
		updateData["title"] = dailySentenceRequest.Title
	}
	rows, err := apis.UpdateDailySentenceTransaction(condition, updateData)
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"rows": rows})
}
