package database

import (
	"fmt"
	"js_statistics/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newPostgreDriver(config *config.Config) *gorm.DB {
	var pageDB *gorm.DB
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=prefer&connect_timeout=%d",
		config.DataBase.DSN.Username,
		config.DataBase.DSN.Password,
		config.DataBase.DSN.Addr,
		config.DataBase.DSN.DB,
		config.DataBase.DSN.ConnectTimeout)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("database connect error: ", err.Error())
	}
	pageDB = db

	if config.DebugModel {
		pageDB = pageDB.Debug()
	}

	rawDB, err := pageDB.DB()
	if err != nil {
		return nil
	}
	rawDB.SetMaxIdleConns(config.DataBase.DSN.MaxIdleConns)
	// rawDB.SetMaxOpenConns(config.DataBase.DSN.MaxOpenConns)
	return pageDB
}
