package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	TodoDB *gorm.DB
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
	TodoDB = db.Model(&Data{})
}
