package main

import (
	"fmt"
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
	*Teams
	*Vehicles
	*Words
}

type Opt int

var opt Opt = 4

func (t *TempFaker) GetOpt() Opt {

	temp := t.Opt

	return *temp
}

type TempFaker struct {
	*Opt
}

var tempfaker = TempFaker{

	&opt,
}

var endict = EnDict{
	&animals,
	&colors,
	&dates,
	&locations,
	&lorems,
	&internet,
	&people,
	&phones,
	&teams,
	&vehicles,
	&words,
}

type Enfaker struct {
	*EnDict
	*TempFaker
}

func NewEnglishFaker(dict *EnDict) Enfaker {

	return Enfaker{
		EnDict:    dict,
		TempFaker: &tempfaker,
	}

}

// func NewDict2[D *EnDict | *Animals](dict D, base *Dict) {

// 	dictType := strings.Split(reflect.TypeOf(dict).String(), ".")[1]

// 	if dictType == "EnDict" {

// 		base.EnDict = any(dict).(*EnDict)

// 	}

// }

func main() {

	enfk := NewEnglishFaker(&endict)

	fmt.Println(enfk.GetOpt())
	println(enfk.Cat())
	println(enfk.After())

}
