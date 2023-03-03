package english

import (
	"strings"
)

type Lorems struct {
	loremwords *[]loremword
}

var lorems = Lorems{
	loremwords: &loremwords,
}

func (e *EnDict) Lorem(amountofwords int, seed_optional ...int64) string {

	if amountofwords <= 0 {

		return ""
	}

	var sb strings.Builder

	var stringtoappend string

	for i := 0; i < amountofwords; i++ {

		if i != 0 {

			stringtoappend = RandomEntryFromSlice(e.loremwords)

		} else {

			stringtoappend = RandomEntryFromSlice(e.loremwords, seed_optional...)

		}

		sb.WriteString(stringtoappend + " ")

	}

	return sb.String()

}
