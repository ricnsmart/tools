package util

import (
	"errors"
	"fmt"
)

type Exception uint16

var (
	Success Exception = 2000

	InvalidParams Exception = 3000

	UnmarshalFailed Exception = 6000
	MarshalFailed   Exception = 6001

	redisPublishFailed Exception = 7000
	redisSetFailed     Exception = 7001
)

func (m Exception) Error() error {
	return errors.New(m.String())
}

func (m Exception) String() string {
	var str string

	switch m {
	case Success:
		str = fmt.Sprintf("成功")
	case InvalidParams:
		str = fmt.Sprintf("非法参数")
	case UnmarshalFailed:
		str = fmt.Sprintf("解码失败")
	case MarshalFailed:
		str = fmt.Sprintf("编码失败")
	case redisPublishFailed:
		str = fmt.Sprintf("使用Redis发布消息失败")
	case redisSetFailed:
		str = fmt.Sprintf("设置Redis缓存失败")
	default:
		str = fmt.Sprintf("未知的错误类型")
	}

	return str
}
