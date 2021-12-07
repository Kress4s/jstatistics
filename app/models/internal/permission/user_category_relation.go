package permission

import (
	"js_statistics/app/models/internal/common"
	"js_statistics/app/models/tables"
)

type UserCategoryRelation struct {
	common.Base `gorm:"embedded"`
	UserID      int64 `gorm:"column:user_id;primaryKey;type:bigint;not null;comment:用户id"`
	CategoryID  int64 `gorm:"column:category_id;primaryKey;type:bigint;not null;comment:js分类id"`
}

func (UserCategoryRelation) TableName() string {
	return tables.UserCategoryRelation
}
