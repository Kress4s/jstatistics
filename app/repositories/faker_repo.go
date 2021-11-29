package repositories

import (
	"js_statistics/app/models"
	"js_statistics/app/response"
	"js_statistics/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	fakerRepoInstance FakerRepo
	fakerOnce         sync.Once
)

type fakerRepoImpl struct{}

func GetFakerRepo() FakerRepo {
	fakerOnce.Do(func() {
		fakerRepoInstance = &fakerRepoImpl{}
	})
	return fakerRepoInstance
}

type FakerRepo interface {
	Create(db *gorm.DB, ip *models.Faker) exception.Exception
	Get(db *gorm.DB, id int64) (*models.Faker, exception.Exception)
	GetByJsID(db *gorm.DB, JsID int64) (*models.Faker, exception.Exception)
	Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
}

func (fri *fakerRepoImpl) Create(db *gorm.DB, faker *models.Faker) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(faker).Error)
}

func (fri *fakerRepoImpl) Get(db *gorm.DB, id int64) (*models.Faker, exception.Exception) {
	faker := models.Faker{}
	res := db.Where(&models.Faker{ID: id}).Find(&faker)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &faker, nil
}

func (fri *fakerRepoImpl) GetByJsID(db *gorm.DB, JsID int64) (*models.Faker, exception.Exception) {
	faker := models.Faker{}
	res := db.Where(&models.Faker{JsID: JsID}).Find(&faker)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &faker, nil
}

func (fri *fakerRepoImpl) Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.Faker{}).Where(&models.Faker{ID: id}).Updates(param).Error)
}
