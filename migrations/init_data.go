package migrations

import (
	"js_statistics/app/models"
	"time"
)

func InitUser() *models.User {
	return &models.User{ID: 1, Username: "admin", Password: "MTIzNDU2", IsAdmin: true, Status: true, Base: models.Base{
		CreateBy: "admin",
		CreateAt: time.Now(),
		UpdateBy: "admin",
		UpdateAt: time.Now(),
	}}
}

func InitPermissions() []models.Permission {
	p0 := models.Permission{ID: 1, Name: "顶级权限", MenuName: "顶级权限", Route: "/", Identify: "/", Type: 0, ParentID: 0, Index: 0}
	p1 := models.Permission{ID: 2, Name: "权限管理", MenuName: "权限管理", Route: "/admin", Identify: "/admin", Type: 0, ParentID: 1, Index: 0}
	p2 := models.Permission{ID: 3, Name: "操作日志", MenuName: "操作日志", Route: "/admin/log", Identify: "/admin/log", Type: 0, ParentID: 2, Index: 0}
	p3 := models.Permission{ID: 4, Name: "权限规则", MenuName: "权限规则", Route: "/admin/permission", Identify: "/admin/permission", Type: 0, ParentID: 2, Index: 0}
	p4 := models.Permission{ID: 5, Name: "权限规则列表", MenuName: "权限规则列表", Route: "admin.adminpermission.index", Identify: "admin.adminpermission.index", Type: 0, ParentID: 4, Index: 0}
	p5 := models.Permission{ID: 6, Name: "权限规则详情", MenuName: "权限规则详情", Route: "admin.adminpermission.show", Identify: "admin.adminpermission.show", Type: 0, ParentID: 4, Index: 0}
	p6 := models.Permission{ID: 7, Name: "权限规则删除", MenuName: "权限规则删除", Route: "admin.adminpermission.destroy", Identify: "admin.adminpermission.destroy", Type: 0, ParentID: 4, Index: 0}
	p7 := models.Permission{ID: 8, Name: "权限规则添加", MenuName: "权限规则添加", Route: "admin.adminpermission.create", Identify: "admin.adminpermission.create", Type: 0, ParentID: 4, Index: 0}
	p8 := models.Permission{ID: 9, Name: "管理组", ParentID: 2, MenuName: "管理组", Route: "/admin/group", Identify: "/admin/group", Type: 0, Index: 0}
	p9 := models.Permission{ID: 10, Name: "管理组列表", ParentID: 9, MenuName: "管理组列表", Route: "admin.adminrole.index", Identify: "admin.adminrole.index", Type: 0, Index: 0}
	p10 := models.Permission{ID: 11, Name: "管理组详情", ParentID: 9, MenuName: "管理组详情", Route: "admin.adminrole.show", Identify: "admin.adminrole.show", Type: 0, Index: 0}
	p11 := models.Permission{ID: 12, Name: "管理组删除", ParentID: 9, MenuName: "管理组删除", Route: "admin.adminrole.destroy", Identify: "admin.adminrole.destroy", Type: 0, Index: 0}
	p12 := models.Permission{ID: 13, Name: "管理组编辑", ParentID: 9, MenuName: "管理组编辑", Route: "admin.adminrole.update", Identify: "admin.adminrole.update", Type: 0, Index: 0}
	p13 := models.Permission{ID: 14, Name: "管理组添加", ParentID: 9, MenuName: "管理组添加", Route: "admin.adminrole.create", Identify: "admin.adminrole.create", Type: 0, Index: 0}
	p14 := models.Permission{ID: 15, Name: "管理员", ParentID: 2, MenuName: "管理员", Route: "/admin/user", Identify: "/admin/user", Type: 0, Index: 0}
	p15 := models.Permission{ID: 16, Name: "管理员列表", ParentID: 15, MenuName: "管理员列表", Route: "admin.admin.index", Identify: "admin.admin.index", Type: 0, Index: 0}
	p16 := models.Permission{ID: 17, Name: "管理员详情", ParentID: 15, MenuName: "管理员详情", Route: "admin.admin.show", Identify: "admin.admin.show", Type: 0, Index: 0}
	p17 := models.Permission{ID: 18, Name: "管理员删除", ParentID: 15, MenuName: "管理员删除", Route: "admin.admin.destroy", Identify: "admin.admin.destroy", Type: 0, Index: 0}
	p18 := models.Permission{ID: 19, Name: "管理员编辑", ParentID: 15, MenuName: "管理员编辑", Route: "admin.admin.update", Identify: "admin.admin.update", Type: 0, Index: 0}
	p19 := models.Permission{ID: 20, Name: "管理员添加", ParentID: 15, MenuName: "管理员添加", Route: "admin.admin.create", Identify: "admin.admin.create", Type: 0, Index: 0}
	p20 := models.Permission{ID: 21, Name: "数据统计", ParentID: 2, MenuName: "数据统计", Route: "/data", Identify: "/data", Type: 0, Index: 0}
	p21 := models.Permission{ID: 22, Name: "流量统计", ParentID: 21, MenuName: "流量统计", Route: "/data/flow", Identify: "/data/flow", Type: 0, Index: 0}
	p22 := models.Permission{ID: 23, Name: "流量统计流量数据", ParentID: 22, MenuName: "流量统计流量数据", Route: "admin.statics.flow_data", Identify: "admin.statics.flow_data", Type: 0, Index: 0}
	p23 := models.Permission{ID: 24, Name: "流量统计访问量", ParentID: 22, MenuName: "流量统计访问量", Route: "admin.statics.visits", Identify: "admin.statics.visits", Type: 0, Index: 0}
	p24 := models.Permission{ID: 25, Name: "流量统计头部", ParentID: 22, MenuName: "流量统计头部", Route: "admin.statics.accesstop", Identify: "admin.statics.accesstop", Type: 0, Index: 0}
	p25 := models.Permission{ID: 26, Name: "来路统计", ParentID: 21, MenuName: "来路统计", Route: "/data/from", Identify: "/data/from", Type: 0, Index: 0}
	p26 := models.Permission{ID: 27, Name: "来路统计导出", ParentID: 26, MenuName: "来路统计导出", Route: "admin.traffic_log.export", Identify: "admin.traffic_log.export", Type: 0, Index: 0}
	p27 := models.Permission{ID: 28, Name: "来路统计列表", ParentID: 26, MenuName: "来路统计列表", Route: "admin.trafficlog.index", Identify: "admin.trafficlog.index", Type: 0, Index: 0}
	p28 := models.Permission{ID: 29, Name: "应用管理", ParentID: 2, MenuName: "应用管理", Route: "/application", Identify: "/application", Type: 0, Index: 0}
	p29 := models.Permission{ID: 30, Name: "CDN白名单", ParentID: 29, MenuName: "CDN白名单", Route: "/application/cdn-white", Identify: "/application/cdn-white", Type: 0, Index: 0}
	p30 := models.Permission{ID: 31, Name: "CDN白名单列表", ParentID: 30, MenuName: "CDN白名单列表", Route: "admin.cdnwhitelist.index", Identify: "admin.cdnwhitelist.index", Type: 0, Index: 0}
	return []models.Permission{p0, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15, p16, p17, p18, p19, p20, p21, p22, p23, p24, p25, p26, p27, p28, p29, p30}

}
