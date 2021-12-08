package service

import (
	"js_statistics/app/repositories"
	"js_statistics/app/response"
	"js_statistics/app/vo"
	"js_statistics/commom/drivers/database"
	"js_statistics/commom/tools"
	"js_statistics/constant"
	"js_statistics/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	loginInstance LoginService
	loginOnce     sync.Once
)

type loginServiceImpl struct {
	db   *gorm.DB
	repo repositories.UserRepo
}

func GetLoginService() LoginService {
	loginOnce.Do(func() {
		loginInstance = &loginServiceImpl{
			db:   database.GetDriver(),
			repo: repositories.GetUserRepo(),
		}
	})
	return loginInstance
}

type LoginService interface {
	Login(username, password string) (*vo.LoginResponse, exception.Exception)
}

func (ls *loginServiceImpl) Login(username, password string) (*vo.LoginResponse, exception.Exception) {
	password = string(tools.Base64Encode([]byte(password)))
	ok, status, userID, ex := ls.repo.CheckPassword(ls.db, username, password)
	if ex != nil || !ok {
		return nil, ex
	}
	if !status {
		return nil, exception.New(response.ExceptionUserClose, "对不起 您登录权限已被关闭")
	}
	// token
	token, exp := tools.Token(userID, username)
	return &vo.LoginResponse{
		AccessToken: token,
		TokenType:   constant.Authorization,
		Expiry:      exp,
	}, nil
}
