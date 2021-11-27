package versions

import (
	"js_statistics/app/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// V0001InitTables init tables
var V0001InitTables = &gormigrate.Migration{
	ID: "0001_init_tables",
	Migrate: func(tx *gorm.DB) error {
		// 创建 操作人员表，角色表, 操作人员角色关联表，用户登录记录表
		if err := tx.AutoMigrate(
			// 权限管理
			&models.User{},
			&models.Role{},
			&models.UserRoleRelation{},
			&models.Permission{},
			&models.RolePermissionRelation{},

			// 应用管理
			&models.DomainMgr{},
			&models.BlackIPMgr{},
			&models.WhiteIP{},
			&models.CDN{},
			&models.JsPrimary{},
			&models.JsCategory{},
			&models.JsManage{},
			&models.RedirectManage{},
			&models.Faker{},

			// 统计表
			&models.IPStatistics{},
			&models.UVStatistics{},
			&models.IPRecode{},

			// 文件表
			&models.Object{},

			// 日志表
			&models.SystemLog{},
		); err != nil {
			return err
		}
		return nil
	},
}
