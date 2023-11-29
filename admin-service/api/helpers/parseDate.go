package helpers

import "time"

func ParseDate(dateStr string) (time.Time, error) {
	const layout = "2006-01-02"
	return time.Parse(layout, dateStr)
}
