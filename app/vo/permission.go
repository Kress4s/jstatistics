package vo

import (
	"js_statistics/app/models"
	"time"
)

type PermissionReq struct {
	// 名称
	Name string `json:"name"`
	// 菜单名称
	MenuName string `json:"menu_name"`
	// 路由
	Route string `json:"route"`
	// 权限标识
	Identify string `json:"identify"`
	// 权限类型   1: 菜单权限 0: 操作权限
	Type int `json:"type"`
	// 父级菜单的ID,最高级为 0
	ParentID int64 `json:"parent_id"`
}

func (p PermissionReq) ToModel(openID string) models.Permission {
	return models.Permission{
		Name:     p.Name,
		MenuName: p.MenuName,
		Route:    p.Route,
		Identify: p.Identify,
		Type:     p.Type,
		ParentID: p.ParentID,
		Base: models.Base{
			CreateBy: openID,
			UpdateBy: openID,
		},
	}
}

type PermissionResp struct {
	Name     string `json:"name"`
	MenuName string `json:"menu_name"`
	Route    string `json:"route"`
	Identify string `json:"identify"`
	ID       int64  `json:"id"`
	Type     int    `json:"type"`
	Index    int    `json:"index"`
	ParentID int64  `json:"parent_id"`
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
	ParentID int64 `json:"parent_id"`
}

func (pur *PermissionUpdateReq) ToMap(openID string) map[string]interface{} {
	return map[string]interface{}{
		"name":      pur.Name,
		"menu_name": pur.MenuName,
		"route":     pur.Route,
		"identify":  pur.Identify,
		"type":      pur.Type,
		// "index": pur.Index,
		"parent_id": pur.ParentID,
		"update_by": openID,
		"update_at": time.Now(),
	}
}

type PermissionTree struct {
	Name           string            `json:"name"`
	MenuName       string            `json:"menu_name"`
	Identify       string            `json:"identify"`
	Route          string            `json:"route"`
	SubPermissions []*PermissionTree `json:"sub_permissions"`
	Index          int               `json:"index"`
	Type           int               `json:"type"`
	ParentID       int64             `json:"parent_id"`
	ID             int64             `json:"id"`
}

func NewPermissionTree(p *models.Permission) *PermissionTree {
	return &PermissionTree{
		ID:       p.ID,
		Name:     p.Name,
		MenuName: p.MenuName,
		Route:    p.Route,
		Type:     p.Type,
		Identify: p.Identify,
		Index:    p.Index,
		ParentID: p.ParentID,
	}
}
