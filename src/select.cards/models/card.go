package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Card struct {
	Id           int    "orm:\"auto\""
	Front        string "orm:\"null\""
	Back         string "orm:\"null\""
	Setid        int
	Lastedittime time.Time "orm:\"auto_now;type(datetime)\""
}

func init() {
	orm.RegisterModel(new(Card))
}
