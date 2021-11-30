package vo

import (
	"js_statistics/app/models"
	"time"
)

type RoleReq struct {
	// 角色命
	Name string `json:"name"`
	// 标识符
	Identify string `json:"identify"`
	// 说明
	Description string `json:"description"`
	// 权限ids
	Permissions []int64 `json:"permission"`
}

type RoleUpdateReq struct {
	// 角色命
	Name string `json:"name"`
	// 标识符
	Identify string `json:"identify"`
	// 说明
	Description string `json:"description"`
	// 权限ids
	Permissions []int64 `json:"permission"`
}

func (rup *RoleUpdateReq) ToMap(openID string) map[string]interface{} {
	return map[string]interface{}{
		"name":        rup.Name,
		"identify":    rup.Identify,
		"description": rup.Description,
		"update_by":   openID,
		"update_at":   time.Now(),
	}
}

func (r *RoleReq) ToModel(openID string) *models.Role {
	return &models.Role{
		Name:        r.Name,
		Identify:    r.Identify,
		Description: r.Description,
		Base: models.Base{
			CreateBy: openID,
			UpdateBy: openID,
		},
	}
}

type RoleResp struct {
	Name        string `json:"name"`
	Identify    string `json:"identify"`
	Description string `json:"description"`
	Permissions []int64
	ID          int64 `json:"id"`
}

type RoleBriefResp struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}
