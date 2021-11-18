package dash_board

import "time"

type Statistics struct {
	ID         int64     `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	JsID       int64     `gorm:"column:js_id;type:bigint;not null;comment:js管理的id"`
	CategoryID int64     `gorm:"column:category_id;type:bigint;not null;comment:js子分类的id"`
	PrimaryID  int64     `gorm:"column:primary_id;type:bigint;not null;comment:js分类的id"`
	VisitType  int       `gorm:"column:visit_type;type:integer;not null;comment:访问类型0:ip; 1:uv"`
	IP         string    `gorm:"column:ip;type:varchar(50);not null;comment:ip地址"`
	VisitTime  time.Time `gorm:"column:create_at;type:date;not null;comment:访问时间"`
}
