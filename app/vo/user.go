package vo

import "js_statistics/app/models"

type UserReq struct {
	// 用户名
	UserName string `json:"user_name"`
	// 密码
	Password string `json:"password"`
	// 是否管理员
	IsAdmin bool `json:"is_admin"`
	// 状态
	Status bool `json:"status"`
}

type UserUpdateReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
	Status   bool   `json:"status"`
}

func (u UserReq) ToModel(openID string) models.User {
	return models.User{
		Username: u.UserName,
		Password: u.Password,
		IsAdmin:  u.IsAdmin,
		Status:   u.Status,
		Base: models.Base{
			CreateBy: openID,
			UpdateBy: openID,
		},
	}
}

type LoginReq struct {
	// 用户名
	UserName string `json:"user_name"`
	// 密码
	Password string `json:"password"`
}

type ProfileResp struct {
	Name  string `json:"name"`
	ID    int64  `json:"id"`
	Admin bool   `json:"admin"`
}

type UserResp struct {
	UserName string `json:"user_name"`
	ID       int64  `json:"id"`
	Admin    bool   `json:"admin"`
	Status   bool   `json:"status"`
}

type UserToMenusResp struct {
	MenuName string `json:"menu_name"`
	Route    string `json:"router"`
	MenuID   int64  `json:"menu_id"`
	Identify string `json:"identify"`
}
