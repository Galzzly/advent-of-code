package main

import (
	"aocli/utils/maps"
	"strings"
)

func doPartOne(input string) int {
	mapper, rect := maps.MakeImagePointMapRect(strings.Split(strings.TrimSpace(input), "\n"))
	return countRemovable(mapper, rect)
}
