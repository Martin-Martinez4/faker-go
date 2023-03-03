package english

import (
	"fmt"
	"testing"
)

func TestEmoji(t *testing.T) {

	eng := NewEnglishFaker()

	emojislen := len(*eng.emojis)

	var testcases = []struct {
		seed     int64
		expected string
	}{
		{
			seed:     1,
			expected: (*eng.emojis)[GenerateRandIntBelowMaximum(emojislen, 1)],
		},
		{
			seed:     2,
			expected: (*eng.emojis)[GenerateRandIntBelowMaximum(emojislen, 2)],
		},
		{
			seed:     500,
			expected: (*eng.emojis)[GenerateRandIntBelowMaximum(emojislen, 500)],
		},
		{
			seed:     600,
			expected: (*eng.emojis)[GenerateRandIntBelowMaximum(emojislen, 600)],
		},
		{
			seed:     700,
			expected: (*eng.emojis)[GenerateRandIntBelowMaximum(emojislen, 700)],
		},
		{
			seed:     851,
			expected: (*eng.emojis)[GenerateRandIntBelowMaximum(emojislen, 851)],
		},
	}

	for testnumber, tc := range testcases {

		got := eng.Emoji(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("test number %d failed: got %v, expected %v", testnumber, got, tc.expected)

			t.Error(StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
func TestEmail(t *testing.T) {

	eng := NewEnglishFaker()

	usernameslen := len(*eng.usernames)
	emaildomainslen := len(*eng.emaildomains)
	extensionslen := len(*eng.extensions)

	var testcases = []struct {
		seed     int64
		expected string
	}{
		{
			seed:     1,
			expected: fmt.Sprintf("%s@%s%s", (*eng.usernames)[GenerateRandIntBelowMaximum(usernameslen, 1)], (*eng.emaildomains)[GenerateRandIntBelowMaximum(emaildomainslen, 1)], (*eng.extensions)[GenerateRandIntBelowMaximum(extensionslen, 1)]),
		},
		{
			seed:     2,
			expected: fmt.Sprintf("%s@%s%s", (*eng.usernames)[GenerateRandIntBelowMaximum(usernameslen, 2)], (*eng.emaildomains)[GenerateRandIntBelowMaximum(emaildomainslen, 2)], (*eng.extensions)[GenerateRandIntBelowMaximum(extensionslen, 2)]),
		},
		{
			seed:     500,
			expected: fmt.Sprintf("%s@%s%s", (*eng.usernames)[GenerateRandIntBelowMaximum(usernameslen, 500)], (*eng.emaildomains)[GenerateRandIntBelowMaximum(emaildomainslen, 500)], (*eng.extensions)[GenerateRandIntBelowMaximum(extensionslen, 500)]),
		},
		{
			seed:     600,
			expected: fmt.Sprintf("%s@%s%s", (*eng.usernames)[GenerateRandIntBelowMaximum(usernameslen, 600)], (*eng.emaildomains)[GenerateRandIntBelowMaximum(emaildomainslen, 600)], (*eng.extensions)[GenerateRandIntBelowMaximum(extensionslen, 600)]),
		},
		{
			seed:     700,
			expected: fmt.Sprintf("%s@%s%s", (*eng.usernames)[GenerateRandIntBelowMaximum(usernameslen, 700)], (*eng.emaildomains)[GenerateRandIntBelowMaximum(emaildomainslen, 700)], (*eng.extensions)[GenerateRandIntBelowMaximum(extensionslen, 700)]),
		},
		{
			seed:     851,
			expected: fmt.Sprintf("%s@%s%s", (*eng.usernames)[GenerateRandIntBelowMaximum(usernameslen, 851)], (*eng.emaildomains)[GenerateRandIntBelowMaximum(emaildomainslen, 851)], (*eng.extensions)[GenerateRandIntBelowMaximum(extensionslen, 851)]),
		},
	}

	for testnumber, tc := range testcases {

		got := eng.Email(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("test number %d failed: got %v, expected %v", testnumber, got, tc.expected)

			t.Error(StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
func TestUsername(t *testing.T) {

	eng := NewEnglishFaker()

	usernameslen := len(*eng.usernames)

	var testcases = []struct {
		seed     int64
		expected string
	}{
		{
			seed:     1,
			expected: (*eng.usernames)[GenerateRandIntBelowMaximum(usernameslen, 1)],
		},
		{
			seed:     2,
			expected: (*eng.usernames)[GenerateRandIntBelowMaximum(usernameslen, 2)],
		},
		{
			seed:     500,
			expected: (*eng.usernames)[GenerateRandIntBelowMaximum(usernameslen, 500)],
		},
		{
			seed:     600,
			expected: (*eng.usernames)[GenerateRandIntBelowMaximum(usernameslen, 600)],
		},
		{
			seed:     700,
			expected: (*eng.usernames)[GenerateRandIntBelowMaximum(usernameslen, 700)],
		},
		{
			seed:     851,
			expected: (*eng.usernames)[GenerateRandIntBelowMaximum(usernameslen, 851)],
		},
	}

	for testnumber, tc := range testcases {

		got := eng.Username(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("test number %d failed: got %v, expected %v", testnumber, got, tc.expected)

			t.Error(StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
func TestWebsite(t *testing.T) {

	eng := NewEnglishFaker()

	webdomainslen := len(*eng.websitedomains)
	extensionslen := len(*eng.extensions)

	var testcases = []struct {
		seed     int64
		expected string
	}{
		{
			seed:     1,
			expected: fmt.Sprintf("https://www.%s%s", (*eng.websitedomains)[GenerateRandIntBelowMaximum(webdomainslen, 1)], (*eng.extensions)[GenerateRandIntBelowMaximum(extensionslen, 1)]),
		},
		{
			seed:     2,
			expected: fmt.Sprintf("https://www.%s%s", (*eng.websitedomains)[GenerateRandIntBelowMaximum(webdomainslen, 2)], (*eng.extensions)[GenerateRandIntBelowMaximum(extensionslen, 2)]),
		},
		{
			seed:     500,
			expected: fmt.Sprintf("https://www.%s%s", (*eng.websitedomains)[GenerateRandIntBelowMaximum(webdomainslen, 500)], (*eng.extensions)[GenerateRandIntBelowMaximum(extensionslen, 500)]),
		},
		{
			seed:     600,
			expected: fmt.Sprintf("https://www.%s%s", (*eng.websitedomains)[GenerateRandIntBelowMaximum(webdomainslen, 600)], (*eng.extensions)[GenerateRandIntBelowMaximum(extensionslen, 600)]),
		},
		{
			seed:     700,
			expected: fmt.Sprintf("https://www.%s%s", (*eng.websitedomains)[GenerateRandIntBelowMaximum(webdomainslen, 700)], (*eng.extensions)[GenerateRandIntBelowMaximum(extensionslen, 700)]),
		},
		{
			seed:     851,
			expected: fmt.Sprintf("https://www.%s%s", (*eng.websitedomains)[GenerateRandIntBelowMaximum(webdomainslen, 851)], (*eng.extensions)[GenerateRandIntBelowMaximum(extensionslen, 851)]),
		},
	}

	for testnumber, tc := range testcases {

		got := eng.Website(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("test number %d failed: got %v, expected %v", testnumber, got, tc.expected)

			t.Error(StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
