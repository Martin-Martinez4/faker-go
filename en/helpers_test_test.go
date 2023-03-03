package english

import (
	"fmt"
	"testing"
)

func TestStyleTerminalOuput(t *testing.T) {

	var escapesequence = "\033["
	var resetstyles = " \033[0m"

	var testcases = []struct {
		basestring string
		options    []string
		expected   string
	}{
		{
			basestring: "test",
			options:    []string{},
			expected:   "test",
		},
		{
			basestring: "test",
			options:    []string{ANSItextcolors.red},
			expected:   fmt.Sprintf(escapesequence + ANSItextcolors.red + "m" + " test" + resetstyles),
		},
		{
			basestring: "test",
			options:    []string{ANSItextcolors.red, ANSIbackgroundcolors.white},
			expected:   fmt.Sprintf(escapesequence + ANSItextcolors.red + ";" + ANSIbackgroundcolors.white + "m" + " test" + resetstyles),
		},
	}

	for _, tc := range testcases {

		got := StyleTerminalOuput(tc.basestring, tc.options...)
		if got != tc.expected {

			println(got == tc.expected, tc.options)

			t.Errorf("got: %s, expected %s", got, tc.expected)
		}

	}

}
