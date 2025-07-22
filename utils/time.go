package utils

import "time"

const (
	NormalActionDuration time.Duration = time.Minute * 15   // 15'
	AccessDuration       time.Duration = time.Hour * 24     // 1 ngày
	RefreshDuration      time.Duration = AccessDuration * 7 // 1 tuần
)

func GetPrimitiveTime() time.Time {
	// 1/1/1900 - 00:00:00
	return time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)
}

func IsActionExpired(exp time.Time) bool {
	return time.Now().After(exp)
}
