package service

import (
	repositories "js_statistics/app/repositories"
	"js_statistics/app/vo"
	"js_statistics/commom/drivers/database"
	"js_statistics/commom/tools"
	"js_statistics/exception"
	"sync"
	"time"

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
	Profile(id uint) (*vo.ProfileResp, exception.Exception)
	Create(openID string, params *vo.UserReq) exception.Exception
	Get(id uint) (*vo.ProfileResp, exception.Exception)
	Update(openID string, id uint, params *vo.UserUpdateReq) exception.Exception
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

func (us *userServiceImpl) Profile(id uint) (*vo.ProfileResp, exception.Exception) {
	user, ex := us.repo.Profile(us.db, id)
	if ex != nil {
		return nil, ex
	}
	return &vo.ProfileResp{
		ID:    user.ID,
		Name:  user.Username,
		Admin: user.IsAdmin,
	}, nil
}

func (us *userServiceImpl) Create(openID string, params *vo.UserReq) exception.Exception {
	// password
	params.Password = string(tools.Base64Encode([]byte(params.Password)))
	user := params.ToModel(openID)
	return us.repo.Create(us.db, &user)
}

func (us *userServiceImpl) Get(id uint) (*vo.ProfileResp, exception.Exception) {
	user, ex := us.repo.Profile(us.db, id)
	if ex != nil {
		return nil, ex
	}
	return &vo.ProfileResp{
		ID:    user.ID,
		Name:  user.Username,
		Admin: user.IsAdmin,
	}, nil
}

func (us *userServiceImpl) Update(openID string, id uint, params *vo.UserUpdateReq) exception.Exception {
	r := make(map[string]interface{})
	// password is nil, declear not change
	if len(params.Password) != 0 {
		r["password"] = string(tools.Base64Encode([]byte(params.Password)))
	}
	r["user_name"] = params.UserName
	r["is_admin"] = params.IsAdmin
	r["update_by"] = openID
	r["update_at"] = time.Now()
	return us.repo.Update(us.db, id, r)
}
