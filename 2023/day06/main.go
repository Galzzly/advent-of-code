package main

import (
	_ "embed"
	"os"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var inputTest string

func main() {
	// Check argv if we use test input or not
	if len(os.Args) > 1 && os.Args[1] == "test" {
		input = inputTest
	}

	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}

type Race struct {
	Time, Distance int
}

func getWays(r Race) int {
	max := getVal(r, r.Time, -1)
	min := getVal(r, 0, 1)
	return max - min + 1
}

func getVal(r Race, S, M int) int {
	for i := 0; i < r.Time; i++ {
		runtime := r.Time - (S + i*M)
		D := (S + i*M) * runtime
		if D > r.Distance {
			return (S + i*M)
		}
	}
	return r.Time - S
}
