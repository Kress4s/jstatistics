package permission

import (
	"js_statistics/app/models/internal/common"
	"js_statistics/app/models/tables"
)

type RolePermissionRelation struct {
	PermissionID uint `gorm:"column:permission_id;primaryKey;type:bigint;not null;comment:权限id"`
	RoleID       uint `gorm:"column:role_id;primaryKey;type:bigint;not null;comment:角色id"`
	common.Base  `gorm:"embedded"`
}

func (RolePermissionRelation) TableName() string {
	return tables.RolePermissionRelation
}
