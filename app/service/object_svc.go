package service

import (
	"fmt"
	"io"
	"js_statistics/app/models"
	"js_statistics/app/repositories"
	"js_statistics/app/response"
	"js_statistics/app/vo"
	"js_statistics/commom/drivers/database"
	"js_statistics/exception"
	"sync"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ObjectService interface {
	UploadFromReader(openID string, filename string, filesize int64, reader io.Reader) (string, exception.Exception)
	Download(id string) (*vo.ObjectResp, exception.Exception)
	Delete(id string) exception.Exception
}

type objectServiceImpl struct {
	db   *gorm.DB
	repo repositories.ObjectRepo
}

var (
	objectInstance ObjectService
	objectOnce     sync.Once
)

func GetObjectService() ObjectService {
	objectOnce.Do(func() {
		objectInstance = &objectServiceImpl{
			db:   database.GetDriver(),
			repo: repositories.GetObjectRepo(),
		}
	})
	return objectInstance
}

func (osi *objectServiceImpl) UploadFromReader(
	openID string, filename string, filesize int64, reader io.Reader,
) (string, exception.Exception) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return "", exception.Wrap(response.ExceptionGenerateID, err)
	}

	id := uid.String()

	now := time.Now().UTC()

	if ex := osi.repo.UploadFromReader(osi.db, &models.Object{
		ID:       id,
		Filename: filename,
		Path:     fmt.Sprintf("%s/%s", id, filename),
		Size:     filesize,
		Buff:     nil,
		Base: models.Base{
			CreateBy: openID,
			CreateAt: now,
			UpdateBy: openID,
			UpdateAt: now,
		},
	}, reader); ex != nil {
		return "", ex
	}
	return id, nil
}

func (osi *objectServiceImpl) Delete(id string) exception.Exception {
	return osi.repo.Delete(osi.db, id)
}

func (osi *objectServiceImpl) Download(id string) (*vo.ObjectResp, exception.Exception) {
	obj, ex := osi.repo.Download(osi.db, id)
	if ex != nil {
		return nil, ex
	}

	return &vo.ObjectResp{
		ID:       obj.ID,
		Filename: obj.Filename,
		Content:  string(obj.Buff),
	}, nil
}
