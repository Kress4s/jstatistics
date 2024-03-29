package repositories

import (
	"fmt"
	"js_statistics/app/models"
	"js_statistics/app/models/tables"
	"js_statistics/app/response"
	"js_statistics/app/vo"
	"js_statistics/exception"
	"sync"

	"gorm.io/gorm"
)

var (
	userRepoInstance UserRepo
	userOnce         sync.Once
)

type UserRepoImpl struct{}

func GetUserRepo() UserRepo {
	userOnce.Do(func() {
		userRepoInstance = &UserRepoImpl{}
	})
	return userRepoInstance
}

type UserRepo interface {
	Profile(db *gorm.DB, id int64) (*models.User, exception.Exception)
	CheckPassword(db *gorm.DB, account, password string) (bool, bool, int64, exception.Exception)
	Create(db *gorm.DB, user *models.User) exception.Exception
	Get(db *gorm.DB, id int64) (*models.User, exception.Exception)
	List(db *gorm.DB, pageInfo *vo.PageInfo, id int64) (int64, []models.UserListInfo, exception.Exception)
	Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
	Delete(db *gorm.DB, id int64) exception.Exception
	MultiDelete(db *gorm.DB, ids []int64) exception.Exception
	GetUserMenus(db *gorm.DB, userID int64) ([]models.UserToMenus, exception.Exception)
	StatusChange(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception
}

func (u *UserRepoImpl) Profile(db *gorm.DB, id int64) (*models.User, exception.Exception) {
	user := models.User{}
	res := db.Where(&models.User{ID: id}).Find(&user)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "用户未找到")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &user, nil
}

func (u *UserRepoImpl) CheckPassword(db *gorm.DB, username, password string) (bool, bool, int64, exception.Exception) {
	user := &models.User{}
	res := db.Where(&models.User{Username: username, Password: password}).Find(user)
	if res.Error != nil {
		return false, false, 0, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	if res.RowsAffected == 0 {
		return false, false, 0, exception.New(response.ExceptionInvalidUserPassword, "用户名/密码错误")
	}
	return true, user.Status, user.ID, nil
}

func (u *UserRepoImpl) Create(db *gorm.DB, user *models.User) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Create(user).Error)
}

func (u *UserRepoImpl) Get(db *gorm.DB, id int64) (*models.User, exception.Exception) {
	user := models.User{}
	res := db.Where(&models.User{ID: id}).Find(&user)
	if res.RowsAffected == 0 {
		return nil, exception.New(response.ExceptionRecordNotFound, "用户未找到")
	}
	if res.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, res.Error)
	}
	return &user, nil
}

func (u *UserRepoImpl) List(db *gorm.DB, pageInfo *vo.PageInfo, id int64) (int64, []models.UserListInfo, exception.Exception) {
	users := make([]models.UserListInfo, 0)
	tx := db.Table(tables.User + " AS u").
		Select(`u.id AS id, u.user_name AS user_name, u.is_admin AS is_admin, u.status AS status, string_agg(r.name, ',') AS role_names`).
		Joins(fmt.Sprintf("LEFT JOIN %s as ur ON ur.user_id = u.id", tables.UserRoleRelation)).
		Joins(fmt.Sprintf("LEFT JOIN %s as r ON ur.role_id = r.id", tables.Role)).
		Group("u.id, u.user_name, u.is_admin, u.status")
	if pageInfo.Keywords != "" {
		tx = tx.Scopes(vo.FuzzySearch(pageInfo.Keywords, "user_name"))
	}
	if id > 0 {
		tx.Where("id = ?", id)
	}
	tx = tx.Order("u.id").Limit(pageInfo.PageSize).Offset(pageInfo.Offset()).Find(&users)
	count := int64(0)
	res := tx.Limit(-1).Offset(-1).Count(&count)
	return count, users, exception.Wrap(response.ExceptionDatabase, res.Error)
}

func (u *UserRepoImpl) Update(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.User{}).Where(&models.User{ID: id}).Updates(param).Error)
}

func (u *UserRepoImpl) Delete(db *gorm.DB, id int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.User{}, id).Error)
}

func (u *UserRepoImpl) MultiDelete(db *gorm.DB, ids []int64) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase, db.Delete(&models.User{}, ids).Error)
}

func (u *UserRepoImpl) GetUserMenus(db *gorm.DB, userID int64) ([]models.UserToMenus, exception.Exception) {
	menus := make([]models.UserToMenus, 0)
	tx := db.Table(tables.User+" as u").
		Select(
			"u.id as uid, p.id as menu_id, p.menu_name as menu_name, p.route as route, p.identify as identify, p.type as type, p.parent_id as parent_id",
		).
		Joins(fmt.Sprintf("INNER JOIN %s AS ur_rel ON ur_rel.user_id = u.id", tables.UserRoleRelation)).
		Joins(fmt.Sprintf("INNER JOIN %s as rp_rel on rp_rel.role_id = ur_rel.role_id", tables.RolePermissionRelation)).
		Joins(fmt.Sprintf("INNER JOIN %s AS p ON p.id = rp_rel.permission_id", tables.Permission)).
		Group("u.id,p.id,p.menu_name,p.route").
		Having("u.id = ?", userID).
		Scan(&menus)
	if tx.Error != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return menus, nil
}

func (u *UserRepoImpl) StatusChange(db *gorm.DB, id int64, param map[string]interface{}) exception.Exception {
	return exception.Wrap(response.ExceptionDatabase,
		db.Model(&models.User{}).Where(&models.User{ID: id}).Updates(param).Error)
}
