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
	jsmRepoInstance JsmRepo
	jsmOnce         sync.Once
)

type JsmRepoImpl struct{}

func GetJsmRepo() JsmRepo {
	jsmOnce.Do(func() {
		jsmRepoInstance = &JsmRepoImpl{}
	})
	return jsmRepoInstance
}

type JsmRepo interface {
	Create(db *gorm.DB, jsm *models.JsManage) exception.Exception
	ListByCategoryID(db *gorm.DB, pageInfo *vo.PageInfo, cid int64) (int64, []models.JsManage, exception.Exception)
	Get(db *gorm.DB, id int64) (*models.JsManage, exception.Exception)
	Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
	Delete(db *gorm.DB, id int64) exception.Exception
	MultiDelete(db *gorm.DB, ids []int64) exception.Exception
}

func (jsi *JsmRepoImpl) Create(db *gorm.DB, jsm *models.JsManage) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(jsm).Error)
}

func (jsi *JsmRepoImpl) ListByCategoryID(db *gorm.DB, pageInfo *vo.PageInfo, pid int64) (int64, []models.JsManage, exception.Exception) {
	jsms := make([]models.JsManage, 0)
	tx := db.Table(tables.JsManage)
	if pageInfo.Keywords != "" {
		tx = tx.Scopes(vo.FuzzySearch(pageInfo.Keywords, "title"))
	}
	tx.Where("category_id = ?", pid).Order("id").Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).Find(&jsms)
	count := int64(0)
	res := tx.Limit(-1).Offset(-1).Count(&count)
	return count, jsms, exception.Wrap(response.ExceptionDatabase, res.Error)
}

func (jsi *JsmRepoImpl) Get(db *gorm.DB, id int64) (*models.JsManage, exception.Exception) {
	jsCategory := models.JsManage{}
	res := db.Where(&models.JsManage{ID: id}).Find(&jsCategory)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &jsCategory, nil
}

func (jsi *JsmRepoImpl) Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.JsManage{}).Where(&models.JsManage{ID: id}).Updates(param).Error)
}

func (jsi *JsmRepoImpl) Delete(db *gorm.DB, id int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.JsManage{}, id).Error)
}

func (jsi *JsmRepoImpl) MultiDelete(db *gorm.DB, ids []int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.JsManage{}, ids).Error)
}
