package util

import "regexp"

// 匹配符合设备类型+设备序列号的域名解析记录值
var DeviceReg = regexp.MustCompile("^[a-z]([0-9]{5,})$")

// 匹配11位手机
var MobileReg = regexp.MustCompile("^1\\d{10}$")

// 匹配邮箱
var EmailReg = regexp.MustCompile("[a-zA-Z0-9_.-]+@[a-zA-Z0-9-]+(\\.[a-zA-Z0-9-]+)*\\.[a-zA-Z0-9]{2,6}")

// 匹配正整数
var UintReg = regexp.MustCompile("^[1-9]\\d*$")
