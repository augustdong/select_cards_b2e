package filters

import (
	"fmt"
	beecontext "github.com/astaxie/beego/context"
	"reflect"
	"select.cards/controllers"
)

var Auth = func(ctx *beecontext.Context) {
	// 验证是否是合法用户，不合法返回错误信息
	ctx.Input.RunController = reflect.TypeOf(&controllers.AuthController{}).Elem()
	ctx.Input.RunMethod = "HandleAuth"
}
