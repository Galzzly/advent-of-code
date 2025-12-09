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

type mapper struct {
	dest, src, size int
}

var (
	seeds []int
	maps  [][]mapper
)

func main() {
	// Check argv if we use test input or not
	if len(os.Args) > 1 && os.Args[1] == "test" {
		input = inputTest
	}

	parseInput(input)

	answer := doPartOne()
	println(answer)

	answer = doPartTwo()
	println(answer)
}

func parseInput(input string) {
	sections := strings.Split(strings.TrimSpace(input), "\n\n")

	// Parse seeds
	seedLine := sections[0][7:] // Skip "seeds: "
	seeds = parseNumbers(seedLine)

	// Parse maps
	maps = make([][]mapper, len(sections)-1)
	for i := 1; i < len(sections); i++ {
		lines := strings.Split(sections[i], "\n")[1:] // Skip header line
		maps[i-1] = make([]mapper, len(lines))

		for j, line := range lines {
			nums := parseNumbers(line)
			maps[i-1][j] = mapper{dest: nums[0], src: nums[1], size: nums[2]}
		}
	}
}

func parseNumbers(s string) []int {
	var nums []int
	i := 0
	for i < len(s) {
		if s[i] >= '0' && s[i] <= '9' {
			num := 0
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			nums = append(nums, num)
		} else {
			i++
		}
	}
	return nums
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
