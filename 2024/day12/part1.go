package main

import (
	"aocli/utils/maps"
	"aocli/utils/reader"
	"image"
)

func doPartOne(input string) int {
	mapper := maps.MakeImagePointMap(reader.FileLineByLine(input))
	seen := make(map[image.Point]bool)
	var delta = []image.Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var res int
	for i, r := range mapper {
		if seen[i] {
			continue
		}
		NP := []image.Point{i}
		var area, perim int
		for len(NP) > 0 {
			P := NP[0]
			NP = NP[1:]
			if seen[P] {
				continue
			}
			area++

			seen[P] = true
			for _, d := range delta {
				np := P.Add(d)
				if _, ok := mapper[np]; mapper[np] == r && ok {
					NP = append(NP, np)
					continue
				}
				perim++
			}
		}
		res += area * perim
	}
	return res
}
