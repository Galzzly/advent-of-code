package main

import (
	"aocli/utils/reader"
	"image"
)

var mdelta = []image.Point{
	{-1, -1}, {-1, 1},
}

var aletters = [][][]rune{
	{{'M', 'S'}, {'S', 'M'}},
	{{'M', 'S'}, {'S', 'M'}},
}

func doPartTwo(input string) int {
	var mapper Mapper
	mapper, allA := makeMap(reader.FileLineByLine(input), 'A')
	var res int
	for _, i := range allA {
		res += mapper.checkMAS(i)
	}
	return res
}

func (m Mapper) checkMAS(XY image.Point) int {
	for _, a := range aletters {
		succ := true
		for i, d := range mdelta {
			if !((m[XY.Add(d)] == a[i][0] && m[XY.Sub(d)] == a[i][1]) || (m[XY.Add(d)] == a[i][1] && m[XY.Sub(d)] == a[i][0])) {
				// succ++
				succ = false
				break
			}
		}
		if succ {
			return 1
		}
	}
	return 0
}
