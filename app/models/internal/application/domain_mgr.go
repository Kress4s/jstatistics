package application

import (
	"js_statistics/app/models/internal/common"
	"js_statistics/app/models/tables"
)

type DomainMgr struct {
	common.Base `gorm:"embedded"`
	Title       string `gorm:"column:title;type:varchar(50);not null;comment:标题"`
	Domain      string `gorm:"column:domain;type:varchar(50);uniqueIndex;not null;comment:域名"`
	Certificate string `gorm:"column:certificate;type:varchar(200);comment:证书"`
	SecretKey   string `gorm:"column:secret_key;type:varchar(200);comment:秘钥"`
	ID          int64  `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	SSL         bool   `gorm:"column:ssl;type:boolean;default:false;not null;comment:是否使用ssl"`
}

func (DomainMgr) TableName() string {
	return tables.DomainMgr
}
