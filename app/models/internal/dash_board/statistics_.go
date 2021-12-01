package dash_board

import (
	"js_statistics/app/models/tables"
	"js_statistics/app/models/views"
	"time"
)

type IPStatistics struct {
	VisitTime  time.Time `gorm:"column:visit_time;type:date;index;not null;comment:访问时间"`
	IP         string    `gorm:"column:ip;type:varchar(50);not null;comment:ip地址"`
	ID         int64     `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	JsID       int64     `gorm:"column:js_id;type:bigint;index:i_p_c_j,priority:3;not null;comment:js管理的id"`
	CategoryID int64     `gorm:"column:category_id;type:bigint;index:i_p_c_j,priority:2;not null;comment:js子分类的id"`
	PrimaryID  int64     `gorm:"column:primary_id;type:bigint;index:i_p_c_j,priority:1;not null;comment:js分类的id"`
}

func (IPStatistics) TableName() string {
	return tables.IPStatistics
}

type UVStatistics struct {
	VisitTime  time.Time `gorm:"column:visit_time;type:date;index;not null;comment:访问时间"`
	IP         string    `gorm:"column:ip;type:varchar(50);not null;comment:ip地址"`
	Cookie     string    `gorm:"column:cookie;type:varchar(50);not null;comment:cookie"`
	JsID       int64     `gorm:"column:js_id;type:bigint;index:u_p_c_j,priority:3;not null;comment:js管理的id"`
	CategoryID int64     `gorm:"column:category_id;type:bigint;index:u_p_c_j,priority:2;not null;comment:js子分类的id"`
	PrimaryID  int64     `gorm:"column:primary_id;type:bigint;index:u_p_c_j,priority:1;not null;comment:js分类的id"`
	ID         int64     `gorm:"column:id;primaryKey;unique;not null;comment:id"`
}

func (UVStatistics) TableName() string {
	return tables.UVStatistics
}

type IPRecode struct {
	VisitTime  time.Time `gorm:"column:visit_time;type:date;index;not null;comment:访问时间"`
	ToURL      string    `gorm:"column:to_url;type:varchar(100);not null;comment:去路url"`
	IP         string    `gorm:"column:ip;type:varchar(50);not null;comment:ip地址"`
	Region     string    `gorm:"column:region;type:varchar(30);comment:地域名称"`
	FromURL    string    `gorm:"column:from_url;type:varchar(100);not null;comment:来路url"`
	RegionCode string    `gorm:"column:region_code;type:varchar(50);comment:地域编码"`
	CategoryID int64     `gorm:"column:category_id;type:bigint;index:r_p_c_j,priority:2;not null;comment:js子分类的id"`
	ID         int64     `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	JsID       int64     `gorm:"column:js_id;type:bigint;index:r_p_c_j,priority:3;not null;comment:js管理的id"`
	VisitType  int       `gorm:"column:visit_type;type:integer;not null;comment:访问类型 0：ip, 1:UV"`
	PrimaryID  int64     `gorm:"column:primary_id;type:bigint;index:r_p_c_j,priority:1;not null;comment:js分类的id"`
}

func (IPRecode) TableName() string {
	return tables.IPRecode
}

type IPVisitStatistic struct {
	VisitTime time.Time `gorm:"column:visit_time"`
	Count     int64     `gorm:"column:count"`
}

type UVisitStatistic struct {
	VisitTime time.Time `gorm:"column:visit_time"`
	Count     int64     `gorm:"column:count"`
}

// ip地区分布统计
type RegionStatistic struct {
	Region string `gorm:"column:region"`
	Count  int64  `gorm:"column:count"`
}

// JS流量统计
type JSVisitStatistic struct {
	Title string `gorm:"column:title"`
	JsID  int64  `gorm:"column:js_id"`
	Count int64  `gorm:"column:count"`
}

// 数据统计 -- 流量统计 -- 流量数据
type FlowDataView struct {
	UVTime     time.Time `gorm:"column:uv_time"`
	IPTime     time.Time `gorm:"column:ip_time"`
	Title      string    `gorm:"column:title"`
	PrimaryID  int64     `gorm:"column:primary_id"`
	IPCount    int64     `gorm:"column:ip_count"`
	UVCount    int64     `gorm:"column:uv_count"`
	JSid       int64     `gorm:"column:js_id"`
	CategoryID int64     `gorm:"column:category_id"`
}

func (FlowDataView) TableName() string {
	return views.FlowDataView
}

type FlowDataStatistic struct {
	Title   string `gorm:"column:title"`
	IPCount int64  `gorm:"column:ip_count"`
	UVCount int64  `gorm:"column:uv_count"`
}

type FromAnalysisView struct {
	Title   string `gorm:"title"`
	FromURL string `gorm:"from_url"`
	ToUrl   string `gorm:"to_url"`
	Count   int64  `gorm:"count"`
}

// type IPFlowDataStatisticView struct {

// }
