package application

import (
	"js_statistics/app/models/internal/common"
	"js_statistics/app/models/tables"
)

type JsCategory struct {
	ID        int64  `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	Title     string `gorm:"column:title;type:varchar(50);not null;comment:标题"`
	PrimaryID int64  `gorm:"column:primary_id;type:bigint;not null;comment:js主分类id"`
	DomainID  int64  `gorm:"column:domain_id;type:bigint;comment:域名id"`
	common.Base
}

func (JsCategory) TableName() string {
	return tables.JsCategory
}
