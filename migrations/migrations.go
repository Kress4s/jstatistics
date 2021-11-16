package migrations

import (
	// "js_statistics/app/models/internal/users"
	"js_statistics/app/models"
	"js_statistics/commom/drivers/database"
)

func Migrate() error {
	db := database.GetDriver()
	return db.AutoMigrate(
		// 权限管理
		&models.User{}, &models.Role{}, &models.UserRoleRelation{}, &models.Permission{},
		&models.RolePermissionRelation{},

		// 应用管理
		&models.DomainMgr{}, &models.BlackIPMgr{}, &models.WhiteIP{}, &models.CDN{},
		&models.JsPrimary{}, &models.JsCategory{}, &models.JsManage{}, &models.RedirectManage{},
	)
}
