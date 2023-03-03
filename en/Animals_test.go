package english

import (
	"fmt"
	"testing"
)

func TestAnimal(t *testing.T) {

	var eng = NewEnglishFaker()
	var testcases = []struct {
		expected string
		seed     int64
	}{

		{
			expected: (*eng.genericanimals)[GenerateRandIntBelowMaximum(len(*eng.genericanimals), 1)],
			seed:     1,
		},
		{
			expected: (*eng.genericanimals)[GenerateRandIntBelowMaximum(len(*eng.genericanimals), 2)],
			seed:     2,
		},
		{
			expected: (*eng.genericanimals)[GenerateRandIntBelowMaximum(len(*eng.genericanimals), 3)],
			seed:     3,
		},
		{
			expected: (*eng.genericanimals)[GenerateRandIntBelowMaximum(len(*eng.genericanimals), 4)],
			seed:     4,
		},
		{
			expected: (*eng.genericanimals)[GenerateRandIntBelowMaximum(len(*eng.genericanimals), 5)],
			seed:     5,
		},
		{
			expected: (*eng.genericanimals)[GenerateRandIntBelowMaximum(len(*eng.genericanimals), 6)],
			seed:     6,
		},
		{
			expected: (*eng.genericanimals)[GenerateRandIntBelowMaximum(len(*eng.genericanimals), 10000)],
			seed:     10000,
		},
	}

	for _, tc := range testcases {

		got := eng.Animal(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v, expected %v", got, tc.expected)

			t.Error(StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
func TestBear(t *testing.T) {

	var eng = NewEnglishFaker()
	var testcases = []struct {
		expected string
		seed     int64
	}{

		{
			expected: (*eng.bears)[GenerateRandIntBelowMaximum(len(*eng.bears), 1)],
			seed:     1,
		},
		{
			expected: (*eng.bears)[GenerateRandIntBelowMaximum(len(*eng.bears), 2)],
			seed:     2,
		},
		{
			expected: (*eng.bears)[GenerateRandIntBelowMaximum(len(*eng.bears), 3)],
			seed:     3,
		},
		{
			expected: (*eng.bears)[GenerateRandIntBelowMaximum(len(*eng.bears), 4)],
			seed:     4,
		},
		{
			expected: (*eng.bears)[GenerateRandIntBelowMaximum(len(*eng.bears), 5)],
			seed:     5,
		},
		{
			expected: (*eng.bears)[GenerateRandIntBelowMaximum(len(*eng.bears), 6)],
			seed:     6,
		},
		{
			expected: (*eng.bears)[GenerateRandIntBelowMaximum(len(*eng.bears), 10000)],
			seed:     10000,
		},
	}

	for _, tc := range testcases {

		got := eng.Bear(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v, expected %v", got, tc.expected)

			t.Error(StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
func TestBird(t *testing.T) {

	var eng = NewEnglishFaker()
	var testcases = []struct {
		expected string
		seed     int64
	}{

		{
			expected: (*eng.birds)[GenerateRandIntBelowMaximum(len(*eng.birds), 1)],
			seed:     1,
		},
		{
			expected: (*eng.birds)[GenerateRandIntBelowMaximum(len(*eng.birds), 2)],
			seed:     2,
		},
		{
			expected: (*eng.birds)[GenerateRandIntBelowMaximum(len(*eng.birds), 3)],
			seed:     3,
		},
		{
			expected: (*eng.birds)[GenerateRandIntBelowMaximum(len(*eng.birds), 4)],
			seed:     4,
		},
		{
			expected: (*eng.birds)[GenerateRandIntBelowMaximum(len(*eng.birds), 5)],
			seed:     5,
		},
		{
			expected: (*eng.birds)[GenerateRandIntBelowMaximum(len(*eng.birds), 6)],
			seed:     6,
		},
		{
			expected: (*eng.birds)[GenerateRandIntBelowMaximum(len(*eng.birds), 10000)],
			seed:     10000,
		},
	}

	for _, tc := range testcases {

		got := eng.Bird(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v, expected %v", got, tc.expected)

			t.Error(StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
func TestCat(t *testing.T) {

	var eng = NewEnglishFaker()
	var testcases = []struct {
		expected string
		seed     int64
	}{

		{
			expected: (*eng.cats)[GenerateRandIntBelowMaximum(len(*eng.cats), 1)],
			seed:     1,
		},
		{
			expected: (*eng.cats)[GenerateRandIntBelowMaximum(len(*eng.cats), 2)],
			seed:     2,
		},
		{
			expected: (*eng.cats)[GenerateRandIntBelowMaximum(len(*eng.cats), 3)],
			seed:     3,
		},
		{
			expected: (*eng.cats)[GenerateRandIntBelowMaximum(len(*eng.cats), 4)],
			seed:     4,
		},
		{
			expected: (*eng.cats)[GenerateRandIntBelowMaximum(len(*eng.cats), 5)],
			seed:     5,
		},
		{
			expected: (*eng.cats)[GenerateRandIntBelowMaximum(len(*eng.cats), 6)],
			seed:     6,
		},
		{
			expected: (*eng.cats)[GenerateRandIntBelowMaximum(len(*eng.cats), 10000)],
			seed:     10000,
		},
	}

	for _, tc := range testcases {

		got := eng.Cat(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v, expected %v", got, tc.expected)

			t.Error(StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
func TestCow(t *testing.T) {

	var eng = NewEnglishFaker()
	var testcases = []struct {
		expected string
		seed     int64
	}{

		{
			expected: (*eng.cows)[GenerateRandIntBelowMaximum(len(*eng.cows), 1)],
			seed:     1,
		},
		{
			expected: (*eng.cows)[GenerateRandIntBelowMaximum(len(*eng.cows), 2)],
			seed:     2,
		},
		{
			expected: (*eng.cows)[GenerateRandIntBelowMaximum(len(*eng.cows), 3)],
			seed:     3,
		},
		{
			expected: (*eng.cows)[GenerateRandIntBelowMaximum(len(*eng.cows), 4)],
			seed:     4,
		},
		{
			expected: (*eng.cows)[GenerateRandIntBelowMaximum(len(*eng.cows), 5)],
			seed:     5,
		},
		{
			expected: (*eng.cows)[GenerateRandIntBelowMaximum(len(*eng.cows), 6)],
			seed:     6,
		},
		{
			expected: (*eng.cows)[GenerateRandIntBelowMaximum(len(*eng.cows), 10000)],
			seed:     10000,
		},
	}

	for _, tc := range testcases {

		got := eng.Cow(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v, expected %v", got, tc.expected)

			t.Error(StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
func TestCetacean(t *testing.T) {

	var eng = NewEnglishFaker()
	var testcases = []struct {
		expected string
		seed     int64
	}{

		{
			expected: (*eng.cetaceans)[GenerateRandIntBelowMaximum(len(*eng.cetaceans), 1)],
			seed:     1,
		},
		{
			expected: (*eng.cetaceans)[GenerateRandIntBelowMaximum(len(*eng.cetaceans), 2)],
			seed:     2,
		},
		{
			expected: (*eng.cetaceans)[GenerateRandIntBelowMaximum(len(*eng.cetaceans), 3)],
			seed:     3,
		},
		{
			expected: (*eng.cetaceans)[GenerateRandIntBelowMaximum(len(*eng.cetaceans), 4)],
			seed:     4,
		},
		{
			expected: (*eng.cetaceans)[GenerateRandIntBelowMaximum(len(*eng.cetaceans), 5)],
			seed:     5,
		},
		{
			expected: (*eng.cetaceans)[GenerateRandIntBelowMaximum(len(*eng.cetaceans), 6)],
			seed:     6,
		},
		{
			expected: (*eng.cetaceans)[GenerateRandIntBelowMaximum(len(*eng.cetaceans), 10000)],
			seed:     10000,
		},
	}

	for _, tc := range testcases {

		got := eng.Cetacean(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v, expected %v", got, tc.expected)

			t.Error(StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
func TestDog(t *testing.T) {

	var eng = NewEnglishFaker()
	var testcases = []struct {
		expected string
		seed     int64
	}{

		{
			expected: (*eng.dogs)[GenerateRandIntBelowMaximum(len(*eng.dogs), 1)],
			seed:     1,
		},
		{
			expected: (*eng.dogs)[GenerateRandIntBelowMaximum(len(*eng.dogs), 2)],
			seed:     2,
		},
		{
			expected: (*eng.dogs)[GenerateRandIntBelowMaximum(len(*eng.dogs), 3)],
			seed:     3,
		},
		{
			expected: (*eng.dogs)[GenerateRandIntBelowMaximum(len(*eng.dogs), 4)],
			seed:     4,
		},
		{
			expected: (*eng.dogs)[GenerateRandIntBelowMaximum(len(*eng.dogs), 5)],
			seed:     5,
		},
		{
			expected: (*eng.dogs)[GenerateRandIntBelowMaximum(len(*eng.dogs), 6)],
			seed:     6,
		},
		{
			expected: (*eng.dogs)[GenerateRandIntBelowMaximum(len(*eng.dogs), 10000)],
			seed:     10000,
		},
	}

	for _, tc := range testcases {

		got := eng.Dog(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v, expected %v", got, tc.expected)

			t.Error(StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
func TestFish(t *testing.T) {

	var eng = NewEnglishFaker()
	var testcases = []struct {
		expected string
		seed     int64
	}{

		{
			expected: (*eng.fishes)[GenerateRandIntBelowMaximum(len(*eng.fishes), 1)],
			seed:     1,
		},
		{
			expected: (*eng.fishes)[GenerateRandIntBelowMaximum(len(*eng.fishes), 2)],
			seed:     2,
		},
		{
			expected: (*eng.fishes)[GenerateRandIntBelowMaximum(len(*eng.fishes), 3)],
			seed:     3,
		},
		{
			expected: (*eng.fishes)[GenerateRandIntBelowMaximum(len(*eng.fishes), 4)],
			seed:     4,
		},
		{
			expected: (*eng.fishes)[GenerateRandIntBelowMaximum(len(*eng.fishes), 5)],
			seed:     5,
		},
		{
			expected: (*eng.fishes)[GenerateRandIntBelowMaximum(len(*eng.fishes), 6)],
			seed:     6,
		},
		{
			expected: (*eng.fishes)[GenerateRandIntBelowMaximum(len(*eng.fishes), 10000)],
			seed:     10000,
		},
	}

	for _, tc := range testcases {

		got := eng.Fish(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v, expected %v", got, tc.expected)

			t.Error(StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
func TestInsect(t *testing.T) {

	var eng = NewEnglishFaker()
	var testcases = []struct {
		expected string
		seed     int64
	}{

		{
			expected: (*eng.insects)[GenerateRandIntBelowMaximum(len(*eng.insects), 1)],
			seed:     1,
		},
		{
			expected: (*eng.insects)[GenerateRandIntBelowMaximum(len(*eng.insects), 2)],
			seed:     2,
		},
		{
			expected: (*eng.insects)[GenerateRandIntBelowMaximum(len(*eng.insects), 3)],
			seed:     3,
		},
		{
			expected: (*eng.insects)[GenerateRandIntBelowMaximum(len(*eng.insects), 4)],
			seed:     4,
		},
		{
			expected: (*eng.insects)[GenerateRandIntBelowMaximum(len(*eng.insects), 5)],
			seed:     5,
		},
		{
			expected: (*eng.insects)[GenerateRandIntBelowMaximum(len(*eng.insects), 6)],
			seed:     6,
		},
		{
			expected: (*eng.insects)[GenerateRandIntBelowMaximum(len(*eng.insects), 10000)],
			seed:     10000,
		},
	}

	for _, tc := range testcases {

		got := eng.Insect(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v, expected %v", got, tc.expected)

			t.Error(StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
func TestMonkey(t *testing.T) {

	var eng = NewEnglishFaker()
	var testcases = []struct {
		expected string
		seed     int64
	}{

		{
			expected: (*eng.monkeys)[GenerateRandIntBelowMaximum(len(*eng.monkeys), 1)],
			seed:     1,
		},
		{
			expected: (*eng.monkeys)[GenerateRandIntBelowMaximum(len(*eng.monkeys), 2)],
			seed:     2,
		},
		{
			expected: (*eng.monkeys)[GenerateRandIntBelowMaximum(len(*eng.monkeys), 3)],
			seed:     3,
		},
		{
			expected: (*eng.monkeys)[GenerateRandIntBelowMaximum(len(*eng.monkeys), 4)],
			seed:     4,
		},
		{
			expected: (*eng.monkeys)[GenerateRandIntBelowMaximum(len(*eng.monkeys), 5)],
			seed:     5,
		},
		{
			expected: (*eng.monkeys)[GenerateRandIntBelowMaximum(len(*eng.monkeys), 6)],
			seed:     6,
		},
		{
			expected: (*eng.monkeys)[GenerateRandIntBelowMaximum(len(*eng.monkeys), 10000)],
			seed:     10000,
		},
	}

	for _, tc := range testcases {

		got := eng.Monkey(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v, expected %v", got, tc.expected)

			t.Error(StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
func TestRodent(t *testing.T) {

	var eng = NewEnglishFaker()
	var testcases = []struct {
		expected string
		seed     int64
	}{

		{
			expected: (*eng.rodents)[GenerateRandIntBelowMaximum(len(*eng.rodents), 1)],
			seed:     1,
		},
		{
			expected: (*eng.rodents)[GenerateRandIntBelowMaximum(len(*eng.rodents), 2)],
			seed:     2,
		},
		{
			expected: (*eng.rodents)[GenerateRandIntBelowMaximum(len(*eng.rodents), 3)],
			seed:     3,
		},
		{
			expected: (*eng.rodents)[GenerateRandIntBelowMaximum(len(*eng.rodents), 4)],
			seed:     4,
		},
		{
			expected: (*eng.rodents)[GenerateRandIntBelowMaximum(len(*eng.rodents), 5)],
			seed:     5,
		},
		{
			expected: (*eng.rodents)[GenerateRandIntBelowMaximum(len(*eng.rodents), 6)],
			seed:     6,
		},
		{
			expected: (*eng.rodents)[GenerateRandIntBelowMaximum(len(*eng.rodents), 10000)],
			seed:     10000,
		},
	}

	for _, tc := range testcases {

		got := eng.Rodent(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v, expected %v", got, tc.expected)

			t.Error(StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
func TestSnake(t *testing.T) {

	var eng = NewEnglishFaker()
	var testcases = []struct {
		expected string
		seed     int64
	}{

		{
			expected: (*eng.snakes)[GenerateRandIntBelowMaximum(len(*eng.snakes), 1)],
			seed:     1,
		},
		{
			expected: (*eng.snakes)[GenerateRandIntBelowMaximum(len(*eng.snakes), 2)],
			seed:     2,
		},
		{
			expected: (*eng.snakes)[GenerateRandIntBelowMaximum(len(*eng.snakes), 3)],
			seed:     3,
		},
		{
			expected: (*eng.snakes)[GenerateRandIntBelowMaximum(len(*eng.snakes), 4)],
			seed:     4,
		},
		{
			expected: (*eng.snakes)[GenerateRandIntBelowMaximum(len(*eng.snakes), 5)],
			seed:     5,
		},
		{
			expected: (*eng.snakes)[GenerateRandIntBelowMaximum(len(*eng.snakes), 6)],
			seed:     6,
		},
		{
			expected: (*eng.snakes)[GenerateRandIntBelowMaximum(len(*eng.snakes), 10000)],
			seed:     10000,
		},
	}

	for _, tc := range testcases {

		got := eng.Snake(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v, expected %v", got, tc.expected)

			t.Error(StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
func TestTurtle(t *testing.T) {

	var eng = NewEnglishFaker()
	var testcases = []struct {
		expected string
		seed     int64
	}{

		{
			expected: (*eng.turtles)[GenerateRandIntBelowMaximum(len(*eng.turtles), 1)],
			seed:     1,
		},
		{
			expected: (*eng.turtles)[GenerateRandIntBelowMaximum(len(*eng.turtles), 2)],
			seed:     2,
		},
		{
			expected: (*eng.turtles)[GenerateRandIntBelowMaximum(len(*eng.turtles), 3)],
			seed:     3,
		},
		{
			expected: (*eng.turtles)[GenerateRandIntBelowMaximum(len(*eng.turtles), 4)],
			seed:     4,
		},
		{
			expected: (*eng.turtles)[GenerateRandIntBelowMaximum(len(*eng.turtles), 5)],
			seed:     5,
		},
		{
			expected: (*eng.turtles)[GenerateRandIntBelowMaximum(len(*eng.turtles), 6)],
			seed:     6,
		},
		{
			expected: (*eng.turtles)[GenerateRandIntBelowMaximum(len(*eng.turtles), 10000)],
			seed:     10000,
		},
	}

	for _, tc := range testcases {

		got := eng.Turtle(tc.seed)
		if got != tc.expected {

			errormessage := fmt.Sprintf("got %v, expected %v", got, tc.expected)

			t.Error(StyleTerminalOuput(errormessage, ANSItextcolors.red))

		}

	}

}
