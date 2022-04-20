package util

import "encoding/hex"

func StrHex(str []byte) string {
	return hex.EncodeToString(str)
}

func HexStr(h string) string {
	d, _ := hex.DecodeString(h)
	return string(d)
}
