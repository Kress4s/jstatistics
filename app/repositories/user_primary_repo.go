package repositories

import (
	"js_statistics/app/models"
	"js_statistics/app/response"
	"js_statistics/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	userPrimaryRepoInstance UserPrimaryRepo
	userPrimaryOnce         sync.Once
)

type UserPrimaryRepoImpl struct{}

func GetUserPrimaryRepo() UserPrimaryRepo {
	userPrimaryOnce.Do(func() {
		userPrimaryRepoInstance = &UserPrimaryRepoImpl{}
	})
	return userPrimaryRepoInstance
}

type UserPrimaryRepo interface {
	Create(db *gorm.DB, rps []models.UserPrimaryRelation) exception.Exception
	GetByUserID(db *gorm.DB, userID int64) ([]models.UserPrimaryRelation, exception.Exception)
	DeleteByUserID(db *gorm.DB, userID int64) exception.Exception
	DeleteByUsersID(db *gorm.DB, usersID ...int64) exception.Exception
	DeleteByPrimaryID(db *gorm.DB, pid int64) exception.Exception
	DeleteByPrimariesID(db *gorm.DB, pids ...int64) exception.Exception
}

func (uri *UserPrimaryRepoImpl) Create(db *gorm.DB, urs []models.UserPrimaryRelation) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(&urs).Error)
}

func (uri *UserPrimaryRepoImpl) DeleteByUserID(db *gorm.DB, userID int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Where("user_id = ?", userID).Delete(models.UserPrimaryRelation{}).Error)
}

func (uri *UserPrimaryRepoImpl) DeleteByUsersID(db *gorm.DB, usersID ...int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Where("user_id in (?)", usersID).Delete(models.UserPrimaryRelation{}).Error)
}

func (uri *UserPrimaryRepoImpl) DeleteByPrimaryID(db *gorm.DB, pid int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Where("primary_id = ?", pid).Delete(models.UserPrimaryRelation{}).Error)
}

func (uri *UserPrimaryRepoImpl) DeleteByPrimariesID(db *gorm.DB, pids ...int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Where("primary_id in (?)", pids).Delete(models.UserPrimaryRelation{}).Error)
}

func (uri *UserPrimaryRepoImpl) GetByUserID(db *gorm.DB, userID int64) ([]models.UserPrimaryRelation,
	exception.Exception) {
	urs := make([]models.UserPrimaryRelation, 0)
	tx := db.Where(&models.UserPrimaryRelation{UserID: userID}).Find(&urs)
	if tx.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return urs, nil
}
