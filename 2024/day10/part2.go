package main

import "image"

func doPartTwo() int {
	var res int
	for _, P := range startPoints {
		res += checkDistinctTrail(P)
	}
	return res
}

func checkDistinctTrail(p image.Point) int {
	NP := []image.Point{p}
	var res int
	for len(NP) > 0 {
		P := NP[0]
		NP = NP[1:]
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
