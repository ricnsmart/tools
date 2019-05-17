package util

import (
	"fmt"
	"strconv"
	"strings"
)

// 银行家舍入法
// 不适用于取整，特别是1，会输出0
// 具体规则：
//1. 被修约的数字小于5时，该数字舍去；
//2. 被修约的数字大于5时，则进位；
//3. 被修约的数字等于5时，要看5前面的数字，若是奇数则进位，若是偶数则将5舍掉，即修约后末尾数字都成为偶数；若5的后面还有不为“0”的任何数，则此时无论5的前面是奇数还是偶数，均应进位。
func BankerRounding(num interface{}, accuracy int) (f float64) {

	f, _ = strconv.ParseFloat(fmt.Sprintf("%0."+strconv.Itoa(accuracy)+"f", num), 64)

	return
}

// 将map中的float值进行舍入
func MapRound(m map[string]interface{}, accuracy int) {
	for key, value := range m {
		switch value.(type) {
		case float64:
			m[key] = BankerRounding(value.(float64), accuracy)
		case float32:
			m[key] = float32(BankerRounding(value.(float32), accuracy))
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
