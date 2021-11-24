package repositories

import (
	"fmt"
	"js_statistics/app/models"
	"js_statistics/app/models/tables"
	"js_statistics/app/response"
	"js_statistics/app/vo"
	"js_statistics/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	faRepoInstance FaRepo
	faOnce         sync.Once
)

type FaRepoImpl struct{}

func GetFaRepo() FaRepo {
	faOnce.Do(func() {
		faRepoInstance = &FaRepoImpl{}
	})
	return faRepoInstance
}

type FaRepo interface {
	FromStatistic(db *gorm.DB, param *vo.JSFilterParams, pageInfo *vo.PageInfo,
		beginAt, endAt string) (int64, []models.FromAnalysisView, exception.Exception)
}

func (fri *FaRepoImpl) FromStatistic(db *gorm.DB, param *vo.JSFilterParams, pageInfo *vo.PageInfo,
	beginAt, endAt string) (int64, []models.FromAnalysisView, exception.Exception) {
	data := make([]models.FromAnalysisView, 0)
	subTx := db.Table(tables.IPRecode).Select("js_id, from_url, to_url").
		Where("visit_time >= ? and visit_time <= ?", beginAt, endAt)
	if param.PrimaryID == 0 {
		return 0, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	subTx = subTx.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		subTx = subTx.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		subTx = subTx.Where("js_id = ?", param.JsID)
	}
	count := int64(0)
	tx := db.Table("(?) AS sub", subTx).
		Select("js.title as title, sub.js_id as js_id, sub.from_url as from_url, sub.to_url as to_url, count(*) as count").
		Joins(fmt.Sprintf("INNER JOIN %s AS js ON js.id = sub.js_id", tables.JsManage)).
		Group("js.title, sub.js_id, sub.from_url, sub.to_url").
		Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).
		Scan(&data).Limit(-1).Offset(-1).Count(&count)
	return count, data, exception.Wrap(response.ExceptionDatabase, tx.Error)
}
