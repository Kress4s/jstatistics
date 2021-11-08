package service

import (
	repositories "js_statistics/app/repositories"
	"js_statistics/app/vo"
	"js_statistics/commom/drivers/database"
	"js_statistics/commom/tools"
	"js_statistics/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	userServiceInstance UserService
	userOnce            sync.Once
)

type userServiceImpl struct {
	db   *gorm.DB
	repo repositories.UserRepo
}

type UserService interface {
	Profile(id uint) (*vo.Profile, exception.Exception)
	Create(params *vo.UserReq) exception.Exception
}

func GetUserService() UserService {
	userOnce.Do(func() {
		userServiceInstance = &userServiceImpl{
			db:   database.GetDriver(),
			repo: repositories.GetUserRepo(),
		}
	})
	return userServiceInstance
}

func (us *userServiceImpl) Profile(id uint) (*vo.Profile, exception.Exception) {
	user, ex := us.repo.Profile(us.db, id)
	if ex != nil {
		return nil, ex
	}
	return &vo.Profile{
		ID:    user.ID,
		Name:  user.Username,
		Admin: user.IsAdmin,
	}, nil
}

func (us *userServiceImpl) Create(params *vo.UserReq) exception.Exception {
	// password
	params.Password = string(tools.Base64Encode([]byte(params.Password)))
	user := params.ToModel()
	return us.repo.Create(us.db, &user)
}
