package main

import (
	"math/rand"
	"time"
)

func GenerateRandIntBelowMaximum(maximum int, seed_optional ...int64) int {

	var seedValue int64

	if len(seed_optional) > 0 {
		seedValue = seed_optional[0]
	} else {

		seedValue = time.Now().UnixNano()

	}

	rand.Seed(seedValue)

	return rand.Intn(maximum)
}

func RandomIntBetweenTwoNumbers(minimum int, maximum int, seed_optional ...int64) int {

	var seedValue int64

	if len(seed_optional) > 0 {
		seedValue = seed_optional[0]
	} else {

		seedValue = time.Now().UnixNano()

	}

	var aFloat = RandomFloatBetweenTwoNumbers(0, 1, seedValue)

	return int(float64(minimum) + aFloat*(float64(maximum)-float64(minimum)))
}

func RandomFloatBetweenTwoNumbers(minimum int, maximum int, seed_optional ...int64) float64 {

	var seedValue int64

	if len(seed_optional) > 0 {
		seedValue = seed_optional[0]
	} else {

		seedValue = time.Now().UnixNano()

	}

	rand.Seed(seedValue)

	return float64(minimum) + rand.Float64()*(float64(maximum)-float64(minimum))
}
