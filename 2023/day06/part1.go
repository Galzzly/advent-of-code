package main

import (
	"aocli/utils"
	"strings"
)

func doPartOne(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	// Build up the races
	T := strings.Fields(lines[0])[1:]
	D := strings.Fields(lines[1])[1:]
	R := make([]Race, len(T))
	for i := range R {
		R[i] = Race{utils.Atoi(T[i]), utils.Atoi(D[i])}
	}

	// Get the ways to win
	ways := make([]int, len(R))
	for i, r := range R {
		ways[i] = getWays(r)
	}
	return utils.MultiplyArray(ways)
}
