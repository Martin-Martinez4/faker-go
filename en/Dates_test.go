package english

import (
	"fmt"
	"testing"
	"time"
)

func TestNow(t *testing.T) {

	var eng = NewEnglishFaker()

	var testcases = []struct {
		dateformat string
	}{
		{dateformat: string(Dateformats.ANSIC)},
		{dateformat: string(Dateformats.Kitchen)},
		{dateformat: string(Dateformats.mmddyy)},
		{dateformat: string(Dateformats.mmmddyyyy)},
		{dateformat: string(Dateformats.mmddyyyy)},
		{dateformat: string(Dateformats.mmmddyyyy)},
		{dateformat: string(Dateformats.RFC1123)},
		{dateformat: string(Dateformats.RFC1123Z)},
		{dateformat: string(Dateformats.RFC3339)},
		{dateformat: string(Dateformats.RFC3339Nano)},
		{dateformat: string(Dateformats.RFC822)},
		{dateformat: string(Dateformats.RFC822Z)},
		{dateformat: string(Dateformats.RFC850)},
		{dateformat: string(Dateformats.RubyDate)},
		{dateformat: string(Dateformats.UnixDate)},
	}

	for testnumber, tc := range testcases {

		now := time.Now().Format(string(tc.dateformat))
		got := eng.Now(Dateformat(tc.dateformat))
		if got != now {

			errormessage := fmt.Sprintf("Test Number %d failed.  got %v , expected %v", testnumber, got, now)

			t.Error("\n" + StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}

func TestAfter(t *testing.T) {

	var eng = NewEnglishFaker()

	var testcases = []struct {
		dateformat string
	}{
		{dateformat: string(Dateformats.ANSIC)},
		{dateformat: string(Dateformats.mmddyy)},
		{dateformat: string(Dateformats.mmmddyyyy)},
		{dateformat: string(Dateformats.mmddyyyy)},
		{dateformat: string(Dateformats.mmmddyyyy)},
		{dateformat: string(Dateformats.RFC1123)},
		{dateformat: string(Dateformats.RFC1123Z)},
		{dateformat: string(Dateformats.RFC3339)},
		{dateformat: string(Dateformats.RFC3339Nano)},
		{dateformat: string(Dateformats.RFC822)},
		{dateformat: string(Dateformats.RFC822Z)},
		{dateformat: string(Dateformats.RFC850)},
		{dateformat: string(Dateformats.RubyDate)},
		{dateformat: string(Dateformats.UnixDate)},
	}

	for testnumber, tc := range testcases {

		now := time.Now()
		got := eng.After(Dateformat(tc.dateformat))
		gottime, error := time.Parse(tc.dateformat, got)
		if error != nil {

			errormessage := fmt.Sprintf("Test Number %d failed.  Error parsing time value, expected %v", testnumber, got)

			t.Error("\n" + StyleTerminalOuput(errormessage, ANSItextcolors.red))
		}
		if !now.Before(gottime) {

			errormessage := fmt.Sprintf("Test Number %d failed.  got %v , expected %v", testnumber, got, now)

			t.Error("\n" + StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}

func TestBefore(t *testing.T) {

	var eng = NewEnglishFaker()

	var testcases = []struct {
		dateformat string
	}{
		{dateformat: string(Dateformats.ANSIC)},
		{dateformat: string(Dateformats.mmddyy)},
		{dateformat: string(Dateformats.mmmddyyyy)},
		{dateformat: string(Dateformats.mmddyyyy)},
		{dateformat: string(Dateformats.mmmddyyyy)},
		{dateformat: string(Dateformats.RFC1123)},
		{dateformat: string(Dateformats.RFC1123Z)},
		{dateformat: string(Dateformats.RFC3339)},
		{dateformat: string(Dateformats.RFC3339Nano)},
		{dateformat: string(Dateformats.RFC822)},
		{dateformat: string(Dateformats.RFC822Z)},
		{dateformat: string(Dateformats.RFC850)},
		{dateformat: string(Dateformats.RubyDate)},
		{dateformat: string(Dateformats.UnixDate)},
	}

	for testnumber, tc := range testcases {

		now := time.Now()
		got := eng.Before(Dateformat(tc.dateformat))
		gottime, error := time.Parse(tc.dateformat, got)
		if error != nil {

			errormessage := fmt.Sprintf("Test Number %d failed.  Error parsing time value, expected %v", testnumber, got)

			t.Error("\n" + StyleTerminalOuput(errormessage, ANSItextcolors.red))
		}
		if now.Before(gottime) {

			errormessage := fmt.Sprintf("Test Number %d failed.  got %v , expected %v", testnumber, got, now)

			t.Error("\n" + StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
