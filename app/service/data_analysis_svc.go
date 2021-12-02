package service

import (
	"js_statistics/app/repositories"
	"js_statistics/app/response"
	"js_statistics/app/vo"
	"js_statistics/commom/drivers/database"
	"js_statistics/commom/tools"
	"js_statistics/constant"
	"js_statistics/exception"
	"sync"
	"time"

	"gorm.io/gorm"
)

var (
	daServiceInstance DaService
	daOnce            sync.Once
)

type daServiceImpl struct {
	db   *gorm.DB
	repo repositories.DaRepo
}

func GetDaService() DaService {
	daOnce.Do(func() {
		daServiceInstance = &daServiceImpl{
			db:   database.GetDriver(),
			repo: repositories.GetDaRepo(),
		}
	})
	return daServiceInstance
}

type DaService interface {
	TodayIP(param *vo.JSFilterParams) (*vo.TodayIP, exception.Exception)
	YesterdayIP(param *vo.JSFilterParams) (*vo.YesterdayIP, exception.Exception)
	ThisMonthIP(param *vo.JSFilterParams) (*vo.ThisMonthIP, exception.Exception)
	LastMonthIP(param *vo.JSFilterParams) (*vo.LastMonthIP, exception.Exception)
	IPAndUVisit(param *vo.JSFilterParams, beginAt, endAt string) (*vo.HomeIPAndUVisit, exception.Exception)
	TodayIPAndUVisit(param *vo.JSFilterParams) (*vo.HomeIPAndUVisit, exception.Exception)
	YesterdayIPAndUVisit(param *vo.JSFilterParams) (*vo.HomeIPAndUVisit, exception.Exception)
	FromNowIPAndUVisit(param *vo.JSFilterParams) (*vo.HomeIPAndUVisit, exception.Exception)
	TodayFlowData(param *vo.JSFilterParams, pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception)
	YesterdayFlowData(param *vo.JSFilterParams, pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception)
	TimeScopeFlowData(param *vo.JSFilterParams, pageInfo *vo.PageInfo, beginAt, endAt string,
	) (*vo.DataPagination, exception.Exception)
	FromNowFlowData(param *vo.JSFilterParams, pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception)
}

func (dsi *daServiceImpl) TodayIP(param *vo.JSFilterParams) (*vo.TodayIP, exception.Exception) {
	count, ex := dsi.repo.TodayIP(dsi.db, param)
	if ex != nil {
		return nil, ex
	}
	return &vo.TodayIP{Count: count}, nil
}

func (dsi *daServiceImpl) YesterdayIP(param *vo.JSFilterParams) (*vo.YesterdayIP, exception.Exception) {
	count, ex := dsi.repo.YesterdayIP(dsi.db, param)
	if ex != nil {
		return nil, ex
	}
	return &vo.YesterdayIP{Count: count}, nil
}

func (dsi *daServiceImpl) ThisMonthIP(param *vo.JSFilterParams) (*vo.ThisMonthIP, exception.Exception) {
	beginAt, endAt := tools.GetThisMonthTimeScope(time.Now())
	count, ex := dsi.repo.ThisMonthIP(dsi.db, param, beginAt.Format(constant.DateFormat),
		endAt.Format(constant.DateFormat))
	if ex != nil {
		return nil, ex
	}
	return &vo.ThisMonthIP{Count: count}, nil
}

func (dsi *daServiceImpl) LastMonthIP(param *vo.JSFilterParams) (*vo.LastMonthIP, exception.Exception) {
	beginAt, endAt := tools.GetLastMonthTimeScope(time.Now())
	count, ex := dsi.repo.LastMonthIP(dsi.db, param, beginAt.Format(constant.DateFormat),
		endAt.Format(constant.DateFormat))
	if ex != nil {
		return nil, ex
	}
	return &vo.LastMonthIP{Count: count}, nil
}

func (dsi *daServiceImpl) IPAndUVisit(param *vo.JSFilterParams, beginAt, endAt string) (*vo.HomeIPAndUVisit,
	exception.Exception) {
	ipData, uvData, ex := dsi.repo.IPAndUVisit(dsi.db, param, beginAt, endAt)
	if ex != nil {
		return nil, ex
	}
	ipVisit := make([]vo.IPVisit, 0, len(ipData))
	uvVisit := make([]vo.UVVisit, 0, len(ipData))
	// 生成连续时间
	begin, err := time.Parse(constant.DateFormat, beginAt)
	end, _err := time.Parse(constant.DateFormat, endAt)
	if err != nil || _err != nil {
		return nil, exception.Wrap(response.ExceptionParseDate, _err)
	}
	buckets := tools.DayIterator(begin, end)
	for i := range buckets {
		isExist := false
		for j := range ipData {
			if ipData[j].VisitTime.Format(constant.DateFormat) == buckets[i] {
				ipVisit = append(ipVisit, vo.IPVisit{
					Count:  ipData[j].Count,
					Bucket: ipData[j].VisitTime.Format(constant.DateFormat),
				})
				isExist = true
				break
			}
		}
		if !isExist {
			ipVisit = append(ipVisit, vo.IPVisit{
				Count:  0,
				Bucket: buckets[i],
			})
		}
	}

	for i := range buckets {
		isExist := false
		for j := range uvData {
			if uvData[j].VisitTime.Format(constant.DateFormat) == buckets[i] {
				uvVisit = append(uvVisit, vo.UVVisit{
					Count:  uvData[j].Count,
					Bucket: uvData[j].VisitTime.Format(constant.DateFormat),
				})
				isExist = true
				break
			}
		}
		if !isExist {
			uvVisit = append(uvVisit, vo.UVVisit{
				Count:  0,
				Bucket: buckets[i],
			})
		}
	}
	return &vo.HomeIPAndUVisit{IP: ipVisit, UV: uvVisit}, nil
}

func (dsi *daServiceImpl) TodayIPAndUVisit(param *vo.JSFilterParams) (*vo.HomeIPAndUVisit, exception.Exception) {
	ipData, uvData, ex := dsi.repo.TodayIPAndUVisit(dsi.db, param)
	if ex != nil {
		return nil, ex
	}
	ipVisit := make([]vo.IPVisit, 0, len(ipData))
	uvVisit := make([]vo.UVVisit, 0, len(ipData))
	for i := range ipData {
		ipVisit = append(ipVisit, vo.IPVisit{
			Count:  ipData[i].Count,
			Bucket: ipData[i].VisitTime.Format(constant.DateFormat),
		})
	}
	if len(ipVisit) == 0 {
		ipVisit = append(ipVisit, vo.IPVisit{
			Count:  0,
			Bucket: time.Now().Format(constant.DateFormat),
		})
	}
	for j := range uvData {
		uvVisit = append(uvVisit, vo.UVVisit{
			Count:  uvData[j].Count,
			Bucket: uvData[j].VisitTime.Format(constant.DateFormat),
		})
	}
	if len(uvData) == 0 {
		uvVisit = append(uvVisit, vo.UVVisit{
			Count:  0,
			Bucket: time.Now().Format(constant.DateFormat),
		})
	}
	return &vo.HomeIPAndUVisit{IP: ipVisit, UV: uvVisit}, nil
}

func (dsi *daServiceImpl) YesterdayIPAndUVisit(param *vo.JSFilterParams) (*vo.HomeIPAndUVisit, exception.Exception) {
	ipData, uvData, ex := dsi.repo.YesterdayIPAndUVisit(dsi.db, param)
	if ex != nil {
		return nil, ex
	}
	ipVisit := make([]vo.IPVisit, 0, len(ipData))
	uvVisit := make([]vo.UVVisit, 0, len(ipData))
	for i := range ipData {
		ipVisit = append(ipVisit, vo.IPVisit{
			Count:  ipData[i].Count,
			Bucket: ipData[i].VisitTime.Format(constant.DateFormat),
		})
	}
	if len(ipVisit) == 0 {
		ipVisit = append(ipVisit, vo.IPVisit{
			Count:  0,
			Bucket: time.Now().Format(constant.DateFormat),
		})
	}
	for j := range uvData {
		uvVisit = append(uvVisit, vo.UVVisit{
			Count:  uvData[j].Count,
			Bucket: uvData[j].VisitTime.Format(constant.DateFormat),
		})
	}
	if len(uvData) == 0 {
		uvVisit = append(uvVisit, vo.UVVisit{
			Count:  0,
			Bucket: time.Now().Format(constant.DateFormat),
		})
	}
	return &vo.HomeIPAndUVisit{IP: ipVisit, UV: uvVisit}, nil
}

func (dsi *daServiceImpl) FromNowIPAndUVisit(param *vo.JSFilterParams) (*vo.HomeIPAndUVisit, exception.Exception) {
	ipData, uvData, ex := dsi.repo.FromNowIPAndUVisit(dsi.db, param)
	if ex != nil {
		return nil, ex
	}
	ipVisit := make([]vo.IPVisit, 0, len(ipData))
	uvVisit := make([]vo.UVVisit, 0, len(ipData))
	// 生成连续时间
	IPBuckets := make([]string, 0)
	if len(ipData) > 0 {
		IPBuckets = tools.DayIterator(ipData[0].VisitTime, time.Now())
	}
	for i := range IPBuckets {
		isExist := false
		for j := range ipData {
			if ipData[j].VisitTime.Format(constant.DateFormat) == IPBuckets[i] {
				ipVisit = append(ipVisit, vo.IPVisit{
					Count:  ipData[j].Count,
					Bucket: ipData[j].VisitTime.Format(constant.DateFormat),
				})
				isExist = true
				break
			}
		}
		if !isExist {
			ipVisit = append(ipVisit, vo.IPVisit{
				Count:  0,
				Bucket: IPBuckets[i],
			})
		}
	}

	// 生成连续时间
	UVBuckets := make([]string, 0)
	if len(uvData) > 0 {
		UVBuckets = tools.DayIterator(uvData[0].VisitTime, time.Now())
	}
	for i := range UVBuckets {
		isExist := false
		for j := range uvData {
			if uvData[j].VisitTime.Format(constant.DateFormat) == UVBuckets[i] {
				uvVisit = append(uvVisit, vo.UVVisit{
					Count:  uvData[j].Count,
					Bucket: uvData[j].VisitTime.Format(constant.DateFormat),
				})
				isExist = true
				break
			}
		}
		if !isExist {
			uvVisit = append(uvVisit, vo.UVVisit{
				Count:  0,
				Bucket: UVBuckets[i],
			})
		}
	}
	return &vo.HomeIPAndUVisit{IP: ipVisit, UV: uvVisit}, nil
}

func (dsi *daServiceImpl) TodayFlowData(param *vo.JSFilterParams, pageInfo *vo.PageInfo) (*vo.DataPagination,
	exception.Exception) {
	count, data, ex := dsi.repo.TodayFlowData(dsi.db, param, pageInfo)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.FlowDataResp, 0, len(data))
	for i := range data {
		resp = append(resp, vo.FlowDataResp{
			Title: data[i].Title,
			IP:    data[i].IPCount,
			UV:    data[i].UVCount,
		})
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}

func (dsi *daServiceImpl) YesterdayFlowData(param *vo.JSFilterParams, pageInfo *vo.PageInfo) (*vo.DataPagination,
	exception.Exception) {
	count, data, ex := dsi.repo.YesterdayFlowData(dsi.db, param, pageInfo)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.FlowDataResp, 0, len(data))
	for i := range data {
		resp = append(resp, vo.FlowDataResp{
			Title: data[i].Title,
			IP:    data[i].IPCount,
			UV:    data[i].UVCount,
		})
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}

func (dsi *daServiceImpl) TimeScopeFlowData(param *vo.JSFilterParams, pageInfo *vo.PageInfo, beginAt, endAt string,
) (*vo.DataPagination, exception.Exception) {
	count, data, ex := dsi.repo.TimeScopeFlowData(dsi.db, param, pageInfo, beginAt, endAt)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.FlowDataResp, 0, len(data))
	for i := range data {
		resp = append(resp, vo.FlowDataResp{
			Title: data[i].Title,
			IP:    data[i].IPCount,
			UV:    data[i].UVCount,
		})
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}

func (dsi *daServiceImpl) FromNowFlowData(param *vo.JSFilterParams, pageInfo *vo.PageInfo) (*vo.DataPagination,
	exception.Exception) {
	count, data, ex := dsi.repo.FromNowFlowData(dsi.db, param, pageInfo)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.FlowDataResp, 0, len(data))
	for i := range data {
		resp = append(resp, vo.FlowDataResp{
			Title: data[i].Title,
			IP:    data[i].IPCount,
			UV:    data[i].UVCount,
		})
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}
