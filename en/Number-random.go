package english

import (
	"math/rand"
	"time"
)

func GenerateRandIntBelowMaximum(maximum int, seed_optional ...int64) int {

	var randomInt int

	if len(seed_optional) > 0 {

		// Sets the random seed to the value provided
		// Creates predicatble returns which is useful for testing
		seedValue := seed_optional[0]
		rand.Seed(seedValue)

		randomInt = rand.Intn(maximum)

		// Sets the random seed tp a value based on the current time in order to make the sand random again
		rand.Seed(time.Now().UnixNano())

	} else {

		randomInt = rand.Intn(maximum)
	}

	return randomInt
}

func RandomIntBetweenTwoNumbers(minimum int, maximum int, seed_optional ...int64) int {

	var aFloat float64

	if len(seed_optional) > 0 {

		// RandomFloatBetweenTwoNumbers already resets the seed to time.Now().UnixNano(), so it is skipped here
		aFloat = RandomFloatBetweenTwoNumbers(0, 1, seed_optional[0])

	} else {

		aFloat = RandomFloatBetweenTwoNumbers(0, 1)
	}

	return int(float64(minimum) + aFloat*(float64(maximum)-float64(minimum)))
}

func RandomFloatBetweenTwoNumbers(minimum int, maximum int, seed_optional ...int64) float64 {

	var randomFloat64 float64

	if len(seed_optional) > 0 {

		seedValue := seed_optional[0]
		rand.Seed(seedValue)
		randomFloat64 = rand.Float64()

		rand.Seed(time.Now().Unix())

	} else {

		randomFloat64 = rand.Float64()

	}

	return float64(minimum) + randomFloat64*(float64(maximum)-float64(minimum))
}
