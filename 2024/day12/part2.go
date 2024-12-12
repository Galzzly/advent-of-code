package main

import (
	"aocli/utils/maps"
	"aocli/utils/reader"
	"image"
	"slices"
)

func doPartTwo(input string) int {
	mapper := maps.MakeImagePointMap(reader.FileLineByLine(input))
	seen := make(map[image.Point]bool)
	var delta = []image.Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var res int
	for i, r := range mapper {
		if seen[i] {
			continue
		}
		NP := []image.Point{i}
		sides := make(map[image.Point][]image.Point)
		var area int
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
				if _, ok := sides[d]; !ok {
					sides[d] = []image.Point{}
				}
				sides[d] = append(sides[d], P)
			}
		}
		var S int
		for _, s := range sides {
			sideseen := make(map[image.Point]bool)
			for _, i := range s {
				if _, ok := sideseen[i]; ok {
					continue
				}
				S++
				NP := []image.Point{i}
				for len(NP) > 0 {
					P := NP[0]
					NP = NP[1:]
					if sideseen[P] {
						continue
					}
					sideseen[P] = true
					for _, dd := range delta {
						np := P.Add(dd)
						if slices.Contains(s, np) {
							NP = append(NP, np)
						}
					}
				}
			}
		}
		res += area * S
	}
	return res
}
