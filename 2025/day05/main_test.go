package main

import (
	"strings"
	"testing"
)

func BenchmarkPartOne(b *testing.B) {
	// Parse input once before benchmark
	sections = strings.Split(strings.TrimSpace(input), "\n\n")
	parseRanges()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		doPartOne()
	}
}

func BenchmarkPartTwo(b *testing.B) {
	// Parse input once before benchmark
	sections = strings.Split(strings.TrimSpace(input), "\n\n")
	parseRanges()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		doPartTwo()
	}
}
