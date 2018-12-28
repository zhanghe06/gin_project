package models

import (
	"time"
)

type DailySentence struct {
	Id             int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Author         string    `json:"author" xorm:"not null default '' comment('作者') VARCHAR(20)"`
	Title          string    `json:"title" xorm:"not null default '' comment('标题') unique VARCHAR(100)"`
	Classification string    `json:"classification" xorm:"not null default '' comment('类型') VARCHAR(20)"`
	Score          int       `json:"score" xorm:"not null default 0 comment('分数（0-5）') TINYINT(4)"`
	StatusDelete   int       `json:"status_delete" xorm:"not null default 0 comment('删除状态（0:未删除,1:已删除）') TINYINT(4)"`
	DeleteTime     time.Time `json:"delete_time" xorm:"comment('删除时间') TIMESTAMP"`
	CreateTime     time.Time `json:"create_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') TIMESTAMP"`
	UpdateTime     time.Time `json:"update_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') TIMESTAMP"`
}
