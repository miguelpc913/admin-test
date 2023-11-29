package helpers

import "time"

func ParseDateTime(dateStr string) (time.Time, error) {
	const layout = "2006-01-02 15:04:00"
	return time.Parse(layout, dateStr)
}
