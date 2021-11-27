package common

import (
	"js_statistics/app/models/tables"
	"time"
)

type SystemLog struct {
	ID          int64     `gorm:"column:id;primaryKey;unique;not null;comment:系统操作日志ID"`
	UserName    string    `gorm:"column:username;type:varchar(30);not null;comment:用户"`
	IP          string    `gorm:"column:ip;type:varchar(30);not null;comment:操作ip"`
	Address     string    `gorm:"column:address;type:varchar(60);not null;comment:操作地址"`
	Description string    `gorm:"column:description;type:varchar(60);not null;comment:描述"`
	OperateAt   time.Time `gorm:"column:operate_at;type:timestamptz;not null;comment:操作时间"`
}

func (SystemLog) TableName() string {
	return tables.SystemLog
}
