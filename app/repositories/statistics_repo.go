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
	DeleteByPrimaryID(db *gorm.DB, pid int64) exception.Exception
	DeleteByCategoryID(db *gorm.DB, cid int64) exception.Exception
	DeleteByCategoriesID(db *gorm.DB, cids ...int64) exception.Exception
	DeleteByJsID(db *gorm.DB, jsID int64) exception.Exception
	DeleteByJsIDs(db *gorm.DB, jsID ...int64) exception.Exception
}

func (sri *StcRepoImpl) DeleteByPrimaryID(db *gorm.DB, pid int64) exception.Exception {
	if res := db.Where("primary_id = ?", pid).Delete(&models.IPStatistics{}); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	if res := db.Where("primary_id = ?", pid).Delete(&models.UVStatistics{}); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	if res := db.Where("primary_id = ?", pid).Delete(&models.IPRecode{}); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return nil
}

func (sri *StcRepoImpl) DeleteByCategoryID(db *gorm.DB, cid int64) exception.Exception {
	if res := db.Where("category_id = ?", cid).Delete(&models.IPStatistics{}); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	if res := db.Where("category_id = ?", cid).Delete(&models.UVStatistics{}); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	if res := db.Where("category_id = ?", cid).Delete(&models.IPRecode{}); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return nil
}

func (sri *StcRepoImpl) DeleteByCategoriesID(db *gorm.DB, cids ...int64) exception.Exception {
	if res := db.Where("category_id in (?)", cids).Delete(&models.IPStatistics{}); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	if res := db.Where("category_id in (?)", cids).Delete(&models.UVStatistics{}); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	if res := db.Where("category_id in (?)", cids).Delete(&models.IPRecode{}); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return nil
}

func (sri *StcRepoImpl) DeleteByJsID(db *gorm.DB, jsID int64) exception.Exception {
	if res := db.Where("js_id = ?", jsID).Delete(&models.IPStatistics{}); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	if res := db.Where("js_id = ?", jsID).Delete(&models.UVStatistics{}); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	if res := db.Where("js_id = ?", jsID).Delete(&models.IPRecode{}); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return nil
}

func (sri *StcRepoImpl) DeleteByJsIDs(db *gorm.DB, jsID ...int64) exception.Exception {
	if res := db.Where("js_id in (?)", jsID).Delete(&models.IPStatistics{}); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	if res := db.Where("js_id in (?)", jsID).Delete(&models.UVStatistics{}); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	if res := db.Where("js_id in (?)", jsID).Delete(&models.IPRecode{}); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return nil
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
