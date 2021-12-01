package repositories

import (
	"fmt"
	"js_statistics/app/models"
	"js_statistics/app/models/tables"
	"js_statistics/app/models/views"
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
	subIP := db.Table(tables.IPStatistics).Select("ip, visit_time, count(*) as count").
		Where("visit_time BETWEEN ? AND ?", beginAt, endAt)
	if param.PrimaryID == 0 {
		return nil, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	subIP = subIP.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		subIP = subIP.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		subIP = subIP.Where("js_id = ?", param.JsID)
	}
	subIP = subIP.Group("ip, visit_time")
	txIP := db.Table("(?) as sub", subIP).Select("sub.visit_time, sum(sub.count) as count").
		Group("sub.visit_time").Scan(&ip)
	if txIP.Error != nil {
		return nil, nil, exception.Wrap(response.ExceptionDatabase, txIP.Error)
	}
	subUV := db.Table(tables.UVStatistics).Select("cookie, visit_time, count(*) as count").
		Where("visit_time BETWEEN ? AND ?", beginAt, endAt)
	if param.PrimaryID == 0 {
		return nil, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	subUV = subUV.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		subUV = subUV.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		subUV = subUV.Where("js_id = ?", param.JsID)
	}
	subUV = subUV.Group("visit_time, cookie")
	txUV := db.Table("(?) as sub", subUV).Select("sub.visit_time, sum(sub.count) as count").
		Group("sub.visit_time").Scan(&uv)
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
	subIP := db.Table(tables.IPStatistics).Select("ip, visit_time, count(*) as count").
		Where("visit_time = ?", today)
	if param.PrimaryID == 0 {
		return nil, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	subIP = subIP.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		subIP = subIP.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		subIP = subIP.Where("js_id = ?", param.JsID)
	}
	subIP = subIP.Group("ip, visit_time")
	txIP := db.Table("(?) as sub", subIP).Select("sub.visit_time, sum(sub.count) as count").
		Group("sub.visit_time").Scan(&ip)
	if txIP.Error != nil {
		return nil, nil, exception.Wrap(response.ExceptionDatabase, txIP.Error)
	}
	subUV := db.Table(tables.UVStatistics).Select("cookie, visit_time, count(*) as count").
		Where("visit_time = ?", today)
	if param.PrimaryID == 0 {
		return nil, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	subUV = subUV.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		subUV = subUV.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		subUV = subUV.Where("js_id = ?", param.JsID)
	}
	subUV = subUV.Group("visit_time, cookie")
	txUV := db.Table("(?) as sub", subUV).Select("sub.visit_time, sum(sub.count) as count").
		Group("sub.visit_time").Scan(&uv)
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
	subIP := db.Table(tables.IPStatistics).Select("ip, visit_time, count(*) as count").
		Where("visit_time = ?", yesterday)
	if param.PrimaryID == 0 {
		return nil, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	subIP = subIP.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		subIP = subIP.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		subIP = subIP.Where("js_id = ?", param.JsID)
	}
	subIP = subIP.Group("ip, visit_time")
	txIP := db.Table("(?) as sub", subIP).Select("sub.visit_time, sum(sub.count) as count").
		Group("sub.visit_time").Scan(&ip)
	if txIP.Error != nil {
		return nil, nil, exception.Wrap(response.ExceptionDatabase, txIP.Error)
	}
	subUV := db.Table(tables.UVStatistics).Select("cookie, visit_time, count(*) as count").
		Where("visit_time = ?", yesterday)
	if param.PrimaryID == 0 {
		return nil, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	subUV = subUV.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		subUV = subUV.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		subUV = subUV.Where("js_id = ?", param.JsID)
	}
	subUV = subUV.Group("visit_time, cookie")
	txUV := db.Table("(?) as sub", subUV).Select("sub.visit_time, sum(sub.count) as count").
		Group("sub.visit_time").Scan(&uv)
	if txUV.Error != nil {
		return nil, nil, exception.Wrap(response.ExceptionDatabase, txUV.Error)
	}
	return ip, uv, nil
}

func (dri *daRepoImpl) FromNowIPAndUVisit(db *gorm.DB, param *vo.JSFilterParams) ([]models.IPVisitStatistic,
	[]models.UVisitStatistic, exception.Exception) {
	ip := make([]models.IPVisitStatistic, 0)
	uv := make([]models.UVisitStatistic, 0)
	subIP := db.Table(tables.IPStatistics).Select("ip, visit_time, count(*) as count")
	if param.PrimaryID == 0 {
		return nil, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	subIP = subIP.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		subIP = subIP.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		subIP = subIP.Where("js_id = ?", param.JsID)
	}
	subIP = subIP.Group("ip, visit_time")
	txIP := db.Table("(?) as sub", subIP).Select("sub.visit_time, sum(sub.count) as count").
		Group("sub.visit_time").Scan(&ip)
	if txIP.Error != nil {
		return nil, nil, exception.Wrap(response.ExceptionDatabase, txIP.Error)
	}
	subUV := db.Table(tables.UVStatistics).Select("cookie, visit_time, count(*) as count")
	if param.PrimaryID == 0 {
		return nil, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	subUV = subUV.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		subUV = subUV.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		subUV = subUV.Where("js_id = ?", param.JsID)
	}
	subUV = subUV.Group("visit_time, cookie")
	txUV := db.Table("(?) as sub", subUV).Select("sub.visit_time, sum(sub.count) as count").
		Group("sub.visit_time").Scan(&uv)
	if txUV.Error != nil {
		return nil, nil, exception.Wrap(response.ExceptionDatabase, txUV.Error)
	}
	return ip, uv, nil
}

func (dri *daRepoImpl) TodayFlowData(db *gorm.DB, param *vo.JSFilterParams, pageInfo *vo.PageInfo) (int64,
	[]models.FlowDataStatistic, exception.Exception) {
	today := time.Now().Format(constant.DateFormat)
	flowData := make([]models.FlowDataStatistic, 0)
	subIPTx := db.Table(views.IPFlowDataView).Select("js_id, sum(count) as count")
	if param.PrimaryID == 0 {
		return 0, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	subIPTx = subIPTx.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		subIPTx = subIPTx.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		subIPTx = subIPTx.Where("js_id = ?", param.JsID)
	}
	subIPTx = subIPTx.Where("visit_time = ?", today)
	subIPTx = subIPTx.Group("js_id")

	subUVTx := db.Table(views.UVFlowDataView).Select("js_id, sum(count) as count")
	subUVTx = subUVTx.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		subUVTx = subUVTx.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		subUVTx = subUVTx.Where("js_id = ?", param.JsID)
	}
	subUVTx = subUVTx.Where("visit_time = ?", today)
	subUVTx = subUVTx.Group("js_id")

	count := int64(0)
	jcpTX := db.Table(tables.JsManage + " AS js").
		Select("js.id AS js_id, js.title AS title, jc.id AS category_id, jc.primary_id").
		Joins(fmt.Sprintf("INNER JOIN %s AS jc ON js.category_id = jc.id", tables.JsCategory))
	jcpTX = jcpTX.Where("jc.primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		jcpTX = jcpTX.Where("jc.id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		jcpTX = jcpTX.Where("js.id = ?", param.JsID)
	}
	tx := db.Table("(?) AS jst", jcpTX).
		Select("jst.title AS title, ip_flow.count AS ip_count, uv_flow.count AS uv_count").
		Joins("LEFT JOIN (?) AS ip_flow on jst.js_id = ip_flow.js_id", subIPTx).
		Joins("LEFT JOIN (?) AS uv_flow on jst.js_id = uv_flow.js_id", subUVTx).
		Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).
		Scan(&flowData).Limit(-1).Offset(-1).Count(&count)
	return count, flowData, exception.Wrap(response.ExceptionDatabase, tx.Error)
}

func (dri *daRepoImpl) YesterdayFlowData(db *gorm.DB, param *vo.JSFilterParams, pageInfo *vo.PageInfo) (int64,
	[]models.FlowDataStatistic, exception.Exception) {
	yesterday := time.Now().AddDate(0, 0, -1).Format(constant.DateFormat)
	flowData := make([]models.FlowDataStatistic, 0)
	subIPTx := db.Table(views.IPFlowDataView).Select("js_id, sum(count) as count")
	if param.PrimaryID == 0 {
		return 0, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	subIPTx = subIPTx.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		subIPTx = subIPTx.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		subIPTx = subIPTx.Where("js_id = ?", param.JsID)
	}
	subIPTx = subIPTx.Where("visit_time = ?", yesterday)
	subIPTx = subIPTx.Group("js_id")

	subUVTx := db.Table(views.UVFlowDataView).Select("js_id, sum(count) as count")
	subUVTx = subUVTx.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		subUVTx = subUVTx.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		subUVTx = subUVTx.Where("js_id = ?", param.JsID)
	}
	subUVTx = subUVTx.Where("visit_time = ?", yesterday)
	subUVTx = subUVTx.Group("js_id")

	count := int64(0)
	jcpTX := db.Table(tables.JsManage + " AS js").
		Select("js.id AS js_id, js.title AS title, jc.id AS category_id, jc.primary_id").
		Joins(fmt.Sprintf("INNER JOIN %s AS jc ON js.category_id = jc.id", tables.JsCategory))
	jcpTX = jcpTX.Where("jc.primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		jcpTX = jcpTX.Where("jc.id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		jcpTX = jcpTX.Where("js.id = ?", param.JsID)
	}
	tx := db.Table("(?) AS jst", jcpTX).
		Select("jst.title AS title, ip_flow.count AS ip_count, uv_flow.count AS uv_count").
		Joins("LEFT JOIN (?) AS ip_flow on jst.js_id = ip_flow.js_id", subIPTx).
		Joins("LEFT JOIN (?) AS uv_flow on jst.js_id = uv_flow.js_id", subUVTx).
		Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).
		Scan(&flowData).Limit(-1).Offset(-1).Count(&count)
	return count, flowData, exception.Wrap(response.ExceptionDatabase, tx.Error)
}

func (dri *daRepoImpl) TimeScopeFlowData(db *gorm.DB, param *vo.JSFilterParams, pageInfo *vo.PageInfo,
	beginAt, endAt string) (int64, []models.FlowDataStatistic, exception.Exception) {
	flowData := make([]models.FlowDataStatistic, 0)
	subIPTx := db.Table(views.IPFlowDataView).Select("js_id, sum(count) as count")
	if param.PrimaryID == 0 {
		return 0, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	subIPTx = subIPTx.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		subIPTx = subIPTx.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		subIPTx = subIPTx.Where("js_id = ?", param.JsID)
	}
	subIPTx = subIPTx.Where("visit_time >= ? and visit_time < ?", beginAt, endAt)
	subIPTx = subIPTx.Group("js_id")

	subUVTx := db.Table(views.UVFlowDataView).Select("js_id, sum(count) as count")
	subUVTx = subUVTx.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		subUVTx = subUVTx.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		subUVTx = subUVTx.Where("js_id = ?", param.JsID)
	}
	subUVTx = subUVTx.Where("visit_time >= ? and visit_time < ?", beginAt, endAt)
	subUVTx = subUVTx.Group("js_id")

	count := int64(0)
	jcpTX := db.Table(tables.JsManage + " AS js").
		Select("js.id AS js_id, js.title AS title, jc.id AS category_id, jc.primary_id").
		Joins(fmt.Sprintf("INNER JOIN %s AS jc ON js.category_id = jc.id", tables.JsCategory))
	jcpTX = jcpTX.Where("jc.primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		jcpTX = jcpTX.Where("jc.id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		jcpTX = jcpTX.Where("js.id = ?", param.JsID)
	}
	tx := db.Table("(?) AS jst", jcpTX).
		Select("jst.title AS title, ip_flow.count AS ip_count, uv_flow.count AS uv_count").
		Joins("LEFT JOIN (?) AS ip_flow on jst.js_id = ip_flow.js_id", subIPTx).
		Joins("LEFT JOIN (?) AS uv_flow on jst.js_id = uv_flow.js_id", subUVTx).
		Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).
		Scan(&flowData).Limit(-1).Offset(-1).Count(&count)
	return count, flowData, exception.Wrap(response.ExceptionDatabase, tx.Error)
}

func (dri *daRepoImpl) FromNowFlowData(db *gorm.DB, param *vo.JSFilterParams, pageInfo *vo.PageInfo) (int64,
	[]models.FlowDataStatistic, exception.Exception) {
	flowData := make([]models.FlowDataStatistic, 0)
	subIPTx := db.Table(views.IPFlowDataView).Select("js_id, sum(count) as count")
	if param.PrimaryID == 0 {
		return 0, nil, exception.New(response.ExceptionInvalidRequestParameters, "primary_id must choose")
	}
	subIPTx = subIPTx.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		subIPTx = subIPTx.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		subIPTx = subIPTx.Where("js_id = ?", param.JsID)
	}
	subIPTx = subIPTx.Group("js_id")

	subUVTx := db.Table(views.UVFlowDataView).Select("js_id, sum(count) as count")
	subUVTx = subUVTx.Where("primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		subUVTx = subUVTx.Where("category_id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		subUVTx = subUVTx.Where("js_id = ?", param.JsID)
	}
	subUVTx = subUVTx.Group("js_id")
	count := int64(0)
	jcpTX := db.Table(tables.JsManage + " AS js").
		Select("js.id AS js_id, js.title AS title, jc.id AS category_id, jc.primary_id").
		Joins(fmt.Sprintf("INNER JOIN %s AS jc ON js.category_id = jc.id", tables.JsCategory))
	jcpTX = jcpTX.Where("jc.primary_id = ?", param.PrimaryID)
	if param.CategoryID != 0 {
		jcpTX = jcpTX.Where("jc.id = ?", param.CategoryID)
	}
	if param.JsID != 0 {
		jcpTX = jcpTX.Where("js.id = ?", param.JsID)
	}
	tx := db.Table("(?) AS jst", jcpTX).
		Select("jst.title AS title, ip_flow.count AS ip_count, uv_flow.count AS uv_count").
		Joins("LEFT JOIN (?) AS ip_flow on jst.js_id = ip_flow.js_id", subIPTx).
		Joins("LEFT JOIN (?) AS uv_flow on jst.js_id = uv_flow.js_id", subUVTx).
		Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).
		Scan(&flowData).Limit(-1).Offset(-1).Count(&count)
	return count, flowData, exception.Wrap(response.ExceptionDatabase, tx.Error)
}
