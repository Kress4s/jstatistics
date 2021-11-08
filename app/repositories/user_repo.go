package repositories

import (
	"js_statistics/app/models"
	"js_statistics/app/response"
	"js_statistics/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	userRepoInstace UserRepo
	userOnce        sync.Once
)

type UserRepoImpl struct{}

func GetUserRepo() UserRepo {
	userOnce.Do(func() {
		userRepoInstace = &UserRepoImpl{}
	})
	return userRepoInstace
}

type UserRepo interface {
	Profile(db *gorm.DB, id uint) (*models.User, exception.Exception)
	CheckPassword(db *gorm.DB, account, password string) (bool, uint, exception.Exception)
	Create(db *gorm.DB, user *models.User) exception.Exception
}

func (u *UserRepoImpl) Profile(db *gorm.DB, id uint) (*models.User, exception.Exception) {
	user := models.User{}
	res := db.Where(&models.User{ID: id}).Find(&user)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &user, nil
}

func (u *UserRepoImpl) CheckPassword(db *gorm.DB, username, password string) (bool, uint, exception.Exception) {
	user := &models.User{}
	res := db.Where(&models.User{Username: username, Password: password}).Find(user)
	if res.Error != nil {
		return false, 0, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	if res.RowsAffected == 0 {
		return false, 0, exception.New(response.ExceptionInvalidUserPassword, "user or password is wrong")
	}
	return true, user.ID, nil
}

func (u *UserRepoImpl) Create(db *gorm.DB, user *models.User) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(user).Error)
}
