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
	roleRepoInstance RoleRepo
	roleOnce         sync.Once
)

type RoleRepoImpl struct{}

func GetRoleRepo() RoleRepo {
	roleOnce.Do(func() {
		roleRepoInstance = &RoleRepoImpl{}
	})
	return roleRepoInstance
}

type RoleRepo interface {
	Creat(db *gorm.DB, role *models.Role) exception.Exception
	Get(db *gorm.DB, id uint) (*models.Role, exception.Exception)
	List(db *gorm.DB, page *vo.PageInfo) (int64, []models.Role, exception.Exception)
	Update(db *gorm.DB, id uint, param map[string]interface{}) exception.Exception
	Delete(db *gorm.DB, id uint) exception.Exception
	GetByIDs(db *gorm.DB, ids []uint) ([]models.Role, exception.Exception)
}

func (rri *RoleRepoImpl) Creat(db *gorm.DB, role *models.Role) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(role).Error)
}

func (rri *RoleRepoImpl) Get(db *gorm.DB, id uint) (*models.Role, exception.Exception) {
	role := models.Role{}
	res := db.Where(&models.Role{ID: id}).Find(&role)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &role, nil
}

func (rri *RoleRepoImpl) List(db *gorm.DB, pageInfo *vo.PageInfo) (int64, []models.Role, exception.Exception) {
	roles := make([]models.Role, 0)
	tx := db.Table(tables.Role)
	if pageInfo.Keywords != "" {
		tx = tx.Scopes(vo.FuzzySearch(pageInfo.Keywords, "name", "identify"))
	}
	tx.Order("id").Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).Find(&roles)
	count := int64(0)
	res := tx.Limit(-1).Offset(-1).Count(&count)
	return count, roles, exception.Wrap(response.ExceptionDatabase, res.Error)
}

func (rri *RoleRepoImpl) Update(db *gorm.DB, id uint, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.Role{}).Where(&models.Role{ID: id}).Updates(param).Error)
}

func (rri *RoleRepoImpl) Delete(db *gorm.DB, id uint) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.Role{}, id).Error)
}

func (rri *RoleRepoImpl) GetByIDs(db *gorm.DB, ids []uint) ([]models.Role, exception.Exception) {
	roles := make([]models.Role, 0, len(ids))
	tx := db.Find(&roles, ids)
	if tx.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return roles, nil
}
