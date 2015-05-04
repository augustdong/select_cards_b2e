package defs

const (
	Auth_RegisterDefault   = "0"  // 注册默认返回值
	Auth_RegisterOk        = "1"  // 注册成功
	Auth_RegisterUserExist = "-1" // 注册用户已存在
	Auth_RegisterFail      = "-2" // 注册失败
)

const (
	Auth_LoginDefault = "0"  // 登陆默认返回值
	Auth_LoginOk      = "1"  // 登陆成功
	Auth_LoginFail    = "-1" // 登陆失败
)
