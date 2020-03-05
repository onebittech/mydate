package mydate

import "time"

// Date is a date in MySQL format
type Date string

const dateLayout = "2006-01-02"

// DateFormatLayout for use in Time.Format
var DateFormatLayout = "02/01/2006"

// Location by default for generating time.Time objects
var Location, _ = time.LoadLocation("Local")

// CurrentDate returns the current date
func CurrentDate() Date {
	t := time.Now()
	return Date(t.Format(dateLayout))
}

// Time generates a time.Time object from a Date
func (d Date) Time() (time.Time, error) {
	date := string(d)

	if date == "" {
		var t time.Time
		return t, nil
	}

	return time.ParseInLocation(dateLayout, date, Location)
}

// Format returns a textual representation of the time value formatted
// according to layout, which defines the format by showing how the reference
// time, defined to be
//	Mon Jan 2 15:04:05 -0700 MST 2006
// would be displayed if it were the value; it serves as an example of the
// desired output. The same display rules will then be applied to the time
// value.
func (d Date) Format() string {
	date := string(d)

	if date == "" {
		return ""
	}

	t, err := d.Time()
	if err != nil {
		return ""
	}

	return t.Format(DateFormatLayout)
}