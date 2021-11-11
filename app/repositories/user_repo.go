package repositories

import (
	"js_statistics/app/models"
	"js_statistics/app/models/tables"
	"js_statistics/app/response"
	"js_statistics/app/vo"
	"js_statistics/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	userRepoInstance UserRepo
	userOnce         sync.Once
)

type UserRepoImpl struct{}

func GetUserRepo() UserRepo {
	userOnce.Do(func() {
		userRepoInstance = &UserRepoImpl{}
	})
	return userRepoInstance
}

type UserRepo interface {
	Profile(db *gorm.DB, id uint) (*models.User, exception.Exception)
	CheckPassword(db *gorm.DB, account, password string) (bool, uint, exception.Exception)
	Create(db *gorm.DB, user *models.User) exception.Exception
	Get(db *gorm.DB, id uint) (*models.User, exception.Exception)
	List(db *gorm.DB, pageInfo *vo.PageInfo) (int64, []models.User, exception.Exception)
	Update(db *gorm.DB, id uint, param map[string]interface{}) exception.Exception
	Delete(db *gorm.DB, id uint) exception.Exception
	MultiDelete(db *gorm.DB, ids []uint) exception.Exception
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

func (u *UserRepoImpl) Get(db *gorm.DB, id uint) (*models.User, exception.Exception) {
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

func (u *UserRepoImpl) List(db *gorm.DB, pageInfo *vo.PageInfo) (int64, []models.User, exception.Exception) {
	users := make([]models.User, 0)
	tx := db.Table(tables.User)
	if pageInfo.Keywords != "" {
		tx = tx.Scopes(vo.FuzzySearch(pageInfo.Keywords, "user_name"))
	}
	tx.Order("id").Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).Find(&users)
	count := int64(0)
	res := tx.Limit(-1).Offset(-1).Count(&count)
	return count, users, exception.Wrap(response.ExceptionDatabase, res.Error)
}

func (u *UserRepoImpl) Update(db *gorm.DB, id uint, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.User{}).Where(&models.User{ID: id}).Updates(param).Error)
}

func (u *UserRepoImpl) Delete(db *gorm.DB, id uint) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.User{}, id).Error)
}

func (u *UserRepoImpl) MultiDelete(db *gorm.DB, ids []uint) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.User{}, ids).Error)
}
