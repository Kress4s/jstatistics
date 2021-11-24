package service

import (
	"js_statistics/app/repositories"
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
	homeServiceInstance HomeService
	homeOnce            sync.Once
)

type homeServiceImpl struct {
	db   *gorm.DB
	repo repositories.HomeRepo
}

func GetHomeService() HomeService {
	homeOnce.Do(func() {
		homeServiceInstance = &homeServiceImpl{
			db:   database.GetDriver(),
			repo: repositories.GetHomeRepo(),
		}
	})
	return homeServiceInstance
}

type HomeService interface {
	TodayIP() (*vo.TodayIP, exception.Exception)
	YesterdayIP() (*vo.YesterdayIP, exception.Exception)
	ThisMonthIP() (*vo.ThisMonthIP, exception.Exception)
	LastMonthIP() (*vo.LastMonthIP, exception.Exception)
	IPAndUVisit() (*vo.HomeIPAndUVisit, exception.Exception)
	RegionStatistic() ([]vo.RegionStatisticResp, exception.Exception)
	JSVisitStatistic(pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception)
}

func (hsi *homeServiceImpl) TodayIP() (*vo.TodayIP, exception.Exception) {
	count, ex := hsi.repo.TodayIP(hsi.db)
	if ex != nil {
		return nil, ex
	}
	return &vo.TodayIP{Count: count}, nil
}

func (hsi *homeServiceImpl) YesterdayIP() (*vo.YesterdayIP, exception.Exception) {
	count, ex := hsi.repo.YesterdayIP(hsi.db)
	if ex != nil {
		return nil, ex
	}
	return &vo.YesterdayIP{Count: count}, nil
}

func (hsi *homeServiceImpl) ThisMonthIP() (*vo.ThisMonthIP, exception.Exception) {
	beginAt, endAt := tools.GetThisMonthTimeScope(time.Now())
	count, ex := hsi.repo.ThisMonthIP(hsi.db, beginAt, endAt)
	if ex != nil {
		return nil, ex
	}
	return &vo.ThisMonthIP{Count: count}, nil
}

func (hsi *homeServiceImpl) LastMonthIP() (*vo.LastMonthIP, exception.Exception) {
	beginAt, endAt := tools.GetLastMonthTimeScope(time.Now())
	count, ex := hsi.repo.LastMonthIP(hsi.db, beginAt, endAt)
	if ex != nil {
		return nil, ex
	}
	return &vo.LastMonthIP{Count: count}, nil
}

func (hsi *homeServiceImpl) IPAndUVisit() (*vo.HomeIPAndUVisit, exception.Exception) {
	beginAt, endAt := tools.GetLastMonthTimeScope(time.Now())
	ipData, uvData, ex := hsi.repo.IPAndUVisit(hsi.db, beginAt, endAt)
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
	for j := range uvData {
		uvVisit = append(uvVisit, vo.UVVisit{
			Count:  uvData[j].Count,
			Bucket: uvData[j].VisitTime.Format(constant.DateFormat),
		})
	}

	return &vo.HomeIPAndUVisit{IP: ipVisit, UV: uvVisit}, nil
}

func (hsi *homeServiceImpl) RegionStatistic() ([]vo.RegionStatisticResp, exception.Exception) {
	data, ex := hsi.repo.RegionStatistic(hsi.db)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.RegionStatisticResp, 0, len(data))
	for i := range data {
		resp = append(resp, vo.RegionStatisticResp{
			Region: data[i].Region,
			Count:  data[i].Count,
		})
	}
	return resp, nil
}

func (hsi *homeServiceImpl) JSVisitStatistic(pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception) {
	count, data, ex := hsi.repo.JSVisitStatistic(hsi.db, pageInfo)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.JSVisitStatisticResp, 0, len(data))
	for i, j := range resp {
		resp = append(resp, vo.JSVisitStatisticResp{
			Rank:  pageInfo.Offset() + 1 + i,
			Title: j.Title,
			Count: j.Count,
		})
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}