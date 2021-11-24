package repositories

import (
	"fmt"
	"js_statistics/app/models"
	"js_statistics/app/models/tables"
	"js_statistics/app/response"
	"js_statistics/app/vo"
	"js_statistics/constant"
	"js_statistics/exception"
	"sync"
	"time"

	"gorm.io/gorm"
)

var (
	homeRepoInstance HomeRepo
	homeOnce         sync.Once
)

type homeRepoImpl struct{}

func GetHomeRepo() HomeRepo {
	homeOnce.Do(func() {
		homeRepoInstance = &homeRepoImpl{}
	})
	return homeRepoInstance
}

type HomeRepo interface {
	TodayIP(db *gorm.DB) (int64, exception.Exception)
	YesterdayIP(db *gorm.DB) (int64, exception.Exception)
	ThisMonthIP(db *gorm.DB, beginAt, endAt string) (int64, exception.Exception)
	LastMonthIP(db *gorm.DB, beginAt, endAt string) (int64, exception.Exception)
	IPAndUVisit(db *gorm.DB, beginAt, endAt string) ([]models.IPVisitStatistic, []models.UVisitStatistic,
		exception.Exception)
	RegionStatistic(db *gorm.DB) ([]models.RegionStatistic, exception.Exception)
	JSVisitStatistic(db *gorm.DB, pageInfo *vo.PageInfo) (int64, []models.JSVisitStatistic, exception.Exception)
}

func (hri *homeRepoImpl) TodayIP(db *gorm.DB) (int64, exception.Exception) {
	count := int64(0)
	tx := db.Table(tables.IPRecode).Select("count(distinct(ip))").
		Where("visit_time = ?", time.Now().Format(constant.DateFormat)).Count(&count)
	if tx.Error != nil {
		return 0, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return count, nil
}

func (hri *homeRepoImpl) YesterdayIP(db *gorm.DB) (int64, exception.Exception) {
	count := int64(0)
	tx := db.Table(tables.IPRecode).Select("count(distinct(ip))").
		Where("visit_time = ?", time.Now().AddDate(0, 0, -1).Format(constant.DateFormat)).Count(&count)
	if tx.Error != nil {
		return 0, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return count, nil
}

func (hri *homeRepoImpl) ThisMonthIP(db *gorm.DB, beginAt, endAt string) (int64, exception.Exception) {
	count := int64(0)
	tx := db.Table(tables.IPRecode).Select("count(distinct(ip))").
		Where("visit_time BETWEEN ? AND ?", beginAt, endAt).Count(&count)
	if tx.Error != nil {
		return 0, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return count, nil
}

func (hri *homeRepoImpl) LastMonthIP(db *gorm.DB, beginAt, endAt string) (int64, exception.Exception) {
	count := int64(0)
	tx := db.Table(tables.IPRecode).Select("count(distinct(ip))").
		Where("visit_time >= ? AND visit_time < ?", beginAt, endAt).Count(&count)
	if tx.Error != nil {
		return 0, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return count, nil
}

func (hri *homeRepoImpl) IPAndUVisit(db *gorm.DB, beginAt, endAt string) ([]models.IPVisitStatistic,
	[]models.UVisitStatistic, exception.Exception) {
	ip := make([]models.IPVisitStatistic, 0)
	uv := make([]models.UVisitStatistic, 0)
	txIP := db.Table(tables.IPStatistics).Select("ip, visit_time, count(*) as count").
		Where("visit_time BETWEEN ? AND ?", beginAt, endAt).Group("ip, visit_time").Scan(&ip)
	if txIP.Error != nil {
		return nil, nil, exception.Wrap(response.ExceptionDatabase, txIP.Error)
	}
	txUV := db.Table(tables.UVStatistics).Select("cookie, visit_time, count(*) as count").
		Where("visit_time BETWEEN ? AND ?", beginAt, endAt).Group("visit_time, cookie").Scan(&uv)
	if txUV.Error != nil {
		return nil, nil, exception.Wrap(response.ExceptionDatabase, txUV.Error)
	}
	return ip, uv, nil
}

func (hri *homeRepoImpl) RegionStatistic(db *gorm.DB) ([]models.RegionStatistic, exception.Exception) {
	res := make([]models.RegionStatistic, 0)
	tx := db.Table(tables.IPRecode).Select("region, count(*) as count").Group("region").Scan(&res)
	if tx.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return res, nil
}

func (hri *homeRepoImpl) JSVisitStatistic(db *gorm.DB, pageInfo *vo.PageInfo) (int64, []models.JSVisitStatistic,
	exception.Exception) {
	statics := make([]models.JSVisitStatistic, 0)
	tx := db.Table(tables.IPRecode).Select("ir.js_id AS js_id, count(*) AS count, js.title AS title").
		Joins(fmt.Sprintf("INNER JOIN %s ON js.id = ir.js_id", tables.JsManage)).
		Group("ir.js_id, js.title").
		Order("ir.js_id, js.title, count DESC").
		Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).
		Scan(&statics)
	count := int64(0)
	res := tx.Limit(-1).Offset(-1).Count(&count)
	return count, statics, exception.Wrap(response.ExceptionDatabase, res.Error)
}
