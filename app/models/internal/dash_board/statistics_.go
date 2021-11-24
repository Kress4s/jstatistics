package dash_board

import (
	"js_statistics/app/models/tables"
	"js_statistics/app/models/views"
	"time"
)

type IPStatistics struct {
	ID         int64     `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	JsID       int64     `gorm:"column:js_id;type:bigint;index:i_p_c_j,priority:3;not null;comment:js管理的id"`
	CategoryID int64     `gorm:"column:category_id;type:bigint;index:i_p_c_j,priority:2;not null;comment:js子分类的id"`
	PrimaryID  int64     `gorm:"column:primary_id;type:bigint;index:i_p_c_j,priority:1;not null;comment:js分类的id"`
	IP         string    `gorm:"column:ip;type:varchar(50);not null;comment:ip地址"`
	VisitTime  time.Time `gorm:"column:visit_time;type:date;index;not null;comment:访问时间"`
}

func (IPStatistics) TableName() string {
	return tables.IPStatistics
}

type UVStatistics struct {
	ID         int64     `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	JsID       int64     `gorm:"column:js_id;type:bigint;index:u_p_c_j,priority:3;not null;comment:js管理的id"`
	CategoryID int64     `gorm:"column:category_id;type:bigint;index:u_p_c_j,priority:2;not null;comment:js子分类的id"`
	PrimaryID  int64     `gorm:"column:primary_id;type:bigint;index:u_p_c_j,priority:1;not null;comment:js分类的id"`
	IP         string    `gorm:"column:ip;type:varchar(50);not null;comment:ip地址"`
	Cookie     string    `gorm:"column:cookie;type:varchar(50);not null;comment:cookie"`
	VisitTime  time.Time `gorm:"column:visit_time;type:date;index;not null;comment:访问时间"`
}

func (UVStatistics) TableName() string {
	return tables.UVStatistics
}

type IPRecode struct {
	ID         int64     `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	IP         string    `gorm:"column:ip;type:varchar(50);not null;comment:ip地址"`
	JsID       int64     `gorm:"column:js_id;type:bigint;index:r_p_c_j,priority:3;not null;comment:js管理的id"`
	CategoryID int64     `gorm:"column:category_id;type:bigint;index:r_p_c_j,priority:2;not null;comment:js子分类的id"`
	PrimaryID  int64     `gorm:"column:primary_id;type:bigint;index:r_p_c_j,priority:1;not null;comment:js分类的id"`
	FromURL    string    `gorm:"column:from_url;type:varchar(100);not null;comment:来路url"`
	ToURL      string    `gorm:"column:to_url;type:varchar(100);not null;comment:去路url"`
	RegionCode string    `gorm:"column:region_code;type:varchar(50);comment:地域编码"`
	Region     string    `gorm:"column:region;type:varchar(30);comment:地域名称"`
	VisitType  int       `gorm:"column:visit_type;type:integer;not null;comment:访问类型 0：ip, 1:UV"`
	VisitTime  time.Time `gorm:"column:visit_time;type:date;index;not null;comment:访问时间"`
}

func (IPRecode) TableName() string {
	return tables.IPRecode
}

type IPVisitStatistic struct {
	Count     int64     `gorm:"column:count"`
	VisitTime time.Time `gorm:"column:visit_time"`
}

type UVisitStatistic struct {
	Count     int64     `gorm:"column:count"`
	VisitTime time.Time `gorm:"column:visit_time"`
}

// ip地区分布统计
type RegionStatistic struct {
	Region string `gorm:"column:region"`
	Count  int64  `gorm:"column:count"`
}

// JS流量统计
type JSVisitStatistic struct {
	JsID  int64  `gorm:"column:js_id"`
	Title string `gorm:"column:title"`
	Count int64  `gorm:"column:count"`
}

// 数据统计 -- 流量统计 -- 流量数据
type FlowDataView struct {
	JSid       int64     `gorm:"column:js_id"`
	CategoryID int64     `gorm:"column:category_id"`
	PrimaryID  int64     `gorm:"column:primary_id"`
	IPCount    int64     `gorm:"column:ip_count"`
	IPTime     time.Time `gorm:"column:ip_time"`
	UVCount    int64     `gorm:"column:uv_count"`
	UVTime     time.Time `gorm:"column:uv_time"`
	Title      string    `gorm:"column:title"`
}

func (FlowDataView) TableName() string {
	return views.FlowDataView
}

type FlowDataStatistic struct {
	IPCount int64  `gorm:"column:ip_count"`
	UVCount int64  `gorm:"column:uv_count"`
	Title   string `gorm:"column:title"`
}

type FromAnalysisView struct {
	Title   string `gorm:"title"`
	FromURL string `gorm:"from_url"`
	ToUrl   string `gorm:"to_url"`
	Count   int64  `gorm:"count"`
}
