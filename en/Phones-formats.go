package english

type Phoneformat struct {
	Domestic                  string
	DomesticParthenses        string
	DomesticParthensesNoHypen string
	Local                     string
	LocalNoHypen              string
	International             string
	InternationalPlus         string
}

var Phoneformats = Phoneformat{

	Domestic:                  "###-###-####",
	DomesticParthenses:        "(###)-###-####",
	DomesticParthensesNoHypen: "(###) ###-####",
	Local:                     "###-####",
	LocalNoHypen:              "#######",
	International:             "1-###-####",
	InternationalPlus:         "+1-###-####",
}
