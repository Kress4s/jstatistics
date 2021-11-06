package migrations

import (
	// "js_statistics/app/models/internal/users"
	"js_statistics/app/models"
	"js_statistics/commom/drivers/database"
)

func Migrate() error {
	db := database.GetDriver()
	return db.AutoMigrate(
		&models.User{}, &models.Role{}, &models.UserRoleRelation{},
	)
}
