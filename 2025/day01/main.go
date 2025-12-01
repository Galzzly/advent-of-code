package main

import (
	"container/ring"
	_ "embed"
	"os"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var inputTest string

var r *ring.Ring

func init() {
	r = ring.New(100)
	n := r.Len()
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
}

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
