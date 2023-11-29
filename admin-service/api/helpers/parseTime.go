package helpers

import "time"

func ParseTime(dateStr string) (time.Time, error) {
	const layout = "15:04:05"
	return time.Parse(layout, dateStr)
}
