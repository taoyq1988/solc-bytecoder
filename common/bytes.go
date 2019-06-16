package common

import "encoding/hex"

func Bytes2Hex(data []byte) string {
	return hex.EncodeToString(data)
}
