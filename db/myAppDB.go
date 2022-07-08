package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const dbName = "/Users/tahoa/webdev/golang-app/myDatabase.db"

var isMigrated = false

func DB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("Fail to connect database.")
	}
	return db
}

func Migrate() {
	db := DB()
	db.AutoMigrate(&UserDB{})
}
