package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"select.cards/defs"
	"select.cards/helpers"
	"select.cards/models"
)

const (
	RegisterOk        = "1"
	RegisterUserExist = "-1"
	RegisterFail      = "-2"
)

// 注册相关
type UserRegisterController struct {
	beego.Controller
}

func (this *UserRegisterController) HandlePost() {
	user := models.User{Account: this.GetString("account")}

	o := orm.NewOrm()
	err := o.Read(&user)

	if err == nil {
		// 该用户已存在
		result := helpers.JsonResult{Ret: defs.Server_Ok}
		data := map[string]string{
			"ret": RegisterUserExist}
		result.Data = data
		this.Data["json"] = result.GetData()
	} else {

		// 不存在该用户，可以进行注册
		user.Pwd = this.GetString("pwd")
		_, err = o.Insert(&user)

		if err == nil {
			// 注册成功
			result := helpers.JsonResult{Ret: defs.Server_Ok}
			data := map[string]string{
				"ret": RegisterOk}
			result.Data = data
			this.Data["json"] = result.GetData()
		} else {
			// 注册失败
			result := helpers.JsonResult{Ret: defs.Server_Ok}
			data := map[string]string{
				"ret": RegisterFail}
			result.Data = data
			this.Data["json"] = result.GetData()
		}
	}

	this.ServeJson()
}

// 登陆相关
type UserLoginController struct {
	beego.Controller
}

func (this *UserLoginController) HandlePost() {
	user := models.User{Account: this.GetString("account"), Pwd: this.GetString("pwd")}

	o := orm.NewOrm()
	err := o.Read(&user)

	if err == nil {
		// 存在
		fmt.Println("cunzai")
	} else {
		// 不存在
		fmt.Println("buuuuuuuuuucunzai")
	}

	this.ServeJson()
}
