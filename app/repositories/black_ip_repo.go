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
	blackIPRepoInstance BlackIPRepo
	blackIPOnce         sync.Once
)

type BlackIPRepoImpl struct{}

func GetBlackIPRepo() BlackIPRepo {
	blackIPOnce.Do(func() {
		blackIPRepoInstance = &BlackIPRepoImpl{}
	})
	return blackIPRepoInstance
}

type BlackIPRepo interface {
	Create(db *gorm.DB, ip *models.BlackIPMgr) exception.Exception
	List(db *gorm.DB, pageInfo *vo.PageInfo) (int64, []models.BlackIPMgr, exception.Exception)
	Get(db *gorm.DB, id int64) (*models.BlackIPMgr, exception.Exception)
	Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
	Delete(db *gorm.DB, id int64) exception.Exception
	MultiDelete(db *gorm.DB, ids []int64) exception.Exception
	IsExistByIP(db *gorm.DB, ip string) (bool, exception.Exception)
}

func (dri *BlackIPRepoImpl) Create(db *gorm.DB, domain *models.BlackIPMgr) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(domain).Error)
}

func (dri *BlackIPRepoImpl) List(db *gorm.DB, pageInfo *vo.PageInfo) (int64, []models.BlackIPMgr, exception.Exception) {
	ips := make([]models.BlackIPMgr, 0)
	tx := db.Table(tables.BlackIPMgr)
	if pageInfo.Keywords != "" {
		tx = tx.Scopes(vo.FuzzySearch(pageInfo.Keywords, "ip"))
	}
	tx.Order("id").Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).Find(&ips)
	count := int64(0)
	res := tx.Limit(-1).Offset(-1).Count(&count)
	return count, ips, exception.Wrap(response.ExceptionDatabase, res.Error)
}

func (dri *BlackIPRepoImpl) Get(db *gorm.DB, id int64) (*models.BlackIPMgr, exception.Exception) {
	domain := models.BlackIPMgr{}
	res := db.Where(&models.BlackIPMgr{ID: id}).Find(&domain)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &domain, nil
}

func (dri *BlackIPRepoImpl) Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.BlackIPMgr{}).Where(&models.BlackIPMgr{ID: id}).Updates(param).Error)
}

func (dri *BlackIPRepoImpl) Delete(db *gorm.DB, id int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.BlackIPMgr{}, id).Error)
}

func (dri *BlackIPRepoImpl) MultiDelete(db *gorm.DB, ids []int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.BlackIPMgr{}, ids).Error)
}

func (dri *BlackIPRepoImpl) IsExistByIP(db *gorm.DB, ip string) (bool, exception.Exception) {
	wip := models.BlackIPMgr{}
	res := db.Where(&models.BlackIPMgr{IP: ip}).Find(&wip)
	if res.RowsAffected == 0 {
		return false, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return false, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return true, nil
}
