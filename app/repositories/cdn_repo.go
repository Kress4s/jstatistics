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
	cdnRepoInstance CdnRepo
	cdnOnce         sync.Once
)

type CdnRepoImpl struct{}

func GetCdnRepo() CdnRepo {
	cdnOnce.Do(func() {
		cdnRepoInstance = &CdnRepoImpl{}
	})
	return cdnRepoInstance
}

type CdnRepo interface {
	Create(db *gorm.DB, cdn *models.CDN) exception.Exception
	List(db *gorm.DB, pageInfo *vo.PageInfo) (int64, []models.CDN, exception.Exception)
	Get(db *gorm.DB, id int64) (*models.CDN, exception.Exception)
	Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
	Delete(db *gorm.DB, id int64) exception.Exception
	MultiDelete(db *gorm.DB, ids []int64) exception.Exception
	IsExistByIP(db *gorm.DB, ip string) (bool, exception.Exception)
}

func (cri *CdnRepoImpl) Create(db *gorm.DB, cdn *models.CDN) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(cdn).Error)
}

func (cri *CdnRepoImpl) List(db *gorm.DB, pageInfo *vo.PageInfo) (int64, []models.CDN, exception.Exception) {
	cdns := make([]models.CDN, 0)
	tx := db.Table(tables.CDN)
	if pageInfo.Keywords != "" {
		tx = tx.Scopes(vo.FuzzySearch(pageInfo.Keywords, "title", "ip"))
	}
	tx.Order("id").Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).Find(&cdns)
	count := int64(0)
	res := tx.Limit(-1).Offset(-1).Count(&count)
	return count, cdns, exception.Wrap(response.ExceptionDatabase, res.Error)
}

func (cri *CdnRepoImpl) Get(db *gorm.DB, id int64) (*models.CDN, exception.Exception) {
	cdn := models.CDN{}
	res := db.Where(&models.CDN{ID: id}).Find(&cdn)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &cdn, nil
}

func (cri *CdnRepoImpl) Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.CDN{}).Where(&models.CDN{ID: id}).Updates(param).Error)
}

func (cri *CdnRepoImpl) Delete(db *gorm.DB, id int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.CDN{}, id).Error)
}

func (cri *CdnRepoImpl) MultiDelete(db *gorm.DB, ids []int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.CDN{}, ids).Error)
}

func (cri *CdnRepoImpl) IsExistByIP(db *gorm.DB, ip string) (bool, exception.Exception) {
	cdn := models.CDN{}
	res := db.Where(&models.CDN{IP: ip}).Find(&cdn)
	if res.RowsAffected == 0 {
		return false, nil
	}
	if res.Error != nil {
		return false, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return true, nil
}
