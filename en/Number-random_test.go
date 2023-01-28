package main

import (
	"fmt"
	"testing"
)

// Need to test edge cases

func TestGenerateRandIntBelowMaximum(t *testing.T) {

	var testcase = []struct {
		maximum int
		seed    int64
	}{
		{maximum: 1, seed: 1},
		{maximum: 10, seed: 1},
		{maximum: 100, seed: 10},
		{maximum: 10, seed: 1243},
		{maximum: 10000, seed: 1564},
		{maximum: 50, seed: 100000000},
	}

	for _, tc := range testcase {

		got := GenerateRandIntBelowMaximum(tc.maximum, tc.seed)

		if got >= tc.maximum || got < 0 || got != GenerateRandIntBelowMaximum(tc.maximum, tc.seed) {

			t.Errorf("\n\033[31m"+"got %v", got)

		} else {

			fmt.Printf("got: %v, max was: %v\n", got, tc.maximum)

		}

	}

}

func TestRandomFloatBetweenTwoNumbers(t *testing.T) {

	var testcase = []struct {
		maximum int
		minimum int
		seed    int64
	}{
		{minimum: 0, maximum: 1, seed: 1},
		{minimum: 0, maximum: 10, seed: 1},
		{minimum: 0, maximum: 100, seed: 1},
		{minimum: 100, maximum: 100, seed: 1},
		{minimum: 50, maximum: 1000, seed: 1},
		{minimum: 0, maximum: 10, seed: 1},
	}

	for _, tc := range testcase {

		got := RandomFloatBetweenTwoNumbers(tc.minimum, tc.maximum, tc.seed)

		if got > float64(tc.maximum) || got < float64(tc.minimum) || got != RandomFloatBetweenTwoNumbers(tc.minimum, tc.maximum, tc.seed) {

			t.Errorf("\n\033[31m"+"got: %v, min was: %v, max was: %v\n", got, tc.minimum, tc.maximum)

		} else {

			fmt.Printf("got: %v, min was: %v, max was: %v\n", got, tc.minimum, tc.maximum)

		}

	}

}
