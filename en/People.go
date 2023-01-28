package main

type Person struct {
	MaleFirstNames *MaleFirstNames
}

var people = Person{
	MaleFirstNames: &maleFirstNames,
}

func (p *Person) Name() string {

	index := GenerateRandIntBelowMaximum(len(*p.MaleFirstNames))

	return (*p.MaleFirstNames)[index]

}
