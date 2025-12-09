package main

import (
	"aocli/utils/polyfence"
	"image"
	"slices"
)

func doPartTwo(input string) int {
	coords := parseLines(input)
	pf := polyfence.NewPolyfence(coords)

	// Get unique sorted rows and columns for coordinate compression
	rowSet := make(map[int]bool)
	colSet := make(map[int]bool)
	for _, c := range coords {
		rowSet[c.Y] = true
		colSet[c.X] = true
	}

	rows := make([]int, 0, len(rowSet))
	for r := range rowSet {
		rows = append(rows, r)
	}
	slices.Sort(rows)

	cols := make([]int, 0, len(colSet))
	for c := range colSet {
		cols = append(cols, c)
	}
	slices.Sort(cols)

	// Create index maps
	rowIdx := make(map[int]int)
	for i, r := range rows {
		rowIdx[r] = i
	}
	colIdx := make(map[int]int)
	for i, c := range cols {
		colIdx[c] = i
	}

	// Build compressed grid - mark all points inside polygon
	grid := make([][]bool, len(rows))
	for i := range grid {
		grid[i] = make([]bool, len(cols))
	}

	for r := range rows {
		for c := range cols {
			if pf.Inside(image.Point{X: cols[c], Y: rows[r]}) {
				grid[r][c] = true
			}
		}
	}

	// Build 2D prefix sum
	prefixSum := make([][]int, len(rows)+1)
	for i := range prefixSum {
		prefixSum[i] = make([]int, len(cols)+1)
	}

	for r := 1; r <= len(rows); r++ {
		for c := 1; c <= len(cols); c++ {
			val := 0
			if grid[r-1][c-1] {
				val = 1
			}
			prefixSum[r][c] = val + prefixSum[r-1][c] + prefixSum[r][c-1] - prefixSum[r-1][c-1]
		}
	}

	// Find max rectangle
	maxArea := 0
	for i := 0; i < len(coords); i++ {
		c1 := coords[i]
		r1 := rowIdx[c1.Y] + 1
		col1 := colIdx[c1.X] + 1

		for j := i + 1; j < len(coords); j++ {
			c2 := coords[j]

			// Skip lines
			if c1.X == c2.X || c1.Y == c2.Y {
				continue
			}

			// Get compressed indices for c2
			r2 := rowIdx[c2.Y] + 1
			col2 := colIdx[c2.X] + 1

			// Calculate min/max
			var minR, maxR, minC, maxC int
			if r1 < r2 {
				minR, maxR = r1, r2
			} else {
				minR, maxR = r2, r1
			}
			if col1 < col2 {
				minC, maxC = col1, col2
			} else {
				minC, maxC = col2, col1
			}

			// Calculate compressed area (number of grid cells)
			compressedArea := (maxR - minR + 1) * (maxC - minC + 1)

			// Check if all points in rectangle are valid using prefix sum
			sum := prefixSum[maxR][maxC] - prefixSum[minR-1][maxC] - prefixSum[maxR][minC-1] + prefixSum[minR-1][minC-1]

			if sum == compressedArea {
				// Calculate actual area
				dx := c2.X - c1.X
				dy := c2.Y - c1.Y
				if dx < 0 {
					dx = -dx
				}
				if dy < 0 {
					dy = -dy
				}
				area := (dx + 1) * (dy + 1)

				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}
