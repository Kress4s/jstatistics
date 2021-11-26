package migrations

import (
	"js_statistics/commom/drivers/database"
	"js_statistics/migrations/versions"

	"github.com/go-gormigrate/gormigrate/v2"
)

// func Migrate() error {
// 	db := database.GetDriver()
// 	// 初始化管理员
// 	db.Create(InitUser())
// 	// 初始化权限菜单
// 	db.Create(InitPermissions())
// 	return db.AutoMigrate(
// 		// 权限管理
// 		&models.User{}, &models.Role{}, &models.UserRoleRelation{}, &models.Permission{},
// 		&models.RolePermissionRelation{},

// 		// 应用管理
// 		&models.DomainMgr{}, &models.BlackIPMgr{}, &models.WhiteIP{}, &models.CDN{},
// 		&models.JsPrimary{}, &models.JsCategory{}, &models.JsManage{}, &models.RedirectManage{},
// 		&models.Faker{},

// 		// 统计表
// 		&models.IPStatistics{}, &models.UVStatistics{}, &models.IPRecode{},

// 		// 文件表
// 		&models.Object{},
// 	)
// }

var migrations = []*gormigrate.Migration{
	// init table
	versions.V0001InitTables,
	// init views
	versions.V0002InitViews,
	// init product data
	versions.V0003InitProjectData,
}

func Migrate() error {
	return gormigrate.New(database.GetDriver(), &gormigrate.Options{
		TableName:                 "mems_migrations",
		IDColumnName:              "id",
		IDColumnSize:              255,
		UseTransaction:            true,
		ValidateUnknownMigrations: true,
	}, migrations).Migrate()
}
