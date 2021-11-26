package common

import "js_statistics/app/models/tables"

// Object 文件对象定义
type Object struct {
	Base     `gorm:"embedded"`
	ID       string `gorm:"column:id;primaryKey;type:varchar(40);not null;comment:文件ID"`
	Filename string `gorm:"column:filename;type:varchar(255);not null;comment:文件名称"`
	Path     string `gorm:"column:path;type:varchar(255);not null;comment:文件路径"`
	Buff     []byte `gorm:"-"`
	Size     int64  `gorm:"column:size;type:bigint;not null;comment:文件大小"`
}

func (Object) TableName() string {
	return tables.Object
}
