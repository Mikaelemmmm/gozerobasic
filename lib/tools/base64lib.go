package tools

import "encoding/base64"

// Base64EncodeString 编码 []byte(str)
func Base64EncodeString(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Base64DecodeString 解码
func Base64DecodeString(str string) (string, []byte) {
	resBytes, _ := base64.StdEncoding.DecodeString(str)
	return string(resBytes), resBytes
}