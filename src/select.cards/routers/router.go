package routers

import (
	"github.com/astaxie/beego"
	"select.cards/controllers"
	controllersSet "select.cards/controllers/set"
	"select.cards/filters"
)

func init() {
	// 配置过滤器
	// 用户基础权限验证
	beego.InsertFilter("/*", beego.BeforeRouter, filters.Auth)

	// 配置路由
	ns := beego.NewNamespace("/cgi/",
		// 用户注册
		beego.NSRouter("user/register", &controllers.UserRegisterController{}, "post:HandlePost"),
		// 用户登录
		beego.NSRouter("user/login", &controllers.UserLoginController{}, "post:HandlePost"),
		// 获取set详细信息
		beego.NSRouter("set/details", &controllersSet.SetDetailController{}, "get:HandleGet"))
	beego.AddNamespace(ns)

	// 主页逻辑
	beego.Router("/", &controllers.MainController{}, "get:HandleGet")
}
