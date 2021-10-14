package base64

import "encoding/base64"

func DecodeToString(s string) string {
	b, _ := base64.StdEncoding.DecodeString(s)
	return string(b)
}

func EncodeToString(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
