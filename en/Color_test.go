package english

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {

	eng := NewEnglishFaker()

	var testcases = []struct {
		seed     int64
		expected string
	}{

		{
			seed:     1,
			expected: (*eng.colornames)[GenerateRandIntBelowMaximum(len(*eng.colornames), 1)],
		},
		{
			seed:     0,
			expected: (*eng.colornames)[GenerateRandIntBelowMaximum(len(*eng.colornames), 0)],
		},
		{
			seed:     2,
			expected: (*eng.colornames)[GenerateRandIntBelowMaximum(len(*eng.colornames), 2)],
		},
		{
			seed:     10,
			expected: (*eng.colornames)[GenerateRandIntBelowMaximum(len(*eng.colornames), 10)],
		},
		{
			seed:     100,
			expected: (*eng.colornames)[GenerateRandIntBelowMaximum(len(*eng.colornames), 100)],
		},
		{
			seed:     500,
			expected: (*eng.colornames)[GenerateRandIntBelowMaximum(len(*eng.colornames), 500)],
		},
	}

	for testnumber, tc := range testcases {

		got := eng.Color(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("Test Number %d failed.  got %v , expected %v", testnumber, got, tc.expected)

			t.Error(StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}

func TestColorhex(t *testing.T) {

	eng := NewEnglishFaker()

	var testcases = []struct {
		seed     int64
		expected string
	}{

		{
			seed:     15,
			expected: fmt.Sprintf("#%x%x%x", GenerateRandIntBelowMaximum(256, 15), GenerateRandIntBelowMaximum(256, 15), GenerateRandIntBelowMaximum(256, 15)),
		},
		{
			seed:     0,
			expected: fmt.Sprintf("#%x%x%x", GenerateRandIntBelowMaximum(256, 0), GenerateRandIntBelowMaximum(256, 0), GenerateRandIntBelowMaximum(256, 0)),
		},
		{
			seed:     2,
			expected: fmt.Sprintf("#%x%x%x", GenerateRandIntBelowMaximum(256, 2), GenerateRandIntBelowMaximum(256, 2), GenerateRandIntBelowMaximum(256, 2)),
		},
		{
			seed:     10,
			expected: fmt.Sprintf("#%x%x%x", GenerateRandIntBelowMaximum(256, 10), GenerateRandIntBelowMaximum(256, 10), GenerateRandIntBelowMaximum(256, 10)),
		},
		{
			seed:     100,
			expected: "#707070",
		},
		{
			seed:     500,
			expected: fmt.Sprintf("#%x%x%x", GenerateRandIntBelowMaximum(256, 500), GenerateRandIntBelowMaximum(256, 500), GenerateRandIntBelowMaximum(256, 500)),
		},
	}

	for testnumber, tc := range testcases {

		got := eng.Colorhex(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("Test Number %d failed.  got %v , expected %v", testnumber, got, tc.expected)

			t.Error("\n" + StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}

func TestColorrgb(t *testing.T) {

	eng := NewEnglishFaker()

	var testcases = []struct {
		seed     int64
		expected string
	}{

		{
			seed:     65,
			expected: fmt.Sprintf("%s(%d,%d,%d)", eng.ColorSpaces.rgb, GenerateRandIntBelowMaximum(255, 65), GenerateRandIntBelowMaximum(255, 65), GenerateRandIntBelowMaximum(255, 65)),
		},
		{
			seed:     0,
			expected: fmt.Sprintf("%s(%d,%d,%d)", eng.ColorSpaces.rgb, GenerateRandIntBelowMaximum(255, 0), GenerateRandIntBelowMaximum(255, 0), GenerateRandIntBelowMaximum(255, 0)),
		},
		{
			seed:     2,
			expected: fmt.Sprintf("%s(%d,%d,%d)", eng.ColorSpaces.rgb, GenerateRandIntBelowMaximum(255, 2), GenerateRandIntBelowMaximum(255, 2), GenerateRandIntBelowMaximum(255, 2)),
		},
		{
			seed:     10,
			expected: fmt.Sprintf("%s(%d,%d,%d)", eng.ColorSpaces.rgb, GenerateRandIntBelowMaximum(255, 10), GenerateRandIntBelowMaximum(255, 10), GenerateRandIntBelowMaximum(255, 10)),
		},
		{
			seed:     100,
			expected: fmt.Sprintf("%s(%d,%d,%d)", eng.ColorSpaces.rgb, GenerateRandIntBelowMaximum(255, 100), GenerateRandIntBelowMaximum(255, 100), GenerateRandIntBelowMaximum(255, 100)),
		},
		{
			seed:     500,
			expected: fmt.Sprintf("%s(%d,%d,%d)", eng.ColorSpaces.rgb, GenerateRandIntBelowMaximum(255, 500), GenerateRandIntBelowMaximum(255, 500), GenerateRandIntBelowMaximum(255, 500)),
		},
	}

	for testnumber, tc := range testcases {

		got := eng.Colorrgb(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("Test Number %d failed.  got %v , expected %v", testnumber, got, tc.expected)

			t.Error("\n" + StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
func TestColorrgba(t *testing.T) {

	eng := NewEnglishFaker()

	var testcases = []struct {
		seed     int64
		expected string
	}{

		{
			seed:     65,
			expected: fmt.Sprintf("%s(%d,%d,%d,%f)", eng.ColorSpaces.rgba, GenerateRandIntBelowMaximum(255, 65), GenerateRandIntBelowMaximum(255, 65), GenerateRandIntBelowMaximum(255, 65), RandomFloatBetweenTwoNumbers(0, 1, 65)),
		},
		{
			seed:     0,
			expected: fmt.Sprintf("%s(%d,%d,%d,%f)", eng.ColorSpaces.rgba, GenerateRandIntBelowMaximum(255, 0), GenerateRandIntBelowMaximum(255, 0), GenerateRandIntBelowMaximum(255, 0), RandomFloatBetweenTwoNumbers(0, 1, 0)),
		},
		{
			seed:     2,
			expected: fmt.Sprintf("%s(%d,%d,%d,%f)", eng.ColorSpaces.rgba, GenerateRandIntBelowMaximum(255, 2), GenerateRandIntBelowMaximum(255, 2), GenerateRandIntBelowMaximum(255, 2), RandomFloatBetweenTwoNumbers(0, 1, 2)),
		},
		{
			seed:     10,
			expected: fmt.Sprintf("%s(%d,%d,%d,%f)", eng.ColorSpaces.rgba, GenerateRandIntBelowMaximum(255, 10), GenerateRandIntBelowMaximum(255, 10), GenerateRandIntBelowMaximum(255, 10), RandomFloatBetweenTwoNumbers(0, 1, 10)),
		},
		{
			seed:     100,
			expected: fmt.Sprintf("%s(%d,%d,%d,%f)", eng.ColorSpaces.rgba, GenerateRandIntBelowMaximum(255, 100), GenerateRandIntBelowMaximum(255, 100), GenerateRandIntBelowMaximum(255, 100), RandomFloatBetweenTwoNumbers(0, 1, 100)),
		},
		{
			seed:     500,
			expected: fmt.Sprintf("%s(%d,%d,%d,%f)", eng.ColorSpaces.rgba, GenerateRandIntBelowMaximum(255, 500), GenerateRandIntBelowMaximum(255, 500), GenerateRandIntBelowMaximum(255, 500), RandomFloatBetweenTwoNumbers(0, 1, 500)),
		},
	}

	for testnumber, tc := range testcases {

		got := eng.Colorrgba(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("Test Number %d failed.  got %v , expected %v", testnumber, got, tc.expected)

			t.Error("\n" + StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
func TestColorhsl(t *testing.T) {

	eng := NewEnglishFaker()

	var testcases = []struct {
		seed     int64
		expected string
	}{

		{
			seed:     65,
			expected: fmt.Sprintf("%s(%d,%d%%,%d%%)", eng.ColorSpaces.hsl, GenerateRandIntBelowMaximum(359, 65), GenerateRandIntBelowMaximum(101, 65), GenerateRandIntBelowMaximum(101, 65)),
		},
		{
			seed:     0,
			expected: fmt.Sprintf("%s(%d,%d%%,%d%%)", eng.ColorSpaces.hsl, GenerateRandIntBelowMaximum(359, 0), GenerateRandIntBelowMaximum(101, 0), GenerateRandIntBelowMaximum(101, 0)),
		},
		{
			seed:     2,
			expected: fmt.Sprintf("%s(%d,%d%%,%d%%)", eng.ColorSpaces.hsl, GenerateRandIntBelowMaximum(359, 2), GenerateRandIntBelowMaximum(101, 2), GenerateRandIntBelowMaximum(101, 2)),
		},
		{
			seed:     10,
			expected: fmt.Sprintf("%s(%d,%d%%,%d%%)", eng.ColorSpaces.hsl, GenerateRandIntBelowMaximum(359, 10), GenerateRandIntBelowMaximum(101, 10), GenerateRandIntBelowMaximum(101, 10)),
		},
		{
			seed:     100,
			expected: fmt.Sprintf("%s(%d,%d%%,%d%%)", eng.ColorSpaces.hsl, GenerateRandIntBelowMaximum(359, 100), GenerateRandIntBelowMaximum(101, 100), GenerateRandIntBelowMaximum(101, 100)),
		},
		{
			seed:     500,
			expected: fmt.Sprintf("%s(%d,%d%%,%d%%)", eng.ColorSpaces.hsl, GenerateRandIntBelowMaximum(359, 500), GenerateRandIntBelowMaximum(101, 500), GenerateRandIntBelowMaximum(101, 500)),
		},
	}

	for testnumber, tc := range testcases {

		got := eng.Colorhsl(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("Test Number %d failed.  got %v , expected %v", testnumber, got, tc.expected)

			t.Error("\n" + StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
func TestColorhsla(t *testing.T) {

	eng := NewEnglishFaker()

	var testcases = []struct {
		seed     int64
		expected string
	}{

		{
			seed:     65,
			expected: fmt.Sprintf("%s(%d,%d%%,%d%%,%f)", eng.ColorSpaces.hsla, GenerateRandIntBelowMaximum(359, 65), GenerateRandIntBelowMaximum(101, 65), GenerateRandIntBelowMaximum(101, 65), RandomFloatBetweenTwoNumbers(0, 1, 65)),
		},
		{
			seed:     0,
			expected: fmt.Sprintf("%s(%d,%d%%,%d%%,%f)", eng.ColorSpaces.hsla, GenerateRandIntBelowMaximum(359, 0), GenerateRandIntBelowMaximum(101, 0), GenerateRandIntBelowMaximum(101, 0), RandomFloatBetweenTwoNumbers(0, 1, 0)),
		},
		{
			seed:     2,
			expected: fmt.Sprintf("%s(%d,%d%%,%d%%,%f)", eng.ColorSpaces.hsla, GenerateRandIntBelowMaximum(359, 2), GenerateRandIntBelowMaximum(101, 2), GenerateRandIntBelowMaximum(101, 2), RandomFloatBetweenTwoNumbers(0, 1, 2)),
		},
		{
			seed:     10,
			expected: fmt.Sprintf("%s(%d,%d%%,%d%%,%f)", eng.ColorSpaces.hsla, GenerateRandIntBelowMaximum(359, 10), GenerateRandIntBelowMaximum(101, 10), GenerateRandIntBelowMaximum(101, 10), RandomFloatBetweenTwoNumbers(0, 1, 10)),
		},
		{
			seed:     100,
			expected: fmt.Sprintf("%s(%d,%d%%,%d%%,%f)", eng.ColorSpaces.hsla, GenerateRandIntBelowMaximum(359, 100), GenerateRandIntBelowMaximum(101, 100), GenerateRandIntBelowMaximum(101, 100), RandomFloatBetweenTwoNumbers(0, 1, 100)),
		},
		{
			seed:     500,
			expected: fmt.Sprintf("%s(%d,%d%%,%d%%,%f)", eng.ColorSpaces.hsla, GenerateRandIntBelowMaximum(359, 500), GenerateRandIntBelowMaximum(101, 500), GenerateRandIntBelowMaximum(101, 500), RandomFloatBetweenTwoNumbers(0, 1, 500)),
		},
	}

	for testnumber, tc := range testcases {

		got := eng.Colorhsla(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("Test Number %d failed.  got %v , expected %v", testnumber, got, tc.expected)

			t.Error("\n" + StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
