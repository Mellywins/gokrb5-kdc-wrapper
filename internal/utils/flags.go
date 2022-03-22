package utils

import "strings"

func SanitizeFlagValue(val string) string {
	if strings.Contains(val, " ") {
		quotedStr := "\"" + val + "\""
		return quotedStr
	}
	return val
}
func AppendGenericFlag(flag string, val string) string {
	return " " + flag + " " + SanitizeFlagValue(val)
}
func AppendValuelessFlag(flag string) string {
	return " " + flag
}
