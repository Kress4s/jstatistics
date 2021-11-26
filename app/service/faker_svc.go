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
	fakerServiceInstance FakerService
	fakerOnce            sync.Once
)

type fakerServiceImpl struct {
	db   *gorm.DB
	repo repositories.FakerRepo
}

func GetFakerService() FakerService {
	fakerOnce.Do(func() {
		fakerServiceInstance = &fakerServiceImpl{
			db:   database.GetDriver(),
			repo: repositories.GetFakerRepo(),
		}
	})
	return fakerServiceInstance
}

type FakerService interface {
	Create(openID string, param *vo.FakerReq) exception.Exception
	Get(id int64) (*vo.FakerResp, exception.Exception)
	Update(openID string, id int64, param *vo.FakerUpdateReq) exception.Exception
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

func (fsi *fakerServiceImpl) Update(openID string, id int64, param *vo.FakerUpdateReq) exception.Exception {
	return fsi.repo.Update(fsi.db, id, param.ToMap(openID))
}
