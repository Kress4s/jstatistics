package permission

import "js_statistics/app/models/tables"

type UserRoleRelation struct {
	UserID int64 `gorm:"column:user_id;primaryKey;type:bigint;not null;comment:用户id"`
	RoleID int64 `gorm:"column:role_id;primaryKey;type:bigint;not null;comment:角色id"`
}

func (UserRoleRelation) TableName() string {
	return tables.UserRoleRelation
}
