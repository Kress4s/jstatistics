package repositories

import (
	"fmt"
	"js_statistics/app/models"
	"js_statistics/app/models/tables"
	"js_statistics/app/response"
	"js_statistics/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	jspRepoInstance JspRepo
	jspOnce         sync.Once
)

type JspRepoImpl struct{}

func GetJspRepo() JspRepo {
	jspOnce.Do(func() {
		jspRepoInstance = &JspRepoImpl{}
	})
	return jspRepoInstance
}

type JspRepo interface {
	Create(db *gorm.DB, jsp *models.JsPrimary) exception.Exception
	ListByUserID(db *gorm.DB, userID int64) ([]models.JsPrimary, exception.Exception)
	List(db *gorm.DB) ([]models.JsPrimary, exception.Exception)
	Get(db *gorm.DB, id int64) (*models.JsPrimary, exception.Exception)
	Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
	Delete(db *gorm.DB, id int64) exception.Exception
	GetAllsCategoryTree(db *gorm.DB) ([]models.AllsCategory, exception.Exception)
}

func (jri *JspRepoImpl) Create(db *gorm.DB, jsp *models.JsPrimary) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(jsp).Error)
}

func (jri *JspRepoImpl) ListByUserID(db *gorm.DB, userID int64) ([]models.JsPrimary, exception.Exception) {
	jsps := make([]models.JsPrimary, 0)
	sub := db.Table(tables.UserPrimaryRelation).Where("user_id = ?", userID)
	tx := db.Table(tables.JsPrimary+" AS p").Select("p.*").
		Joins("INNER JOIN (?) AS up ON up.primary_id = p.id", sub)
	tx.Order("p.id").Find(&jsps)
	return jsps, exception.Wrap(response.ExceptionDatabase, tx.Error)
}

func (jri *JspRepoImpl) List(db *gorm.DB) ([]models.JsPrimary, exception.Exception) {
	jsps := make([]models.JsPrimary, 0)
	tx := db.Table(tables.JsPrimary).Order("id").Find(&jsps)
	return jsps, exception.Wrap(response.ExceptionDatabase, tx.Error)
}

func (jri *JspRepoImpl) Get(db *gorm.DB, id int64) (*models.JsPrimary, exception.Exception) {
	jsp := models.JsPrimary{}
	res := db.Where(&models.JsPrimary{ID: id}).Find(&jsp)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &jsp, nil
}

func (jri *JspRepoImpl) Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.JsPrimary{}).Where(&models.JsPrimary{ID: id}).Updates(param).Error)
}

func (jri *JspRepoImpl) Delete(db *gorm.DB, id int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.JsPrimary{}, id).Error)
}

func (jri *JspRepoImpl) GetAllsCategoryTree(db *gorm.DB) ([]models.AllsCategory, exception.Exception) {
	allCategories := make([]models.AllsCategory, 0)
	res := db.Table(tables.JsPrimary + " AS p").
		Select("p.id AS id, p.title AS title, c.id AS cid, c.title AS c_title, c.primary_id AS pid").
		Joins(fmt.Sprintf("LEFT JOIN %s AS c ON c.primary_id = p.id", tables.JsCategory)).
		Scan(&allCategories)
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return allCategories, nil
}
