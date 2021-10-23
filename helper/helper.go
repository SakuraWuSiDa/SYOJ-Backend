package helper

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5OfStr(str string) string { //获取一个字符串的MD5
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}