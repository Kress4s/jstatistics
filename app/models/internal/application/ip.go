package application

import (
	"js_statistics/app/models/internal/common"
	"js_statistics/app/models/tables"
)

type WhiteIP struct {
	common.Base `gorm:"embedded"`
	Title       string `gorm:"column:title;type:varchar(50);not null;comment:标题"`
	IP          string `gorm:"column:ip;type:varchar(50);uniqueIndex;not null;comment:ip"`
	ID          int64  `gorm:"column:id;primaryKey;unique;not null;comment:id"`
}

func (WhiteIP) TableName() string {
	return tables.WhiteIP
}
