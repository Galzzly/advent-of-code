package main

import "testing"

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solveAll(nums, 25)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solveAll(nums, 75)
	}
}
