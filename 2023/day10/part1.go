package main

import (
	"aocli/utils/maps"
	"aocli/utils/reader"
)

func doPartOne(input string) int {
	pipemap := maps.MakeImagePointMap(reader.FileLineByLine(input))
	startpoint, _ := maps.MapKey(pipemap, 'S')

	return len(getPointsInLoop(pipemap, startpoint)) / 2
}
