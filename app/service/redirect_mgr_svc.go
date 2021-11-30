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
	"time"

	"gorm.io/gorm"
)

var (
	rmServiceInstance RmService
	rmOnce            sync.Once
)

type rmServiceImpl struct {
	db     *gorm.DB
	repo   repositories.RmRepo
	rlRepo repositories.RlRepo
}

func GetRmService() RmService {
	rmOnce.Do(func() {
		rmServiceInstance = &rmServiceImpl{
			db:     database.GetDriver(),
			repo:   repositories.GetRmRepo(),
			rlRepo: repositories.GetRlRepo(),
		}
	})
	return rmServiceInstance
}

type RmService interface {
	Create(openID string, param *vo.RedirectManageReq) exception.Exception
	Get(id int64) (*vo.RedirectManageResp, exception.Exception)
	ListByCategoryID(page *vo.PageInfo, cid int64) (*vo.DataPagination, exception.Exception)
	Update(openID string, id int64, param *vo.RedirectManageUpdateReq) exception.Exception
	Delete(id int64) exception.Exception
	MultiDelete(ids string) exception.Exception
}

func (rsi *rmServiceImpl) Create(openID string, param *vo.RedirectManageReq) exception.Exception {
	rmMgr := param.ToModel(openID)
	tx := rsi.db.Begin()
	defer tx.Callback()
	if tx.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	ex := rsi.repo.Create(tx, rmMgr)
	if ex != nil {
		return ex
	}
	// 添加跳转管理的日志
	redirectLog := vo.RedirectLog(rmMgr)
	if ex := rsi.rlRepo.Create(tx, redirectLog); ex != nil {
		return ex
	}
	if res := tx.Commit(); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return nil
}

func (rsi *rmServiceImpl) Get(id int64) (*vo.RedirectManageResp, exception.Exception) {
	rm, ex := rsi.repo.Get(rsi.db, id)
	if ex != nil {
		return nil, ex
	}
	return vo.NewRedirectManageResponse(rm), nil
}

func (rsi *rmServiceImpl) ListByCategoryID(pageInfo *vo.PageInfo, cid int64) (*vo.DataPagination, exception.Exception) {
	count, rms, ex := rsi.repo.ListByCategoryID(rsi.db, pageInfo, cid)
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
	tx := rsi.db.Begin()
	defer tx.Callback()
	if tx.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	if ex := rsi.repo.Update(tx, id, param.ToMap(openID)); ex != nil {
		return ex
	}
	redirect, ex := rsi.repo.Get(tx, id)
	if ex != nil {
		return ex
	}
	if ex := rsi.rlRepo.Update(tx, id, param.CategoryID, map[string]interface{}{
		"pc":          param.PC,
		"android":     param.Android,
		"ios":         param.IOS,
		"old_pc":      redirect.PC,
		"old_android": redirect.Android,
		"old_ios":     redirect.IOS,
		"type":        "修改",
		"update_at":   time.Now(),
	}); ex != nil {
		return ex
	}
	if res := tx.Commit(); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return nil
}

func (rsi *rmServiceImpl) Delete(id int64) exception.Exception {
	tx := rsi.db.Begin()
	defer tx.Callback()
	if tx.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	redirect, ex := rsi.repo.Get(tx, id)
	if ex != nil {
		return ex
	}
	if ex := rsi.repo.Delete(tx, id); ex != nil {
		return ex
	}
	if ex := rsi.rlRepo.Delete(tx, id, redirect.CategoryID); ex != nil {
		return nil
	}
	if res := tx.Commit(); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return nil
}

func (rsi *rmServiceImpl) MultiDelete(ids string) exception.Exception {
	idslice := strings.Split(ids, ",")
	if len(idslice) == 0 {
		return exception.New(response.ExceptionInvalidRequestParameters, "无效参数")
	}
	rids := make([]int64, 0, len(idslice))
	for i := range idslice {
		id, err := strconv.ParseUint(idslice[i], 10, 0)
		if err != nil {
			return exception.Wrap(response.ExceptionParseStringToInt64Error, err)
		}
		rids = append(rids, int64(id))
	}
	tx := rsi.db.Begin()
	defer tx.Callback()
	if tx.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	var categoryID int64
	if len(rids) > 0 {
		rdm, ex := rsi.repo.Get(tx, rids[0])
		if ex != nil {
			return ex
		}
		categoryID = rdm.CategoryID
	}
	if ex := rsi.repo.MultiDelete(rsi.db, rids); ex != nil {
		return ex
	}
	if ex := rsi.rlRepo.MultiDelete(tx, rids, categoryID); ex != nil {
		return ex
	}
	if res := tx.Commit(); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return nil
}
