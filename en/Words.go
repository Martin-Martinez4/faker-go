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

func (e *EnDict) Adjective(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.adjectives, seed_optional...)

}
func (e *EnDict) Adverb(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.adverbs, seed_optional...)

}
func (e *EnDict) Noun(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.nouns, seed_optional...)

}
func (e *EnDict) Verb(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.verbs.baseverbs, seed_optional...)

}
func (e *EnDict) VerbPastTense(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.verbs.simplepastverbs, seed_optional...)

}
func (e *EnDict) VerbPastParticiple(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.verbs.pastparticipleverbs, seed_optional...)

}
func (e *EnDict) VerbThirdPersonSingular(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.verbs.thirdpersonsingularverbs, seed_optional...)

}
func (e *EnDict) VerbPresentParticiple(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.verbs.presentparticleverbs, seed_optional...)

}
