package permission

import (
	"js_statistics/app/models/internal/common"
	"js_statistics/app/models/tables"
)

type UserRoleRelation struct {
	UserID      uint `gorm:"column:user_id;primaryKey;type:bigint;not null;comment:用户id"`
	RoleID      uint `gorm:"column:role_id;primaryKey;type:bigint;not null;comment:角色id"`
	common.Base `gorm:"embedded"`
}

func (UserRoleRelation) TableName() string {
	return tables.UserRoleRelation
}
