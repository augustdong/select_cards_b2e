package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Set struct {
	Id           int "orm:\"auto\""
	Name         string
	Description  string
	Ownerid      int
	Createtime   time.Time "orm:\"auto_now_add;type(datetime)\""
	Lastedittime time.Time "orm:\"auto_now;type(datetime)\""
}

func init() {
	orm.RegisterModel(new(Set))
}
