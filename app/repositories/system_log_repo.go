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
	syslogRepoInstance SyslogRepo
	syslogOnce         sync.Once
)

type syslogRepoImpl struct{}

func GetSyslogRepo() SyslogRepo {
	syslogOnce.Do(func() {
		syslogRepoInstance = &syslogRepoImpl{}
	})
	return syslogRepoInstance
}

type SyslogRepo interface {
	Create(db *gorm.DB, log *models.SystemLog) exception.Exception
	List(db *gorm.DB, pageInfo *vo.PageInfo) (int64, []models.SystemLog, exception.Exception)
}

func (sli *syslogRepoImpl) Create(db *gorm.DB, log *models.SystemLog) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(log).Error)
}

func (sli *syslogRepoImpl) List(db *gorm.DB, pageInfo *vo.PageInfo) (int64, []models.SystemLog, exception.Exception) {
	logs := make([]models.SystemLog, 0)
	tx := db.Table(tables.SystemLog)
	tx.Order("id").Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).Find(&logs)
	count := int64(0)
	res := tx.Limit(-1).Offset(-1).Count(&count)
	return count, logs, exception.Wrap(response.ExceptionDatabase, res.Error)
}
