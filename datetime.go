package mydate

import (
	"strings"
	"time"
)

// Datetime is a date and time in MySQL format
type Datetime string

const datetimeLayout = "2006-01-02 15:04:05.000000"

// DatetimeFormatLayout for use in Time.Format
var DatetimeFormatLayout = "02/01/2006 15:04:05"

// NewDatetime creates a new Datetime from a time.Time object
func NewDatetime(t time.Time) Datetime {
	return Datetime(t.Format(datetimeLayout))
}

// Now returns the current time
func Now() Datetime {
	t := time.Now()
	return Datetime(t.Format(datetimeLayout))
}

// Time generates a time.Time object from a Datetime
func (d Datetime) Time() (time.Time, error) {
	date := string(d)

	if date == "" {
		var t time.Time
		return t, nil
	}

	if !strings.Contains(date, ".") {
		date += ".000000"
	}

	return time.ParseInLocation(datetimeLayout, date, Location)
}

// Format returns a textual representation of the time value formatted
// according to layout, which defines the format by showing how the reference
// time, defined to be
//	Mon Jan 2 15:04:05 -0700 MST 2006
// would be displayed if it were the value; it serves as an example of the
// desired output. The same display rules will then be applied to the time
// value.
func (d Datetime) Format() string {
	date := string(d)

	if date == "" {
		return ""
	}

	t, err := d.Time()
	if err != nil {
		return ""
	}

	return t.Format(DatetimeFormatLayout)
}
