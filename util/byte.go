package util

import "bytes"

// 去除[]byte中的0
// 适用于提取ascII码
func WipeOutZero(buf []byte) []byte {
	index := bytes.IndexByte(buf, 0)
	// 排除不存在0的情况
	if index == -1 {
		return buf
	}
	return buf[0:index]
}
