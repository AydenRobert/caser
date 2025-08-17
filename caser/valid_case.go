package caser

type ValidCase string

const (
	Title          ValidCase = "title"
	Sentence       ValidCase = "sentence"
	Upper          ValidCase = "upper"
	Lower          ValidCase = "lower"
	Pascal         ValidCase = "pascal"
	Camel          ValidCase = "camel"
	TitleSnake     ValidCase = "title_snake"
	Snake          ValidCase = "snake"
	ScreamingSnake ValidCase = "screaming_snake"
	Train          ValidCase = "train"
	Kebab          ValidCase = "kebab"
)

func IsValidCase(s string) (caser ValidCase, ok bool) {
	ok = true
	switch s {
	case "title":
		caser = Title
	case "sentence":
		caser = Sentence
	case "upper":
		caser = Upper
	case "lower":
		caser = Lower
	case "pascal":
		caser = Pascal
	case "camel":
		caser = Camel
	case "title_snake":
		caser = TitleSnake
	case "snake":
		caser = Snake
	case "screaming_snake":
		caser = ScreamingSnake
	case "train":
		caser = Train
	case "kebab":
		caser = Kebab
	default:
		ok = false
	}
	return
}
