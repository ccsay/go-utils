package stringutil

import (
	"strings"
)

// 去掉首尾空格
func Trim(value string) string {
	return strings.Trim(value, " ")
}

// 判断是否为空
func IsEmpty(value string) bool {
	return len(Trim(value)) == 0
}

// 判断指针字符串为空
func IsPointerEmpty(s *string) bool {
	return s == nil || len(Trim(*s)) == 0
}

// 判断指针字符串不为空
func IsNotPointerEmpty(s *string) bool {
	return s != nil && len(Trim(*s)) > 0
}
