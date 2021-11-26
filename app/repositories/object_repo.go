package repositories

import (
	"errors"
	"fmt"
	"io"
	"js_statistics/app/models/tables"
	"js_statistics/commom/drivers/minio"
	"js_statistics/exception"
	"js_statistics/minio_sdk"
	"sync"
	"time"

	"js_statistics/app/models"
	"js_statistics/app/response"

	"gorm.io/gorm"
)

type ObjectRepo interface {
	Upload(db *gorm.DB, o *models.Object) exception.Exception
	UploadFromReader(db *gorm.DB, o *models.Object, reader io.Reader) exception.Exception
	Download(db *gorm.DB, id string) (*models.Object, exception.Exception)
	Delete(db *gorm.DB, id string) exception.Exception
	Upsert(db *gorm.DB, id string, o *models.Object) error
	Import(db *gorm.DB, id string, o *models.Object) error
}

type objectRepositoryImpl struct {
	minio  minio_sdk.Client
	bucket string
}

var (
	objectRepoInstance ObjectRepo
	objectRepoOnce     sync.Once
)

func GetObjectRepo() ObjectRepo {
	objectRepoOnce.Do(func() {
		objectRepoInstance = &objectRepositoryImpl{
			bucket: "jstatistic",
			minio:  minio.GetDriver(),
		}
	})
	return objectRepoInstance
}

func (ori *objectRepositoryImpl) Upload(db *gorm.DB, o *models.Object) exception.Exception {
	if ex := exception.Wrap(
		response.ExceptionUploadObject,
		ori.minio.UploadObject(ori.bucket, o.Path, o.Buff)); ex != nil {
		return ex
	}
	return exception.Wrap(response.ExceptionDatabase, db.Create(o).Error)
}

func (ori *objectRepositoryImpl) UploadFromReader(db *gorm.DB, o *models.Object, reader io.Reader) exception.Exception {
	if ex := exception.Wrap(
		response.ExceptionUploadObject,
		ori.minio.UploadObjectFromReader(ori.bucket, o.Path, reader, o.Size)); ex != nil {
		return ex
	}
	return exception.Wrap(response.ExceptionDatabase, db.Create(o).Error)
}

func (ori *objectRepositoryImpl) Download(db *gorm.DB, id string) (*models.Object, exception.Exception) {
	var obj models.Object
	err := db.Where(models.Object{
		ID: id,
	}).First(&obj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, exception.Wrap(response.ExceptionRecordNotFound, err)
	}
	if err != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, err)
	}

	content, err := ori.minio.DownloadObject(ori.bucket, obj.Path)
	if err != nil {
		return nil, exception.Wrap(response.ExceptionDownloadObject, err)
	}

	obj.Buff = content

	return &obj, nil
}

func (ori *objectRepositoryImpl) Delete(db *gorm.DB, id string) exception.Exception {
	var obj models.Object
	db.Where(models.Object{
		ID: id,
	}).Find(&obj)
	if ex := exception.Wrap(response.ExceptionDeleteObject, ori.minio.DeleteObject(ori.bucket, obj.Path)); ex != nil {
		return ex
	}
	return exception.Wrap(response.ExceptionDatabase, db.Where(models.Object{
		ID: obj.ID,
	}).Delete(models.Object{}).Error)
}

func (ori *objectRepositoryImpl) Upsert(db *gorm.DB, id string, o *models.Object) error {
	now := time.Now().UTC()

	if ex := exception.Wrap(
		response.ExceptionUploadObject,
		ori.minio.UploadObject(ori.bucket, o.Path, o.Buff)); ex != nil {
		return ex
	}

	sqlStr := fmt.Sprintf(`INSERT INTO %s (id, filename, path, size, create_by, update_by, create_at,
update_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?) on conflict (id)
do update set filename = ?, path = ?, size = ?`,
		tables.Object)
	result := db.Exec(
		sqlStr,
		id,
		o.Filename,
		o.Path,
		len(o.Buff),
		o.CreateBy,
		o.UpdateBy,
		now,
		now,
		o.Filename,
		o.Path,
		len(o.Buff))
	return exception.Wrap(response.ExceptionDatabase, result.Error)
}

func (ori *objectRepositoryImpl) Import(db *gorm.DB, id string, o *models.Object) error {
	now := time.Now().UTC()
	if ex := exception.Wrap(
		response.ExceptionUploadObject,
		ori.minio.UploadObject(ori.bucket, o.Path, o.Buff)); ex != nil {
		return ex
	}
	sqlStr := fmt.Sprintf(`INSERT INTO %s (id, filename, path, size, create_by, update_by, create_at,
update_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`, tables.Object)
	result := db.Exec(
		sqlStr,
		id,
		o.Filename,
		o.Path,
		len(o.Buff),
		o.CreateBy,
		o.UpdateBy,
		now,
		now)
	return exception.Wrap(response.ExceptionDatabase, result.Error)
}
