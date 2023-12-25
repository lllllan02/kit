package kit

import (
	"strings"
	"unicode"
)

// CamelCase 将字符串转换为小驼峰格式。
func CamelCase(s string) string {
	var builder strings.Builder

	strs := splitToStrings(s, false)
	for i, str := range strs {
		if i == 0 {
			builder.WriteString(strings.ToLower(str))
		} else {
			builder.WriteString(Capitalize(str))
		}
	}
	return builder.String()
}

// Capitalize 保证字符串的首字母大写，其他字母小写。
func Capitalize(s string) string {
	result := make([]rune, len(s))
	for i, r := range s {
		if i == 0 {
			result[i] = unicode.ToUpper(r)
		} else {
			result[i] = unicode.ToLower(r)
		}
	}
	return string(result)
}

// SnakeCase 将字符串转换为下划线格式。
func SnakeCase(s string) string {
	result := splitToStrings(s, false)
	return strings.Join(result, "_")
}

// UpperSnakeCase 将字符串转换为大写的下划线格式。
func UpperSnakeCase(s string) string {
	result := splitToStrings(s, true)
	return strings.Join(result, "_")
}

// ReverseString 翻转字符串
func ReverseString(s string) string {
	runes := []rune(s)
	return string(Reverse(runes))
}
