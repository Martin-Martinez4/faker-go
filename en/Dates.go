package english

import (
	"time"
)

type Dates struct {
}

var dates = Dates{}

func (d *Dates) Now(format ...Dateformat) string {

	if len(format) > 0 {

		return time.Now().Format(string(format[0]))
	} else {
		return time.Now().String()
	}

}

func (d *Dates) After(format ...Dateformat) string {

	now := time.Now()

	hoursInAYear := 8761

	hoursToAdd := GenerateRandIntBelowMaximum(hoursInAYear)

	if len(format) > 0 {

		after := now.Add(time.Hour * time.Duration(hoursToAdd))
		return after.Format(string(format[0]))

	} else {

		return now.Add(time.Hour * time.Duration(hoursToAdd)).String()
	}
}
func (d *Dates) Before(format ...Dateformat) string {

	now := time.Now()

	hoursInAYear := 8761

	hoursToSubtract := GenerateRandIntBelowMaximum(hoursInAYear)

	if len(format) > 0 {

		before := now.Add(time.Hour * time.Duration(-hoursToSubtract))
		return before.Format(string(format[0]))

	} else {

		return now.Add((time.Hour) * time.Duration(-hoursToSubtract)).String()
	}

}
