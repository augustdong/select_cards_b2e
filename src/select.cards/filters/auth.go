package filters

import (
	beecontext "github.com/astaxie/beego/context"
	"reflect"
	"select.cards/controllers"
)

const (
	registerURI = "/cgi/user/register"
	loginURI    = "/cgi/user/login"
)

var Auth = func(ctx *beecontext.Context) {
	requestURI := ctx.Request.RequestURI
	method := ctx.Request.Method
	// 权限验证API在路由里处理
	if (requestURI == registerURI && method == "POST") || (requestURI == loginURI && method == "POST") {
		// 此处不需要验证
	} else {
		// 验证是否是登陆用户，没登陆返回错误信息
		account := ctx.Input.Session("user_account")
		if account == nil {
			ctx.Input.RunController = reflect.TypeOf(&controllers.AuthFailController{}).Elem()
			ctx.Input.RunMethod = "HandleRequest"
		}
	}
}
