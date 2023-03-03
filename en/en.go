package english

import (
	"math/rand"
	"time"
)

type EnDict struct {
	*Animals
	*Colors
	*Dates
	*Locations
	*Lorems
	*Internet
	*Person
	*Phones
	*Tech
	*Words
}

type Enfaker struct {
	*EnDict
}

func NewEnglishFaker() Enfaker {

	var endict = EnDict{
		&animals,
		&colors,
		&dates,
		&locations,
		&lorems,
		&internet,
		&people,
		&phones,
		&tech,
		&words,
	}

	rand.Seed(time.Now().UnixNano())

	return Enfaker{
		EnDict: &endict,
	}

}
