package repositories

import (
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
	daRepoInstance DaRepo
	daOnce         sync.Once
)

type daRepoImpl struct{}

func GetDaRepo() DaRepo {
	daOnce.Do(func() {
		daRepoInstance = &daRepoImpl{}
	})
	return daRepoInstance
}

type DaRepo interface {
	TodayIP(db *gorm.DB, param *vo.JSFilterParams) (int64, exception.Exception)
	YesterdayIP(db *gorm.DB, param *vo.JSFilterParams) (int64, exception.Exception)
	ThisMonthIP(db *gorm.DB, param *vo.JSFilterParams, beginAt, endAt string) (int64, exception.Exception)
	LastMonthIP(db *gorm.DB, param *vo.JSFilterParams, beginAt, endAt string) (int64, exception.Exception)
	// time scope
	IPAndUVisit(db *gorm.DB, param *vo.JSFilterParams, beginAt, endAt string) ([]models.IPVisitStatistic,
		[]models.UVisitStatistic, exception.Exception)
	TodayIPAndUVisit(db *gorm.DB, param *vo.JSFilterParams) ([]models.IPVisitStatistic,
		[]models.UVisitStatistic, exception.Exception)
	YesterdayIPAndUVisit(db *gorm.DB, param *vo.JSFilterParams) ([]models.IPVisitStatistic,
		[]models.UVisitStatistic, exception.Exception)
	FromNowIPAndUVisit(db *gorm.DB, param *vo.JSFilterParams) ([]models.IPVisitStatistic, []models.UVisitStatistic,
		exception.Exception)
	TodayFlowData(db *gorm.DB, param *vo.JSFilterParams, pageInfo *vo.PageInfo) (int64,
		[]models.FlowDataStatistic, exception.Exception)
	YesterdayFlowData(db *gorm.DB, param *vo.JSFilterParams, pageInfo *vo.PageInfo) (int64,
		[]models.FlowDataStatistic, exception.Exception)
	// time scope
	TimeScopeFlowData(db *gorm.DB, param *vo.JSFilterParams, pageInfo *vo.PageInfo, beginAt, endAt string) (int64,
		[]models.FlowDataStatistic, exception.Exception)
	FromNowFlowData(db *gorm.DB, param *vo.JSFilterParams, pageInfo *vo.PageInfo) (int64, []models.FlowDataStatistic,
		exception.Exception)
}

func (dri *daRepoImpl) TodayIP(db *gorm.DB, param *vo.JSFilterParams) (int64, exception.Exception) {
	tx := db.Table(tables.IPRecode).Select("count(distinct(ip))").
		Where("visit_time = ?", time.Now().Format(constant.DateFormat))
	if param.PrimaryID == 0 {
		return 0, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	if param.CategoryID != 0 {
		tx.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		tx.Where("js_id = ?", param.JsID)
	}
	count := int64(0)
	res := tx.Count(&count)
	return count, exception.Wrap(response.ExceptionDatabase, res.Error)
}

func (dri *daRepoImpl) YesterdayIP(db *gorm.DB, param *vo.JSFilterParams) (int64, exception.Exception) {
	tx := db.Table(tables.IPRecode).Select("count(distinct(ip))").
		Where("visit_time = ?", time.Now().AddDate(0, 0, -1).Format(constant.DateFormat))
	if param.PrimaryID == 0 {
		return 0, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	if param.CategoryID != 0 {
		tx.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		tx.Where("js_id = ?", param.JsID)
	}
	count := int64(0)
	res := tx.Count(&count)
	return count, exception.Wrap(response.ExceptionDatabase, res.Error)
}

func (dri *daRepoImpl) ThisMonthIP(db *gorm.DB, param *vo.JSFilterParams, beginAt, endAt string) (int64,
	exception.Exception) {

	tx := db.Table(tables.IPRecode).Select("count(distinct(ip))").
		Where("visit_time BETWEEN ? AND ?", beginAt, endAt)
	if param.PrimaryID == 0 {
		return 0, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	if param.CategoryID != 0 {
		tx = tx.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		tx = tx.Where("js_id = ?", param.JsID)
	}
	count := int64(0)
	res := tx.Count(&count)
	return count, exception.Wrap(response.ExceptionDatabase, res.Error)
}

func (dri *daRepoImpl) LastMonthIP(db *gorm.DB, param *vo.JSFilterParams, beginAt, endAt string) (int64,
	exception.Exception) {
	tx := db.Table(tables.IPRecode).Select("count(distinct(ip))").
		Where("visit_time >= ? and visit_time < ?", beginAt, endAt)
	if param.PrimaryID == 0 {
		return 0, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	tx = tx.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		tx = tx.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		tx = tx.Where("js_id = ?", param.JsID)
	}
	count := int64(0)
	res := tx.Count(&count)
	return count, exception.Wrap(response.ExceptionDatabase, res.Error)
}

func (dri *daRepoImpl) IPAndUVisit(db *gorm.DB, param *vo.JSFilterParams, beginAt, endAt string,
) ([]models.IPVisitStatistic, []models.UVisitStatistic, exception.Exception) {
	ip := make([]models.IPVisitStatistic, 0)
	uv := make([]models.UVisitStatistic, 0)
	txIP := db.Table(tables.IPStatistics).Select("ip, visit_time, count(*) as count").
		Where("visit_time BETWEEN ? AND ?", beginAt, endAt)
	if param.PrimaryID == 0 {
		return nil, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	txIP = txIP.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		txIP = txIP.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		txIP = txIP.Where("js_id = ?", param.JsID)
	}
	txIP = txIP.Group("ip, visit_time").Scan(&ip)
	if txIP.Error != nil {
		return nil, nil, exception.Wrap(response.ExceptionDatabase, txIP.Error)
	}
	txUV := db.Table(tables.UVStatistics).Select("cookie, visit_time, count(*) as count").
		Where("visit_time BETWEEN ? AND ?", beginAt, endAt)
	if param.PrimaryID == 0 {
		return nil, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	txUV = txUV.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		txUV = txUV.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		txUV = txUV.Where("js_id = ?", param.JsID)
	}
	txUV = txUV.Group("visit_time, cookie").Scan(&uv)
	if txUV.Error != nil {
		return nil, nil, exception.Wrap(response.ExceptionDatabase, txUV.Error)
	}
	return ip, uv, nil
}

func (dri *daRepoImpl) TodayIPAndUVisit(db *gorm.DB, param *vo.JSFilterParams) ([]models.IPVisitStatistic,
	[]models.UVisitStatistic, exception.Exception) {
	ip := make([]models.IPVisitStatistic, 0)
	uv := make([]models.UVisitStatistic, 0)
	today := time.Now().Format(constant.DateFormat)
	txIP := db.Table(tables.IPStatistics).Select("ip, visit_time, count(*) as count").
		Where("visit_time = ?", today)
	if param.PrimaryID == 0 {
		return nil, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	txIP = txIP.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		txIP = txIP.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		txIP = txIP.Where("js_id = ?", param.JsID)
	}
	txIP = txIP.Group("ip, visit_time").Scan(&ip)
	if txIP.Error != nil {
		return nil, nil, exception.Wrap(response.ExceptionDatabase, txIP.Error)
	}
	txUV := db.Table(tables.UVStatistics).Select("cookie, visit_time, count(*) as count").
		Where("visit_time = ?", today)
	if param.PrimaryID == 0 {
		return nil, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	txUV = txUV.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		txUV = txUV.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		txUV = txUV.Where("js_id = ?", param.JsID)
	}
	txUV = txUV.Group("visit_time, cookie").Scan(&uv)
	if txUV.Error != nil {
		return nil, nil, exception.Wrap(response.ExceptionDatabase, txUV.Error)
	}
	return ip, uv, nil
}

func (dri *daRepoImpl) YesterdayIPAndUVisit(db *gorm.DB, param *vo.JSFilterParams) ([]models.IPVisitStatistic,
	[]models.UVisitStatistic, exception.Exception) {
	ip := make([]models.IPVisitStatistic, 0)
	uv := make([]models.UVisitStatistic, 0)
	yesterday := time.Now().AddDate(0, 0, -1).Format(constant.DateFormat)
	txIP := db.Table(tables.IPStatistics).Select("ip, visit_time, count(*) as count").
		Where("visit_time = ?", yesterday)
	if param.PrimaryID == 0 {
		return nil, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	txIP = txIP.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		txIP = txIP.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		txIP = txIP.Where("js_id = ?", param.JsID)
	}
	txIP = txIP.Group("ip, visit_time").Scan(&ip)
	if txIP.Error != nil {
		return nil, nil, exception.Wrap(response.ExceptionDatabase, txIP.Error)
	}
	txUV := db.Table(tables.UVStatistics).Select("cookie, visit_time, count(*) as count").
		Where("visit_time = ?", yesterday)
	if param.PrimaryID == 0 {
		return nil, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	txUV = txUV.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		txUV = txUV.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		txUV = txUV.Where("js_id = ?", param.JsID)
	}
	txUV = txUV.Group("visit_time, cookie").Scan(&uv)
	if txUV.Error != nil {
		return nil, nil, exception.Wrap(response.ExceptionDatabase, txUV.Error)
	}
	return ip, uv, nil
}

func (dri *daRepoImpl) FromNowIPAndUVisit(db *gorm.DB, param *vo.JSFilterParams) ([]models.IPVisitStatistic,
	[]models.UVisitStatistic, exception.Exception) {
	ip := make([]models.IPVisitStatistic, 0)
	uv := make([]models.UVisitStatistic, 0)
	txIP := db.Table(tables.IPStatistics).Select("ip, visit_time, count(*) as count")
	if param.PrimaryID == 0 {
		return nil, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	txIP = txIP.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		txIP = txIP.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		txIP = txIP.Where("js_id = ?", param.JsID)
	}
	txIP = txIP.Group("ip, visit_time").Scan(&ip)
	if txIP.Error != nil {
		return nil, nil, exception.Wrap(response.ExceptionDatabase, txIP.Error)
	}
	txUV := db.Table(tables.UVStatistics).Select("cookie, visit_time, count(*) as count")
	if param.PrimaryID == 0 {
		return nil, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	txUV = txUV.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		txUV = txUV.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		txUV = txUV.Where("js_id = ?", param.JsID)
	}
	txUV = txUV.Group("visit_time, cookie").Scan(&uv)
	if txUV.Error != nil {
		return nil, nil, exception.Wrap(response.ExceptionDatabase, txUV.Error)
	}
	return ip, uv, nil
}

func (dri *daRepoImpl) TodayFlowData(db *gorm.DB, param *vo.JSFilterParams, pageInfo *vo.PageInfo) (int64,
	[]models.FlowDataStatistic, exception.Exception) {
	flowData := make([]models.FlowDataStatistic, 0)
	today := time.Now().Format(constant.DateFormat)
	subTx := db.Model(&models.FlowDataView{}).
		Where("ip_time = ? and uv_time = ?", today, today)
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
	tx := db.Table("(?) as sub", subTx).Select(
		"sub.title as title, sum(sub.ip_count) as ip_count, sum(sub.uv_count) as uv_count").
		Group("sub.title").
		Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).
		Scan(&flowData).Limit(-1).Offset(-1).Count(&count)
	return count, flowData, exception.Wrap(response.ExceptionDatabase, tx.Error)
}

func (dri *daRepoImpl) YesterdayFlowData(db *gorm.DB, param *vo.JSFilterParams, pageInfo *vo.PageInfo) (int64,
	[]models.FlowDataStatistic, exception.Exception) {
	flowData := make([]models.FlowDataStatistic, 0)
	yesterday := time.Now().AddDate(0, 0, -1).Format(constant.DateFormat)
	subTx := db.Model(&models.FlowDataView{}).
		Where("ip_time = ? and uv_time = ?", yesterday, yesterday)
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
	tx := db.Table("(?) as sub", subTx).Select(
		"sub.title as title, sum(sub.ip_count) as ip_count, sum(sub.uv_count) as uv_count").
		Group("sub.title").
		Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).
		Scan(&flowData).Limit(-1).Offset(-1).Count(&count)
	return count, flowData, exception.Wrap(response.ExceptionDatabase, tx.Error)
}

func (dri *daRepoImpl) TimeScopeFlowData(db *gorm.DB, param *vo.JSFilterParams, pageInfo *vo.PageInfo,
	beginAt, endAt string) (int64, []models.FlowDataStatistic, exception.Exception) {
	flowData := make([]models.FlowDataStatistic, 0)
	subTx := db.Model(&models.FlowDataView{}).
		Where("ip_time >= ? and ip_time <= ? and uv_time >= ? and uv_time <= ?", beginAt, endAt, beginAt, endAt)
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
	tx := db.Table("(?) as sub", subTx).Select(
		"sub.title as title, sum(sub.ip_count) as ip_count, sum(sub.uv_count) as uv_count").
		Group("sub.title").
		Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).
		Scan(&flowData).Limit(-1).Offset(-1).Count(&count)
	return count, flowData, exception.Wrap(response.ExceptionDatabase, tx.Error)
}

func (dri *daRepoImpl) FromNowFlowData(db *gorm.DB, param *vo.JSFilterParams, pageInfo *vo.PageInfo) (int64,
	[]models.FlowDataStatistic, exception.Exception) {
	flowData := make([]models.FlowDataStatistic, 0)
	subTx := db.Model(&models.FlowDataView{})
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
	tx := db.Table("(?) as sub", subTx).Select(
		"sub.title as title, sum(sub.ip_count) as ip_count, sum(sub.uv_count) as uv_count").
		Group("sub.title").
		Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).
		Scan(&flowData).Limit(-1).Offset(-1).Count(&count)
	return count, flowData, exception.Wrap(response.ExceptionDatabase, tx.Error)
}
