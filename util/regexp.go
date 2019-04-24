package util

import "regexp"

var (
	// 域名，只匹配数字、字母和.
	OnlyDomainRecord = regexp.MustCompile(`^(\w|\.)+$`)

	HasDomainRecord = regexp.MustCompile(`[0-9]+`)

	HasNum = regexp.MustCompile(`[0-9]+`)
)
