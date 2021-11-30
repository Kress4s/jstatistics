package application

import (
	"js_statistics/app/models/internal/common"
	"js_statistics/app/models/tables"
)

type Faker struct {
	common.Base
	Text    string `gorm:"column:text;type:varchar(100);comment:文本内容"`
	ObjID   string `gorm:"column:obj_id;type:varchar(40);comment:文件ID"`
	ID      int64  `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	ReqType int    `gorm:"column:req_type;type:integer;comment:请求类型;0:text/html,1:text/plain;2:text/xml,3:application/json"`
	JsID    int64  `gorm:"column:js_id;type:bigint;not null;comment:js的id"`
	Type    int    `gorm:"column:type;type:integer;not null;comment:伪装内容类型;0:文本,1:图片;2:音频，3:视频"`
	Status  bool   `gorm:"column:status;type:boolean;default:true;not null;comment:状态"`
}

func (Faker) TableName() string {
	return tables.Faker
}
