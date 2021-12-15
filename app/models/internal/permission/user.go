package permission

import (
	"js_statistics/app/models/internal/common"
	"js_statistics/app/models/tables"
)

type User struct {
	common.Base `gorm:"embedded"`
	Username    string `gorm:"column:user_name;type:varchar(50);not null;comment:用户名"`
	Password    string `gorm:"column:password;type:varchar(60);not null;comment:密码"`
	ID          int64  `gorm:"column:id;primaryKey;unique;not null;comment:id"`
	IsAdmin     bool   `gorm:"column:is_admin;type:boolean;not null;comment:是否是超管"`
	Status      bool   `gorm:"column:status;type:boolean;comment:状态"`
}

func (User) TableName() string {
	return tables.User
}

type UserToMenus struct {
	MenuName string `gorm:"column:menu_name"`
	Route    string `gorm:"column:router"`
	Identify string `gorm:"column:identify"`
	Type     int    `gorm:"column:type"`
	MenuID   int64  `gorm:"column:menu_id"`
	ParentID int64  `gorm:"column:parent_id"`
}

type UserListInfo struct {
	Username  string `gorm:"column:user_name"`
	RoleNames string `gorm:"column:role_names"`
	ID        int64  `gorm:"column:id"`
	IsAdmin   bool   `gorm:"column:is_admin;type:boolean;not null;comment:是否是超管"`
	Status    bool   `gorm:"column:status"`
}
