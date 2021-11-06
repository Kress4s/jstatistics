package users

import "js_statistics/app/models/tables"

type User struct {
	ID       uint   `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	Username string `gorm:"column:user_name;type:varchar(50);not null;comment:用户名"`
	Password string `gorm:"column:password;type:varchar(60);not null;comment:密码"`
	IsAdmin  bool   `gorm:"column:is_admin;type:boolean;not null;comment:是否是超管"`
}

func (User) TableName() string {
	return tables.User
}
