package english

type Animals struct {
	bears          *[]Bear
	birds          *[]Bird
	cats           *[]Cat
	cetaceans      *[]Cetacean
	cows           *[]Cow
	dogs           *[]Dog
	fishes         *[]Fish
	genericanimals *[]GenericAnimal
	insects        *[]Insect
	monkeys        *[]Monkey
	rodents        *[]Rodent
	snakes         *[]Snake
	turtles        *[]Turtle
}

var animals = Animals{
	bears:          &bears,
	birds:          &birds,
	cats:           &cats,
	cetaceans:      &cetaceans,
	cows:           &cows,
	dogs:           &dogs,
	fishes:         &fishes,
	genericanimals: &genericanimals,
	insects:        &insects,
	monkeys:        &monkeys,
	rodents:        &rodents,
	snakes:         &snakes,
	turtles:        &turtles,
}

func (e *EnDict) Animal(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.genericanimals, seed_optional...)

}

func (e *EnDict) Bear(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.bears, seed_optional...)
}

func (e *EnDict) Bird(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.birds, seed_optional...)
}

func (e *EnDict) Cat(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.cats, seed_optional...)

}

func (e *EnDict) Cow(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.cows, seed_optional...)

}

func (e *EnDict) Cetacean(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.cetaceans, seed_optional...)

}

func (e *EnDict) Dog(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.dogs, seed_optional...)

}

func (e *EnDict) Fish(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.fishes, seed_optional...)

}

func (e *EnDict) Insect(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.insects, seed_optional...)

}

func (e *EnDict) Monkey(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.monkeys, seed_optional...)

}

func (e *EnDict) Rodent(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.rodents, seed_optional...)

}

func (e *EnDict) Snake(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.snakes, seed_optional...)

}

func (e *EnDict) Turtle(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.turtles, seed_optional...)

}
