package main

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

	if len(format) > 0 {

		after := now.Add(time.Hour * 1)
		return after.Format(string(format[0]))

	} else {

		return now.Add(time.Hour * 1).String()
	}
}
func (d *Dates) Before(format ...Dateformat) string {

	now := time.Now()

	if len(format) > 0 {

		before := now.Add(time.Hour * 1)
		return before.Format(string(format[0]))

	} else {

		return now.Add(-time.Hour * 1).String()
	}

}
