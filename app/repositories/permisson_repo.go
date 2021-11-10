package repositories

import (
	"js_statistics/app/models"
	"js_statistics/app/response"
	"js_statistics/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	permissionRepoInstace PermissionRepo
	permissionOnce        sync.Once
)

type PermissionRepoImpl struct{}

func GetPermissionRepo() PermissionRepo {
	permissionOnce.Do(func() {
		permissionRepoInstace = &PermissionRepoImpl{}
	})
	return permissionRepoInstace
}

type PermissionRepo interface {
	Create(db *gorm.DB, p *models.Permission) exception.Exception
	Get(db *gorm.DB, id uint) (*models.Permission, exception.Exception)
	GetAll(db *gorm.DB) ([]models.Permission, exception.Exception)
	GetTop(db *gorm.DB) (*models.Permission, exception.Exception)
	Update(db *gorm.DB, id uint, param map[string]interface{}) exception.Exception
	Delete(db *gorm.DB, ids []uint) exception.Exception
}

func (u *PermissionRepoImpl) Create(db *gorm.DB, p *models.Permission) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(p).Error)
}

func (u *PermissionRepoImpl) Get(db *gorm.DB, id uint) (*models.Permission, exception.Exception) {
	p := models.Permission{}
	res := db.Where(&models.Permission{ID: id}).Find(&p)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &p, nil
}

func (u *PermissionRepoImpl) GetAll(db *gorm.DB) ([]models.Permission, exception.Exception) {
	tp := make([]models.Permission, 0)
	tx := db.Find(&tp)
	if tx.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return tp, nil
}

func (u *PermissionRepoImpl) GetTop(db *gorm.DB) (*models.Permission, exception.Exception) {
	tp := models.Permission{}
	tx := db.Where(&models.Permission{ParentID: 0}).Find(&tp)
	if tx.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if tx.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return &tp, nil
}

func (u *PermissionRepoImpl) Update(db *gorm.DB, id uint, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.Permission{}).Where(&models.Permission{ID: id}).Updates(param).Error)
}

func (u *PermissionRepoImpl) Delete(db *gorm.DB, ids []uint) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.Permission{}, ids).Error)
}
