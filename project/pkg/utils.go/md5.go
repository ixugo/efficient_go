package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// EncryptMD5 获取 md5
func EncryptMD5(c []byte) string {
	if len(c) == 0 {
		return ""
	}
	m := md5.New()
	_, _ = m.Write(c)
	return hex.EncodeToString(m.Sum(nil))
}
