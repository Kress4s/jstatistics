package repositories

import (
	"js_statistics/app/models"
	"js_statistics/app/response"
	"js_statistics/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	stcRepoInstance StcRepo
	stcOnce         sync.Once
)

type StcRepoImpl struct{}

func GetStcRepo() StcRepo {
	stcOnce.Do(func() {
		stcRepoInstance = &StcRepoImpl{}
	})
	return stcRepoInstance
}

type StcRepo interface {
	CreateIPStatistics(db *gorm.DB, jsm *models.IPStatistics) exception.Exception
	CreateUVStatistics(db *gorm.DB, jsm *models.UVStatistics) exception.Exception
	CreateIPRecode(db *gorm.DB, jsm *models.IPRecode) exception.Exception
}

func (sri *StcRepoImpl) CreateIPStatistics(db *gorm.DB, jsm *models.IPStatistics) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(jsm).Error)
}

func (sri *StcRepoImpl) CreateUVStatistics(db *gorm.DB, jsm *models.UVStatistics) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(jsm).Error)
}

func (sri *StcRepoImpl) CreateIPRecode(db *gorm.DB, jsm *models.IPRecode) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(jsm).Error)
}
