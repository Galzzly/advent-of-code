package main

import (
	"aocli/utils/reader"
	"image"
)

func doPartOne(input string) int {
	mapper, start := makeMap(reader.FileLineByLine(input))
	dir := 0
	P := start
	visited := make(map[image.Point]interface{})
	for {
		visited[P] = nil
		NP := P.Add(delta[dir])
		if _, ok := mapper[NP]; !ok {
			break
		}
		if mapper[NP] {
			dir = (dir + 1) % 4
			NP = P.Add(delta[dir])
		}
		P = NP
	}
	return len(visited)
}

func makeMap(input []string) (map[image.Point]bool, image.Point) {
	mapper := make(map[image.Point]bool)
	var start image.Point
	for y, s := range input {
		for x, r := range s {
			mapper[image.Point{x, y}] = r == '#'
			if r == '^' {
				start = image.Point{x, y}
			}
		}
	}
	return mapper, start
}
