package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id         int
	Account    string "orm:\"pk\""
	Pwd        string
	Createtime time.Time "orm:\"auto_now_add;type(datetime)\""
}

func init() {
	orm.RegisterModel(new(User))
}
