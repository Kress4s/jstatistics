package vo

import "js_statistics/app/models"

type PermissionReq struct {
	// 名称
	Name string `json:"name"`
	// 菜单名称
	MenuName string `json:"menu_name"`
	// 路由
	Route string `json:"route"`
	// 权限标识
	Identify string `json:"identify"`
	// 权限类型   0: 菜单权限 1: 操作权限
	Type int `json:"type"`
	// 父级菜单的ID,最高级为 0
	ParentID uint `json:"parent_id"`
}

func (p PermissionReq) ToModel() models.Permission {
	return models.Permission{
		Name:     p.Name,
		MenuName: p.MenuName,
		Route:    p.Route,
		Identify: p.Identify,
		Type:     p.Type,
		ParentID: p.ParentID,
	}
}

type PermissionResp struct {
	// ID
	ID uint `json:"id"`
	// 名称
	Name string `json:"name"`
	// 菜单名称
	MenuName string `json:"menu_name"`
	// 路由
	Route string `json:"route"`
	// 权限标识
	Identify string `json:"identify"`
	// 权限类型   0: 菜单权限 1: 操作权限
	Type int `json:"type"`
	// 索引
	Index int `json:"index"`
	// 父级菜单的ID,最高级为 0
	ParentID uint `json:"parent_id"`
}

type PermissionUpdateReq struct {
	// 名称
	Name string `json:"name"`
	// 菜单名称
	MenuName string `json:"menu_name"`
	// 路由
	Route string `json:"route"`
	// 权限标识
	Identify string `json:"identify"`
	// 权限类型   0: 菜单权限 1: 操作权限
	Type int `json:"type"`
	// 索引
	Index int `json:"index"`
	// 父级菜单的ID,最高级为 0
	ParentID uint `json:"parent_id"`
}

func (pur *PermissionUpdateReq) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":      pur.Name,
		"menu_name": pur.MenuName,
		"route":     pur.Route,
		"identify":  pur.Identify,
		"type":      pur.Type,
		// TODO index is not known
		// "index": pur.Index,
		"parent_id": pur.ParentID,
	}
}

type PermissionTree struct {
	// ID
	ID uint `json:"id"`
	// 名称
	Name string `json:"name"`
	// 菜单名称
	MenuName string `json:"menu_name"`
	// 路由
	Route string `json:"route"`
	// 索引
	Index int `json:"index"`
	// 父级ID
	ParentID uint `json:"parent_id"`
	// 子级权限
	SubPermissions []*PermissionTree `json:"sub_permissions"`
}

func NewPermissionTree(p *models.Permission) *PermissionTree {
	return &PermissionTree{
		ID:       p.ID,
		Name:     p.Name,
		MenuName: p.MenuName,
		Route:    p.Route,
		Index:    p.Index,
		ParentID: p.ParentID,
	}
}
