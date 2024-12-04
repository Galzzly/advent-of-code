package main

import (
	"aocli/utils/reader"
	"image"
)

type Mapper map[image.Point]rune

var delta = []image.Point{
	{-1, 1}, {0, 1}, {1, 1},
	{-1, 0}, {1, 0},
	{-1, -1}, {0, -1}, {1, -1},
}

var letters []rune = []rune("MAS")

func doPartOne(input string) int {
	var mapper Mapper
	mapper, allX := makeMap(reader.FileLineByLine(input), 'X')
	var res int
	for _, i := range allX {
		for _, d := range delta {
			res += mapper.checkXmas(i, d)
		}
	}
	return res
}

func (m Mapper) checkXmas(XY image.Point, d image.Point) int {
	P := XY

	for i := range letters {
		NP := P.Add(d)
		if !m.findNextLetter(NP, letters[i]) {
			return 0
		}
		P = NP
	}
	return 1
}

func (m Mapper) findNextLetter(XY image.Point, char rune) bool {
	return m[XY] == char
}

func makeMap(lines []string, c rune) (mapping map[image.Point]rune, allC []image.Point) {
	mapping = make(map[image.Point]rune)
	for y, s := range lines {
		for x, r := range s {
			mapping[image.Point{x, y}] = r
			if r == c {
				allC = append(allC, image.Point{x, y})
			}
		}
	}
	return
}
