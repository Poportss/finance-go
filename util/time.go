package util

import "time"

const layout = "2006-01-02T15:04:05"

func StringToTime(value string) time.Time {
	convertTime, _ := time.Parse(layout, value)

	return convertTime
}
