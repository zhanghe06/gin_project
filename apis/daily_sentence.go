package apis

import (
	"fmt"
	"github.com/zhanghe06/gin_project/dbs"
	"github.com/zhanghe06/gin_project/models"
)

func UpdateDailySentenceTransaction(condition map[string]interface{}, updateData map[string]interface{}) (rows int64, err error) {
	tx := dbs.DbClient.Begin()
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
			fmt.Println(err.Error())
			tx.Rollback()
		}
	}()
	var dailySentences []*models.DailySentence

	resultFind := tx.Set("gorm:query_option", "FOR UPDATE").Where(condition).Find(&dailySentences)
	if resultFind.Error  != nil {
		panic(resultFind.Error)
	}

	resultUpdate := tx.Model(&dailySentences).Updates(updateData)
	if resultUpdate.Error  != nil {
		panic(resultUpdate.Error)
	}
	tx.Commit()
	return resultUpdate.RowsAffected, nil
}
