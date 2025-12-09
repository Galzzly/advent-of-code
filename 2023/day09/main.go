package main

import (
	_ "embed"
	"os"
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

var sequences [][]int

func parseInput(input string) {
	sequences = make([][]int, 0, 200)

	i := 0
	for i < len(input) {
		if input[i] == '\n' || input[i] == '\r' {
			i++
			continue
		}

		nums := make([]int, 0, 21)
		for i < len(input) && input[i] != '\n' {
			// Parse number (can be negative)
			negative := false
			if input[i] == '-' {
				negative = true
				i++
			}

			num := 0
			for i < len(input) && input[i] >= '0' && input[i] <= '9' {
				num = num*10 + int(input[i]-'0')
				i++
			}
			if negative {
				num = -num
			}
			nums = append(nums, num)

			// Skip space
			if i < len(input) && input[i] == ' ' {
				i++
			}
		}
		sequences = append(sequences, nums)
	}
}

func extrapolate(nums []int) (int, int) {
	// Build difference pyramid and compute both extrapolations simultaneously
	n := len(nums)
	diffs := make([]int, n)
	copy(diffs, nums)

	var nextVal, prevVal int
	sign := 1

	for length := n; length > 1; length-- {
		// Add current rightmost value for part 1
		nextVal += diffs[length-1]

		// Add current leftmost value for part 2 (alternating signs)
		prevVal += sign * diffs[0]
		sign = -sign

		// Compute next level of differences
		allZero := true
		for i := 0; i < length-1; i++ {
			diffs[i] = diffs[i+1] - diffs[i]
			if diffs[i] != 0 {
				allZero = false
			}
		}

		if allZero {
			break
		}
	}

	return nextVal, prevVal
}
