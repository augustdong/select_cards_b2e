package helpers

import (
	"strconv"
)

type JsonResult struct {
	Ret  int // 0为默认值、1代表成功、负数为错误
	Data interface{}
	Msg  string
}

type JsonResultInterface interface {
	GetData()
}

func (r *JsonResult) GetData() interface{} {
	tmp := make(map[string]interface{})
	tmp["ret"] = strconv.Itoa(r.Ret)

	if r.Msg != "" {
		tmp["msg"] = r.Msg
	}
	if r.Data != nil {
		tmp["data"] = r.Data
	}

	return tmp
}
