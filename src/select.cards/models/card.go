package models

import (
	"github.com/astaxie/beego/orm"
)

type Card struct {
	Id    int
	Front string
	Back  string
}

func init() {
	orm.RegisterModel(new(Card))
}
