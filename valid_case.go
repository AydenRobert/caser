package main

type validCase string

const (
	title          validCase = "title"
	sentence       validCase = "sentence"
	upper          validCase = "upper"
	lower          validCase = "lower"
	pascal         validCase = "pascal"
	camel          validCase = "camel"
	titleSnake     validCase = "title_snake"
	snake          validCase = "snake"
	screamingSnake validCase = "screaming_snake"
	train          validCase = "train"
	kebab          validCase = "kebab"
)

func isValidCase(s string) (caser validCase, ok bool) {
	ok = true
	switch s {
	case "title":
		caser = title
	case "sentence":
		caser = sentence
	case "upper":
		caser = upper
	case "lower":
		caser = lower
	case "pascal":
		caser = pascal
	case "camel":
		caser = camel
	case "title_snake":
		caser = titleSnake
	case "snake":
		caser = snake
	case "screaming_snake":
		caser = screamingSnake
	case "train":
		caser = train
	case "kebab":
		caser = kebab
	default:
		ok = false
	}
	return
}
