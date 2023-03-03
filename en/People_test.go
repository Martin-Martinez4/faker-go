package english

import (
	"fmt"
	"testing"
)

func TestMaleName(t *testing.T) {
	var eng = NewEnglishFaker()

	var testcases = []struct {
		seed     int64
		expected string
	}{
		{seed: 1, expected: "Everett"},
		{seed: 2, expected: "Aryan"},
		{seed: 3, expected: "Henry"},
		{seed: 99, expected: "Julien"},
	}

	for _, tc := range testcases {

		got := eng.MaleName(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v expected: %v", got, tc.expected)

			t.Error(StyleTerminalOuput("\n"+errormessage, ANSItextcolors.red))

		}

	}

}
func TestFemaleName(t *testing.T) {

	var eng = NewEnglishFaker()

	var testcases = []struct {
		seed     int64
		expected string
	}{
		{seed: 1, expected: "Kathryn"},
		{seed: 2, expected: "Saundra"},
		{seed: 3, expected: "Margaret"},
		{seed: 99, expected: "Daphne"},
	}

	for _, tc := range testcases {

		got := eng.FemaleName(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v expected: %v", got, tc.expected)

			t.Error(StyleTerminalOuput("\n"+errormessage, ANSItextcolors.red))

		}

	}

}
func TestLastName(t *testing.T) {

	var eng = NewEnglishFaker()

	var testcases = []struct {
		seed     int64
		expected string
	}{
		{seed: 1, expected: "Brooks"},
		{seed: 2, expected: "Blankenship"},
		{seed: 3, expected: "Rodriguez"},
		{seed: 99, expected: "Anthony"},
	}

	for _, tc := range testcases {

		got := eng.LastName(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v expected: %v", got, tc.expected)

			t.Error(StyleTerminalOuput("\n"+errormessage, ANSItextcolors.red))

		}

	}

}
