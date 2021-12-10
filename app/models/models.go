package models

import (
	"js_statistics/app/models/internal/application"
	"js_statistics/app/models/internal/common"
	"js_statistics/app/models/internal/dash_board"
	"js_statistics/app/models/internal/permission"
)

type (
	// 权限管理
	Base                   = common.Base
	User                   = permission.User
	UserListInfo           = permission.UserListInfo
	Role                   = permission.Role
	UserRoleRelation       = permission.UserRoleRelation
	UserCategoryRelation   = permission.UserCategoryRelation
	UserPrimaryRelation    = permission.UserPrimaryRelation
	Permission             = permission.Permission
	RolePermissionRelation = permission.RolePermissionRelation
	UserToMenus            = permission.UserToMenus

	// 应用管理
	DomainMgr        = application.DomainMgr
	BlackIPMgr       = application.BlackIPMgr
	WhiteIP          = application.WhiteIP
	CDN              = application.CDN
	JsPrimary        = application.JsPrimary
	AllsCategory     = application.AllsCategory
	JsCategory       = application.JsCategory
	JsManage         = application.JsManage
	JsManageListView = application.JsManageListView
	RedirectManage   = application.RedirectManage
	IPStatistics     = dash_board.IPStatistics
	UVStatistics     = dash_board.UVStatistics
	IPRecode         = dash_board.IPRecode
	Object           = common.Object
	Faker            = application.Faker
	SystemLog        = common.SystemLog
	RedirectLog      = application.RedirectLog

	// view
	IPVisitStatistic = dash_board.IPVisitStatistic
	UVisitStatistic  = dash_board.UVisitStatistic
	RegionStatistic  = dash_board.RegionStatistic
	JSVisitStatistic = dash_board.JSVisitStatistic

	// data flow view
	FlowDataView      = dash_board.FlowDataView
	FlowDataStatistic = dash_board.FlowDataStatistic

	FromAnalysisView = dash_board.FromAnalysisView
)
