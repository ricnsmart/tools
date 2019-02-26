package util

import (
	"errors"
)

type (
	Message interface {
		Error() error
		Name() string
		String() string
	}

	Exception int
)

var (
	Success Exception = 2000

	//请求类
	RequestTimeOut Exception = 3000
	InvalidParams  Exception = 3001

	//序列化类
	UnmarshalFailed Exception = 6000
	MarshalFailed   Exception = 6001

	//消息处理类
	RedisPublishFailed   Exception = 7000
	RedisSubscribeFailed Exception = 7002
	SetCacheFailed       Exception = 7003
	GetCacheFailed       Exception = 7004

	MQPublishFailed Exception = 7005

	//文件表格操作类
	CreateXLSXFailed Exception = 8000
	InvalidExcel     Exception = 8001
	UploadFileFailed Exception = 8002
	OpenFileFailed   Exception = 8003
	CreateFileFailed Exception = 8004
	WriteFileFailed  Exception = 8005
	ReadFileFailed   Exception = 8006

	//微服务类
	DomainRecordNotExist        Exception = 9000
	GetDomainRecordsFailed      Exception = 9001
	UpdateDomainRecordFailed    Exception = 9002
	CheckDomainRecordFailed     Exception = 9003
	SetDomainRecordStatusFailed Exception = 9004

	EnableThingFailed     Exception = 9010
	DisableThingFailed    Exception = 9011
	CheckDeviceNameFailed Exception = 9012
	RegisterDeviceFailed  Exception = 9013
	GetDeviceSecretFailed Exception = 9014
	SendCaptchaFailed     Exception = 9020

	// influx
	WriteInfluxFailed Exception = 9030
	QueryInfluxFailed Exception = 9031

	// 通用
	UnknownType Exception = 10000
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
	case RequestTimeOut:
		str = "请求超时"
	case UnmarshalFailed:
		str = "解码失败"
	case MarshalFailed:
		str = "编码失败"
	case RedisPublishFailed:
		str = "向Redis发布消息失败"
	case RedisSubscribeFailed:
		str = "订阅Redis主题失败"
	case SetCacheFailed:
		str = "设置Redis缓存失败"
	case GetCacheFailed:
		str = "获取Redis缓存失败"
	case MQPublishFailed:
		str = "向RabbitMQ投递消息失败"
	case CreateXLSXFailed:
		str = "创建表格失败"
	case InvalidExcel:
		str = "表格中没有有效数据"
	case UploadFileFailed:
		str = "上传文件失败"
	case OpenFileFailed:
		str = "打开文件失败"
	case CreateFileFailed:
		str = "创建文件失败"
	case WriteFileFailed:
		str = "写入文件失败"
	case ReadFileFailed:
		str = "读取文件失败"
	case GetDomainRecordsFailed:
		str = "查找域名记录失败"
	case DomainRecordNotExist:
		str = "不存在解析记录"
	case SetDomainRecordStatusFailed:
		str = "设置解析记录状态失败"
	case UpdateDomainRecordFailed:
		str = "更新域名解析记录失败"
	case CheckDomainRecordFailed:
		str = "检查域名解析失败"
	case EnableThingFailed:
		str = "启用设备失败"
	case DisableThingFailed:
		str = "禁用设备失败"
	case CheckDeviceNameFailed:
		str = "注册设备名称失败"
	case RegisterDeviceFailed:
		str = "批量注册设备失败"
	case GetDeviceSecretFailed:
		str = "获取设备密钥失败"
	case SendCaptchaFailed:
		str = "发送验证码失败"
	case WriteInfluxFailed:
		str = "写入Influx失败"
	case QueryInfluxFailed:
		str = "查询Influx失败"
	case UnknownType:
		str = "未知的类型"
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
	case RequestTimeOut:
		str = "RequestTimeOut"
	case UnmarshalFailed:
		str = "UnmarshalFailed"
	case MarshalFailed:
		str = "MarshalFailed"
	case RedisPublishFailed:
		str = "RedisPublishFailed"
	case RedisSubscribeFailed:
		str = "RedisSubscribeFailed"
	case SetCacheFailed:
		str = "SetCacheFailed"
	case GetCacheFailed:
		str = "GetCacheFailed"
	case CreateXLSXFailed:
		str = "CreateXLSXFailed"
	case InvalidExcel:
		str = "InvalidExcel"
	case UploadFileFailed:
		str = "UploadFileFailed"
	case OpenFileFailed:
		str = "OpenFileFailed"
	case CreateFileFailed:
		str = "CreateFileFailed"
	case WriteFileFailed:
		str = "WriteFileFailed"
	case ReadFileFailed:
		str = "ReadFileFailed"
	case DomainRecordNotExist:
		str = "DomainRecordNotExist"
	case GetDomainRecordsFailed:
		str = "GetDomainRecordsFailed"
	case SetDomainRecordStatusFailed:
		str = "SetDomainRecordStatusFailed"
	case UpdateDomainRecordFailed:
		str = "UpdateDomainRecordFailed"
	case CheckDomainRecordFailed:
		str = "CheckDomainRecordFailed"
	case EnableThingFailed:
		str = "EnableThingFailed"
	case DisableThingFailed:
		str = "DisableThingFailed"
	case CheckDeviceNameFailed:
		str = "CheckDeviceNameFailed"
	case RegisterDeviceFailed:
		str = "RegisterDeviceFailed"
	case GetDeviceSecretFailed:
		str = "GetDeviceSecretFailed"
	case SendCaptchaFailed:
		str = "SendCaptchaFailed"
	case WriteInfluxFailed:
		str = "WriteInfluxFailed"
	case QueryInfluxFailed:
		str = "QueryInfluxFailed"
	case UnknownType:
		str = "UnknownType"
	default:
		str = "UnknownReason"
	}

	return str
}
