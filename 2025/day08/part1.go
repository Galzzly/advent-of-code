package main

import (
	"aocli/utils/union"
	"slices"
)

type coords struct {
	x, y, z int
}

type distance struct {
	d    int
	a, b int
}

func doPartOne(k int) int {
	// Sort by distance (shortest first)
	slices.SortFunc(distances, func(a, b distance) int {
		return a.d - b.d
	})

	// Connect the k shortest pairs
	uf := union.NewUnionFind(len(junctions))
	for i := 0; i < k && i < len(distances); i++ {
		uf.Union(distances[i].a, distances[i].b)
	}

	// Count component sizes
	componentSizes := make(map[int]int)
	for j := range junctions {
		root := uf.Find(j)
		componentSizes[root]++
	}

	// Get the sizes and sort them (largest first)
	sizes := make([]int, 0, len(componentSizes))
	for _, size := range componentSizes {
		sizes = append(sizes, size)
	}

	// Sort descending
	slices.SortFunc(sizes, func(a, b int) int {
		return b - a
	})

	// Multiply the three largest
	if len(sizes) < 3 {
		return 0
	}
	return sizes[0] * sizes[1] * sizes[2]
}
