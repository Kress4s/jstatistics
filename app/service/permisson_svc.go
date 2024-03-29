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
	permissionServiceInstance PermissionService
	permissionOnce            sync.Once
)

type permissionServiceImpl struct {
	db     *gorm.DB
	repo   repositories.PermissionRepo
	rpRepo repositories.RolePermissionRepo
}

type PermissionService interface {
	Create(openID string, params *vo.PermissionReq) exception.Exception
	Get(id int64) (*vo.PermissionResp, exception.Exception)
	GetAll() ([]vo.PermissionResp, exception.Exception)
	Update(openID string, id int64, params *vo.PermissionUpdateReq) exception.Exception
	GetPermissionTree() (*vo.PermissionTree, exception.Exception)
	Delete(openID string, id int64) exception.Exception
}

func GetPermissionService() PermissionService {
	permissionOnce.Do(func() {
		permissionServiceInstance = &permissionServiceImpl{
			db:     database.GetDriver(),
			repo:   repositories.GetPermissionRepo(),
			rpRepo: repositories.GetRolePermissionRepo(),
		}
	})
	return permissionServiceInstance
}

func (ps *permissionServiceImpl) Create(openID string, params *vo.PermissionReq) exception.Exception {
	p := params.ToModel(openID)
	return ps.repo.Create(ps.db, &p)
}

func (ps *permissionServiceImpl) Get(id int64) (*vo.PermissionResp, exception.Exception) {
	p, ex := ps.repo.Get(ps.db, id)
	if ex != nil {
		return nil, ex
	}
	return &vo.PermissionResp{
		ID:       p.ID,
		Name:     p.Name,
		MenuName: p.MenuName,
		Route:    p.Route,
		Identify: p.Identify,
		Type:     p.Type,
		Index:    p.Index,
		ParentID: p.ParentID,
	}, nil
}

func (ps *permissionServiceImpl) GetAll() ([]vo.PermissionResp, exception.Exception) {
	p, ex := ps.repo.GetAll(ps.db)
	if ex != nil {
		return nil, ex
	}
	resp := make([]vo.PermissionResp, 0, len(p))
	for i := range p {
		resp = append(resp, vo.PermissionResp{
			ID:       p[i].ID,
			Name:     p[i].Name,
			MenuName: p[i].MenuName,
			Identify: p[i].Identify,
			Route:    p[i].Route,
			ParentID: p[i].ParentID,
			Type:     p[i].Type,
		})
	}
	return resp, nil
}

func (ps *permissionServiceImpl) GetPermissionTree() (*vo.PermissionTree, exception.Exception) {
	top, ex := ps.repo.GetTop(ps.db)
	if ex != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, ex)
	}
	permissions, ex := ps.repo.GetAll(ps.db)
	if ex != nil {
		return nil, exception.Wrap(response.ExceptionDatabase, ex)
	}
	topsPermission := vo.NewPermissionTree(top)
	allPermission := make([]*vo.PermissionTree, 0, len(permissions))
	for i := range permissions {
		if permissions[i].ParentID == 0 {
			continue
		}
		allPermission = append(allPermission, vo.NewPermissionTree(&permissions[i]))
	}
	ps.makePermissionTree(allPermission, topsPermission, nil)
	return topsPermission, nil
}

func (ps *permissionServiceImpl) makePermissionTree(allPermission []*vo.PermissionTree,
	topermission *vo.PermissionTree, ids *[]int64) {
	children, _ := ps.haveChild(allPermission, topermission)
	if len(children) != 0 {
		if ids != nil {
			for i := range children {
				*ids = append(*ids, children[i].ID)
			}
		}
		topermission.SubPermissions = append(topermission.SubPermissions, children...)
		for _, v := range children {
			_, yes := ps.haveChild(allPermission, v)
			if yes {
				ps.makePermissionTree(allPermission, v, ids)
			}
		}
	}
}

func (ps *permissionServiceImpl) haveChild(allPermissions []*vo.PermissionTree, topPermission *vo.PermissionTree,
) (children []*vo.PermissionTree, yes bool) {
	for _, v := range allPermissions {
		if v.ParentID == topPermission.ID {
			children = append(children, v)
		}
	}
	if children != nil {
		yes = true
	}
	return
}

func (us *permissionServiceImpl) Update(openID string, id int64, params *vo.PermissionUpdateReq) exception.Exception {
	return us.repo.Update(us.db, id, params.ToMap(openID))
}

func (ps *permissionServiceImpl) Delete(openID string, id int64) exception.Exception {
	/*
		1. 查出节点
		2. 查出该节点下所有子节点的id
		3. id汇总下
	*/
	p, ex := ps.repo.Get(ps.db, id)
	if ex != nil {
		return ex
	}
	dp := vo.NewPermissionTree(p)
	permissions, ex := ps.repo.GetAll(ps.db)
	if ex != nil {
		return ex
	}
	allPermission := make([]*vo.PermissionTree, 0, len(permissions))
	ids := new([]int64)
	*ids = append(*ids, id)
	for i := range permissions {
		if permissions[i].ParentID == 0 || permissions[i].ID == dp.ID {
			continue
		}
		allPermission = append(allPermission, vo.NewPermissionTree(&permissions[i]))
	}
	ps.makePermissionTree(allPermission, dp, ids)
	tx := ps.db.Begin()
	defer tx.Rollback()
	if tx.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	if ex = ps.repo.Delete(tx, *ids); ex != nil {
		return ex
	}
	if ex = ps.rpRepo.DeleteByPermissionID(tx, id); ex != nil {
		return ex
	}
	if res := tx.Commit(); res.Error != nil {
		return exception.Wrap(response.ExceptionDatabase, tx.Error)
	}
	return nil
}
