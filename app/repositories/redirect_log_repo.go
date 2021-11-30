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
	rlRepoInstance RlRepo
	rlOnce         sync.Once
)

type RlRepoImpl struct{}

func GetRlRepo() RlRepo {
	rlOnce.Do(func() {
		rlRepoInstance = &RlRepoImpl{}
	})
	return rlRepoInstance
}

type RlRepo interface {
	Create(db *gorm.DB, rm *models.RedirectLog) exception.Exception
	ListByCategoryID(db *gorm.DB, pageInfo *vo.PageInfo, cid int64) (int64, []models.RedirectLog, exception.Exception)
	GetByCategoryID(db *gorm.DB, id int64) (*models.RedirectLog, exception.Exception)
	Update(db *gorm.DB, redirectID, categoryID int64, param map[string]interface{}) exception.Exception
	Delete(db *gorm.DB, redirectID, categoryID int64) exception.Exception
	MultiDelete(db *gorm.DB, rmIDs []int64, categoryID int64) exception.Exception
}

func (rri *RlRepoImpl) Create(db *gorm.DB, rl *models.RedirectLog) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(rl).Error)
}

func (rri *RlRepoImpl) GetByCategoryID(db *gorm.DB, id int64) (*models.RedirectLog, exception.Exception) {
	rm := models.RedirectLog{}
	res := db.Where(&models.RedirectLog{ID: id}).Find(&rm)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &rm, nil
}

func (rri *RlRepoImpl) ListByCategoryID(db *gorm.DB, pageInfo *vo.PageInfo, cid int64) (int64, []models.RedirectLog, exception.Exception) {
	rls := make([]models.RedirectLog, 0)
	tx := db.Table(tables.RedirectLog)
	if pageInfo.Keywords != "" {
		tx = tx.Scopes(vo.FuzzySearch(pageInfo.Keywords, "title"))
	}
	tx.Where("category_id = ?", cid).Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).Find(&rls)
	count := int64(0)
	res := tx.Limit(-1).Offset(-1).Count(&count)
	return count, rls, exception.Wrap(response.ExceptionDatabase, res.Error)
}

func (rri *RlRepoImpl) Update(db *gorm.DB, redirectID, categoryID int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.RedirectLog{}).
			Where(&models.RedirectLog{CategoryID: categoryID, RedirectID: redirectID}).Updates(param).Error)
}

func (rri *RlRepoImpl) Delete(db *gorm.DB, redirectID, categoryID int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.
		Where("category_id = ? and redirect_id = ?", categoryID, redirectID).Delete(&models.RedirectLog{}).Error)
}

func (rri *RlRepoImpl) MultiDelete(db *gorm.DB, rmIDs []int64, categoryID int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Where("redirect_id in ? and category_id = ?", rmIDs, categoryID).Delete(&models.RedirectLog{}).Error)
}
