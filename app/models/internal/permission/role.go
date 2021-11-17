package permission

import (
	"js_statistics/app/models/internal/common"
	"js_statistics/app/models/tables"
)

type Role struct {
	ID          int64  `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	Name        string `gorm:"column:name;type:varchar(50);not null;comment:角色名"`
	Identify    string `gorm:"column:identify;type:varchar(60);not null;comment:标识符"`
	Description string `gorm:"column:description;type:varchar(60);comment:标识符"`
	common.Base `gorm:"embedded"`
}

func (Role) TableName() string {
	return tables.Role
}
