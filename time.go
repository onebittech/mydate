package mydate

import "time"

// Time represents a time in MySQL format
type Time string

const timeLayout = "15:04:05"

// CurrentTime returns the current time in MySQL format
func CurrentTime() Time {
	t := time.Now()
	return Time(t.Format(timeLayout))
}
