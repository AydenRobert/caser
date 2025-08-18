package caser_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/AydenRobert/caser/caser"
)

func TestPascalToSnake(t *testing.T) {
	pascalStr := "ThisIsAPascalString"
	expList := []string{
		"this",
		"is",
		"a",
		"pascal",
		"string",
	}
	expOut := "this_is_a_pascal_string"
	dataSep := caser.SepDataIn(pascalStr, caser.Pascal)
	result := caser.JoinDataOut(dataSep, caser.Snake)

	if len(expList) != len(dataSep) {
		t.Errorf(
			"Length of generated list is different from length of expected\n%v\n%v\n",
			expList,
			dataSep,
		)
		return
	}

	wrongIndexes := []int{}
	for i := range expList {
		if expList[i] != dataSep[i] {
			wrongIndexes = append(wrongIndexes, i)
		}
	}

	if len(wrongIndexes) != 0 {
		sb := strings.Builder{}
		sb.WriteString("Expected List | Generated List\n")
		sb.WriteString("==============================\n")
		for _, i := range wrongIndexes {
			sb.WriteString(fmt.Sprintf("%s | %s\n", expList[i], dataSep[i]))
		}
		t.Error(sb.String())
		return
	}

	if expOut != result {
		t.Errorf("Expected Output: %s\nGenerated Output: %s\n", expOut, result)
	}
}
