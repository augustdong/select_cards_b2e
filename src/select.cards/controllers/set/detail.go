package set

import (
	"github.com/astaxie/beego"
)

type SetDetailController struct {
	beego.Controller
}

func (c *SetDetailController) HandleGet() {
	c.Data["json"] = "Are'u finding me?"
	c.ServeJson()
}
