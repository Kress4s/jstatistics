package permission

import (
	"js_statistics/app/models/internal/common"
	"js_statistics/app/models/tables"
)

type Permission struct {
	ID          uint   `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	Name        string `gorm:"column:name;type:varchar(50);not null;comment:名称"`
	MenuName    string `gorm:"column:menu_name;type:varchar(50);not null;comment:菜单名称"`
	Route       string `gorm:"column:route;type:varchar(80);not null;comment:路由"`
	Identify    string `gorm:"column:identify;type:varchar(80);not null;comment:权限表示"`
	Type        int    `gorm:"column:type;type:smallint;not null;comment:权限类型(0: 菜单权限  1:操作权限)"`
	Index       int    `gorm:"column:index;type:smallint;comment:排序"`
	ParentID    uint   `gorm:"column:parent_id;not null;comment:父级id"`
	common.Base `gorm:"embedded"`
}

func (Permission) TableName() string {
	return tables.Permission
}
