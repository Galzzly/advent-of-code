package main

import (
	"aocli/utils/union"
	"slices"
)

func doPartTwo() int {
	// Sort by distance (shortest first)
	slices.SortFunc(distances, func(a, b distance) int {
		return a.d - b.d
	})

	// Connect pairs until we have a spanning tree (n-1 connections)
	uf := union.NewUnionFind(len(junctions))
	ans := 0
	connection := 0

	for _, d := range distances {
		if uf.Find(d.a) != uf.Find(d.b) {
			connection++
			if connection == len(junctions)-1 {
				ans = junctions[d.a].x * junctions[d.b].x
				break
			}
			uf.Union(d.a, d.b)
		}
	}

	return ans
}
