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
	syslogServiceInstance SyslogService
	syslogOnce            sync.Once
)

type syslogServiceImpl struct {
	db   *gorm.DB
	repo repositories.SyslogRepo
}

func GetSyslogService() SyslogService {
	syslogOnce.Do(func() {
		syslogServiceInstance = &syslogServiceImpl{
			db:   database.GetDriver(),
			repo: repositories.GetSyslogRepo(),
		}
	})
	return syslogServiceInstance
}

type SyslogService interface {
	Create(param *vo.SystemLogReq) exception.Exception
	List(page *vo.PageInfo) (*vo.DataPagination, exception.Exception)
}

func (sli *syslogServiceImpl) Create(param *vo.SystemLogReq) exception.Exception {
	log := param.ToModel()
	return sli.repo.Create(sli.db, log)
}

func (sli *syslogServiceImpl) List(pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception) {
	count, logs, ex := sli.repo.List(sli.db, pageInfo)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.SystemLogResp, 0, len(logs))
	for i := range logs {
		resp = append(resp, *vo.NewSysLogResp(&logs[i]))
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}
