package defs

const (
	Server_Default     = "0"  // 服务器默认返回值
	Server_Ok          = "1"  // 服务器正确处理
	Server_Error       = "-1" // 服务器内部错误
	Server_AuthFail    = "-2" // 服务器权限验证失败
	Server_ParamsError = "-3" // 服务器请求参数错误
)
