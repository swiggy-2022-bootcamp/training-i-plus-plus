package utils

import "time"

func ISODateToTime(date string) (time.Time, error) {
	return time.Parse("2006-01-02T15:04:05.000Z", date)
}
