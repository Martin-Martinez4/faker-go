package english

import "fmt"

type Internet struct {
	emojis         *[]Emoji
	emaildomains   *[]EmailDomain
	extensions     *[]Extension
	usernames      *[]Username
	websitedomains *[]WebsiteDomain
}

var internet = Internet{

	emojis:         &emojis,
	emaildomains:   &emaildomains,
	extensions:     &extensions,
	usernames:      &usernames,
	websitedomains: &websitedomains,
}

func (e *EnDict) Emoji(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.emojis, seed_optional...)

}

func (e *EnDict) Email(seed_optional ...int64) string {

	domain := RandomEntryFromSlice(e.emaildomains, seed_optional...)
	extension := RandomEntryFromSlice(e.extensions, seed_optional...)
	username := RandomEntryFromSlice(e.usernames, seed_optional...)

	return fmt.Sprintf("%s@%s%s", username, domain, extension)

}

func (e *EnDict) Username(seed_optional ...int64) string {

	return RandomEntryFromSlice(e.usernames, seed_optional...)

}

func (e *EnDict) Website(seed_optional ...int64) string {

	domain := RandomEntryFromSlice(e.websitedomains, seed_optional...)
	extensions := RandomEntryFromSlice(e.extensions, seed_optional...)

	return fmt.Sprintf("https://www.%s%s", domain, extensions)

}
