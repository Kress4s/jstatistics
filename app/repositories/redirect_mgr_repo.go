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
	rmRepoInstance RmRepo
	rmOnce         sync.Once
)

type RmRepoImpl struct{}

func GetRmRepo() RmRepo {
	rmOnce.Do(func() {
		rmRepoInstance = &RmRepoImpl{}
	})
	return rmRepoInstance
}

type RmRepo interface {
	Create(db *gorm.DB, rm *models.RedirectManage) exception.Exception
	List(db *gorm.DB, pageInfo *vo.PageInfo) (int64, []models.RedirectManage, exception.Exception)
	Get(db *gorm.DB, id int64) (*models.RedirectManage, exception.Exception)
	Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
	Delete(db *gorm.DB, id int64) exception.Exception
	MultiDelete(db *gorm.DB, ids []int64) exception.Exception
	GetUsefulByCategoryID(db *gorm.DB, id int64) (*models.RedirectManage, exception.Exception)
}

func (jsi *RmRepoImpl) Create(db *gorm.DB, rm *models.RedirectManage) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(rm).Error)
}

func (jsi *RmRepoImpl) List(db *gorm.DB, pageInfo *vo.PageInfo) (int64, []models.RedirectManage, exception.Exception) {
	rms := make([]models.RedirectManage, 0)
	tx := db.Table(tables.RedirectManage)
	if pageInfo.Keywords != "" {
		tx = tx.Scopes(vo.FuzzySearch(pageInfo.Keywords, "title"))
	}
	tx.Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).Find(&rms)
	count := int64(0)
	res := tx.Limit(-1).Offset(-1).Count(&count)
	return count, rms, exception.Wrap(response.ExceptionDatabase, res.Error)
}

func (jsi *RmRepoImpl) Get(db *gorm.DB, id int64) (*models.RedirectManage, exception.Exception) {
	jsCategory := models.RedirectManage{}
	res := db.Where(&models.RedirectManage{ID: id}).Find(&jsCategory)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &jsCategory, nil
}

func (jsi *RmRepoImpl) Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.RedirectManage{}).Where(&models.RedirectManage{ID: id}).Updates(param).Error)
}

func (jsi *RmRepoImpl) Delete(db *gorm.DB, id int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.RedirectManage{}, id).Error)
}

func (jsi *RmRepoImpl) MultiDelete(db *gorm.DB, ids []int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.RedirectManage{}, ids).Error)
}

func (jsi *RmRepoImpl) GetUsefulByCategoryID(db *gorm.DB, id int64) (*models.RedirectManage, exception.Exception) {
	jsCategory := make([]models.RedirectManage, 0)
	res := db.Where(&models.RedirectManage{CategoryID: id, Status: true}).Find(&jsCategory)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &(jsCategory[0]), nil
}
