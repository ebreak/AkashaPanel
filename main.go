package main

import (
	_ "AkashaPanel/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.Listen.HTTPPort = 6244
	beego.Run()
}
