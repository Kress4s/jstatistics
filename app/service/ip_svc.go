package service

import (
	"js_statistics/app/repositories"
	"js_statistics/app/response"
	"js_statistics/app/vo"
	"js_statistics/commom/drivers/database"
	"js_statistics/exception"
	"strconv"
	"strings"
	"sync"

	"gorm.io/gorm"
)

var (
	ipServiceInstance WhiteIPService
	ipOnce            sync.Once
)

type ipServiceImpl struct {
	db   *gorm.DB
	repo repositories.IPRepo
}

func GetWhiteIPService() WhiteIPService {
	ipOnce.Do(func() {
		ipServiceInstance = &ipServiceImpl{
			db:   database.GetDriver(),
			repo: repositories.GetIPRepo(),
		}
	})
	return ipServiceInstance
}

type WhiteIPService interface {
	Create(openID string, param *vo.IPReq) exception.Exception
	Get(id int64) (*vo.IPResp, exception.Exception)
	List(page *vo.PageInfo) (*vo.DataPagination, exception.Exception)
	Update(openID string, id int64, param *vo.IPUpdateReq) exception.Exception
	Delete(id int64) exception.Exception
	MultiDelete(ids string) exception.Exception
}

func (dsi *ipServiceImpl) Create(openID string, param *vo.IPReq) exception.Exception {
	ipMgr := param.ToModel(openID)
	return dsi.repo.Create(dsi.db, ipMgr)
}

func (dsi *ipServiceImpl) Get(id int64) (*vo.IPResp, exception.Exception) {
	ipMgr, ex := dsi.repo.Get(dsi.db, id)
	if ex != nil {
		return nil, ex
	}
	return vo.NewIPResponse(ipMgr), nil
}

func (dsi *ipServiceImpl) List(pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception) {
	count, ips, ex := dsi.repo.List(dsi.db, pageInfo)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.IPResp, 0, len(ips))
	for i := range ips {
		resp = append(resp, *vo.NewIPResponse(&ips[i]))
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}

func (dsi *ipServiceImpl) Update(openID string, id int64, param *vo.IPUpdateReq) exception.Exception {
	return dsi.repo.Update(dsi.db, id, param.ToMap(openID))
}

func (dsi *ipServiceImpl) Delete(id int64) exception.Exception {
	return dsi.repo.Delete(dsi.db, id)
}

func (dsi *ipServiceImpl) MultiDelete(ids string) exception.Exception {
	idslice := strings.Split(ids, ",")
	if len(idslice) == 0 {
		return exception.New(response.ExceptionInvalidRequestParameters, "无效参数")
	}
	did := make([]int64, 0, len(idslice))
	for i := range idslice {
		id, err := strconv.ParseUint(idslice[i], 10, 0)
		if err != nil {
			return exception.Wrap(response.ExceptionParseStringToUintError, err)
		}
		did = append(did, int64(id))
	}
	return dsi.repo.MultiDelete(dsi.db, did)
}
