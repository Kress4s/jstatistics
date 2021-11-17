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
	jscServiceInstance JscService
	jscOnce            sync.Once
)

type jscServiceImpl struct {
	db         *gorm.DB
	repo       repositories.JscRepo
	domainRepo repositories.DomainRepo
	jspRepo    repositories.JspRepo
}

func GetJscService() JscService {
	jscOnce.Do(func() {
		jscServiceInstance = &jscServiceImpl{
			db:         database.GetDriver(),
			repo:       repositories.GetJscRepo(),
			domainRepo: repositories.GetDomainRepo(),
			jspRepo:    repositories.GetJspRepo(),
		}
	})
	return jscServiceInstance
}

type JscService interface {
	Create(openID string, param *vo.JsCategoryReq) exception.Exception
	Get(id int64) (*vo.JsCategoryResp, exception.Exception)
	ListByPrimaryID(page *vo.PageInfo, pid int64) (*vo.DataPagination, exception.Exception)
	Update(openID string, id int64, param *vo.JsCategoryUpdateReq) exception.Exception
	Delete(id int64) exception.Exception
	MultiDelete(ids string) exception.Exception
}

func (jsi *jscServiceImpl) Create(openID string, param *vo.JsCategoryReq) exception.Exception {
	jscMgr := param.ToModel(openID)
	return jsi.repo.Create(jsi.db, jscMgr)
}

func (jsi *jscServiceImpl) Get(id int64) (*vo.JsCategoryResp, exception.Exception) {
	jsc, ex := jsi.repo.Get(jsi.db, id)
	if ex != nil {
		return nil, ex
	}
	domain, ex := jsi.domainRepo.Get(jsi.db, jsc.DomainID)
	if ex != nil {
		if ex.Type() == response.ExceptionRecordNotFound {
			domain = nil
		} else {
			return nil, ex
		}
	}
	jsp, ex := jsi.jspRepo.Get(jsi.db, jsc.PrimaryID)
	if ex != nil {
		return nil, ex
	}
	return vo.NewJsCategoryResponse(jsc, domain, jsp), nil
}

func (jsi *jscServiceImpl) ListByPrimaryID(pageInfo *vo.PageInfo, pid int64) (*vo.DataPagination, exception.Exception) {
	count, jscs, ex := jsi.repo.ListByPrimaryID(jsi.db, pageInfo, pid)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.JsCategoryResp, 0, len(jscs))
	for i := range jscs {
		domain, ex := jsi.domainRepo.Get(jsi.db, jscs[i].DomainID)
		if ex.Type() == response.ExceptionRecordNotFound {
			domain = nil
		} else if ex != nil {
			return nil, ex
		}
		jsp, ex := jsi.jspRepo.Get(jsi.db, jscs[i].PrimaryID)
		if ex != nil {
			return nil, ex
		}
		resp = append(resp, *vo.NewJsCategoryResponse(&jscs[i], domain, jsp))
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}

func (jsi *jscServiceImpl) Update(openID string, id int64, param *vo.JsCategoryUpdateReq) exception.Exception {
	return jsi.repo.Update(jsi.db, id, param.ToMap(openID))
}

func (jsi *jscServiceImpl) Delete(id int64) exception.Exception {
	return jsi.repo.Delete(jsi.db, id)
}

func (jsi *jscServiceImpl) MultiDelete(ids string) exception.Exception {
	idslice := strings.Split(ids, ",")
	if len(idslice) == 0 {
		return exception.New(response.ExceptionInvalidRequestParameters, "无效参数")
	}
	jid := make([]int64, 0, len(idslice))
	for i := range idslice {
		id, err := strconv.ParseUint(idslice[i], 10, 0)
		if err != nil {
			return exception.Wrap(response.ExceptionParseStringToUintError, err)
		}
		jid = append(jid, int64(id))
	}
	return jsi.repo.MultiDelete(jsi.db, jid)
}
