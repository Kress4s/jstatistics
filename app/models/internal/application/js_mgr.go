package application

import (
	"js_statistics/app/models/internal/common"
	"js_statistics/app/models/tables"
	"js_statistics/types"
)

type JsManage struct {
	common.Base
	Title         string            `gorm:"column:title;type:varchar(50);not null;comment:标题"`
	ShieldArea    string            `gorm:"column:shield_area;type:varchar(300);comment:屏蔽地区"`
	Sign          string            `gorm:"column:sign;type:varchar(50);not null;comment:js标志字符串"`
	HrefID        string            `gorm:"column:href_id;type:varchar(200);comment:href跳转id"`
	KeyWord       string            `gorm:"column:key_word;type:varchar(30);comment:关键词"`
	ClientType    types.BigintArray `gorm:"column:client_type;type:varchar(20)[];comment:客户端"`
	SearchEngines types.BigintArray `gorm:"column:search_engines;type:varchar(20)[];comment:搜索引擎"`
	FromMode      int               `gorm:"column:from_mode;type:integer;comment:来源类型"`
	RedirectMode  int               `gorm:"column:redirect_mode;type:integer;not null;comment:跳转方式"`
	RedirectCode  int               `gorm:"column:redirect_code;type:integer;comment:跳转代码"`
	ReleaseTime   int               `gorm:"column:release_time;type:integer;not null;comment:封禁时间"`
	WaitTime      int               `gorm:"column:wait_time;type:integer;comment:跳转等待时间"`
	CategoryID    int64             `gorm:"column:category_id;type:bigint;not null;comment:js分类id"`
	RedirectCount int               `gorm:"column:redirect_count;type:integer;not null;comment:跳转次数"`
	ID            int64             `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	Status        bool              `gorm:"column:status;type:boolean;default:true;comment:状态"`
}

type JsManageListView struct {
	common.Base
	Title         string            `gorm:"column:title"`
	ShieldArea    string            `gorm:"column:shield_area"`
	Sign          string            `gorm:"column:sign"`
	HrefID        string            `gorm:"column:href_id"`
	KeyWord       string            `gorm:"column:key_word"`
	ClientType    types.BigintArray `gorm:"column:client_type"`
	SearchEngines types.BigintArray `gorm:"column:search_engines"`
	FromMode      int               `gorm:"column:from_mode"`
	RedirectMode  int               `gorm:"column:redirect_mode"`
	RedirectCode  int               `gorm:"column:redirect_code"`
	ReleaseTime   int               `gorm:"column:release_time"`
	WaitTime      int               `gorm:"column:wait_time"`
	CategoryID    int64             `gorm:"column:category_id"`
	RedirectCount int               `gorm:"column:redirect_count"`
	ID            int64             `gorm:"column:id;primaryKey"`
	IPCount       int64             `gorm:"column:ip_count"`
	Status        bool              `gorm:"column:status"`
}

func (JsManage) TableName() string {
	return tables.JsManage
}
