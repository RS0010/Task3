package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Data{}, &User{})
	if err != nil {
		panic(err)
	}
}

func userDb() *gorm.DB {
	return db.Model(&User{})
}

func todoDb() *gorm.DB {
	return db.Model(&Data{})
}
