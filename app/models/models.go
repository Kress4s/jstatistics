package models

import (
	"js_statistics/app/models/internal/permission"
)

type (
	User                   = permission.User
	Role                   = permission.Role
	UserRoleRelation       = permission.UserRoleRelation
	Permission             = permission.Permission
	RolePermissionRelation = permission.RolePermissionRelation
)
