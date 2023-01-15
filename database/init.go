package database

import (
	"AkashaPanel/objects"
	"AkashaPanel/utils"

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

	// init config
	if !db.Migrator().HasTable(&objects.Config{}) {
		utils.LogInfo("Initializing the database.")
		db.Migrator().CreateTable(&objects.Config{})

		db.Create(&objects.Config{Key: "username", Value: "Nahida"})

		password := utils.RandString(16)
		utils.LogInfo("Default username is [Nahida], password is [" + password + "], please copy that.")
		db.Create(&objects.Config{Key: "password", Value: password})
	}

	// initTable(&objects.User{})
}

func initTable(i interface{}) {
	if !db.Migrator().HasTable(i) {
		db.Migrator().CreateTable(i)
	}
}
