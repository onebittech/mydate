package mydate

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDatetime(t *testing.T) {
	_, err := Now().Time()
	assert.Nil(t, err, "Time now should not give an error")

	dateStringError := Datetime("ANBCWW")
	_, err = dateStringError.Time()
	assert.Error(t, err, "Time should give an error for invalid string")
	emptyFormatError := dateStringError.Format()
	assert.Equal(t, "", emptyFormatError, "Format should give empty string for invalid string")

	dateStringEmpty := Datetime("")
	_, err = dateStringEmpty.Time()
	assert.Nil(t, err, "Time empty string should not give an error")

	formattedEmptyString := dateStringEmpty.Format()
	assert.Equal(t, "", formattedEmptyString, "Empty string formatted should be empty string")

	local, _ := time.LoadLocation("Local")
	timeDate := time.Date(2020, 1, 1, 0, 0, 0, 0, local)
	timeDateString := NewDatetime(timeDate)
	assert.Equal(t, "2020-01-01 00:00:00.000000", string(timeDateString), "Dates should be equal")

	timeDateStringFormat := timeDateString.Format()
	assert.Equal(t, "01/01/2020 00:00:00", timeDateStringFormat, "Formatted dates should be equal")

	dateStringOK := Datetime("2020-01-01 00:00:00")
	dateStringOKTime, err := dateStringOK.Time()
	assert.Nil(t, err, "Time date string should not give an error")
	assert.True(t, dateStringOKTime.Equal(timeDate), "Time should be equal")
}

func TestDate(t *testing.T) {
	_, err := CurrentDate().Time()
	assert.Nil(t, err, "Time now should not give an error")

	dateStringError := Date("ANBCWW")
	_, err = dateStringError.Time()
	assert.Error(t, err, "Time should give an error for invalid string")
	emptyFormatError := dateStringError.Format()
	assert.Equal(t, "", emptyFormatError, "Format should give empty string for invalid string")

	dateStringEmpty := Date("")
	_, err = dateStringEmpty.Time()
	assert.Nil(t, err, "Time empty string should not give an error")

	formattedEmptyString := dateStringEmpty.Format()
	assert.Equal(t, "", formattedEmptyString, "Empty string formatted should be empty string")

	local, _ := time.LoadLocation("Local")
	timeDate := time.Date(2020, 1, 1, 0, 0, 0, 0, local)
	dateStringOK := Date("2020-01-01")
	dateStringOKTime, err := dateStringOK.Time()
	assert.Nil(t, err, "Time date string should not give an error")
	assert.True(t, dateStringOKTime.Equal(timeDate), "Time should be equal")
	assert.Equal(t, "01/01/2020", dateStringOK.Format(), "Formatted dates should be equal")
}

func TestTime(t *testing.T) {
	CurrentTime()
}
