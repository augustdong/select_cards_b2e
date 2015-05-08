package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"select.cards/defs"
	"select.cards/helpers"
	"select.cards/models"
	"select.cards/utils/encrypt"
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
			"ret": defs.Auth_RegisterUserExist}
		result.Data = data
		this.Data["json"] = result.GetData()
	} else {

		// 不存在该用户，可以进行注册
		user.Pwd = encrypt.EncryptPwd(this.GetString("pwd"))
		_, err = o.Insert(&user)

		if err == nil {
			// 注册成功
			result := helpers.JsonResult{Ret: defs.Server_Ok}
			data := map[string]string{
				"ret": defs.Auth_RegisterOk}
			result.Data = data
			this.Data["json"] = result.GetData()
		} else {
			// 注册失败
			result := helpers.JsonResult{Ret: defs.Server_Ok}
			data := map[string]string{
				"ret": defs.Auth_RegisterFail}
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
	user := models.User{Account: this.GetString("account"), Pwd: encrypt.EncryptPwd(this.GetString("pwd"))}

	o := orm.NewOrm()
	err := o.Read(&user, "account", "pwd")

	this.SessionRegenerateID()
	if err == nil {
		this.SetSession("user_account", user.Account)
		this.SetSession("user_id", user.Id)
		// 登陆成功
		result := helpers.JsonResult{Ret: defs.Server_Ok}
		data := map[string]string{
			"ret": defs.Auth_LoginOk}
		result.Data = data
		this.Data["json"] = result.GetData()
	} else {
		// 登陆失败
		result := helpers.JsonResult{Ret: defs.Server_Ok}
		data := map[string]string{
			"ret": defs.Auth_LoginFail}
		result.Data = data
		this.Data["json"] = result.GetData()
	}

	this.ServeJson()
}
