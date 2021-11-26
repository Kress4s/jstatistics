package application

// import (
// 	"js_statistics/app/models/internal/common"
// 	"js_statistics/app/models/tables"
// )

// type Object struct {
// 	ID       string `gorm:"column:id;primaryKey;type:varchar(40);not null;comment:文件ID"`
// 	Type     int    `gorm:"column:type;type:integer;comment:伪装内容类型;0:文本,1:图片;2:音频，3:视频"`
// 	TextType int    `gorm:"column:type;type:integer;comment:伪装内容类型;0:plain,1:html;2:xml,3:json"`
// 	Filename string `gorm:"column:filename;type:varchar(255);not null;comment:文件名称"`
// 	Path     string `gorm:"column:path;type:varchar(255);not null;comment:文件路径"`
// 	Buff     []byte `gorm:"-"`
// 	Size     int64  `gorm:"column:size;type:bigint;not null;comment:文件大小"`
// 	common.Base
// }

// func (Object) TableName() string {
// 	return tables.Object
// }
