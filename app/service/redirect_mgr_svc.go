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
	rmServiceInstance RmService
	rmOnce            sync.Once
)

type rmServiceImpl struct {
	db   *gorm.DB
	repo repositories.RmRepo
}

func GetRmService() RmService {
	rmOnce.Do(func() {
		rmServiceInstance = &rmServiceImpl{
			db:   database.GetDriver(),
			repo: repositories.GetRmRepo(),
		}
	})
	return rmServiceInstance
}

type RmService interface {
	Create(openID string, param *vo.RedirectManageReq) exception.Exception
	Get(id int64) (*vo.RedirectManageResp, exception.Exception)
	List(page *vo.PageInfo) (*vo.DataPagination, exception.Exception)
	Update(openID string, id int64, param *vo.RedirectManageUpdateReq) exception.Exception
	Delete(id int64) exception.Exception
	MultiDelete(ids string) exception.Exception
}

func (rsi *rmServiceImpl) Create(openID string, param *vo.RedirectManageReq) exception.Exception {
	rmMgr := param.ToModel(openID)
	return rsi.repo.Create(rsi.db, rmMgr)
}

func (rsi *rmServiceImpl) Get(id int64) (*vo.RedirectManageResp, exception.Exception) {
	rm, ex := rsi.repo.Get(rsi.db, id)
	if ex != nil {
		return nil, ex
	}
	return vo.NewRedirectManageResponse(rm), nil
}

func (rsi *rmServiceImpl) List(pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception) {
	count, rms, ex := rsi.repo.List(rsi.db, pageInfo)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.RedirectManageResp, 0, len(rms))
	for i := range rms {
		resp = append(resp, *vo.NewRedirectManageResponse(&rms[i]))
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}

func (rsi *rmServiceImpl) Update(openID string, id int64, param *vo.RedirectManageUpdateReq) exception.Exception {
	return rsi.repo.Update(rsi.db, id, param.ToMap(openID))
}

func (rsi *rmServiceImpl) Delete(id int64) exception.Exception {
	return rsi.repo.Delete(rsi.db, id)
}

func (rsi *rmServiceImpl) MultiDelete(ids string) exception.Exception {
	idslice := strings.Split(ids, ",")
	if len(idslice) == 0 {
		return exception.New(response.ExceptionInvalidRequestParameters, "无效参数")
	}
	jid := make([]int64, 0, len(idslice))
	for i := range idslice {
		id, err := strconv.ParseUint(idslice[i], 10, 0)
		if err != nil {
			return exception.Wrap(response.ExceptionParseStringToInt64Error, err)
		}
		jid = append(jid, int64(id))
	}
	return rsi.repo.MultiDelete(rsi.db, jid)
}
