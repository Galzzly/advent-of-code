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

var mapper map[image.Point]rune
var antennas map[rune][]image.Point

func main() {
	// Check argv if we use test input or not
	if len(os.Args) > 1 && os.Args[1] == "test" {
		input = inputTest
	}

	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo()
	println(answer)
}

func makeMap(input []string) {
	mapper = make(map[image.Point]rune)
	antennas = make(map[rune][]image.Point)
	for y, s := range input {
		for x, r := range s {
			mapper[image.Point{x, y}] = r
			if r != '.' {
				antennas[r] = append(antennas[r], image.Point{x, y})
			}
		}
	}
}
