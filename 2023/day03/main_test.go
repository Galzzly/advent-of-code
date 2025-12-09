package main

import (
	"strings"
	"testing"
)

func BenchmarkPartOne(b *testing.B) {
	grid = strings.Split(strings.TrimSpace(input), "\n")
	height = len(grid)
	width = len(grid[0])
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		doPartOne()
	}
}

func BenchmarkPartTwo(b *testing.B) {
	grid = strings.Split(strings.TrimSpace(input), "\n")
	height = len(grid)
	width = len(grid[0])
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		doPartTwo()
	}
}
