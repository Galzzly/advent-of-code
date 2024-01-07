package main

import (
	_ "embed"
	"os"
	"unicode"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var inputTest string

type Digits [2]digit
type digit struct {
	id  int
	num string
}

var digitMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	// Check argv if we use test input or not
	if len(os.Args) > 1 && os.Args[1] == "test" {
		input = inputTest
	}

	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}

func findNums(line string) Digits {
	result := [2]digit{{-1, ""}, {-1, ""}}
	for i, c := range line {
		if unicode.IsDigit(c) {
			if result[0].id == -1 {
				result[0] = digit{i, string(c)}
				result[1] = digit{i, string(c)}
				continue
			}
			result[1] = digit{i, string(c)}
		}
	}
	return result
}
