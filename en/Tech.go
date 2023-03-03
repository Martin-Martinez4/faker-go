package english

type Tech struct {
	techadjectives *[]TechAdjective
}

var tech = Tech{
	techadjectives: &techadjectives,
}

func (e *EnDict) TechAdj() string {

	return RandomEntryFromSlice(e.techadjectives)

}
