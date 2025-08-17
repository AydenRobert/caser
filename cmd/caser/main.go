package main

import (
	"fmt"
	"io"
	"os"

	"github.com/AydenRobert/caser/caser"
)

func main() {
	args := os.Args
	if len(args) > 3 || len(args) < 2 {
		return
	}

	var inDelim, outDelim caser.ValidCase
	var ok bool


	dataIn, err := io.ReadAll(os.Stdin)
	if err != nil {
		return
	}

	if len(args) == 2 {
		inDelim = caser.DetectCase(string(dataIn))
		outDelim, ok = caser.IsValidCase(args[1])
		if !ok {
			return
		}
	} else {
		inDelim, ok = caser.IsValidCase(args[1])
		if !ok {
			return
		}
		outDelim, ok = caser.IsValidCase(args[2])
		if !ok {
			return
		}
	}

	dataSep := caser.SepDataIn(string(dataIn), inDelim)

	outStr := caser.JoinDataOut(dataSep, outDelim)

	fmt.Print(outStr)
}
