package utils

import "regexp"

func IsNumericString(s string) bool {
	return regexp.MustCompile(`^\d+$`).MatchString(s)
}
