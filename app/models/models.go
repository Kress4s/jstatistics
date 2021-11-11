package models

import (
	"js_statistics/app/models/internal/common"
	"js_statistics/app/models/internal/permission"
)

type (
	Base                   = common.Base
	User                   = permission.User
	Role                   = permission.Role
	UserRoleRelation       = permission.UserRoleRelation
	Permission             = permission.Permission
	RolePermissionRelation = permission.RolePermissionRelation
)
