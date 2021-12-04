package migrations

import (
	"js_statistics/app/models"
	"time"
)

func InitUser() *models.User {
	return &models.User{Username: "admin", Password: "MTIzNDU2", IsAdmin: true, Status: true, Base: models.Base{
		CreateBy: "admin",
		CreateAt: time.Now(),
		UpdateBy: "admin",
		UpdateAt: time.Now(),
	}}
}

func InitPermissions() []models.Permission {
	p0 := models.Permission{Name: "顶级权限", MenuName: "顶级权限", Route: "/", Identify: "/", Type: 0, ParentID: 0, Index: 0}
	p1 := models.Permission{Name: "权限管理", MenuName: "权限管理", Route: "/admin", Identify: "/admin", Type: 0, ParentID: 1, Index: 0}
	p2 := models.Permission{Name: "操作日志", MenuName: "操作日志", Route: "/admin/log", Identify: "/admin/log", Type: 0, ParentID: 2, Index: 0}
	p3 := models.Permission{Name: "权限规则", MenuName: "权限规则", Route: "/admin/permission", Identify: "/admin/permission", Type: 0, ParentID: 2, Index: 0}
	p4 := models.Permission{Name: "权限规则列表", MenuName: "权限规则列表", Route: "admin.adminpermission.index", Identify: "admin.adminpermission.index", Type: 0, ParentID: 4, Index: 0}
	p5 := models.Permission{Name: "权限规则详情", MenuName: "权限规则详情", Route: "admin.adminpermission.show", Identify: "admin.adminpermission.show", Type: 0, ParentID: 4, Index: 0}
	p6 := models.Permission{Name: "权限规则删除", MenuName: "权限规则删除", Route: "admin.adminpermission.destroy", Identify: "admin.adminpermission.destroy", Type: 0, ParentID: 4, Index: 0}
	p7 := models.Permission{Name: "权限规则添加", MenuName: "权限规则添加", Route: "admin.adminpermission.create", Identify: "admin.adminpermission.create", Type: 0, ParentID: 4, Index: 0}
	p8 := models.Permission{Name: "管理组", ParentID: 2, MenuName: "管理组", Route: "/admin/group", Identify: "/admin/group", Type: 0, Index: 0}
	p9 := models.Permission{Name: "管理组列表", ParentID: 9, MenuName: "管理组列表", Route: "admin.adminrole.index", Identify: "admin.adminrole.index", Type: 0, Index: 0}
	p10 := models.Permission{Name: "管理组详情", ParentID: 9, MenuName: "管理组详情", Route: "admin.adminrole.show", Identify: "admin.adminrole.show", Type: 0, Index: 0}
	p11 := models.Permission{Name: "管理组删除", ParentID: 9, MenuName: "管理组删除", Route: "admin.adminrole.destroy", Identify: "admin.adminrole.destroy", Type: 0, Index: 0}
	p12 := models.Permission{Name: "管理组编辑", ParentID: 9, MenuName: "管理组编辑", Route: "admin.adminrole.update", Identify: "admin.adminrole.update", Type: 0, Index: 0}
	p13 := models.Permission{Name: "管理组添加", ParentID: 9, MenuName: "管理组添加", Route: "admin.adminrole.create", Identify: "admin.adminrole.create", Type: 0, Index: 0}
	p14 := models.Permission{Name: "管理员", ParentID: 2, MenuName: "管理员", Route: "/admin/user", Identify: "/admin/user", Type: 0, Index: 0}
	p15 := models.Permission{Name: "管理员列表", ParentID: 15, MenuName: "管理员列表", Route: "admin.admin.index", Identify: "admin.admin.index", Type: 0, Index: 0}
	p16 := models.Permission{Name: "管理员详情", ParentID: 15, MenuName: "管理员详情", Route: "admin.admin.show", Identify: "admin.admin.show", Type: 0, Index: 0}
	p17 := models.Permission{Name: "管理员删除", ParentID: 15, MenuName: "管理员删除", Route: "admin.admin.destroy", Identify: "admin.admin.destroy", Type: 0, Index: 0}
	p18 := models.Permission{Name: "管理员编辑", ParentID: 15, MenuName: "管理员编辑", Route: "admin.admin.update", Identify: "admin.admin.update", Type: 0, Index: 0}
	p19 := models.Permission{Name: "管理员添加", ParentID: 15, MenuName: "管理员添加", Route: "admin.admin.create", Identify: "admin.admin.create", Type: 0, Index: 0}
	p20 := models.Permission{Name: "数据统计", ParentID: 2, MenuName: "数据统计", Route: "/data", Identify: "/data", Type: 0, Index: 0}
	p21 := models.Permission{Name: "流量统计", ParentID: 21, MenuName: "流量统计", Route: "/data/flow", Identify: "/data/flow", Type: 0, Index: 0}
	p22 := models.Permission{Name: "流量统计流量数据", ParentID: 22, MenuName: "流量统计流量数据", Route: "admin.statics.flow_data", Identify: "admin.statics.flow_data", Type: 0, Index: 0}
	p23 := models.Permission{Name: "流量统计访问量", ParentID: 22, MenuName: "流量统计访问量", Route: "admin.statics.visits", Identify: "admin.statics.visits", Type: 0, Index: 0}
	p24 := models.Permission{Name: "流量统计头部", ParentID: 22, MenuName: "流量统计头部", Route: "admin.statics.accesstop", Identify: "admin.statics.accesstop", Type: 0, Index: 0}
	p25 := models.Permission{Name: "来路统计", ParentID: 21, MenuName: "来路统计", Route: "/data/from", Identify: "/data/from", Type: 0, Index: 0}
	p26 := models.Permission{Name: "来路统计导出", ParentID: 26, MenuName: "来路统计导出", Route: "admin.traffic_log.export", Identify: "admin.traffic_log.export", Type: 0, Index: 0}
	p27 := models.Permission{Name: "来路统计列表", ParentID: 26, MenuName: "来路统计列表", Route: "admin.trafficlog.index", Identify: "admin.trafficlog.index", Type: 0, Index: 0}
	p28 := models.Permission{Name: "应用管理", ParentID: 2, MenuName: "应用管理", Route: "/application", Identify: "/application", Type: 0, Index: 0}
	p29 := models.Permission{Name: "CDN白名单", ParentID: 29, MenuName: "CDN白名单", Route: "/application/cdn-white", Identify: "/application/cdn-white", Type: 0, Index: 0}
	p30 := models.Permission{Name: "CDN白名单列表", ParentID: 30, MenuName: "CDN白名单列表", Route: "admin.cdnwhitelist.index", Identify: "admin.cdnwhitelist.index", Type: 0, Index: 0}
	p31 := models.Permission{Name: "CDN白名单详情", ParentID: 30, MenuName: "CDN白名单列表", Route: "admin.cdnwhitelist.show", Identify: "admin.cdnwhitelist.show", Type: 0, Index: 0}
	p32 := models.Permission{Name: "CDN白名单删除", ParentID: 30, MenuName: "CDN白名单删除", Route: "admin.cdnwhitelist.destroy", Identify: "admin.cdnwhitelist.destroy", Type: 0, Index: 0}
	p33 := models.Permission{Name: "CDN白名单删除", ParentID: 30, MenuName: "CDN白名单编辑", Route: "admin.cdnwhitelist.update", Identify: "admin.cdnwhitelist.update", Type: 0, Index: 0}
	p34 := models.Permission{Name: "CDN白名单添加", ParentID: 30, MenuName: "CDN白名单添加", Route: "admin.cdnwhitelist.create", Identify: "admin.cdnwhitelist.create", Type: 0, Index: 0}
	p35 := models.Permission{Name: "IP白名单", ParentID: 29, MenuName: "IP白名单", Route: "/application/ip-white", Identify: "/application/ip-white", Type: 0, Index: 0}
	p36 := models.Permission{Name: "IP白名单列表", ParentID: 36, MenuName: "IP白名单", Route: "admin.ipwhitelist.index", Identify: "admin.ipwhitelist.index", Type: 0, Index: 0}
	p37 := models.Permission{Name: "IP白名单详情", ParentID: 36, MenuName: "IP白名单详情", Route: "admin.ipwhitelist.show", Identify: "admin.ipwhitelist.show", Type: 0, Index: 0}
	p38 := models.Permission{Name: "IP白名单删除", ParentID: 36, MenuName: "IP白名单删除", Route: "admin.ipwhitelist.destroy", Identify: "admin.ipwhitelist.destroy", Type: 0, Index: 0}
	p39 := models.Permission{Name: "IP白名单编辑", ParentID: 36, MenuName: "IP白名单编辑", Route: "admin.ipwhitelist.update", Identify: "admin.ipwhitelist.update", Type: 0, Index: 0}
	p40 := models.Permission{Name: "IP白名单添加", ParentID: 36, MenuName: "IP白名单添加", Route: "admin.ipwhitelist.create", Identify: "admin.ipwhitelist.create", Type: 0, Index: 0}
	p41 := models.Permission{Name: "IP库管理", ParentID: 29, MenuName: "IP库管理", Route: "/application/ip", Identify: "/application/ip", Type: 0, Index: 0}
	p42 := models.Permission{Name: "IP库管理列表", ParentID: 42, MenuName: "IP库管理列表", Route: "admin.ipmanager.index", Identify: "admin.ipmanager.index", Type: 0, Index: 0}
	p43 := models.Permission{Name: "IP库管理详情", ParentID: 42, MenuName: "IP库管理详情", Route: "admin.ipmanager.show", Identify: "admin.ipmanager.index", Type: 0, Index: 0}
	p44 := models.Permission{Name: "IP库管理删除", ParentID: 42, MenuName: "IP库管理删除", Route: "admin.ipmanager.destroy", Identify: "admin.ipmanager.index", Type: 0, Index: 0}
	p45 := models.Permission{Name: "IP库管理编辑", ParentID: 42, MenuName: "IP库管理编辑", Route: "admin.ipmanager.update", Identify: "admin.ipmanager.index", Type: 0, Index: 0}
	p46 := models.Permission{Name: "IP库管理新建", ParentID: 42, MenuName: "IP库管理新建", Route: "admin.ipmanager.create", Identify: "admin.ipmanager.index", Type: 0, Index: 0}
	p47 := models.Permission{Name: "域名管理", ParentID: 29, MenuName: "域名管理", Route: "/application/js-domain", Identify: "/application/js-domain", Type: 0, Index: 0}
	p48 := models.Permission{Name: "域名管理列表", ParentID: 48, MenuName: "域名管理列表", Route: "admin.domainmanage.index", Identify: "admin.domainmanage.index", Type: 0, Index: 0}
	p49 := models.Permission{Name: "域名管理详情", ParentID: 48, MenuName: "域名管理详情", Route: "admin.domainmanage.show", Identify: "admin.domainmanage.show", Type: 0, Index: 0}
	p50 := models.Permission{Name: "域名管理删除", ParentID: 48, MenuName: "域名管理删除", Route: "admin.domainmanage.destroy", Identify: "admin.domainmanage.destroy", Type: 0, Index: 0}
	p51 := models.Permission{Name: "域名管理编辑", ParentID: 48, MenuName: "域名管理编辑", Route: "admin.domainmanage.update", Identify: "admin.domainmanage.update", Type: 0, Index: 0}
	p52 := models.Permission{Name: "域名管理新建", ParentID: 48, MenuName: "域名管理新建", Route: "admin.domainmanage.create", Identify: "admin.domainmanage.create", Type: 0, Index: 0}
	p53 := models.Permission{Name: "JS分类", ParentID: 29, MenuName: "JS分类", Route: "/application/js-class", Identify: "/application/js-class", Type: 0, Index: 0}
	p54 := models.Permission{Name: "JS分类排序", ParentID: 54, MenuName: "JS分类排序", Route: "admin.category.sort", Identify: "admin.category.sort", Type: 0, Index: 0}
	p55 := models.Permission{Name: "JS分类列表", ParentID: 54, MenuName: "JS分类列表", Route: "admin.category.index", Identify: "admin.category.index", Type: 0, Index: 0}
	p56 := models.Permission{Name: "JS分类详情", ParentID: 54, MenuName: "JS分类详情", Route: "admin.category.show", Identify: "admin.category.show", Type: 0, Index: 0}
	p57 := models.Permission{Name: "JS分类删除", ParentID: 54, MenuName: "JS分类删除", Route: "admin.category.destroy", Identify: "admin.category.destroy", Type: 0, Index: 0}
	p58 := models.Permission{Name: "JS分类编辑", ParentID: 54, MenuName: "JS分类编辑", Route: "admin.category.update", Identify: "admin.category.update", Type: 0, Index: 0}
	p59 := models.Permission{Name: "JS分类新建", ParentID: 54, MenuName: "JS分类新建", Route: "admin.category.create", Identify: "admin.category.create", Type: 0, Index: 0}
	p60 := models.Permission{Name: "JS管理", ParentID: 29, MenuName: "JS管理", Route: "/application/js", Identify: "/application/js", Type: 0, Index: 0}
	p61 := models.Permission{Name: "伪装内容设置详情", ParentID: 61, MenuName: "伪装内容设置详情", Route: "admin.jsdisguise.show", Identify: "admin.jsdisguise.show", Type: 0, Index: 0}
	p62 := models.Permission{Name: "操作日志列表", ParentID: 61, MenuName: "操作日志列表", Route: "admin.urloperationlog.index", Identify: "admin.urloperationlog.index", Type: 0, Index: 0}
	p63 := models.Permission{Name: "跳转管理详情", ParentID: 61, MenuName: "跳转管理详情", Route: "admin.urlmanager.index", Identify: "admin.urlmanager.index", Type: 0, Index: 0}
	p64 := models.Permission{Name: "跳转管理列表", ParentID: 61, MenuName: "跳转管理列表", Route: "admin.urlmanager.show", Identify: "admin.urlmanager.index", Type: 0, Index: 0}
	p65 := models.Permission{Name: "跳转管理删除", ParentID: 61, MenuName: "跳转管理删除", Route: "admin.urlmanager.destroy", Identify: "admin.urlmanager.destroy", Type: 0, Index: 0}
	p66 := models.Permission{Name: "跳转管理编辑", ParentID: 61, MenuName: "跳转管理编辑", Route: "admin.urlmanager.update", Identify: "admin.urlmanager.update", Type: 0, Index: 0}
	p67 := models.Permission{Name: "跳转管理添加", ParentID: 61, MenuName: "跳转管理添加", Route: "admin.urlmanager.create", Identify: "admin.urlmanager.create", Type: 0, Index: 0}
	p68 := models.Permission{Name: "伪装内容设置编辑", ParentID: 61, MenuName: "伪装内容设置编辑", Route: "admin.jsdisguise.update", Identify: "admin.jsdisguise.update", Type: 0, Index: 0}
	p69 := models.Permission{Name: "伪装内容设置创建", ParentID: 61, MenuName: "伪装内容设置创建", Route: "admin.jsdisguise.create", Identify: "admin.jsdisguise.create", Type: 0, Index: 0}
	p70 := models.Permission{Name: "JS管理列表", ParentID: 61, MenuName: "JS管理列表", Route: "admin.jsmanager.index", Identify: "admin.jsmanager.index", Type: 0, Index: 0}
	p71 := models.Permission{Name: "JS管理详情", ParentID: 61, MenuName: "JS管理详情", Route: "admin.jsmanager.show", Identify: "admin.jsmanager.show", Type: 0, Index: 0}
	p72 := models.Permission{Name: "JS管理删除", ParentID: 61, MenuName: "JS管理删除", Route: "admin.jsmanager.destroy", Identify: "admin.jsmanager.destroy", Type: 0, Index: 0}
	p73 := models.Permission{Name: "JS管理编辑", ParentID: 61, MenuName: "JS管理编辑", Route: "admin.jsmanager.update", Identify: "admin.jsmanager.update", Type: 0, Index: 0}
	p74 := models.Permission{Name: "JS管理添加", ParentID: 61, MenuName: "JS管理添加", Route: "admin.jsmanager.create", Identify: "admin.jsmanager.create", Type: 0, Index: 0}
	p75 := models.Permission{Name: "主页", ParentID: 2, MenuName: "主页", Route: "/dashboard", Identify: "/dashboard", Type: 0, Index: 0}
	p76 := models.Permission{Name: "主页用户构成", ParentID: 76, MenuName: "主页用户构成", Route: "admin.statics.area_distribution", Identify: "admin.statics.area_distribution", Type: 0, Index: 0}
	p77 := models.Permission{Name: "主页JS流量排行榜", ParentID: 76, MenuName: "主页JS流量排行榜", Route: "admin.statics.js_charts", Identify: "admin.statics.js_charts", Type: 0, Index: 0}
	return []models.Permission{p0, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12, p13, p14, p15, p16, p17, p18, p19,
		p20, p21, p22, p23, p24, p25, p26, p27, p28, p29, p30, p31, p32, p33, p34, p35, p36, p37, p38, p39, p40, p41,
		p42, p43, p44, p45, p46, p47, p48, p49, p50, p51, p52, p53, p54, p55, p56, p57, p58, p59, p60, p61, p62, p63,
		p64, p65, p66, p67, p68, p69, p70, p71, p72, p73, p74, p75, p76, p77}

}
