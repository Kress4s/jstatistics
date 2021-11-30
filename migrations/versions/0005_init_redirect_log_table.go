package versions

import (
	"js_statistics/app/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// V0005InitRedirectLogTables init redirect log table
var V0005InitRedirectLogTables = &gormigrate.Migration{
	ID: "0005_init_redirect_log_table",
	Migrate: func(tx *gorm.DB) error {
		// 创建 操作人员表，角色表, 操作人员角色关联表，用户登录记录表
		if err := tx.AutoMigrate(
			&models.RedirectLog{},
		); err != nil {
			return err
		}
		return nil
	},
}
