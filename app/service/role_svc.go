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
	roleServiceInstance RoleService
	roleOnce            sync.Once
)

type roleServiceImpl struct {
	db     *gorm.DB
	repo   repositories.RoleRepo
	rpRepo repositories.RolePermissionRepo
	urRepo repositories.UserRoleRepo
}

func GetRoleService() RoleService {
	roleOnce.Do(func() {
		roleServiceInstance = &roleServiceImpl{
			db:     database.GetDriver(),
			repo:   repositories.GetRoleRepo(),
			rpRepo: repositories.GetRolePermissionRepo(),
			urRepo: repositories.GetUserRoleRepo(),
		}
	})
	return roleServiceInstance
}

type RoleService interface {
	Create(openID string, param *vo.RoleReq) exception.Exception
	Get(id int64) (*vo.RoleResp, exception.Exception)
	List(page *vo.PageInfo) (*vo.DataPagination, exception.Exception)
	Update(openID string, id int64, param *vo.RoleUpdateReq) exception.Exception
	Delete(id int64) exception.Exception
}

func (rsi *roleServiceImpl) Create(openID string, param *vo.RoleReq) exception.Exception {
	role := param.ToModel(openID)
	tx := rsi.db.Begin()
	defer tx.Rollback()
	if tx.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	if ex := rsi.repo.Creat(tx, role); ex != nil {
		return ex
	}
	rp := vo.RolePermissionReq{RoleID: role.ID}
	rpm, err := rp.ToModel(openID, param.Permissions)
	if err != nil {
		return exception.Wrap(response.ExceptionParseStringToInt64Error, err)
	}
	if ex := rsi.rpRepo.Create(tx, rpm); ex != nil {
		return ex
	}
	if res := tx.Commit(); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return nil
}

func (rsi *roleServiceImpl) Get(id int64) (*vo.RoleResp, exception.Exception) {
	role, ex := rsi.repo.Get(rsi.db, id)
	if ex != nil {
		return nil, ex
	}
	resp := &vo.RoleResp{
		ID:          role.ID,
		Name:        role.Name,
		Identify:    role.Identify,
		Description: role.Description,
	}
	psms, ex := rsi.rpRepo.GetByRoleID(rsi.db, role.ID)
	if ex != nil {
		return nil, ex
	}
	psids := make([]int64, 0, len(psms))
	for i := range psms {
		psids = append(psids, psms[i].PermissionID)
	}
	resp.Permissions = psids
	return resp, nil
}

func (rsi *roleServiceImpl) List(page *vo.PageInfo) (*vo.DataPagination, exception.Exception) {
	count, roles, ex := rsi.repo.List(rsi.db, page)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.RoleResp, 0, len(roles))
	for i := range roles {
		resp = append(resp, vo.RoleResp{
			ID:          roles[i].ID,
			Name:        roles[i].Name,
			Identify:    roles[i].Identify,
			Description: roles[i].Description,
		})
	}
	return vo.NewDataPagination(count, resp, page), nil
}

func (rsi *roleServiceImpl) Update(openID string, id int64, param *vo.RoleUpdateReq) exception.Exception {
	tx := rsi.db.Begin()
	defer tx.Rollback()
	if tx.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	ex := rsi.repo.Update(tx, id, param.ToMap(openID))
	if ex != nil {
		return ex
	}
	if ex = rsi.rpRepo.DeleteByRoleID(tx, id); ex != nil {
		return ex
	}
	rps := &vo.RolePermissionReq{RoleID: id}
	rpms, err := rps.ToModel(openID, param.Permissions)
	if err != nil {
		return exception.Wrap(response.ExceptionParseStringToInt64Error, err)
	}
	if ex = rsi.rpRepo.Create(tx, rpms); ex != nil {
		return ex
	}
	if res := tx.Commit(); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return nil
}

func (rsi *roleServiceImpl) Delete(id int64) exception.Exception {
	tx := rsi.db.Begin()
	defer tx.Rollback()
	if tx.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	ex := rsi.repo.Delete(tx, id)
	if ex != nil {
		return ex
	}
	if ex = rsi.rpRepo.DeleteByRoleID(tx, id); ex != nil {
		return ex
	}
	if ex = rsi.urRepo.DeleteByRoleID(tx, id); ex != nil {
		return ex
	}
	if res := tx.Commit(); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return nil
}
