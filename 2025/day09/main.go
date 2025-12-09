package main

import (
	_ "embed"
	"fmt"
	"image"
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

func parseLines(input string) []image.Point {
	res := make([]image.Point, 0, len(input))
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		var c image.Point
		fmt.Sscanf(line, "%d,%d", &c.X, &c.Y)

		res = append(res, c)
	}
	return res
}
