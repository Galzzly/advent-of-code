package main

import (
	"aocli/utils"
	_ "embed"
	"os"
	"strings"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var inputTest string

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

func findLargestDigits(input string, numDigits int) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	ans := 0
	for _, line := range lines {
		result := make([]byte, numDigits)
		start := 0
		lineLen := len(line)

		for pos := 0; pos < numDigits; pos++ {
			// We need to leave enough characters for the remaining digits
			remaining := numDigits - pos
			end := lineLen - remaining + 1

			// Find the largest digit in the valid range
			maxIdx := start
			maxDigit := line[start]
			for i := start + 1; i < end; i++ {
				if line[i] > maxDigit {
					maxDigit = line[i]
					maxIdx = i
				}
			}

			result[pos] = maxDigit
			start = maxIdx + 1
		}

		ans += utils.Atoi(string(result))
	}
	return ans
}
