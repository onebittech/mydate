package mydate

import (
	"errors"
	"regexp"
	"time"
)

// Datetime is a date and time in MySQL format
type Datetime string

const datetimeLayout = "2006-01-02 15:04:05.000000"

var errInvalidDate = errors.New("invalid date format")

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

	dateRegex := `\d\d\d\d-\d\d-\d\d`
	timeRegex := dateRegex + " " + `\d\d:\d\d`
	timestampRegex := timeRegex + `:\d\d`
	timestampMilliRegex := timestampRegex + `.\d\d\d`
	timestampMicroRegex := timestampMilliRegex + `\d\d\d`

	var t time.Time
	var err error

	if re := regexp.MustCompile(timestampMicroRegex); re.MatchString(date) {
		t, err = time.ParseInLocation(datetimeLayout, date, Location)
	} else if re := regexp.MustCompile(timestampMilliRegex); re.MatchString(date) {
		t, err = time.ParseInLocation(datetimeLayout, date+"000", Location)
	} else if re := regexp.MustCompile(timestampRegex); re.MatchString(date) {
		t, err = time.ParseInLocation(datetimeLayout, date+".000000", Location)
	} else if re := regexp.MustCompile(timeRegex); re.MatchString(date) {
		t, err = time.ParseInLocation(datetimeLayout, date+":00.000000", Location)
	} else if re := regexp.MustCompile(dateRegex); re.MatchString(date) {
		t, err = Date(date).Time()
	} else {
		err = errInvalidDate
	}

	return t, err
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
