package main

import (
	"aocli/utils"
	"image"
	"strings"
)

func doPartOne(input string) int {
	mapper = make(map[image.Point]int)
	startPoints = []image.Point{}
	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			num := utils.Atoi(string(char))
			mapper[image.Point{x, y}] = num
			if num == 0 {
				startPoints = append(startPoints, image.Point{x, y})
			}
		}
	}
	var res int
	for _, P := range startPoints {
		res += checkRoute(P)
	}
	return res
}

func checkRoute(p image.Point) int {
	NP := []image.Point{p}
	var res int
	seen := map[image.Point]bool{}
	for len(NP) > 0 {
		P := NP[0]
		NP = NP[1:]
		if seen[P] {
			continue
		}
		seen[P] = true
		val := mapper[P]
		if val == 9 {
			res++
			continue
		}
		for _, D := range deltas {
			if mapper[P.Add(D)] == val+1 {
				NP = append(NP, P.Add(D))
			}
		}
	}
	return res
}
