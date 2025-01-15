package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func InitDB(dbName string) *gorm.DB {
	db.err := gorm.Open(sqlite.Open(dbName),&gorm.config{})

	if err != null {
		panic("failed to connect to database")
	}

	DB = db
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Book{})
}