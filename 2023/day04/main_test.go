package main

import "testing"

func BenchmarkPartOne(b *testing.B) {
	parseCards(input)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		doPartOne()
	}
}

func BenchmarkPartTwo(b *testing.B) {
	parseCards(input)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		doPartTwo()
	}
}
