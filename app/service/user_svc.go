package service

import (
	"js_statistics/app/models"
	repositories "js_statistics/app/repositories"
	"js_statistics/app/response"
	"js_statistics/app/vo"
	"js_statistics/commom/drivers/database"
	"js_statistics/commom/tools"
	"js_statistics/exception"
	"strconv"
	"strings"
	"sync"
	"time"

	"gorm.io/gorm"
)

var (
	userServiceInstance UserService
	userOnce            sync.Once
)

type userServiceImpl struct {
	db       *gorm.DB
	repo     repositories.UserRepo
	roleRepo repositories.RoleRepo
	urRepo   repositories.UserRoleRepo
	pmRepo   repositories.PermissionRepo
	ucRepo   repositories.UserCategoryRepo
	upRepo   repositories.UserPrimaryRepo
	jcRepo   repositories.JscRepo
}

type UserService interface {
	Profile(id int64) (*vo.ProfileResp, exception.Exception)
	Create(openID string, params *vo.UserReq) exception.Exception
	Get(id int64) (*vo.ProfileResp, exception.Exception)
	List(pageInfo *vo.PageInfo, id int64) (*vo.DataPagination, exception.Exception)
	Update(openID string, id int64, params *vo.UserUpdateReq) exception.Exception
	UpdateRoles(openID string, id int64, role *vo.UserUpdateRolesReq) exception.Exception
	UpdateJSCAndJS(openID string, id int64, param *vo.UserUpdateJscAndJsReq) exception.Exception
	Delete(openID string, id int64) exception.Exception
	GetRolesByUserID(openID string, uid int64) ([]vo.RoleBriefResp, exception.Exception)
	GetJscAndJsByUserID(openID string, uid int64) (*vo.JsJscAndJsBriefResp, exception.Exception)
	// MultiDelete(openID string, ids string) exception.Exception
	GetUserMenus(openID int64) ([]vo.UserToMenusResp, exception.Exception)
	StatusChange(openID string, id int64, status bool) exception.Exception
	MultiDelete(ids string) exception.Exception
}

func GetUserService() UserService {
	userOnce.Do(func() {
		userServiceInstance = &userServiceImpl{
			db:       database.GetDriver(),
			repo:     repositories.GetUserRepo(),
			urRepo:   repositories.GetUserRoleRepo(),
			roleRepo: repositories.GetRoleRepo(),
			pmRepo:   repositories.GetPermissionRepo(),
			ucRepo:   repositories.GetUserCategoryRepo(),
			upRepo:   repositories.GetUserPrimaryRepo(),
			jcRepo:   repositories.GetJscRepo(),
		}
	})
	return userServiceInstance
}

func (us *userServiceImpl) Profile(id int64) (*vo.ProfileResp, exception.Exception) {
	user, ex := us.repo.Profile(us.db, id)
	if ex != nil {
		return nil, ex
	}
	return &vo.ProfileResp{
		ID:    user.ID,
		Name:  user.Username,
		Admin: user.IsAdmin,
	}, nil
}

func (us *userServiceImpl) Create(openID string, params *vo.UserReq) exception.Exception {
	// password
	params.Password = string(tools.Base64Encode([]byte(params.Password)))
	user := params.ToModel(openID)
	return us.repo.Create(us.db, &user)
}

func (us *userServiceImpl) Get(id int64) (*vo.ProfileResp, exception.Exception) {
	user, ex := us.repo.Profile(us.db, id)
	if ex != nil {
		return nil, ex
	}
	return &vo.ProfileResp{
		ID:    user.ID,
		Name:  user.Username,
		Admin: user.IsAdmin,
	}, nil
}

func (us *userServiceImpl) List(pageInfo *vo.PageInfo, id int64) (*vo.DataPagination, exception.Exception) {
	count, users, ex := us.repo.List(us.db, pageInfo, id)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.UserResp, 0, len(users))
	for i := range users {
		resp = append(resp, vo.UserResp{
			ID:       users[i].ID,
			UserName: users[i].Username,
			Admin:    users[i].IsAdmin,
			Status:   users[i].Status,
		})
	}
	return vo.NewDataPagination(count, resp, pageInfo), nil
}

func (us *userServiceImpl) Update(openID string, id int64, params *vo.UserUpdateReq) exception.Exception {
	r := make(map[string]interface{})
	// password is nil, declear not change
	if len(params.Password) != 0 {
		r["password"] = string(tools.Base64Encode([]byte(params.Password)))
	}
	r["user_name"] = params.UserName
	r["is_admin"] = params.IsAdmin
	r["status"] = params.Status
	r["update_by"] = openID
	r["update_at"] = time.Now()
	return us.repo.Update(us.db, id, r)
}

func (us *userServiceImpl) Delete(openID string, id int64) exception.Exception {
	tx := us.db.Begin()
	defer tx.Rollback()
	if tx.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	ex := us.repo.Delete(tx, id)
	if ex != nil {
		return ex
	}
	if ex = us.urRepo.DeleteByUserID(tx, id); ex != nil {
		return ex
	}
	if ex = us.ucRepo.DeleteByUserID(tx, id); ex != nil {
		return ex
	}
	if ex = us.upRepo.DeleteByUserID(tx, id); ex != nil {
		return ex
	}
	if res := tx.Commit(); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return nil
}

func (us *userServiceImpl) UpdateRoles(openID string, id int64, param *vo.UserUpdateRolesReq) exception.Exception {
	tx := us.db.Begin()
	defer tx.Rollback()
	if tx.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	ex := us.urRepo.DeleteByUserID(tx, id)
	if ex != nil {
		return ex
	}
	urs := &vo.UserUpdateRolesReq{RoleIDs: param.RoleIDs}
	urms := urs.ToModel(openID, id)
	if ex = us.urRepo.Create(tx, urms); ex != nil {
		return ex
	}
	if res := tx.Commit(); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return nil
}

func (us *userServiceImpl) UpdateJSCAndJS(openID string, id int64, param *vo.UserUpdateJscAndJsReq) exception.Exception {
	tx := us.db.Begin()
	if tx.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	defer tx.Rollback()
	if ex := us.ucRepo.DeleteByUserID(tx, id); ex != nil {
		return ex
	}
	if ex := us.upRepo.DeleteByUserID(tx, id); ex != nil {
		return ex
	}

	urs := &vo.UserUpdateJscAndJsReq{CategoryIDs: param.CategoryIDs, PrimaryIDs: param.PrimaryIDs}
	ucrs, uprs := urs.ToModel(openID, id)
	if len(ucrs) > 0 {
		if ex := us.ucRepo.Create(tx, ucrs); ex != nil {
			return ex
		}
	}
	if len(uprs) > 0 {
		if ex := us.upRepo.Create(tx, uprs); ex != nil {
			return ex
		}
	}
	if res := tx.Commit(); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return nil
}

func (us *userServiceImpl) GetRolesByUserID(openID string, uid int64) ([]vo.RoleBriefResp, exception.Exception) {
	urs, ex := us.urRepo.GetByUserID(us.db, uid)
	if ex != nil {
		return nil, ex
	}
	if len(urs) == 0 {
		return []vo.RoleBriefResp{}, nil
	}
	rolesID := make([]int64, 0, len(urs))
	for i := range urs {
		rolesID = append(rolesID, urs[i].RoleID)
	}
	roles, ex := us.roleRepo.GetByIDs(us.db, rolesID)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.RoleBriefResp, 0, len(roles))
	for i := range roles {
		resp = append(resp, vo.RoleBriefResp{
			ID:   roles[i].ID,
			Name: roles[i].Name,
		})
	}
	return resp, nil
}

func (us *userServiceImpl) GetJscAndJsByUserID(openID string, uid int64) (*vo.JsJscAndJsBriefResp,
	exception.Exception) {
	ucs, ex := us.ucRepo.GetByUserID(us.db, uid)
	if ex != nil {
		return nil, ex
	}
	categoriesID := make([]int64, 0, len(ucs))
	for i := range ucs {
		categoriesID = append(categoriesID, ucs[i].CategoryID)
	}
	ups, ex := us.upRepo.GetByUserID(us.db, uid)
	if ex != nil {
		return nil, ex
	}
	primariesID := make([]int64, 0, len(ups))
	for i := range ups {
		primariesID = append(primariesID, ups[i].PrimaryID)
	}
	return &vo.JsJscAndJsBriefResp{PrimariesID: primariesID, CategoriesID: categoriesID}, nil
}

func (us *userServiceImpl) GetUserMenus(userID int64) ([]vo.UserToMenusResp, exception.Exception) {
	user, ex := us.repo.Get(us.db, userID)
	if ex != nil {
		return nil, ex
	}
	res := make([]models.UserToMenus, 0)
	// 判断是否是超管
	if user.IsAdmin {
		ps, ex := us.pmRepo.GetAll(us.db)
		if ex != nil {
			return nil, ex
		}
		for i := range ps {
			res = append(res, models.UserToMenus{
				MenuID:   ps[i].ID,
				MenuName: ps[i].MenuName,
				Route:    ps[i].Route,
				Identify: ps[i].Identify,
			})
		}
	} else {
		// 非超管
		res, ex = us.repo.GetUserMenus(us.db, userID)
		if ex != nil {
			return nil, ex
		}
	}
	for i := range res {
		if res[i].MenuID == 1 {
			// 顶级权限
			ps, ex := us.pmRepo.GetAll(us.db)
			if ex != nil {
				return nil, ex
			}
			res = make([]models.UserToMenus, 0, len(ps))
			for i := range ps {
				res = append(res, models.UserToMenus{
					MenuID:   ps[i].ID,
					MenuName: ps[i].MenuName,
					Route:    ps[i].Route,
					Identify: ps[i].Identify,
				})
			}
			break
		}
	}
	menus := make([]vo.UserToMenusResp, 0, len(res))
	for i := range res {
		menus = append(menus,
			vo.UserToMenusResp{
				MenuID:   res[i].MenuID,
				MenuName: res[i].MenuName,
				Route:    res[i].Route,
				Identify: res[i].Identify,
			},
		)
	}
	return menus, nil
}

func (us *userServiceImpl) StatusChange(openID string, id int64, status bool) exception.Exception {
	return us.repo.StatusChange(us.db, id, map[string]interface{}{
		"status":    status,
		"update_by": openID,
		"update_at": time.Now(),
	})
}

func (us *userServiceImpl) MultiDelete(ids string) exception.Exception {
	idslice := strings.Split(ids, ",")
	if len(idslice) == 0 {
		return exception.New(response.ExceptionInvalidRequestParameters, "无效参数")
	}
	did := make([]int64, 0, len(idslice))
	for i := range idslice {
		id, err := strconv.ParseInt(idslice[i], 0, 64)
		if err != nil {
			return exception.Wrap(response.ExceptionParseStringToInt64Error, err)
		}
		did = append(did, id)
	}
	tx := us.db.Begin()
	if tx.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	defer tx.Rollback()
	if ex := us.urRepo.DeleteByUsersID(tx, did...); ex != nil {
		return ex
	}
	if ex := us.repo.MultiDelete(tx, did); ex != nil {
		return ex
	}
	if ex := us.ucRepo.DeleteByUsersID(tx, did...); ex != nil {
		return ex
	}
	if ex := us.upRepo.DeleteByUsersID(tx, did...); ex != nil {
		return ex
	}
	if err := tx.Commit(); err.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, err.Error)
	}
	return nil
}
