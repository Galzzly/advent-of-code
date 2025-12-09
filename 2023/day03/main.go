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

var (
	grid   []string
	height int
	width  int
)

func main() {
	// Check argv if we use test input or not
	if len(os.Args) > 1 && os.Args[1] == "test" {
		input = inputTest
	}

	// Parse grid once
	grid = strings.Split(strings.TrimSpace(input), "\n")
	height = len(grid)
	width = len(grid[0])

	answer := doPartOne()
	println(answer)

	answer = doPartTwo()
	println(answer)
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
