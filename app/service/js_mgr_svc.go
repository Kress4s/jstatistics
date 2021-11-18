package service

import (
	"fmt"
	"js_statistics/app/repositories"
	"js_statistics/app/response"
	"js_statistics/app/vo"
	"js_statistics/commom/drivers/database"
	"js_statistics/constant"
	"js_statistics/exception"
	"strconv"
	"strings"
	"sync"

	"gorm.io/gorm"
)

var (
	jsmServiceInstance JsmService
	jsmOnce            sync.Once
)

type jsmServiceImpl struct {
	db         *gorm.DB
	repo       repositories.JsmRepo
	jcRepo     repositories.JscRepo
	domainRepo repositories.DomainRepo
}

func GetJsmService() JsmService {
	jsmOnce.Do(func() {
		jsmServiceInstance = &jsmServiceImpl{
			db:         database.GetDriver(),
			repo:       repositories.GetJsmRepo(),
			jcRepo:     repositories.GetJscRepo(),
			domainRepo: repositories.GetDomainRepo(),
		}
	})
	return jsmServiceInstance
}

type JsmService interface {
	Create(openID string, param *vo.JsManageReq) exception.Exception
	Get(id int64) (*vo.JsManageResp, exception.Exception)
	ListByCategoryID(page *vo.PageInfo, pid int64) (*vo.DataPagination, exception.Exception)
	Update(openID string, id int64, param *vo.JsManageUpdateReq) exception.Exception
	Delete(id int64) exception.Exception
	MultiDelete(ids string) exception.Exception
	GetJSiteByID(id int64) (*vo.JSiteResp, exception.Exception)
}

func (jsi *jsmServiceImpl) Create(openID string, param *vo.JsManageReq) exception.Exception {
	jsmMgr := param.ToModel(openID)
	return jsi.repo.Create(jsi.db, jsmMgr)
}

func (jsi *jsmServiceImpl) Get(id int64) (*vo.JsManageResp, exception.Exception) {
	jsm, ex := jsi.repo.Get(jsi.db, id)
	if ex != nil {
		return nil, ex
	}
	return vo.NewJsManageResponse(jsm), nil
}

func (jsi *jsmServiceImpl) ListByCategoryID(pageInfo *vo.PageInfo, pid int64) (*vo.DataPagination,
	exception.Exception) {
	count, jsms, ex := jsi.repo.ListByCategoryID(jsi.db, pageInfo, pid)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.JsManageResp, 0, len(jsms))
	for i := range jsms {
		resp = append(resp, *vo.NewJsManageResponse(&jsms[i]))
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}

func (jsi *jsmServiceImpl) Update(openID string, id int64, param *vo.JsManageUpdateReq) exception.Exception {
	return jsi.repo.Update(jsi.db, id, param.ToMap(openID))
}

func (jsi *jsmServiceImpl) Delete(id int64) exception.Exception {
	return jsi.repo.Delete(jsi.db, id)
}

func (jsi *jsmServiceImpl) MultiDelete(ids string) exception.Exception {
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
	return jsi.repo.MultiDelete(jsi.db, jid)
}

func (jsi *jsmServiceImpl) GetJSiteByID(id int64) (*vo.JSiteResp, exception.Exception) {
	js, ex := jsi.repo.Get(jsi.db, id)
	if ex != nil {
		return nil, ex
	}
	jc, ex := jsi.jcRepo.Get(jsi.db, js.CategoryID)
	if ex != nil {
		return nil, ex
	}
	if jc.DomainID == 0 {
		return &vo.JSiteResp{Site: fmt.Sprintf(constant.JSiteForm, constant.DefaultJsDomain, js.Sign)}, nil
	}
	domain, ex := jsi.domainRepo.Get(jsi.db, jc.DomainID)
	if ex != nil {
		return nil, ex
	}
	return &vo.JSiteResp{Site: fmt.Sprintf(constant.JSiteForm, domain.Domain, js.Sign)}, nil
}
