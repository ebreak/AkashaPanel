package routers

import (
	"AkashaPanel/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/status", &controllers.ApiController{}, "GET:Status")
}
