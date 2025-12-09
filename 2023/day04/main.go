package main

import (
	_ "embed"
	"os"
	"strings"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var inputTest string

var cardWins []int

func main() {
	// Check argv if we use test input or not
	if len(os.Args) > 1 && os.Args[1] == "test" {
		input = inputTest
	}

	parseCards(input)

	answer := doPartOne()
	println(answer)

	answer = doPartTwo()
	println(answer)
}

func parseCards(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	cardWins = make([]int, len(lines))

	for i, line := range lines {
		colonIdx := strings.Index(line, ": ")
		pipeIdx := strings.Index(line, " | ")

		// Parse winning numbers into a set
		winSet := make(map[int]bool, 10)
		numStr := line[colonIdx+2 : pipeIdx]
		j := 0
		for j < len(numStr) {
			if numStr[j] >= '0' && numStr[j] <= '9' {
				num := 0
				for j < len(numStr) && numStr[j] >= '0' && numStr[j] <= '9' {
					num = num*10 + int(numStr[j]-'0')
					j++
				}
				winSet[num] = true
			} else {
				j++
			}
		}

		// Check your numbers against winning set
		wins := 0
		numStr = line[pipeIdx+3:]
		j = 0
		for j < len(numStr) {
			if numStr[j] >= '0' && numStr[j] <= '9' {
				num := 0
				for j < len(numStr) && numStr[j] >= '0' && numStr[j] <= '9' {
					num = num*10 + int(numStr[j]-'0')
					j++
				}
				if winSet[num] {
					wins++
				}
			} else {
				j++
			}
		}

		cardWins[i] = wins
	}
}
