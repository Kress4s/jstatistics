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
	jspServiceInstance JspService
	jspOnce            sync.Once
)

type jspServiceImpl struct {
	db   *gorm.DB
	repo repositories.JspRepo
}

func GetJspService() JspService {
	jspOnce.Do(func() {
		jspServiceInstance = &jspServiceImpl{
			db:   database.GetDriver(),
			repo: repositories.GetJspRepo(),
		}
	})
	return jspServiceInstance
}

type JspService interface {
	Create(openID string, param *vo.JsPrimaryReq) exception.Exception
	Get(id uint) (*vo.JsPrimaryResp, exception.Exception)
	List(page *vo.PageInfo) (*vo.DataPagination, exception.Exception)
	Update(openID string, id uint, param *vo.JsPrimaryUpdateReq) exception.Exception
	Delete(id uint) exception.Exception
}

func (jsi *jspServiceImpl) Create(openID string, param *vo.JsPrimaryReq) exception.Exception {
	jspMgr := param.ToModel(openID)
	return jsi.repo.Create(jsi.db, jspMgr)
}

func (jsi *jspServiceImpl) Get(id uint) (*vo.JsPrimaryResp, exception.Exception) {
	jspMgr, ex := jsi.repo.Get(jsi.db, id)
	if ex != nil {
		return nil, ex
	}
	return vo.NewJsPrimaryResponse(jspMgr), nil
}

func (jsi *jspServiceImpl) List(pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception) {
	count, jsps, ex := jsi.repo.List(jsi.db, pageInfo)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.JsPrimaryResp, 0, len(jsps))
	for i := range jsps {
		resp = append(resp, *vo.NewJsPrimaryResponse(&jsps[i]))
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}

func (jsi *jspServiceImpl) Update(openID string, id uint, param *vo.JsPrimaryUpdateReq) exception.Exception {
	return jsi.repo.Update(jsi.db, id, param.ToMap(openID))
}

func (jsi *jspServiceImpl) Delete(id uint) exception.Exception {
	return jsi.repo.Delete(jsi.db, id)
}
