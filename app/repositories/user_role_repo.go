package repositories

import (
	"js_statistics/app/models"
	"js_statistics/app/response"
	"js_statistics/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	userRoleRepoInstance UserRoleRepo
	userRoleOnce         sync.Once
)

type UserRoleRepoImpl struct{}

func GetUserRoleRepo() UserRoleRepo {
	userRoleOnce.Do(func() {
		userRoleRepoInstance = &UserRoleRepoImpl{}
	})
	return userRoleRepoInstance
}

type UserRoleRepo interface {
	Create(db *gorm.DB, rps []models.UserRoleRelation) exception.Exception
	GetByUserID(db *gorm.DB, userID int64) ([]models.UserRoleRelation, exception.Exception)
	DeleteByUserID(db *gorm.DB, userID int64) exception.Exception
	DeleteByRoleID(db *gorm.DB, roleID int64) exception.Exception
	DeleteByUsersID(db *gorm.DB, usersID ...int64) exception.Exception
	DeleteByRolesID(db *gorm.DB, rolesID ...int64) exception.Exception
}

func (uri *UserRoleRepoImpl) Create(db *gorm.DB, urs []models.UserRoleRelation) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(&urs).Error)
}

func (uri *UserRoleRepoImpl) DeleteByUserID(db *gorm.DB, userID int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Where("user_id = ?", userID).Delete(models.UserRoleRelation{}).Error)
}

func (uri *UserRoleRepoImpl) DeleteByUsersID(db *gorm.DB, usersID ...int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Where("user_id in (?)", usersID).Delete(models.UserRoleRelation{}).Error)
}

func (uri *UserRoleRepoImpl) DeleteByRoleID(db *gorm.DB, roleID int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Where("role_id = ?", roleID).Delete(models.UserRoleRelation{}).Error)
}

func (uri *UserRoleRepoImpl) DeleteByRolesID(db *gorm.DB, rolesID ...int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Where("role_id in (?)", rolesID).Delete(models.UserRoleRelation{}).Error)
}

func (uri *UserRoleRepoImpl) GetByUserID(db *gorm.DB, userID int64) ([]models.UserRoleRelation,
	exception.Exception) {
	urs := make([]models.UserRoleRelation, 0)
	tx := db.Where(&models.UserRoleRelation{UserID: userID}).Find(&urs)
	if tx.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return urs, nil
}
