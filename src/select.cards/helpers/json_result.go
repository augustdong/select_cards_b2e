package helpers

type JsonResult struct {
	Ret  string
	Data interface{}
	Msg  string
}

type JsonResultInterface interface {
	GetData()
}

func (r *JsonResult) GetData() interface{} {
	tmp := make(map[string]interface{})
	tmp["ret"] = r.Ret

	if r.Msg != "" {
		tmp["msg"] = r.Msg
	}
	if r.Data != nil {
		tmp["data"] = r.Data
	}

	return tmp
}
