package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5(source string) string {
	h := md5.New()
	h.Write([]byte(source))
	return hex.EncodeToString(h.Sum(nil))
}
