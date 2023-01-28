package util

import (
	"crypto/md5"
	"encoding/hex"
)

// EncodeMD5 MD5加密
// 对上传的文件名做加密后进行写入
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
