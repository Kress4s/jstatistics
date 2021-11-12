package models

import (
	"js_statistics/app/models/internal/application"
	"js_statistics/app/models/internal/common"
	"js_statistics/app/models/internal/permission"
)

type (
	// 权限管理
	Base                   = common.Base
	User                   = permission.User
	Role                   = permission.Role
	UserRoleRelation       = permission.UserRoleRelation
	Permission             = permission.Permission
	RolePermissionRelation = permission.RolePermissionRelation

	// 应用管理
	DomainMgr  = application.DomainMgr
	BlackIPMgr = application.BlackIPMgr
	WhiteIP    = application.WhiteIP
)
