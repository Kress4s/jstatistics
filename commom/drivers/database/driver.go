package database

import (
	"js_statistics/config"
	"sync"

	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
	once       sync.Once
)

func GetDriver() *gorm.DB {
	once.Do(func() {
		conf := config.GetConfig()
		switch conf.DataBase.Type {
		case "postgres":
			dbInstance = newPostgreDriver(conf)
		default:
			dbInstance = nil
		}
	})
	return dbInstance
}
