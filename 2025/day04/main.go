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

var delta = []image.Point{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
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

func countRemovable(mapper map[image.Point]rune, rect image.Rectangle) int {
	ans := 0
	for p, r := range mapper {
		if r == '@' && countNeighbors(mapper, rect, p) < 4 {
			ans++
		}
	}
	return ans
}

func countNeighbors(mapper map[image.Point]rune, rect image.Rectangle, p image.Point) int {
	count := 0
	for _, d := range delta {
		np := p.Add(d)
		if np.In(rect) && mapper[np] == '@' {
			count++
		}
	}
	return count
}
