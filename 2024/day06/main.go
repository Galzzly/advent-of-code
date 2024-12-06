package main

import (
	_ "embed"
	"image"
	"os"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var inputTest string

var delta = []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

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
