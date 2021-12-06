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
	domainServiceInstance DomainService
	domainOnce            sync.Once
)

type domainServiceImpl struct {
	db      *gorm.DB
	repo    repositories.DomainRepo
	jscRepo repositories.JscRepo
}

func GetDomainService() DomainService {
	domainOnce.Do(func() {
		domainServiceInstance = &domainServiceImpl{
			db:      database.GetDriver(),
			repo:    repositories.GetDomainRepo(),
			jscRepo: repositories.GetJscRepo(),
		}
	})
	return domainServiceInstance
}

type DomainService interface {
	Create(openID string, param *vo.DomainReq) exception.Exception
	Get(id int64) (*vo.DomainResp, exception.Exception)
	List(page *vo.PageInfo) (*vo.DataPagination, exception.Exception)
	Update(openID string, id int64, param *vo.DomainUpdateReq) exception.Exception
	Delete(id int64) exception.Exception
	MultiDelete(ids string) exception.Exception
	ListAll() ([]vo.DomainResp, exception.Exception)
}

func (dsi *domainServiceImpl) Create(openID string, param *vo.DomainReq) exception.Exception {
	domainMgr := param.ToModel(openID)
	return dsi.repo.Create(dsi.db, domainMgr)
}

func (dsi *domainServiceImpl) Get(id int64) (*vo.DomainResp, exception.Exception) {
	domainMgr, ex := dsi.repo.Get(dsi.db, id)
	if ex != nil {
		return nil, ex
	}
	return vo.NewDomainResponse(domainMgr), nil
}

func (dsi *domainServiceImpl) List(pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception) {
	count, domains, ex := dsi.repo.List(dsi.db, pageInfo)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.DomainResp, 0, len(domains))
	for i := range domains {
		resp = append(resp, *vo.NewDomainResponse(&domains[i]))
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}

func (dsi *domainServiceImpl) Update(openID string, id int64, param *vo.DomainUpdateReq) exception.Exception {
	return dsi.repo.Update(dsi.db, id, param.ToMap(openID))
}

func (dsi *domainServiceImpl) Delete(id int64) exception.Exception {
	tx := dsi.db.Begin()
	if tx.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	defer tx.Rollback()
	if ex := dsi.jscRepo.UpdateDomainIDToO(tx, id); ex != nil {
		return ex
	}
	if ex := dsi.repo.Delete(tx, id); ex != nil {
		return ex
	}
	if err := tx.Commit(); err.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, err.Error)
	}
	return nil
}

func (dsi *domainServiceImpl) MultiDelete(ids string) exception.Exception {
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
	tx := dsi.db.Begin()
	if tx.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	defer tx.Rollback()
	if ex := dsi.jscRepo.UpdateDomainIDToO(tx, did...); ex != nil {
		return ex
	}
	if ex := dsi.repo.MultiDelete(tx, did); ex != nil {
		return ex
	}
	if err := tx.Commit(); err.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, err.Error)
	}
	return nil
}

func (dsi *domainServiceImpl) ListAll() ([]vo.DomainResp, exception.Exception) {
	domains, ex := dsi.repo.ListAll(dsi.db)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.DomainResp, 0, len(domains))
	for i := range domains {
		resp = append(resp, vo.DomainResp{
			ID:    domains[i].ID,
			Title: domains[i].Title,
		})
	}
	return resp, nil
}
