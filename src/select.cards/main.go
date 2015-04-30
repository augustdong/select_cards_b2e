package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"select.cards/controllers"
	_ "select.cards/routers"
)

func init() {
	// 配置错误处理
	beego.ErrorController(&controllers.ErrorController{})

	// 数据库设置
	orm.RegisterDriver("mymysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "augustdong:bo0567BOQ@tcp(52.74.106.80:3306)/select_cards?charset=utf8", 30)

	// session设置
	// 持久化暂时设置为文件方式
	beego.SessionOn = true
	beego.SessionProvider = "file"
	beego.SessionSavePath = "./tmp"
	beego.SessionName = "session"
}

func main() {
	beego.Run()
}
