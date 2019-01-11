package util

import (
	"errors"
	"github.com/labstack/echo"
	"regexp"
)

var ErrHeartBeatPackage = errors.New("收到了心跳包")

var WarnInvalidParams = errors.New("非法请求")

// 域名，只匹配数字、字母和.
var DomainRecordReg = regexp.MustCompile(`^(\w|\.)+$`)

// 包含心跳包正则判断
var HeartBeatReg = regexp.MustCompile("^f\\d{10}")

const (
	Success = 2000

	UserNotExist                 = 4010
	UserNamePasswordNotMatch     = 4011
	RoleExistAssociationUser     = 4012
	RootDeleteProhibited         = 4013
	InvalidParams                = 4014
	InvalidMobile                = 4015
	InvalidCaptcha               = 4016
	AdminDeleteProhibited        = 4017
	RoleNotExist                 = 4018
	OrganizationDeleteProhibited = 4019
	MobileInUsed                 = 4020
	UserNameInUsed               = 4021
	InvalidUser                  = 4022
	OrganizationNotExist         = 4023
	OrganizationNameInUsed       = 4024
	DevicesNotExist              = 4025
	DeviceOffline                = 4026
	RequestTimeOut               = 4027
	InvalidExcel                 = 4028
	DomainRecordNotExist         = 4029
	UnknownBindType              = 4030
	UserPermissionsUndefined     = 4031
	UniqueViolation              = 4032
	UnSupportType                = 4033
	AccountDisabled              = 4034
	AlarmNotExist                = 4035
	UnmarshalFailed              = 6000
	GeneratePasswordFailed       = 6001
	RegisterDevicesFailed        = 6002
	QueryStatusFailed            = 6003
	QuerySettingFailed           = 6004

	ErrorFindUser               = 5001
	ErrorGetDomainRecords       = 5002
	ErrorParseParams            = 5003
	ErrorFindFunctions          = 5007
	ErrorFindChildOrganizations = 5008
	ErrorCreateOrganization     = 5009
	ErrorCreateRole             = 5010
	ErrorCreateUser             = 5011
	ErrorUpdateRole             = 5012
	ErrorDeleteRole             = 5013
	ErrorFindRole               = 5014
	ErrorFindOrganization       = 5015
	ErrorSendCaptcha            = 5016
	ErrorSetRedis               = 5017
	ErrorGetRedis               = 5018
	ErrorDeleteUser             = 5019
	ErrorUpdateUser             = 5020
	ErrorUpdateUserPassword     = 5021
	ErrorDeleteOrganization     = 5022
	ErrorUpdateOrganization     = 5023
	ErrorGenerateJwtToken       = 5024
	ErrorGeneratePassword       = 5025
	ErrorUpdateDomainRecord     = 5026
	ErrorCheckDomainRecord      = 5027
	ErrorFindDevice             = 5028
	ErrorGetDeviceData          = 5031
	ErrorPublishMsg             = 5032
	ErrorSubscribe              = 5033
	ErrorSetDomainRecordStatus  = 5034
	ErrorQueryHistory           = 5035
	ErrorCheckDeviceName        = 5037
	ErrorRegisterDevice         = 5038
	ErrorGetAllDevices          = 5039
	ErrorUploadFile             = 5040
	ErrorOpenFile               = 5041
	ErrorCreateFile             = 5042
	ErrorWriteFile              = 5043
	ErrorReadFile               = 5044
	ErrorGetDeviceState         = 5045
	ErrorCreateResource         = 5047
	ErrorFindResource           = 5048
	ErrorGetDeviceSecret        = 5049
	ErrorCreateXLSX             = 5050
	ErrorCreateDevice           = 5051
	ErrorEnableThing            = 5052
	ErrorDisableThing           = 5053
	ErrorQueryDevice            = 5054
	ErrorInitDevice             = 5055
	ErrorDeleteResource         = 5056
	ErrorUpdateUserStatus       = 5057
	ErrorUpdateUserRole         = 5058
	ErrorUpdateMobile           = 5059
	ErrorFindDeviceDirectors    = 5060
	ErrorFindAlarms             = 5061
	ErrorUpdateAlarm            = 5062
	ErrorUpdateDevice           = 5063
	ErrorDeleteDirector         = 5064
	ErrorDeleteSMSRecord        = 5065
	ErrorSetParams              = 5000
)

var msg = map[int]map[string]string{
	Success: {"message": "成功", "name": "Success"},

	UserNotExist:                 {"message": "不存在该用户", "name": "UserNotExist"},
	UserNamePasswordNotMatch:     {"message": "账号或密码错误！", "name": "UserNamePasswordNotMatch"},
	RoleExistAssociationUser:     {"message": "该角色存在关联用户", "name": "RoleExistAssociationUser"},
	RootDeleteProhibited:         {"message": "禁止删除管理员账户", "name": "RootDeleteProhibited"},
	InvalidParams:                {"message": "非法参数", "name": "InvalidParams"},
	InvalidMobile:                {"message": "手机号码错误", "name": "InvalidMobile"},
	InvalidCaptcha:               {"message": "验证码错误或失效", "name": "InvalidCaptcha"},
	AdminDeleteProhibited:        {"message": "禁止删除管理员用户", "name": "AdminDeleteProhibited"},
	RoleNotExist:                 {"message": "不存在该角色", "name": "RoleNotExist"},
	OrganizationDeleteProhibited: {"message": "无法删除该组织", "name": "OrganizationDeleteProhibited"},
	MobileInUsed:                 {"message": "手机号已被使用", "name": "MobileInUsed"},
	UserNameInUsed:               {"message": "用户名已被使用", "name": "UserNameInUsed"},
	InvalidUser:                  {"message": "无效的用户", "name": "InvalidUser"},
	RequestTimeOut:               {"message": "请求超时", "name": "RequestTimeOut"},
	OrganizationNotExist:         {"message": "组织不存在", "name": "OrganizationNotExist"},
	OrganizationNameInUsed:       {"message": "组织名已被使用", "name": "OrganizationNameInUsed"},
	DevicesNotExist:              {"message": "不存在设备", "name": "DevicesNotExist"},
	DeviceOffline:                {"message": "设备离线或不存在", "name": "DeviceOffline"},
	DomainRecordNotExist:         {"message": "不存在解析记录", "name": "DomainRecordNotExist"},
	UnknownBindType:              {"message": "未知的绑定类型", "name": "UnknownBindType"},
	UserPermissionsUndefined:     {"message": "未定义用户的权限", "name": "UserPermissionsUndefined"},
	UniqueViolation:              {"message": "存在重复记录", "name": "UniqueViolation"},
	UnSupportType:                {"message": "不被支持的类型", "name": "UnSupportType"},
	AlarmNotExist:                {"message": "暂无警报", "name": "AlarmNotExist"},
	UnmarshalFailed:              {"message": "解码失败", "name": "UnmarshalFailed"},
	GeneratePasswordFailed:       {"message": "生成密钥失败", "name": "GeneratePasswordFailed"},
	RegisterDevicesFailed:        {"message": "批量注册设备失败", "name": "RegisterDevicesFailed"},
	QueryStatusFailed:            {"message": "查询设备状态失败", "name": "QueryStatusFailed"},
	QuerySettingFailed:           {"message": "查询设备配置失败", "name": "QuerySettingFailed"},

	ErrorFindUser:               {"message": "查找用户失败", "name": "ErrorFindUser"},
	ErrorGetDomainRecords:       {"message": "查找域名记录失败", "name": "ErrorGetDomainRecords"},
	ErrorParseParams:            {"message": "参数解析失败", "name": "ErrorParseParams"},
	ErrorFindFunctions:          {"message": "查找用户功能权限失败", "name": "ErrorFindFunctions"},
	ErrorFindChildOrganizations: {"message": "查找下级组织失败", "name": "ErrorFindChildOrganizations"},
	ErrorCreateOrganization:     {"message": "创建组织失败", "name": "ErrorCreateOrganization"},
	ErrorCreateRole:             {"message": "创建角色失败", "name": "ErrorCreateRole"},
	ErrorCreateUser:             {"message": "创建用户失败", "name": "ErrorCreateUser"},
	ErrorUpdateRole:             {"message": "更新角色失败", "name": "ErrorUpdateRole"},
	ErrorDeleteRole:             {"message": "删除角色失败", "name": "ErrorDeleteRole"},
	ErrorFindRole:               {"message": "查找角色失败", "name": "ErrorFindRole"},
	ErrorFindOrganization:       {"message": "查找组织失败", "name": "ErrorFindOrganization"},
	ErrorSendCaptcha:            {"message": "发送验证码失败", "name": "ErrorSendCaptcha"},
	ErrorSetRedis:               {"message": "设置redis时发生错误", "name": "ErrorSetRedis"},
	ErrorGetRedis:               {"message": "查询redis时发生错误", "name": "ErrorGetRedis"},
	ErrorDeleteUser:             {"message": "删除用户失败", "name": "ErrorDeleteUser"},
	ErrorUpdateUser:             {"message": "更新用户失败", "name": "ErrorUpdateUser"},
	ErrorUpdateUserPassword:     {"message": "更新密码失败", "name": "ErrorUpdateUserPassword"},
	ErrorDeleteOrganization:     {"message": "删除组织失败", "name": "ErrorDeleteOrganization"},
	ErrorUpdateOrganization:     {"message": "更新组织失败", "name": "ErrorUpdateOrganization"},
	ErrorGenerateJwtToken:       {"message": "创建JWT Token失败", "name": "ErrorGenerateJwtToken"},
	ErrorGeneratePassword:       {"message": "密码加密失败", "name": "ErrorGeneratePassword"},
	ErrorUpdateDomainRecord:     {"message": "更新域名解析记录失败", "name": "ErrorUpdateDomainRecord"},
	ErrorCheckDomainRecord:      {"message": "检查域名解析失败", "name": "ErrorCheckDomainRecord"},
	ErrorFindDevice:             {"message": "查找设备失败", "name": "ErrorFindDevice"},
	ErrorGetDeviceData:          {"message": "获取设备信息失败", "name": "ErrorGetDeviceData"},
	ErrorPublishMsg:             {"message": "发布信息失败", "name": "ErrorPublishMsg"},
	ErrorSubscribe:              {"message": "订阅主题失败", "name": "ErrorSubscribe"},
	ErrorSetDomainRecordStatus:  {"message": "设置解析记录状态失败", "name": "ErrorSetDomainRecordStatus"},
	ErrorQueryHistory:           {"message": "查询历史记录失败", "name": "ErrorQueryHistory"},
	ErrorCheckDeviceName:        {"message": "注册设备名称失败", "name": "ErrorCheckDeviceName"},
	ErrorRegisterDevice:         {"message": "批量注册设备失败", "name": "ErrorRegistDevice"},
	ErrorGetAllDevices:          {"message": "查询全部设备失败", "name": "ErrorGetAllDevices"},
	ErrorUploadFile:             {"message": "上传文件失败", "name": "ErrorUploadFile"},
	ErrorOpenFile:               {"message": "打开文件失败", "name": "ErrorOpenFile"},
	ErrorCreateFile:             {"message": "创建文件失败", "name": "ErrorCreateFile"},
	ErrorWriteFile:              {"message": "写入文件失败", "name": "ErrorWriteFile"},
	ErrorReadFile:               {"message": "读取文件失败", "name": "ErrorReadFile"},
	ErrorGetDeviceState:         {"message": "获取设备状态失败", "name": "ErrorGetDeviceState"},
	ErrorCreateResource:         {"message": "创建资源失败", "name": "ErrorCreateResource"},
	ErrorFindResource:           {"message": "查找资源失败", "name": "ErrorFindResource"},
	ErrorGetDeviceSecret:        {"message": "获取设备密钥失败", "name": "ErrorGetDeviceSecret"},
	ErrorCreateXLSX:             {"message": "创建表格失败", "name": "ErrorCreateXLSX"},
	ErrorCreateDevice:           {"message": "创建设备失败", "name": "ErrorCreateDevice"},
	ErrorEnableThing:            {"message": "启用设备失败", "name": "ErrorEnableThing"},
	ErrorDisableThing:           {"message": "禁用设备失败", "name": "ErrorDisableThing"},
	ErrorQueryDevice:            {"message": "查询设备失败", "name": "ErrorQueryDevice"},
	ErrorInitDevice:             {"message": "初始化设备失败", "name": "ErrorInitDevice"},
	ErrorDeleteResource:         {"message": "解除用户资源绑定失败", "name": "ErrorDeleteResource"},
	ErrorUpdateUserStatus:       {"message": "更新用户状态失败", "name": "ErrorUpdateUserStatus"},
	ErrorUpdateUserRole:         {"message": "更新用户角色失败", "name": "ErrorUpdateUserRole"},
	ErrorUpdateMobile:           {"message": "解绑手机失败", "name": "ErrorUpdateMobile"},
	ErrorFindDeviceDirectors:    {"message": "查询设备主管失败", "name": "ErrorFindDeviceDirectors"},
	ErrorFindAlarms:             {"message": "查询警报记录失败", "name": "ErrorFindAlarms"},
	ErrorUpdateAlarm:            {"message": "更新警报记录失败", "name": "ErrorUpdateAlarm"},
	ErrorUpdateDevice:           {"message": "更新设备信息失败", "name": "ErrorUpdateDevice"},
	ErrorDeleteDirector:         {"message": "删除主管失败", "name": "ErrorDeleteDirector"},
	ErrorDeleteSMSRecord:        {"message": "删除发送短信记录失败", "name": "ErrorDeleteSMSRecord"},
	ErrorSetParams:              {"message": "设置或遥控设备失败", "name": "ErrorSetParams"},
}

func Response(e echo.Context, httpCode, statusCode int, data interface{}) error {
	return e.JSON(httpCode, echo.Map{
		"code":    statusCode,
		"name":    msg[statusCode]["name"],
		"message": msg[statusCode]["message"],
		"data":    data,
	})
}

func Error(statusCode int) error {
	return errors.New(msg[statusCode]["message"])
}
