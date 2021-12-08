package service

import (
	"js_statistics/app/models"
	"js_statistics/app/repositories"
	"js_statistics/app/response"
	"js_statistics/app/vo"
	"js_statistics/commom/drivers/database"
	"js_statistics/exception"
	"sort"
	"sync"

	"gorm.io/gorm"
)

var (
	jspServiceInstance JspService
	jspOnce            sync.Once
)

type jspServiceImpl struct {
	db       *gorm.DB
	repo     repositories.JspRepo
	jscRepo  repositories.JscRepo
	jsmRepo  repositories.JsmRepo
	rmRepo   repositories.RmRepo
	stcRepo  repositories.StcRepo
	upRepo   repositories.UserPrimaryRepo
	userRepo repositories.UserRepo
}

func GetJspService() JspService {
	jspOnce.Do(func() {
		jspServiceInstance = &jspServiceImpl{
			db:       database.GetDriver(),
			repo:     repositories.GetJspRepo(),
			jscRepo:  repositories.GetJscRepo(),
			jsmRepo:  repositories.GetJsmRepo(),
			rmRepo:   repositories.GetRmRepo(),
			stcRepo:  repositories.GetStcRepo(),
			upRepo:   repositories.GetUserPrimaryRepo(),
			userRepo: repositories.GetUserRepo(),
		}
	})
	return jspServiceInstance
}

type JspService interface {
	Create(openID string, param *vo.JsPrimaryReq) exception.Exception
	Get(id int64) (*vo.JsPrimaryResp, exception.Exception)
	List(userID int64) ([]vo.JsPrimaryResp, exception.Exception)
	Update(openID string, id int64, param *vo.JsPrimaryUpdateReq) exception.Exception
	Delete(id int64) exception.Exception
	GetAllsCategoryTree() ([]vo.Primaries, exception.Exception)
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

func (jsi *jspServiceImpl) List(userID int64) ([]vo.JsPrimaryResp, exception.Exception) {
	user, ex := jsi.userRepo.Profile(jsi.db, userID)
	if ex != nil {
		return nil, ex
	}
	var jsps []models.JsPrimary
	if user.IsAdmin {
		jsps, ex = jsi.repo.List(jsi.db)
		if ex != nil {
			return nil, ex
		}
	} else {
		jsps, ex = jsi.repo.ListByUserID(jsi.db, userID)
		if ex != nil {
			return nil, ex
		}
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
		if ex := jsi.rmRepo.DeleteByCategoryIDs(tx, cids...); ex != nil {
			return ex
		}

		if ex := jsi.jscRepo.MultiDelete(tx, cids); ex != nil {
			return ex
		}
	}
	if ex := jsi.stcRepo.DeleteByPrimaryID(tx, id); ex != nil {
		return ex
	}

	if ex := jsi.upRepo.DeleteByPrimaryID(tx, id); ex != nil {
		return ex
	}

	if ex := jsi.repo.Delete(tx, id); ex != nil {
		return ex
	}
	if err := tx.Commit(); err.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, err.Error)
	}
	return nil
}

func (jsi *jspServiceImpl) GetAllsCategoryTree() ([]vo.Primaries, exception.Exception) {
	res, ex := jsi.repo.GetAllsCategoryTree(jsi.db)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.Primaries, 0)
	pcMap := make(map[vo.PrimaryKey][]vo.JsCategoryBrief)
	for i := range res {
		key := vo.PrimaryKey{ID: res[i].ID, Title: res[i].Title}

		if res[i].Pid == nil {
			noData := make([]vo.JsCategoryBrief, 0)
			pcMap[key] = noData
		} else {
			if res[i].ID == *res[i].Pid {
				jcb := vo.JsCategoryBrief{
					CID:   *res[i].Cid,
					Title: *res[i].CTitle,
				}
				pcMap[key] = append(pcMap[key], jcb)
			}
		}
	}
	sortKeyID := make([]int, 0)
	sortKeyTitle := make(map[int64]string)
	for k := range pcMap {
		sortKeyID = append(sortKeyID, int(k.ID))
		sortKeyTitle[k.ID] = k.Title
	}
	sort.Ints(sortKeyID)
	sortKey := make([]vo.PrimaryKey, 0, len(sortKeyID))

	for j := range sortKeyID {
		sortKey = append(sortKey, vo.PrimaryKey{
			ID:    int64(sortKeyID[j]),
			Title: sortKeyTitle[int64(sortKeyID[j])],
		})
	}

	for k := range sortKey {
		ps := vo.Primaries{
			ID:         sortKey[k].ID,
			Title:      sortKey[k].Title,
			Categories: pcMap[sortKey[k]],
		}
		resp = append(resp, ps)
	}
	return resp, nil
}
