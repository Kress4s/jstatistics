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
	jscRepoInstance JscRepo
	jscOnce         sync.Once
)

type JscRepoImpl struct{}

func GetJscRepo() JscRepo {
	jscOnce.Do(func() {
		jscRepoInstance = &JscRepoImpl{}
	})
	return jscRepoInstance
}

type JscRepo interface {
	Create(db *gorm.DB, jsc *models.JsCategory) exception.Exception
	ListByPrimaryID(db *gorm.DB, pageInfo *vo.PageInfo, pid, userID int64, isAdmin bool) (int64, []models.JsCategory, exception.Exception)
	Get(db *gorm.DB, id int64) (*models.JsCategory, exception.Exception)
	Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
	Delete(db *gorm.DB, id int64) exception.Exception
	DeleteByPrimaryID(db *gorm.DB, pid int64) exception.Exception
	MultiDelete(db *gorm.DB, ids []int64) exception.Exception
	ListAllByPrimaryID(db *gorm.DB, pid int64) ([]models.JsCategory, exception.Exception)
	UpdateDomainIDToO(db *gorm.DB, domainIDs ...int64) exception.Exception
	GetByIDs(db *gorm.DB, ids []int64) ([]models.JsCategory, exception.Exception)
	ListAllByPrimaryUserID(db *gorm.DB, pid, userID int64, isAdmin bool) ([]models.JsCategory, exception.Exception)
}

func (jsi *JscRepoImpl) Create(db *gorm.DB, jsc *models.JsCategory) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(jsc).Error)
}

func (jsi *JscRepoImpl) ListByPrimaryID(db *gorm.DB, pageInfo *vo.PageInfo, pid, userID int64, isAdmin bool) (int64,
	[]models.JsCategory, exception.Exception) {
	jscs := make([]models.JsCategory, 0)
	var res *gorm.DB
	count := int64(0)
	if !isAdmin {
		sub := db.Table(tables.UserCategoryRelation).Where("user_id = ?", userID)
		tx := db.Table(tables.JsCategory+" AS jc").Where("jc.primary_id = ?", pid).
			Joins("INNER JOIN (?) AS uc ON jc.id = uc.category_id", sub)
		tx = tx.Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).Find(&jscs)
		res = tx.Limit(-1).Offset(-1).Count(&count)
	} else {
		tx := db.Table(tables.JsCategory)
		if pageInfo.Keywords != "" {
			tx = tx.Scopes(vo.FuzzySearch(pageInfo.Keywords, "title"))
		}
		tx.Where("primary_id = ?", pid).Order("id").Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).Find(&jscs)
		res = tx.Limit(-1).Offset(-1).Count(&count)
	}
	return count, jscs, exception.Wrap(response.ExceptionDatabase, res.Error)
}

func (jsi *JscRepoImpl) Get(db *gorm.DB, id int64) (*models.JsCategory, exception.Exception) {
	jsCategory := models.JsCategory{}
	res := db.Where(&models.JsCategory{ID: id}).Find(&jsCategory)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &jsCategory, nil
}

func (jsi *JscRepoImpl) Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.JsCategory{}).Where(&models.JsCategory{ID: id}).Updates(param).Error)
}

func (jsi *JscRepoImpl) Delete(db *gorm.DB, id int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.JsCategory{}, id).Error)
}

func (jsi *JscRepoImpl) DeleteByPrimaryID(db *gorm.DB, pid int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Where("primary_id = ?", pid).Delete(&models.JsCategory{}).Error)
}

func (jsi *JscRepoImpl) MultiDelete(db *gorm.DB, ids []int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.JsCategory{}, ids).Error)
}

func (jsi *JscRepoImpl) ListAllByPrimaryID(db *gorm.DB, pid int64) ([]models.JsCategory, exception.Exception) {
	categories := make([]models.JsCategory, 0)
	tx := db.Where("primary_id = ?", pid).Order("id").Find(&categories)
	if tx.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return categories, nil
}

func (jsi *JscRepoImpl) UpdateDomainIDToO(db *gorm.DB, domainIDs ...int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.JsCategory{}).Where("domain_id in (?)", domainIDs).UpdateColumn("domain_id", 0).Error)
}

func (jsi *JscRepoImpl) GetByIDs(db *gorm.DB, ids []int64) ([]models.JsCategory, exception.Exception) {
	categories := make([]models.JsCategory, 0, len(ids))
	tx := db.Find(&categories, ids)
	if tx.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return categories, nil
}

func (jsi *JscRepoImpl) ListAllByPrimaryUserID(db *gorm.DB, pid, userID int64, isAdmin bool) ([]models.JsCategory, exception.Exception) {
	categories := make([]models.JsCategory, 0)
	if !isAdmin {
		sub := db.Table(tables.UserCategoryRelation).Where("user_id = ?", userID)
		tx := db.Table(tables.JsCategory+" AS jc").Where("jc.primary_id = ?", pid).Select("jc.*").
			Joins("INNER JOIN (?) AS uc ON jc.id = uc.category_id", sub).Find(&categories)
		if tx.Error != nil {
			return nil, exception.Wrap(response.ExceptionDatabase, tx.Error)
		}
	} else {
		tx := db.Where("primary_id = ?", pid).Order("id").Find(&categories)
		if tx.Error != nil {
			return nil, exception.Wrap(response.ExceptionDatabase, tx.Error)
		}
	}
	return categories, nil
}
