package repositories

import (
	"js_statistics/app/models"
	"js_statistics/app/models/tables"
	"js_statistics/app/response"
	"js_statistics/app/vo"
	"js_statistics/constant"
	"js_statistics/exception"
	"sync"
	"time"

	"gorm.io/gorm"
)

var (
	jsmRepoInstance JsmRepo
	jsmOnce         sync.Once
)

type JsmRepoImpl struct{}

func GetJsmRepo() JsmRepo {
	jsmOnce.Do(func() {
		jsmRepoInstance = &JsmRepoImpl{}
	})
	return jsmRepoInstance
}

type JsmRepo interface {
	Create(db *gorm.DB, jsm *models.JsManage) exception.Exception
	ListByCategoryID(db *gorm.DB, pageInfo *vo.PageInfo, cid int64) (int64, []models.JsManageListView,
		exception.Exception)
	Get(db *gorm.DB, id int64) (*models.JsManage, exception.Exception)
	Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
	Delete(db *gorm.DB, id int64) exception.Exception
	DeleteByCategoryID(db *gorm.DB, cid int64) exception.Exception
	DeleteByCategoryIDs(db *gorm.DB, cids []int64) exception.Exception
	MultiDelete(db *gorm.DB, ids []int64) exception.Exception
	GetBySign(db *gorm.DB, sign string) (*models.JsManage, exception.Exception)
	DecreaseRedirectCount(db *gorm.DB, id int64) exception.Exception
	StatusChange(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
	GetIDsByCategoryID(db *gorm.DB, cids []int64) ([]int64, exception.Exception)
}

func (jsi *JsmRepoImpl) Create(db *gorm.DB, jsm *models.JsManage) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(jsm).Error)
}

func (jsi *JsmRepoImpl) ListByCategoryID(db *gorm.DB, pageInfo *vo.PageInfo, pid int64) (int64,
	[]models.JsManageListView, exception.Exception) {
	jsms := make([]models.JsManageListView, 0)
	countSubTx := db.Table(tables.IPStatistics).Select("distinct(ip) AS ip, js_id").
		Where("visit_time = ?", time.Now().Format(constant.DateFormat))
	countTX := db.Table("(?) AS jis", countSubTx).Select("count(*) as count, jis.js_id AS js_id").Group("jis.js_id")
	tx := db.Table(tables.JsManage+" AS js").
		Select("js.*, ips_count.count as ip_count").
		Joins("LEFT JOIN (?) AS ips_count ON ips_count.js_id = js.id", countTX)
	tx.Where("category_id = ?", pid).Order("id").Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).Scan(&jsms)
	count := int64(0)
	res := tx.Limit(-1).Offset(-1).Count(&count)
	return count, jsms, exception.Wrap(response.ExceptionDatabase, res.Error)
}

func (jsi *JsmRepoImpl) Get(db *gorm.DB, id int64) (*models.JsManage, exception.Exception) {
	jsMgr := models.JsManage{}
	res := db.Where(&models.JsManage{ID: id}).Find(&jsMgr)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &jsMgr, nil
}

func (jsi *JsmRepoImpl) Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.JsManage{}).Where(&models.JsManage{ID: id}).Updates(param).Error)
}

func (jsi *JsmRepoImpl) Delete(db *gorm.DB, id int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.JsManage{}, id).Error)
}

func (jsi *JsmRepoImpl) DeleteByCategoryID(db *gorm.DB, cid int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Where("category_id = ?", cid).Delete(&models.JsManage{}).Error)
}

func (jsi *JsmRepoImpl) DeleteByCategoryIDs(db *gorm.DB, cids []int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Where("category_id in (?)", cids).Delete(&models.JsManage{}).Error)
}

func (jsi *JsmRepoImpl) MultiDelete(db *gorm.DB, ids []int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.JsManage{}, ids).Error)
}

func (jsi *JsmRepoImpl) GetBySign(db *gorm.DB, sign string) (*models.JsManage, exception.Exception) {
	jsMgr := models.JsManage{}
	res := db.Where(&models.JsManage{Sign: sign}).Find(&jsMgr)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "recode not found")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &jsMgr, nil
}

func (jsi *JsmRepoImpl) DecreaseRedirectCount(db *gorm.DB, id int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Model(&models.JsManage{ID: id}).Updates(map[string]interface{}{
		"redirect_count": gorm.Expr("redirect_count - ?", 1),
	}).Error)
}

func (jsi *JsmRepoImpl) StatusChange(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.JsManage{}).Where(&models.JsManage{ID: id}).Updates(param).Error)
}

func (jsi *JsmRepoImpl) GetIDsByCategoryID(db *gorm.DB, cids []int64) ([]int64, exception.Exception) {
	mgrs := make([]models.JsManage, 0)
	res := db.Select("id").Where("category_id in (?)", cids).Find(&mgrs)
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	mgrIDs := make([]int64, 0, len(mgrs))
	for i := range mgrs {
		mgrIDs = append(mgrIDs, mgrs[i].ID)
	}
	return mgrIDs, nil
}
