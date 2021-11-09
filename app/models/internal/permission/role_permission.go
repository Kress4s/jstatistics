package permission

import "js_statistics/app/models/tables"

type RolePermissionRelation struct {
	PermissionID int64 `gorm:"column:permission_id;primaryKey;type:bigint;not null;comment:权限id"`
	RoleID       int64 `gorm:"column:role_id;primaryKey;type:bigint;not null;comment:角色id"`
}

func (RolePermissionRelation) TableName() string {
	return tables.RolePermissionRelation
}
