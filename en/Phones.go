package english

import (
	"fmt"
	"strings"
)

type Phones struct {
}

var phones = Phones{}

// BuildPhoneNumber builds a random phone number from the given format and returns it as a string.
//
// The seed_optional parameter can be left blank, it is used to produce predicatble results
//
// #s in the phonenumberformat parameter will be turned into digits and the rest or the characters will be left alone.
// for example: "(###) ###-####" -> "(111) 111-1111"
func BuildPhoneNumber(phonenumberformat string, seed_optional ...int64) string {

	// Check if a seed value for the random number generator is present
	var seedPresent bool

	if len(seed_optional) > 0 {

		seedPresent = true

	} else {

		seedPresent = false
	}

	var sb strings.Builder

	// Will hold the number that replaces the #(s)
	var number int

	// The random int generated is described by (min, max]
	max := 1
	var min int

	for _, char := range phonenumberformat {

		if char == '#' {

			// Each new # in a row means the upper limit that a value can be, increased by one place
			// ## -> max = 99 = (10^2) -1
			// The maximum of the random value generator is non inclusive, so when the max is 100 the highest number can be returned is 99

			// The min is most that one can subtract from the max without decreasing the number of digits
			// ## -> max = 100; min = 10, any number from 10-99 is two digits
			// The relatoniship between the min and max can be described as min = (max - ((max / 10) * 9))
			// 10 = 100 - ((100/10) *9) = 100 - (10 * 9)

			// This is done to run the random number generator only after a streak of #s has ended
			max *= 10

		} else {

			if max > 1 {

				min = (max - ((max / 10) * 9))

				if seedPresent {

					number = RandomIntBetweenTwoNumbers(min, max, seed_optional[0])

				} else {

					number = RandomIntBetweenTwoNumbers(min, max)
				}

				sb.WriteString(fmt.Sprint(number))

				max = 1

			}

			sb.WriteString(string(char))

		}

	}

	if max > 1 {

		min = (max - ((max / 10) * 9))

		if seedPresent {

			number = RandomIntBetweenTwoNumbers(min, max, seed_optional[0])

		} else {

			number = RandomIntBetweenTwoNumbers(min, max)
		}

		sb.WriteString(fmt.Sprint(number))

	}

	return sb.String()
}
