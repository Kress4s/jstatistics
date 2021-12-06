package service

import (
	"js_statistics/app/repositories"
	"js_statistics/app/response"
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
	db      *gorm.DB
	repo    repositories.JspRepo
	jscRepo repositories.JscRepo
	jsmRepo repositories.JsmRepo
}

func GetJspService() JspService {
	jspOnce.Do(func() {
		jspServiceInstance = &jspServiceImpl{
			db:      database.GetDriver(),
			repo:    repositories.GetJspRepo(),
			jscRepo: repositories.GetJscRepo(),
			jsmRepo: repositories.GetJsmRepo(),
		}
	})
	return jspServiceInstance
}

type JspService interface {
	Create(openID string, param *vo.JsPrimaryReq) exception.Exception
	Get(id int64) (*vo.JsPrimaryResp, exception.Exception)
	List() ([]vo.JsPrimaryResp, exception.Exception)
	Update(openID string, id int64, param *vo.JsPrimaryUpdateReq) exception.Exception
	Delete(id int64) exception.Exception
}

func (jsi *jspServiceImpl) Create(openID string, param *vo.JsPrimaryReq) exception.Exception {
	jspMgr := param.ToModel(openID)
	return jsi.repo.Create(jsi.db, jspMgr)
}

func (jsi *jspServiceImpl) Get(id int64) (*vo.JsPrimaryResp, exception.Exception) {
	jspMgr, ex := jsi.repo.Get(jsi.db, id)
	if ex != nil {
		return nil, ex
	}
	return vo.NewJsPrimaryResponse(jspMgr), nil
}

func (jsi *jspServiceImpl) List() ([]vo.JsPrimaryResp, exception.Exception) {
	jsps, ex := jsi.repo.List(jsi.db)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.JsPrimaryResp, 0, len(jsps))
	for i := range jsps {
		resp = append(resp, *vo.NewJsPrimaryResponse(&jsps[i]))
	}
	return resp, nil
}

func (jsi *jspServiceImpl) Update(openID string, id int64, param *vo.JsPrimaryUpdateReq) exception.Exception {
	return jsi.repo.Update(jsi.db, id, param.ToMap(openID))
}

func (jsi *jspServiceImpl) Delete(id int64) exception.Exception {
	categories, ex := jsi.jscRepo.ListAllByPrimaryID(jsi.db, id)
	if ex != nil {
		return ex
	}
	cids := make([]int64, 0, len(categories))
	for i := range categories {
		cids = append(cids, categories[i].ID)
	}

	jms, ex := jsi.jsmRepo.GetIDsByCategoryID(jsi.db, cids)
	if ex != nil {
		return ex
	}

	tx := jsi.db.Begin()
	if tx.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	defer tx.Rollback()
	if len(jms) > 0 {
		if ex := jsi.jsmRepo.MultiDelete(tx, jms); ex != nil {
			return ex
		}
	}
	if len(cids) > 0 {
		if ex := jsi.jscRepo.MultiDelete(tx, cids); ex != nil {
			return ex
		}
	}
	if ex := jsi.repo.Delete(tx, id); ex != nil {
		return ex
	}
	if err := tx.Commit(); err.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, err.Error)
	}
	return nil
}
