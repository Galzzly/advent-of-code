package main

import (
	"aocli/utils/maps"
	"aocli/utils/polyfence"
	"aocli/utils/reader"
	"image"
	"slices"
)

func doPartTwo(input string) int {
	pipemap := maps.MakeImagePointMap(reader.FileLineByLine(input))
	startpoint, _ := maps.MapKey(pipemap, 'S')
	loop := getPointsInLoop(pipemap, startpoint)
	pf := polyfence.NewPolyfence(loop)
	var res int
	for P := range pipemap {
		if !slices.Contains[[]image.Point](loop, P) && pf.Inside(P) {
			res++
		}
	}
	return res
}
