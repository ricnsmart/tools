package util

import (
	"fmt"
	"strconv"
	"strings"
)

// 将map中的float值进行舍入
func MapRound(m map[string]interface{}, accuracy int) {
	for key, value := range m {
		switch value.(type) {
		case float64:
			m[key] = DownRounding(value.(float64), accuracy)
		case float32:
			m[key] = float32(DownRounding(float64(value.(float32)), accuracy))
		}
	}
}

// 保留小数
// 向下取整
func DownRounding(f float64, accuracy int) (r float64) {
	str := fmt.Sprintf(`%f`, f)
	// 整数
	if i := strings.Index(str, "."); i == -1 {
		r, _ = strconv.ParseFloat(str, 64)

	} else {
		if i+accuracy+1 > len(str) {
			return f
		}

		r, _ = strconv.ParseFloat(str[0:i+accuracy+1], 64)
	}

	return
}
