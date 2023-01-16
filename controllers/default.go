package controllers

import (
	"github.com/astaxie/beego"
	"github.com/pbnjay/memory"
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

	if resp.Code >= 100 {
		c.Ctx.ResponseWriter.WriteHeader(resp.Code)
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

func (c *ApiController) Status() {
	type MemoryStruct struct {
		Total uint64
		Avail uint64
	}
	type StatusStruct struct {
		Memory MemoryStruct
	}

	c.response(0, "Akasha is running.", StatusStruct{
		Memory: MemoryStruct{
			Total: memory.TotalMemory(),
			Avail: memory.FreeMemory(),
		},
	})
}
