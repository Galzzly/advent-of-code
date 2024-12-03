package main

import (
	"regexp"
	"strings"
)

func doPartTwo(input string) int {
	input = strings.ReplaceAll(input, "\n", "")
	do_reg := map[bool]string{false: "do\\(\\)", true: "don't\\(\\)"}
	do := true
	var line string
	for {
		regex := regexp.MustCompile(do_reg[do])
		idx := regex.FindStringIndex(input)
		if idx == nil {
			if do {
				line += input
			}
			break
		}
		if do {
			line += input[:idx[0]]
			input = input[idx[1]:]
			do = !do
			continue
		}
		input = input[idx[1]:]
		do = !do
	}
	return doPartOne(line)
}
