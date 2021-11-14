package application

import (
	"js_statistics/app/models/internal/common"
	"js_statistics/app/models/tables"
)

type JsPrimary struct {
	ID    uint   `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	Title string `gorm:"column:title;type:varchar(50);not null;comment:标题"`
	common.Base
}

func (JsPrimary) TableName() string {
	return tables.JsPrimary
}
