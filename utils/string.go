package utils

import "strings"

func ToCombinedString(src []string, sepChar string) string {
	if len(src) < 1 {
		return ""
	}

	var res string = ""

	for i, v := range src {
		res += v
		if i < len(src)-1 && i > 0 {
			res += sepChar
		}
	}

	return res
}

func ToSliceString(src, sepChar string) []string {
	return strings.Split(src, sepChar)
}

func ToNormalizedString(s string) string {
	s = strings.TrimSpace(s)
	return strings.ToLower(s)
}
