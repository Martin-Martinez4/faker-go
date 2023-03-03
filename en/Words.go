package english

type verbs struct {
	baseverbs                *[]verb
	simplepastverbs          *[]verb
	pastparticipleverbs      *[]verb
	thirdpersonsingularverbs *[]verb
	presentparticleverbs     *[]verb
}

type Words struct {
	adjectives *[]adjecttive
	adverbs    *[]adverb
	nouns      *[]noun
	verbs      *verbs
}

var words = Words{
	adjectives: &adjectives,
	adverbs:    &adverbs,
	nouns:      &nouns,
	verbs: &verbs{
		baseverbs:                &baseverbs,
		simplepastverbs:          &simplepastverbs,
		pastparticipleverbs:      &pastparticipleverbs,
		thirdpersonsingularverbs: &thirdpersonsingularverbs,
		presentparticleverbs:     &presentparticleverbs,
	},
}
