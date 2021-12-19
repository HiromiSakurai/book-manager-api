package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	dbLogger "gorm.io/gorm/logger"

	"github.com/HiromiSakurai/book-manager-api/logger"
)

var (
	DB *gorm.DB
)

func SetUp() {
	var err error
	DB, err = db()
	if err != nil {
		logger.Error("Failed to connect to the database", err)
		panic(err.Error())
	} else {
		logger.Info("Successfully connected to the database")
	}

	DB.Logger.LogMode(dbLogger.LogLevel(4))
}

func db() (database *gorm.DB, err error) {
	dsn := "root:@tcp(localhost:3306)/book_manager_api?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
