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
	ipRepoInstance IPRepo
	ipOnce         sync.Once
)

type ipRepoImpl struct{}

func GetIPRepo() IPRepo {
	ipOnce.Do(func() {
		ipRepoInstance = &ipRepoImpl{}
	})
	return ipRepoInstance
}

type IPRepo interface {
	Create(db *gorm.DB, ip *models.WhiteIP) exception.Exception
	List(db *gorm.DB, pageInfo *vo.PageInfo) (int64, []models.WhiteIP, exception.Exception)
	Get(db *gorm.DB, id uint) (*models.WhiteIP, exception.Exception)
	Update(db *gorm.DB, id uint, param map[string]interface{}) exception.Exception
	Delete(db *gorm.DB, id uint) exception.Exception
	MultiDelete(db *gorm.DB, ids []uint) exception.Exception
}

func (iri *ipRepoImpl) Create(db *gorm.DB, domain *models.WhiteIP) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(domain).Error)
}

func (iri *ipRepoImpl) List(db *gorm.DB, pageInfo *vo.PageInfo) (int64, []models.WhiteIP, exception.Exception) {
	ips := make([]models.WhiteIP, 0)
	tx := db.Table(tables.WhiteIP)
	if pageInfo.Keywords != "" {
		tx = tx.Scopes(vo.FuzzySearch(pageInfo.Keywords, "ip"))
	}
	tx.Order("id").Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).Find(&ips)
	count := int64(0)
	res := tx.Limit(-1).Offset(-1).Count(&count)
	return count, ips, exception.Wrap(response.ExceptionDatabase, res.Error)
}

func (iri *ipRepoImpl) Get(db *gorm.DB, id uint) (*models.WhiteIP, exception.Exception) {
	domain := models.WhiteIP{}
	res := db.Where(&models.WhiteIP{ID: id}).Find(&domain)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &domain, nil
}

func (iri *ipRepoImpl) Update(db *gorm.DB, id uint, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.WhiteIP{}).Where(&models.WhiteIP{ID: id}).Updates(param).Error)
}

func (iri *ipRepoImpl) Delete(db *gorm.DB, id uint) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.WhiteIP{}, id).Error)
}

func (iri *ipRepoImpl) MultiDelete(db *gorm.DB, ids []uint) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.WhiteIP{}, ids).Error)
}
