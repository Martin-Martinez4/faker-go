package english

import (
	"fmt"
	"strings"
	"testing"
)

func lorembuilder(amountofwords int, seed_optional ...int64) string {

	if amountofwords <= 0 {

		return ""
	}

	e := NewEnglishFaker()

	var sb strings.Builder

	var stringtoappend string

	for i := 0; i < amountofwords; i++ {

		if i != 0 {
			index := GenerateRandIntBelowMaximum(len(*e.loremwords))

			stringtoappend = (*e.loremwords)[index]

		} else {

			index := GenerateRandIntBelowMaximum(len(*e.loremwords), seed_optional...)

			stringtoappend = (*e.loremwords)[index]
		}

		sb.WriteString(stringtoappend + " ")

	}

	return sb.String()
}

func TestLorem(t *testing.T) {

	eng := NewEnglishFaker()

	var testcases = []struct {
		seed          int64
		amountofwords int
		expected      string
	}{
		{
			seed:          1,
			amountofwords: 0,
			expected:      "",
		},
		{
			seed:          1,
			amountofwords: 1,
			expected:      lorembuilder(1, 1),
		},
		{
			seed:          100,
			amountofwords: 100,
			expected:      lorembuilder(100, 100),
		},
		{
			seed:          123,
			amountofwords: 123,
			expected:      lorembuilder(123, 123),
		},
		{
			seed:          4564,
			amountofwords: 20,
			expected:      lorembuilder(20, 4564),
		},
		{
			seed:          6464684,
			amountofwords: 10,
			expected:      lorembuilder(10, 6464684),
		},
		{
			seed:          0,
			amountofwords: 0,
			expected:      lorembuilder(0, 0),
		},
	}

	for testnumber, tc := range testcases {

		got := eng.Lorem(tc.amountofwords, tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("Test Number %d failed.  got %v , expected %v", testnumber, got, tc.expected)

			t.Error("\n" + StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
