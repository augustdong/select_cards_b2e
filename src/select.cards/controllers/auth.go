package controllers

import (
	"github.com/astaxie/beego"
)

type AuthFailController struct {
	beego.Controller
}

func (c *AuthFailController) HandleRequest() {
	c.Abort("ServerAuthFailJson")
	c.Data["json"] = "dddddddddddd"
	c.ServeJson()
}
