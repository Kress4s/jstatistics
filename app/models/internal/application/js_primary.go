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
