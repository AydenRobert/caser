package caser

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func SepDataIn(dataIn string, inDelim ValidCase) []string {
	var dataSep []string
	switch inDelim {
	case Title, Sentence, Upper, Lower:
		dataSep = strings.Split(dataIn, " ")
	case TitleSnake, Snake, ScreamingSnake:
		dataSep = strings.Split(dataIn, "_")
	case Train, Kebab:
		dataSep = strings.Split(dataIn, "-")
	case Pascal, Camel:
		dataSep = sepCamel(dataIn)
	}
	return dataSep
}

func JoinDataOut(dataSep []string, outDelim ValidCase) (outStr string) {
	caser := cases.Title(language.English)

	for i := range dataSep {
		dataSep[i] = strings.ToLower(dataSep[i])
	}

	switch outDelim {
	case Title:
		for i := range dataSep {
			dataSep[i] = caser.String(dataSep[i])
		}
		outStr = strings.Join(dataSep, " ")
	case Sentence:
		dataSep[0] = caser.String(dataSep[0])
		outStr = strings.Join(dataSep, " ")
	case Upper:
		for i := range dataSep {
			dataSep[i] = strings.ToUpper(dataSep[i])
		}
		outStr = strings.Join(dataSep, " ")
	case Lower:
		outStr = strings.Join(dataSep, " ")
	case Pascal:
		for i := range dataSep {
			dataSep[i] = caser.String(dataSep[i])
		}
		outStr = strings.Join(dataSep, "")
	case Camel:
		for i := range dataSep {
			dataSep[i] = caser.String(dataSep[i])
		}
		dataSep[0] = strings.ToLower(dataSep[0])
		outStr = strings.Join(dataSep, "")
	case TitleSnake:
		for i := range dataSep {
			dataSep[i] = caser.String(dataSep[i])
		}
		outStr = strings.Join(dataSep, "_")
	case Snake:
		outStr = strings.Join(dataSep, "_")
	case ScreamingSnake:
		for i := range dataSep {
			dataSep[i] = strings.ToUpper(dataSep[i])
		}
		outStr = strings.Join(dataSep, "_")
	case Train:
		for i := range dataSep {
			dataSep[i] = caser.String(dataSep[i])
		}
		outStr = strings.Join(dataSep, "-")
	case Kebab:
		outStr = strings.Join(dataSep, "-")
	}
	return
}

func sepCamel(s string) []string {
	outArr := make([]string, 0, 10)
	j := 0
	for i := 1; i < len(s); i++ {
		if string(s[i]) == strings.ToUpper(string(s[i])) {
			outArr = append(outArr, s[j:i])
			j = i
		}
	}
	outArr = append(outArr, s[j:])
	return outArr
}
