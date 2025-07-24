package utils

import (
	"regexp"
	"strconv"
	"time"
)

func IsNumericString(s string) bool {
	return regexp.MustCompile(`^\d+$`).MatchString(s)
}

func GenerateNumber() int {
	millis := time.Now().UnixNano() / int64(time.Millisecond)
	millisStr := strconv.FormatInt(millis, 10)
	number, _ := strconv.Atoi(millisStr[len(millisStr)-6:])
	return number
}
