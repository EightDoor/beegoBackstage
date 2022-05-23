package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// PasswordEncryption md5加密
func PasswordEncryption(val string) string {
	d := []byte(val)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
