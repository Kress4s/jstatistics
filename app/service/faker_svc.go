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
	fakerServiceInstance FakerService
	fakerOnce            sync.Once
)

type fakerServiceImpl struct {
	db      *gorm.DB
	repo    repositories.FakerRepo
	objRepo repositories.ObjectRepo
}

func GetFakerService() FakerService {
	fakerOnce.Do(func() {
		fakerServiceInstance = &fakerServiceImpl{
			db:      database.GetDriver(),
			repo:    repositories.GetFakerRepo(),
			objRepo: repositories.GetObjectRepo(),
		}
	})
	return fakerServiceInstance
}

type FakerService interface {
	Create(openID string, param *vo.FakerReq) exception.Exception
	Get(id int64) (*vo.FakerResp, exception.Exception)
	Update(openID string, id int64, param *vo.FakerUpdateReq) exception.Exception
	GetByJsID(jsID int64) (*vo.FakerResp, exception.Exception)
}

func (fsi *fakerServiceImpl) Create(openID string, param *vo.FakerReq) exception.Exception {
	faker := param.ToModel(openID)
	return fsi.repo.Create(fsi.db, faker)
}

func (fsi *fakerServiceImpl) Get(id int64) (*vo.FakerResp, exception.Exception) {
	faker, ex := fsi.repo.Get(fsi.db, id)
	if ex != nil {
		return nil, ex
	}
	return vo.NewFakerResponse(faker), nil
}

func (fsi *fakerServiceImpl) GetByJsID(jsID int64) (*vo.FakerResp, exception.Exception) {
	faker, ex := fsi.repo.GetByJsID(fsi.db, jsID)
	if ex != nil {
		if ex.Type() == response.ExceptionRecordNotFound {
			return nil, nil
		}
		return nil, ex
	}
	return vo.NewFakerResponse(faker), nil
}

func (fsi *fakerServiceImpl) Update(openID string, id int64, param *vo.FakerUpdateReq) exception.Exception {
	faker, ex := fsi.repo.Get(fsi.db, id)
	if ex != nil {
		return ex
	}
	if faker.Type != 0 && len(faker.ObjID) > 0 {
		// delete file
		if ex := fsi.objRepo.Delete(fsi.db, faker.ObjID); ex != nil {
			return ex
		}
	}
	return fsi.repo.Update(fsi.db, id, param.ToMap(openID))
}
