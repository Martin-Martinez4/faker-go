package english

import "fmt"

type Person struct {
	malenames   *[]MaleName
	femalenames *[]FemaleName
	lastnames   *[]LastName
}

var people = Person{
	malenames:   &malenames,
	femalenames: &femalenames,
	lastnames:   &lastnames,
}

func (e *EnDict) MaleName(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.malenames, seed_optional...)

}

func (e *EnDict) FemaleName(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.femalenames, seed_optional...)

}

func (e *EnDict) LastName(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.lastnames, seed_optional...)

}

func (e *EnDict) Name(seed_optional ...int64) string {

	randomInt := GenerateRandIntBelowMaximum(2, seed_optional...)

	if randomInt == 1 {

		return fmt.Sprintf("%v", e.MaleName(seed_optional...))

	} else {

		return fmt.Sprintf("%v", e.FemaleName(seed_optional...))
	}

}
func (e *EnDict) FullName(seed_optional ...int64) string {

	randomInt := GenerateRandIntBelowMaximum(2, seed_optional...)

	if randomInt == 1 {

		return fmt.Sprintf("%v %v", e.MaleName(seed_optional...), e.LastName(seed_optional...))

	} else {

		return fmt.Sprintf("%v %v", e.FemaleName(seed_optional...), e.LastName(seed_optional...))
	}

}
