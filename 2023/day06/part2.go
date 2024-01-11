package main

import (
	"aocli/utils"
	"strings"
)

func doPartTwo(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var R Race
	R.Time = utils.Atoi(strings.Join(strings.Fields(lines[0])[1:], ""))
	R.Distance = utils.Atoi(strings.Join(strings.Fields(lines[1])[1:], ""))

	return getWays(R)
}
