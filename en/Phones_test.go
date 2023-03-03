package english

import (
	"fmt"
	"testing"
)

func TestBuildPhoneNumberDomestic(t *testing.T) {

	var testcases = []struct {
		format   string
		seed     int64
		expected string
	}{
		{format: Phoneformats.Domestic, seed: 1, expected: "644-644-6441"},
		{format: Phoneformats.Domestic, seed: 2, expected: "250-250-2505"},
		{format: Phoneformats.Domestic, seed: 10, expected: "609-609-6094"},
		{format: Phoneformats.Domestic, seed: 999, expected: "213-213-2132"},
	}

	for _, tc := range testcases {

		got := BuildPhoneNumber(tc.format, tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v expected: %v", got, tc.expected)

			t.Error(StyleTerminalOuput("\n"+errormessage, ANSItextcolors.red))

		}

	}

}
func TestBuildPhoneNumberLocal(t *testing.T) {

	var testcases = []struct {
		format   string
		seed     int64
		expected string
	}{
		{format: Phoneformats.Local, seed: 1, expected: "644-6441"},
		{format: Phoneformats.Local, seed: 2, expected: "250-2505"},
		{format: Phoneformats.Local, seed: 10, expected: "609-6094"},
		{format: Phoneformats.Local, seed: 999, expected: "213-2132"},
	}

	for _, tc := range testcases {

		got := BuildPhoneNumber(tc.format, tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v expected: %v", got, tc.expected)

			t.Error(StyleTerminalOuput("\n"+errormessage, ANSItextcolors.red))

		}

	}

}
func TestBuildPhoneNumberDomesticParthenses(t *testing.T) {

	var testcases = []struct {
		format   string
		seed     int64
		expected string
	}{
		{format: Phoneformats.DomesticParthenses, seed: 1, expected: "(644)-644-6441"},
		{format: Phoneformats.DomesticParthenses, seed: 2, expected: "(250)-250-2505"},
		{format: Phoneformats.DomesticParthenses, seed: 10, expected: "(609)-609-6094"},
		{format: Phoneformats.DomesticParthenses, seed: 999, expected: "(213)-213-2132"},
	}

	for _, tc := range testcases {

		got := BuildPhoneNumber(tc.format, tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v expected: %v", got, tc.expected)

			t.Error(StyleTerminalOuput("\n"+errormessage, ANSItextcolors.red))

		}

	}

}
func TestBuildPhoneNumberDomesticParthensesNoHypen(t *testing.T) {

	var testcases = []struct {
		format   string
		seed     int64
		expected string
	}{
		{format: Phoneformats.DomesticParthensesNoHypen, seed: 1, expected: "(644) 644-6441"},
		{format: Phoneformats.DomesticParthensesNoHypen, seed: 2, expected: "(250) 250-2505"},
		{format: Phoneformats.DomesticParthensesNoHypen, seed: 10, expected: "(609) 609-6094"},
		{format: Phoneformats.DomesticParthensesNoHypen, seed: 999, expected: "(213) 213-2132"},
	}

	for _, tc := range testcases {

		got := BuildPhoneNumber(tc.format, tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v expected: %v", got, tc.expected)

			t.Error(StyleTerminalOuput("\n"+errormessage, ANSItextcolors.red))

		}

	}

}
func TestBuildPhoneNumberLocalNoHypens(t *testing.T) {

	var testcases = []struct {
		format   string
		seed     int64
		expected string
	}{
		{format: Phoneformats.LocalNoHypen, seed: 1, expected: "6441942"},
		{format: Phoneformats.LocalNoHypen, seed: 2, expected: "2505669"},
		{format: Phoneformats.LocalNoHypen, seed: 10, expected: "6094828"},
		{format: Phoneformats.LocalNoHypen, seed: 999, expected: "2132750"},
	}

	for _, tc := range testcases {

		got := BuildPhoneNumber(tc.format, tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v expected: %v", got, tc.expected)

			t.Error(StyleTerminalOuput("\n"+errormessage, ANSItextcolors.red))

		}

	}

}
func TestBuildPhoneNumberInternational(t *testing.T) {

	var testcases = []struct {
		format   string
		seed     int64
		expected string
	}{
		{format: Phoneformats.International, seed: 1, expected: "1-644-6441"},
		{format: Phoneformats.International, seed: 2, expected: "1-250-2505"},
		{format: Phoneformats.International, seed: 10, expected: "1-609-6094"},
		{format: Phoneformats.International, seed: 999, expected: "1-213-2132"},

		{format: Phoneformats.InternationalPlus, seed: 1, expected: "+1-644-6441"},
		{format: Phoneformats.InternationalPlus, seed: 2, expected: "+1-250-2505"},
		{format: Phoneformats.InternationalPlus, seed: 10, expected: "+1-609-6094"},
		{format: Phoneformats.InternationalPlus, seed: 999, expected: "+1-213-2132"},
	}

	for _, tc := range testcases {

		got := BuildPhoneNumber(tc.format, tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v expected: %v", got, tc.expected)

			t.Error(StyleTerminalOuput("\n"+errormessage, ANSItextcolors.red))

		}

	}

}

var result string

func BenchmarkBuildPhoneNumber(b *testing.B) {

	var r string

	for n := 0; n < b.N; n++ {

		r = BuildPhoneNumber(Phoneformats.International)

	}

	result = r

}
