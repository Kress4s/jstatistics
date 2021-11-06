package vo

import "js_statistics/app/models"

type UserReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func (u UserReq) ToModel() models.User {
	return models.User{
		Username: u.UserName,
		Password: u.Password,
	}
}

type LoginReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
