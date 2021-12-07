package versions

import (
	"js_statistics/app/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// V0006InitUserCategoryPrimaryTables init user category/Primary table
var V0006InitUserCategoryPrimaryTables = &gormigrate.Migration{
	ID: "0006_init_user_c_p_table",
	Migrate: func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(
			&models.UserCategoryRelation{},
			&models.UserPrimaryRelation{},
		); err != nil {
			return err
		}
		return nil
	},
}
