package application

import (
	"js_statistics/app/models/tables"
	"time"
)

type RedirectLog struct {
	UpdateAt   time.Time `gorm:"column:update_at;type:timestamp;comment:更新时间"`
	OldPC      string    `gorm:"column:old_pc;type:varchar(100);comment:原pc跳转地址"`
	OldAndroid string    `gorm:"column:old_android;type:varchar(100);comment:原android跳转地址"`
	OldIOS     string    `gorm:"column:old_ios;type:varchar(100);comment:原ios跳转地址"`
	PC         string    `gorm:"column:pc;type:varchar(100);not null;comment:pc跳转地址"`
	IOS        string    `gorm:"column:ios;type:varchar(100);not null;comment:ios跳转地址"`
	Type       string    `gorm:"column:type;type:varchar(10);not null;comment:操作类型"`
	Android    string    `gorm:"column:android;type:varchar(100);not null;comment:android跳转地址"`
	ID         int64     `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	RedirectID int64     `gorm:"column:redirect_id;type:bigint;not null;comment:跳转管理的id"`
	CategoryID int64     `gorm:"column:category_id;type:bigint;not null;comment:js分类id"`
}

func (RedirectLog) TableName() string {
	return tables.RedirectLog
}
