package main

type delimType int

const (
	space delimType = iota
	underscore
	dash
	none
)

func detectCase(s string) (caser validCase) {
	delim := detectDelim(s)
	switch delim {
	case space:
		caser = lower
	case underscore:
		caser = snake
	case dash:
		caser = kebab
	case none:
		caser = camel
	}
	return
}

func detectDelim(s string) (delim delimType) {
	spaceCount, underscoreCount, dashCount := 0, 0, 0
	for _, c := range s {
		switch c {
			case ' ':
				spaceCount++
			case '_':
				underscoreCount++
			case '-':
				dashCount++
		}
	}
	if spaceCount > 0 {
		delim = space
	} else if underscoreCount > dashCount {
		delim = underscore
	} else if underscoreCount < dashCount {
		delim = dash
	} else {
		delim = none
	}
	return
}
