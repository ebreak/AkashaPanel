package main

import (
	"AkashaPanel/database"
	_ "AkashaPanel/routers"
	"os"

	"github.com/astaxie/beego"
)

func main() {
	os.Mkdir(".AkashaPanel", os.ModePerm)

	database.Init()

	beego.BConfig.Listen.HTTPPort = 6244
	beego.Run()
}
