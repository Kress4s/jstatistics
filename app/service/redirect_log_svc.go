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
	rlServiceInstance RlService
	rlOnce            sync.Once
)

type rlServiceImpl struct {
	db   *gorm.DB
	repo repositories.RlRepo
}

func GetRlService() RlService {
	rlOnce.Do(func() {
		rlServiceInstance = &rlServiceImpl{
			db:   database.GetDriver(),
			repo: repositories.GetRlRepo(),
		}
	})
	return rlServiceInstance
}

type RlService interface {
	ListByCategoryID(page *vo.PageInfo, cid int64) (*vo.DataPagination, exception.Exception)
}

func (rsi *rlServiceImpl) ListByCategoryID(page *vo.PageInfo, cid int64) (*vo.DataPagination, exception.Exception) {
	count, rms, ex := rsi.repo.ListByCategoryID(rsi.db, page, cid)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.RedirectLogResp, 0, len(rms))
	for i := range rms {
		resp = append(resp, *vo.NewRedirectLogResp(&rms[i]))
	}
	return vo.NewDataPagination(count, resp, page), nil
}
