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
	domainRepoInstance DomainRepo
	domainOnce         sync.Once
)

type DomainRepoImpl struct{}

func GetDomainRepo() DomainRepo {
	domainOnce.Do(func() {
		domainRepoInstance = &DomainRepoImpl{}
	})
	return domainRepoInstance
}

type DomainRepo interface {
	Create(db *gorm.DB, domain *models.DomainMgr) exception.Exception
	List(db *gorm.DB, pageInfo *vo.PageInfo) (int64, []models.DomainMgr, exception.Exception)
	Get(db *gorm.DB, id uint) (*models.DomainMgr, exception.Exception)
	Update(db *gorm.DB, id uint, param map[string]interface{}) exception.Exception
	Delete(db *gorm.DB, id uint) exception.Exception
	MultiDelete(db *gorm.DB, ids []uint) exception.Exception
}

func (dri *DomainRepoImpl) Create(db *gorm.DB, domain *models.DomainMgr) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(domain).Error)
}

func (dri *DomainRepoImpl) List(db *gorm.DB, pageInfo *vo.PageInfo) (int64, []models.DomainMgr, exception.Exception) {
	domains := make([]models.DomainMgr, 0)
	tx := db.Table(tables.DomainMgr)
	if pageInfo.Keywords != "" {
		tx = tx.Scopes(vo.FuzzySearch(pageInfo.Keywords, "title", "domain"))
	}
	tx.Order("id").Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).Find(&domains)
	count := int64(0)
	res := tx.Limit(-1).Offset(-1).Count(&count)
	return count, domains, exception.Wrap(response.ExceptionDatabase, res.Error)
}

func (dri *DomainRepoImpl) Get(db *gorm.DB, id uint) (*models.DomainMgr, exception.Exception) {
	domain := models.DomainMgr{}
	res := db.Where(&models.DomainMgr{ID: id}).Find(&domain)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &domain, nil
}

func (dri *DomainRepoImpl) Update(db *gorm.DB, id uint, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.DomainMgr{}).Where(&models.DomainMgr{ID: id}).Updates(param).Error)
}

func (dri *DomainRepoImpl) Delete(db *gorm.DB, id uint) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.DomainMgr{}, id).Error)
}

func (dri *DomainRepoImpl) MultiDelete(db *gorm.DB, ids []uint) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.DomainMgr{}, ids).Error)
}
