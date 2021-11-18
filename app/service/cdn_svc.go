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
	cdnServiceInstance CdnService
	cdnOnce            sync.Once
)

type cdnServiceImpl struct {
	db   *gorm.DB
	repo repositories.CdnRepo
}

func GetCdnService() CdnService {
	cdnOnce.Do(func() {
		cdnServiceInstance = &cdnServiceImpl{
			db:   database.GetDriver(),
			repo: repositories.GetCdnRepo(),
		}
	})
	return cdnServiceInstance
}

type CdnService interface {
	Create(openID string, param *vo.CDNReq) exception.Exception
	Get(id int64) (*vo.CDNResp, exception.Exception)
	List(page *vo.PageInfo) (*vo.DataPagination, exception.Exception)
	Update(openID string, id int64, param *vo.CDNUpdateReq) exception.Exception
	Delete(id int64) exception.Exception
	MultiDelete(ids string) exception.Exception
}

func (csi *cdnServiceImpl) Create(openID string, param *vo.CDNReq) exception.Exception {
	cdnMgr := param.ToModel(openID)
	return csi.repo.Create(csi.db, cdnMgr)
}

func (csi *cdnServiceImpl) Get(id int64) (*vo.CDNResp, exception.Exception) {
	cdnMgr, ex := csi.repo.Get(csi.db, id)
	if ex != nil {
		return nil, ex
	}
	return vo.NewCDNResponse(cdnMgr), nil
}

func (csi *cdnServiceImpl) List(pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception) {
	count, cdns, ex := csi.repo.List(csi.db, pageInfo)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.CDNResp, 0, len(cdns))
	for i := range cdns {
		resp = append(resp, *vo.NewCDNResponse(&cdns[i]))
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}

func (csi *cdnServiceImpl) Update(openID string, id int64, param *vo.CDNUpdateReq) exception.Exception {
	return csi.repo.Update(csi.db, id, param.ToMap(openID))
}

func (csi *cdnServiceImpl) Delete(id int64) exception.Exception {
	return csi.repo.Delete(csi.db, id)
}

func (csi *cdnServiceImpl) MultiDelete(ids string) exception.Exception {
	idslice := strings.Split(ids, ",")
	if len(idslice) == 0 {
		return exception.New(response.ExceptionInvalidRequestParameters, "无效参数")
	}
	did := make([]int64, 0, len(idslice))
	for i := range idslice {
		id, err := strconv.ParseUint(idslice[i], 10, 0)
		if err != nil {
			return exception.Wrap(response.ExceptionParseStringToInt64Error, err)
		}
		did = append(did, int64(id))
	}
	return csi.repo.MultiDelete(csi.db, did)
}
