package tables

const (
	// 用户表
	User = "js_user"
	// 角色表
	Role = "js_role"
	// 用户-角色
	UserRoleRelation = "js_user_role_relation"
	// 权限表
	Permission = "js_permission"
	// 角色-权限
	RolePermissionRelation = "js_role_permission_relation"

	// 域名管理
	DomainMgr = "js_domain_mgr"
	// ip库管理
	BlackIPMgr = "js_black_ip_mgr"
	// ip白名单
	WhiteIP = "js_white_ip"
	// cdn 白名单
	CDN = "js_cdn"

	// js主分类
	JsPrimary = "js_primary"
	// js分类
	JsCategory = "js_category"
	// js管理
	JsManage = "js_manager"
	// 跳转管理
	RedirectManage = "js_redirect_manage"

	// ip 访问类型
	IPStatistics = "js_ip_statistics"
	// UV 访问类型
	UVStatistics = "js_uv_statistics"
	// ip记录
	IPRecode = "js_ip_recode"
)
