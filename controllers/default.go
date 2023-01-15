package controllers

import (
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

type Response struct {
	Code int
	Msg  string
	Data interface{}
}

func (c *ApiController) response(args ...interface{}) {
	resp := Response{Code: 0}
	switch len(args) {
	case 3:
		resp.Data = args[2]
		fallthrough
	case 2:
		resp.Msg = args[1].(string)
		fallthrough
	case 1:
		resp.Code = args[0].(int)
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *ApiController) Status() {
	c.response(0, "Akasha is running correctly.")
}
