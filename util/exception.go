package util

import (
	"errors"
)

type Exception int

var (
	Success Exception = 2000

	InvalidParams Exception = 3000

	UnmarshalFailed Exception = 6000
	MarshalFailed   Exception = 6001

	RedisPublishFailed Exception = 7000
	SetCacheFailed     Exception = 7001
	GetCacheFailed     Exception = 7002
)

func (e Exception) Error() error {
	return errors.New(e.String())
}

func (e Exception) String() string {
	var str string

	switch e {
	case Success:
		str = "成功"
	case InvalidParams:
		str = "非法参数"
	case UnmarshalFailed:
		str = "解码失败"
	case MarshalFailed:
		str = "编码失败"
	case RedisPublishFailed:
		str = "使用Redis发布消息失败"
	case SetCacheFailed:
		str = "设置Redis缓存失败"
	case GetCacheFailed:
		str = "获取Redis缓存失败"
	default:
		str = "未知的错误类型"
	}

	return str
}

func (e Exception) Name() string {
	var str string

	switch e {
	case Success:
		str = "Success"
	case InvalidParams:
		str = "InvalidParams"
	case UnmarshalFailed:
		str = "UnmarshalFailed"
	case MarshalFailed:
		str = "MarshalFailed"
	case RedisPublishFailed:
		str = "RedisPublishFailed"
	case SetCacheFailed:
		str = "SetCacheFailed"
	case GetCacheFailed:
		str = "GetCacheFailed"
	default:
		str = "UnknownReason"
	}

	return str
}
