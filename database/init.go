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
	initTable(&objects.Config{})
	InitConfig("username", "Nahida")
	InitConfig("password", utils.RandString(16))
	utils.LogInfo("Username is [" + GetConfig("username") + "], password is [" + GetConfig("password") + "].")

	// initTable(&objects.User{})
}

func initTable(i interface{}) {
	if !db.Migrator().HasTable(i) {
		db.Migrator().CreateTable(i)
	}
}
