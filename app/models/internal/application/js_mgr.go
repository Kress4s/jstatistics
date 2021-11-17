package application

import (
	"js_statistics/app/models/internal/common"
	"js_statistics/app/models/tables"
	"js_statistics/types"
)

type JsManage struct {
	ID            int64             `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	Title         string            `gorm:"column:title;type:varchar(50);not null;comment:标题"`
	ShieldArea    string            `gorm:"column:shield_area;type:varchar(30);comment:屏蔽地区"`
	ClientType    types.BigintArray `gorm:"column:client_type;type:varchar(20)[];comment:客户端"`
	RedirectCount int               `gorm:"column:redirect_count;type:integer;not null;comment:跳转次数"`
	ReleaseTime   int               `gorm:"column:release_time;type:integer;not null;comment:封禁时间"`
	FromMode      int               `gorm:"column:from_mode;type:integer;comment:来源类型"`
	KeyWord       string            `gorm:"column:key_word;type:varchar(30);comment:关键词"`
	SearchEngines types.BigintArray `gorm:"column:search_engines;type:varchar(20)[];comment:搜索引擎"`
	RedirectMode  int               `gorm:"column:redirect_mode;type:integer;not null;comment:跳转方式"`
	RedirectCode  int               `gorm:"column:redirect_code;type:integer;comment:跳转代码"`
	HrefID        string            `gorm:"column:href_id;type:varchar(30);comment:href跳转id"`
	WaitTime      int               `gorm:"column:wait_time;type:integer;comment:跳转等待时间"`
	Status        bool              `gorm:"column:status;type:boolean;default:true;comment:状态"`
	CategoryID    int64             `gorm:"column:category_id;type:bigint;not null;comment:js分类id"`
	Sign          string            `gorm:"column:sign;type:varchar(50);not null;comment:js标志字符串"`
	common.Base
}

func (JsManage) TableName() string {
	return tables.JsManage
}
