package versions

import (
	"js_statistics/migrations/views"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// V0002InitViews init views
var V0002InitViews = &gormigrate.Migration{
	ID: "0002_init_views",
	Migrate: func(tx *gorm.DB) error {
		if err := AutoMigrateView(
			tx,
			views.CreateFlowDataView,
		); err != nil {
			return err
		}
		return nil
	},
}

func AutoMigrateView(tx *gorm.DB, functions ...func(tx *gorm.DB) error) error {
	for _, fn := range functions {
		if err := fn(tx); err != nil {
			return err
		}
	}
	return nil
}
