package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error

	db, err = gorm.Open(sqlite.Open(".AkashaPanel/database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// initTable(&objects.User{})
}

func initTable(i interface{}) {
	if !db.Migrator().HasTable(i) {
		db.Migrator().CreateTable(i)
	}
}
