package caser

import "github.com/AydenRobert/caser/caser"

func Convert(in, out, data string) string {
	inCase, ok := caser.IsValidCase(in)
	if !ok {
		inCase = caser.DetectCase(data)
	}
	outCase, ok := caser.IsValidCase(out)
	if !ok {
		return ""
	}
	dataSep := caser.SepDataIn(data, inCase)
	return caser.JoinDataOut(dataSep, outCase)
}