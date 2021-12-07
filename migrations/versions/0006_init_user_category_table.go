package versions

import (
	"js_statistics/app/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// V0006InitUserCategoryTables init user category table
var V0006InitUserCategoryTables = &gormigrate.Migration{
	ID: "0006_init_user_category_table",
	Migrate: func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(
			&models.UserCategoryRelation{},
		); err != nil {
			return err
		}
		return nil
	},
}
