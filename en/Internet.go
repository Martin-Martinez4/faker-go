package main

type Internet struct {
	Emojis *Emojis
}

var internet = Internet{

	Emojis: &emojis,
}

type Getter2 interface {
	Emoji()
}

func Emoji(i *Internet) string {

	index := GenerateRandIntBelowMaximum(len(*i.Emojis))

	return (*i.Emojis)[index]

}
