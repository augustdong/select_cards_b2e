package encrypt

const salt = "D3F0A7+^DM9A-"

func EncryptPwd(pwd string) string {
	return GetMD5(GetMD5(pwd) + salt)
}
