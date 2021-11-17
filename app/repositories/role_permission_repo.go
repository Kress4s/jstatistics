package repositories

import (
	"js_statistics/app/models"
	"js_statistics/app/response"
	"js_statistics/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	rolePermissionRepoInstance RolePermissionRepo
	rolePermissionOnce         sync.Once
)

type RolePermissionRepoImpl struct{}

func GetRolePermissionRepo() RolePermissionRepo {
	rolePermissionOnce.Do(func() {
		rolePermissionRepoInstance = &RolePermissionRepoImpl{}
	})
	return rolePermissionRepoInstance
}

type RolePermissionRepo interface {
	Create(db *gorm.DB, rps []models.RolePermissionRelation) exception.Exception
	DeleteByRoleID(db *gorm.DB, roleID int64) exception.Exception
	DeleteByPermissionID(db *gorm.DB, pID int64) exception.Exception
	GetByRoleID(db *gorm.DB, roleID int64) ([]models.RolePermissionRelation, exception.Exception)
}

func (rpr *RolePermissionRepoImpl) Create(db *gorm.DB, rps []models.RolePermissionRelation) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(&rps).Error)
}

func (rpr *RolePermissionRepoImpl) GetByRoleID(db *gorm.DB, roleID int64) ([]models.RolePermissionRelation,
	exception.Exception) {
	rps := make([]models.RolePermissionRelation, 0)
	tx := db.Where(&models.RolePermissionRelation{RoleID: roleID}).Find(&rps)
	if tx.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return rps, nil
}

func (rpr *RolePermissionRepoImpl) DeleteByRoleID(db *gorm.DB, roleID int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Where("role_id = ?", roleID).Delete(models.RolePermissionRelation{}).Error)
}

func (rpr *RolePermissionRepoImpl) DeleteByPermissionID(db *gorm.DB, pID int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Where("permission_id = ?", pID).Delete(models.RolePermissionRelation{}).Error)
}
