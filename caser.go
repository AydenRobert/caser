package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	args := os.Args
	if len(args) > 3 || len(args) < 2 {
		return
	}

	var inDelim, outDelim validCase
	var ok bool


	dataIn, err := io.ReadAll(os.Stdin)
	if err != nil {
		return
	}

	if len(args) == 2 {
		inDelim = detectCase(string(dataIn))
		outDelim, ok = isValidCase(args[1])
		if !ok {
			return
		}
	} else {
		inDelim, ok = isValidCase(args[1])
		if !ok {
			return
		}
		outDelim, ok = isValidCase(args[2])
		if !ok {
			return
		}
	}

	dataSep := sepDataIn(string(dataIn), inDelim)

	outStr := joinDataOut(dataSep, outDelim)

	fmt.Print(outStr)
}

func sepDataIn(dataIn string, inDelim validCase) []string {
	var dataSep []string
	switch inDelim {
	case title, sentence, upper, lower:
		dataSep = strings.Split(dataIn, " ")
	case titleSnake, snake, screamingSnake:
		dataSep = strings.Split(dataIn, "_")
	case train, kebab:
		dataSep = strings.Split(dataIn, "-")
	case pascal, camel:
		dataSep = sepCamel(dataIn)
	}
	return dataSep
}

func joinDataOut(dataSep []string, outDelim validCase) (outStr string) {
	caser := cases.Title(language.English)

	for i := range dataSep {
		dataSep[i] = strings.ToLower(dataSep[i])
	}

	switch outDelim {
	case title:
		for i := range dataSep {
			dataSep[i] = caser.String(dataSep[i])
		}
		outStr = strings.Join(dataSep, " ")
	case sentence:
		dataSep[0] = caser.String(dataSep[0])
		outStr = strings.Join(dataSep, " ")
	case upper:
		for i := range dataSep {
			dataSep[i] = strings.ToUpper(dataSep[i])
		}
		outStr = strings.Join(dataSep, " ")
	case lower:
		outStr = strings.Join(dataSep, " ")
	case pascal:
		for i := range dataSep {
			dataSep[i] = caser.String(dataSep[i])
		}
		outStr = strings.Join(dataSep, "")
	case camel:
		for i := range dataSep {
			dataSep[i] = caser.String(dataSep[i])
		}
		dataSep[0] = strings.ToLower(dataSep[0])
		outStr = strings.Join(dataSep, "")
	case titleSnake:
		for i := range dataSep {
			dataSep[i] = caser.String(dataSep[i])
		}
		outStr = strings.Join(dataSep, "_")
	case snake:
		outStr = strings.Join(dataSep, "_")
	case screamingSnake:
		for i := range dataSep {
			dataSep[i] = strings.ToUpper(dataSep[i])
		}
		outStr = strings.Join(dataSep, "_")
	case train:
		for i := range dataSep {
			dataSep[i] = caser.String(dataSep[i])
		}
		outStr = strings.Join(dataSep, "-")
	case kebab:
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
	outArr = append(outArr, s[j:len(s)-1])
	return outArr
}
