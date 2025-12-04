package main

import (
	"aocli/utils/maps"
	"strings"
	"testing"
)

func BenchmarkPartOne(b *testing.B) {
	// Parse input once before benchmark
	mapper, rect = maps.MakeImagePointMapRect(strings.Split(strings.TrimSpace(input), "\n"))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		doPartOne(input)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	// Parse input once before benchmark
	mapper, rect = maps.MakeImagePointMapRect(strings.Split(strings.TrimSpace(input), "\n"))

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		mapperCopy := maps.CopyMap(mapper)
		doPartTwo(mapperCopy)
	}
}
