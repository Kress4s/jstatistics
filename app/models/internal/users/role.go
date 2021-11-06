package users

import "js_statistics/app/models/tables"

type Role struct {
	ID          uint   `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	Name        string `gorm:"column:name;type:varchar(50);not null;comment:角色名"`
	Identify    string `gorm:"column:identify;type:varchar(60);not null;comment:标识符"`
	Description string `gorm:"column:description;type:varchar(60);not null;comment:标识符"`
}

func (Role) TableName() string {
	return tables.Role
}
