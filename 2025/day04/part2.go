package main

import (
	"image"
)

func doPartTwo(mapper map[image.Point]rune) int {
	ans := 0

	for {
		toRemove := []image.Point{}

		// Find all cells that can be removed
		for p, r := range mapper {
			if r == '@' && countNeighbors(mapper, rect, p) < 4 {
				toRemove = append(toRemove, p)
			}
		}

		// If nothing to remove, we're done
		if len(toRemove) == 0 {
			break
		}

		// Remove all marked cells
		for _, p := range toRemove {
			mapper[p] = '.'
			ans++
		}
	}

	return ans
}
