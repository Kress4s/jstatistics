package versions

import (
	"js_statistics/app/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// V0004InitLogTables init log table
var V0004InitLogTables = &gormigrate.Migration{
	ID: "0004_init_log_table",
	Migrate: func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(
			&models.SystemLog{},
		); err != nil {
			return err
		}
		return nil
	},
}
