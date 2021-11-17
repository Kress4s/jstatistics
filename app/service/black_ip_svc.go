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
	blackIPServiceInstance BlackIPService
	blackIPOnce            sync.Once
)

type blackIPServiceImpl struct {
	db   *gorm.DB
	repo repositories.BlackIPRepo
}

func GetBlackIPService() BlackIPService {
	blackIPOnce.Do(func() {
		blackIPServiceInstance = &blackIPServiceImpl{
			db:   database.GetDriver(),
			repo: repositories.GetBlackIPRepo(),
		}
	})
	return blackIPServiceInstance
}

type BlackIPService interface {
	Create(openID string, param *vo.BlackIPReq) exception.Exception
	Get(id int64) (*vo.BlackIPResp, exception.Exception)
	List(page *vo.PageInfo) (*vo.DataPagination, exception.Exception)
	Update(openID string, id int64, param *vo.BlackIPUpdateReq) exception.Exception
	Delete(id int64) exception.Exception
	MultiDelete(ids string) exception.Exception
}

func (dsi *blackIPServiceImpl) Create(openID string, param *vo.BlackIPReq) exception.Exception {
	blackIPMgr := param.ToModel(openID)
	return dsi.repo.Create(dsi.db, blackIPMgr)
}

func (dsi *blackIPServiceImpl) Get(id int64) (*vo.BlackIPResp, exception.Exception) {
	blackIPMgr, ex := dsi.repo.Get(dsi.db, id)
	if ex != nil {
		return nil, ex
	}
	return vo.NewBlackIPResponse(blackIPMgr), nil
}

func (dsi *blackIPServiceImpl) List(pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception) {
	count, ips, ex := dsi.repo.List(dsi.db, pageInfo)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.BlackIPResp, 0, len(ips))
	for i := range ips {
		resp = append(resp, *vo.NewBlackIPResponse(&ips[i]))
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}

func (dsi *blackIPServiceImpl) Update(openID string, id int64, param *vo.BlackIPUpdateReq) exception.Exception {
	return dsi.repo.Update(dsi.db, id, param.ToMap(openID))
}

func (dsi *blackIPServiceImpl) Delete(id int64) exception.Exception {
	return dsi.repo.Delete(dsi.db, id)
}

func (dsi *blackIPServiceImpl) MultiDelete(ids string) exception.Exception {
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
