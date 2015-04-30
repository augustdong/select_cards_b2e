package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id      int
	Account string "orm:\"pk\""
	Pwd     string
}

func init() {
	orm.RegisterModel(new(User))
}
