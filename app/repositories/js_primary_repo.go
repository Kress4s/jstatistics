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
	List(db *gorm.DB, pageInfo *vo.PageInfo) (int64, []models.JsPrimary, exception.Exception)
	Get(db *gorm.DB, id uint) (*models.JsPrimary, exception.Exception)
	Update(db *gorm.DB, id uint, param map[string]interface{}) exception.Exception
	Delete(db *gorm.DB, id uint) exception.Exception
}

func (jri *JspRepoImpl) Create(db *gorm.DB, jsp *models.JsPrimary) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(jsp).Error)
}

func (jri *JspRepoImpl) List(db *gorm.DB, pageInfo *vo.PageInfo) (int64, []models.JsPrimary, exception.Exception) {
	jsps := make([]models.JsPrimary, 0)
	tx := db.Table(tables.JsPrimary)
	if pageInfo.Keywords != "" {
		tx = tx.Scopes(vo.FuzzySearch(pageInfo.Keywords, "title", "ip"))
	}
	tx.Order("id").Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).Find(&jsps)
	count := int64(0)
	res := tx.Limit(-1).Offset(-1).Count(&count)
	return count, jsps, exception.Wrap(response.ExceptionDatabase, res.Error)
}

func (jri *JspRepoImpl) Get(db *gorm.DB, id uint) (*models.JsPrimary, exception.Exception) {
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

func (jri *JspRepoImpl) Update(db *gorm.DB, id uint, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.JsPrimary{}).Where(&models.JsPrimary{ID: id}).Updates(param).Error)
}

func (jri *JspRepoImpl) Delete(db *gorm.DB, id uint) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.JsPrimary{}, id).Error)
}
