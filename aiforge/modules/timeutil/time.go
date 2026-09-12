package timeutil

import "time"

const DATE_FORMAT = "2006-01-02"

func GetMidNight(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func GetDateString(t time.Time) string {
	return t.Format(DATE_FORMAT)
}
