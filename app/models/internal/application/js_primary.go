package application

import (
	"js_statistics/app/models/internal/common"
	"js_statistics/app/models/tables"
)

type JsPrimary struct {
	common.Base
	Title string `gorm:"column:title;type:varchar(50);not null;comment:标题"`
	ID    int64  `gorm:"column:id;primaryKey;unique;not null;comment:id"`
}

func (JsPrimary) TableName() string {
	return tables.JsPrimary
}

type AllsCategory struct {
	ID     int64   `gorm:"column:id"`
	Title  string  `gorm:"column:title"`
	Cid    *int64  `gorm:"column:cid"`
	CTitle *string `gorm:"column:c_title"`
	Pid    *int64  `gorm:"column:pid"`
}
