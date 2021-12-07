package permission

import (
	"js_statistics/app/models/internal/common"
	"js_statistics/app/models/tables"
)

type UserPrimaryRelation struct {
	common.Base `gorm:"embedded"`
	UserID      int64 `gorm:"column:user_id;primaryKey;type:bigint;not null;comment:用户id"`
	PrimaryID   int64 `gorm:"column:primary_id;primaryKey;type:bigint;not null;comment:js主分类id"`
}

func (UserPrimaryRelation) TableName() string {
	return tables.UserPrimaryRelation
}
