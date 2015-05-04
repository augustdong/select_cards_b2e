package controllers

import (
	"github.com/astaxie/beego"
	"select.cards/defs"
	"select.cards/helpers"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) ErrorServerAuthFailJson() {
	result := helpers.JsonResult{Ret: defs.Server_AuthFail}
	c.Data["json"] = result.GetData()
	c.ServeJson()
	c.Ctx.Output.SetStatus(200)
}

func (c *ErrorController) ErrorServerErrorJson() {
	result := helpers.JsonResult{Ret: defs.Server_Error}
	c.Data["json"] = result.GetData()
	c.ServeJson()
	c.Ctx.Output.SetStatus(200)
}

func (c *ErrorController) ErrorServerParamsError() {
	result := helpers.JsonResult{Ret: defs.Server_ParamsError}
	c.Data["json"] = result.GetData()
	c.ServeJson()
	c.Ctx.Output.SetStatus(200)
}
