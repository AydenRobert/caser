package caser

type delimType int

const (
	space delimType = iota
	underscore
	dash
	none
)

func DetectCase(s string) (caser ValidCase) {
	delim := detectDelim(s)
	switch delim {
	case space:
		caser = Lower
	case underscore:
		caser = Snake
	case dash:
		caser = Kebab
	case none:
		caser = Camel
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
