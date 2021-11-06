package tools

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var coder = base64.NewEncoding(base64Table)

// Base64Encode ...
func Base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

// EncodeMD5 ...
// encodeMD5 md5加密
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
