package service

import (
	"js_statistics/app/repositories"
	"js_statistics/app/vo"
	"js_statistics/commom/drivers/database"
	"js_statistics/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	faServiceInstance FaService
	faOnce            sync.Once
)

type faServiceImpl struct {
	db   *gorm.DB
	repo repositories.FaRepo
}

func GetFaService() FaService {
	faOnce.Do(func() {
		faServiceInstance = &faServiceImpl{
			db:   database.GetDriver(),
			repo: repositories.GetFaRepo(),
		}
	})
	return faServiceInstance
}

type FaService interface {
	FromStatistic(param *vo.JSFilterParams, pageInfo *vo.PageInfo, beginAt, endAt string) (*vo.DataPagination,
		exception.Exception)
}

func (fsi *faServiceImpl) FromStatistic(param *vo.JSFilterParams, pageInfo *vo.PageInfo, beginAt, endAt string,
) (*vo.DataPagination, exception.Exception) {
	count, data, ex := fsi.repo.FromStatistic(fsi.db, param, pageInfo, beginAt, endAt)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.FromAnalysisResp, 0, len(data))
	for i := range data {
		resp = append(resp, vo.FromAnalysisResp{
			Title: data[i].Title,
			From:  data[i].FromURL,
			To:    data[i].ToUrl,
			Count: data[i].Count,
		})
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}
