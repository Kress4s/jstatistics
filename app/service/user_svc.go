package service

import (
	repositories "js_statistics/app/repositories"
	"js_statistics/app/response"
	"js_statistics/app/vo"
	"js_statistics/commom/drivers/database"
	"js_statistics/commom/tools"
	"js_statistics/exception"
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
}

type UserService interface {
	Profile(id uint) (*vo.ProfileResp, exception.Exception)
	Create(openID string, params *vo.UserReq) exception.Exception
	Get(id uint) (*vo.ProfileResp, exception.Exception)
	List(pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception)
	Update(openID string, id uint, params *vo.UserUpdateReq) exception.Exception
	UpdateRoles(openID string, id uint, role *vo.UserUpdateRolesReq) exception.Exception
	Delete(openID string, id uint) exception.Exception
	GetRolesByUserID(openID string, uid uint) ([]vo.RoleBriefResp, exception.Exception)
	// MultiDelete(openID string, ids string) exception.Exception
}

func GetUserService() UserService {
	userOnce.Do(func() {
		userServiceInstance = &userServiceImpl{
			db:       database.GetDriver(),
			repo:     repositories.GetUserRepo(),
			urRepo:   repositories.GetUserRoleRepo(),
			roleRepo: repositories.GetRoleRepo(),
		}
	})
	return userServiceInstance
}

func (us *userServiceImpl) Profile(id uint) (*vo.ProfileResp, exception.Exception) {
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

func (us *userServiceImpl) Get(id uint) (*vo.ProfileResp, exception.Exception) {
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

func (us *userServiceImpl) List(pageInfo *vo.PageInfo) (*vo.DataPagination, exception.Exception) {
	count, users, ex := us.repo.List(us.db, pageInfo)
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

func (us *userServiceImpl) Update(openID string, id uint, params *vo.UserUpdateReq) exception.Exception {
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

func (us *userServiceImpl) Delete(openID string, id uint) exception.Exception {
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
	if res := tx.Commit(); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return nil
}

func (us *userServiceImpl) UpdateRoles(openID string, id uint, param *vo.UserUpdateRolesReq) exception.Exception {
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

func (us *userServiceImpl) GetRolesByUserID(openID string, uid uint) ([]vo.RoleBriefResp, exception.Exception) {
	urs, ex := us.urRepo.GetByUserID(us.db, uid)
	if ex != nil {
		return nil, ex
	}
	if len(urs) == 0 {
		return []vo.RoleBriefResp{}, nil
	}
	rolesID := make([]uint, 0, len(urs))
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
