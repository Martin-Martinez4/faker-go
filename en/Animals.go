package main

import "errors"

type Animals struct {
	Bears     *Bears
	Birds     *Birds
	Cats      *Cats
	Cetaceans *Cetaceans
	Cows      *Cows
	Dogs      *Dogs
	Fishes    *Fishes
	Insects   *Insects
	Monkeys   *Monkeys
	Rodents   *Rodents
	Snakes    *Snakes
	Turtles   *Turtles
}

var animals = Animals{
	Bears:     &bears,
	Birds:     &birds,
	Cats:      &cats,
	Cetaceans: &cetaceans,
	Cows:      &cows,
	Dogs:      &dogs,
	Fishes:    &fishes,
	Insects:   &insects,
	Monkeys:   &monkeys,
	Rodents:   &rodents,
	Snakes:    &snakes,
	Turtles:   &turtles,
}

func (a *Animals) Bear() string {

	index := GenerateRandIntBelowMaximum(len(*a.Bears))
	return (*a.Bears)[index]
}

func (a *Animals) Bird() string {

	index := GenerateRandIntBelowMaximum(len(*a.Birds))
	return (*a.Birds)[index]
}

func (a *Animals) Cat() (string, error) {

	if len(*a.Cats) > 0 {

		index := GenerateRandIntBelowMaximum(len(*a.Cats))
		return (*a.Cats)[index], nil

	} else {

		return "", errors.New("Cat list is empty")
	}

}

func (a *Animals) Cow() string {

	index := GenerateRandIntBelowMaximum(len(*a.Cows))

	return (*a.Cows)[index]

}

func (a *Animals) Cetacean() string {

	index := GenerateRandIntBelowMaximum(len(*a.Cetaceans))

	return (*a.Cetaceans)[index]

}

func (a *Animals) Dog() string {

	index := GenerateRandIntBelowMaximum(len(*a.Dogs))

	return (*a.Dogs)[index]

}

func (a *Animals) Fish() string {

	index := GenerateRandIntBelowMaximum(len(*a.Fishes))

	return (*a.Fishes)[index]

}

func (a *Animals) Insect() string {

	index := GenerateRandIntBelowMaximum(len(*a.Insects))

	return (*a.Insects)[index]

}

func (a *Animals) Monkey() string {

	index := GenerateRandIntBelowMaximum(len(*a.Monkeys))

	return (*a.Monkeys)[index]

}

func (a *Animals) Rodent() string {

	index := GenerateRandIntBelowMaximum(len(*a.Rodents))

	return (*a.Rodents)[index]

}

func (a *Animals) Snake() string {

	index := GenerateRandIntBelowMaximum(len(*a.Snakes))

	return (*a.Snakes)[index]

}

func (a *Animals) Turtle() string {

	index := GenerateRandIntBelowMaximum(len(*a.Turtles))

	return (*a.Turtles)[index]

}
