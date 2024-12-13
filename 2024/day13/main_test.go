package main

import "testing"

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solve(input, 0)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solve(input, 10000000000000)
	}
}
