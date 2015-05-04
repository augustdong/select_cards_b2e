package routers

import (
	"github.com/astaxie/beego"
	"select.cards/controllers"
	//"hello/filters"
)

func init() {
	// 配置过滤器
	//beego.InsertFilter("*", beego.BeforeRouter, filters.Auth)

	// 普通配置路由
	ns := beego.NewNamespace("/cgi",
		// 用户注册
		beego.NSRouter("/user/register", &controllers.UserRegisterController{}, "post:HandlePost"),
		// 用户登录
		beego.NSRouter("/user/login", &controllers.UserLoginController{}, "post:HandlePost"))
	beego.AddNamespace(ns)
	beego.Router("/", &controllers.MainController{})
}
