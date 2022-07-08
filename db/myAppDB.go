package db

import (
	// "gorm.io/driver/sqlite"
	// "gorm.io/driver/sqlserver"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// const sqlitedbName = "/Users/tahoa/webdev/golang-app/myDatabase.db"
// const mssqlconnString = "sqlserver://:@127.0.0.1:1433?database=Golang-db"
const pgConnection = "host = 127.0.0.1 port = 54321	user = tahoad password = TO@ss2020	dbname = pg-golang-db"

var isMigrated = false

func DB() *gorm.DB {
	// db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	// db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	db, err := gorm.Open(postgres.Open(pgConnection), &gorm.Config{})
	if err != nil {
		panic("Fail to connect database.")
	}
	return db
}

func MigrateUser() {
	db := DB()
	db.AutoMigrate(&UserDB{})
}
func MigrateEmp() {
	db := DB()
	db.AutoMigrate(&EmployeeDB{})
}
