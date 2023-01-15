package routers

import (
	"AkashaPanel/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/", &controllers.ApiController{}, "GET:Index")
}
