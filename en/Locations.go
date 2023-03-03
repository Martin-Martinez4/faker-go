package english

import "fmt"

type Locations struct {
	cities *[]city
	states *[]state
}

var locations = Locations{
	cities: &cities,
	states: &states,
}

func (e *EnDict) Location(seed_optional ...int64) string {

	city := RandomEntryFromSlice(e.cities, seed_optional...)
	state := RandomEntryFromSlice(e.states, seed_optional...)

	return fmt.Sprintf("%v, %v", city, state)

}
