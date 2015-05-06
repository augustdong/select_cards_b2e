package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) HandleGet() {
	c.Redirect("/p/index.html", 301)
}
