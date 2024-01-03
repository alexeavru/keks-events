package common

import "time"

func ParseTime(value string) time.Time {
	t, _ := time.Parse(time.RFC3339, value)
	return t
}
