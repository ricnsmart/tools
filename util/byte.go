package util

import "bytes"

// 去除[]byte中的0
func WipeOutZero(buf []byte) []byte {
	index := bytes.IndexByte(buf, 0)
	return buf[0:index]
}
