package english

import (
	"fmt"
	"testing"
)

func TestLocation(t *testing.T) {

	eng := NewEnglishFaker()

	citieslen := len(*eng.cities)
	stateslen := len(*eng.states)

	testcases := []struct {
		seed     int64
		expected string
	}{

		{
			seed:     1,
			expected: fmt.Sprintf("%s, %s", (*eng.cities)[GenerateRandIntBelowMaximum(citieslen, 1)], (*eng.states)[GenerateRandIntBelowMaximum(stateslen, 1)]),
		},
		{
			seed:     10,
			expected: fmt.Sprintf("%s, %s", (*eng.cities)[GenerateRandIntBelowMaximum(citieslen, 10)], (*eng.states)[GenerateRandIntBelowMaximum(stateslen, 10)]),
		},
		{
			seed:     251,
			expected: fmt.Sprintf("%s, %s", (*eng.cities)[GenerateRandIntBelowMaximum(citieslen, 251)], (*eng.states)[GenerateRandIntBelowMaximum(stateslen, 251)]),
		},
		{
			seed:     456,
			expected: fmt.Sprintf("%s, %s", (*eng.cities)[GenerateRandIntBelowMaximum(citieslen, 456)], (*eng.states)[GenerateRandIntBelowMaximum(stateslen, 456)]),
		},
		{
			seed:     40444,
			expected: fmt.Sprintf("%s, %s", (*eng.cities)[GenerateRandIntBelowMaximum(citieslen, 40444)], (*eng.states)[GenerateRandIntBelowMaximum(stateslen, 40444)]),
		},
	}

	for testnumber, tc := range testcases {

		got := eng.Location(tc.seed)
		if got != tc.expected {
			errormessage := fmt.Sprintf("Test Number %d failed.  got %v , expected %v", testnumber, got, tc.expected)

			t.Error("\n" + StyleTerminalOuput(errormessage, ANSItextcolors.red))
		}

	}

}
