package repositories

import (
	"js_statistics/app/models"
	"js_statistics/app/response"
	"js_statistics/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	userCategoryRepoInstance UserCategoryRepo
	userCategoryOnce         sync.Once
)

type UserCategoryRepoImpl struct{}

func GetUserCategoryRepo() UserCategoryRepo {
	userCategoryOnce.Do(func() {
		userCategoryRepoInstance = &UserCategoryRepoImpl{}
	})
	return userCategoryRepoInstance
}

type UserCategoryRepo interface {
	Create(db *gorm.DB, rps []models.UserCategoryRelation) exception.Exception
	GetByUserID(db *gorm.DB, userID int64) ([]models.UserCategoryRelation, exception.Exception)
	DeleteByUserID(db *gorm.DB, userID int64) exception.Exception
	DeleteByCategoryID(db *gorm.DB, roleID int64) exception.Exception
}

func (uri *UserCategoryRepoImpl) Create(db *gorm.DB, urs []models.UserCategoryRelation) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(&urs).Error)
}

func (uri *UserCategoryRepoImpl) DeleteByUserID(db *gorm.DB, userID int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Where("user_id = ?", userID).Delete(models.UserCategoryRelation{}).Error)
}

func (uri *UserCategoryRepoImpl) DeleteByCategoryID(db *gorm.DB, roleID int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Where("category_id = ?", roleID).Delete(models.UserCategoryRelation{}).Error)
}

func (uri *UserCategoryRepoImpl) GetByUserID(db *gorm.DB, userID int64) ([]models.UserCategoryRelation,
	exception.Exception) {
	urs := make([]models.UserCategoryRelation, 0)
	tx := db.Where(&models.UserCategoryRelation{UserID: userID}).Find(&urs)
	if tx.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return urs, nil
}
