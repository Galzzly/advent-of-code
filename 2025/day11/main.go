package main

import (
	_ "embed"
	"os"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var inputTest string

//go:embed input_test2.txt
var inputTest2 string

func main() {
	// Check argv if we use test input or not
	if len(os.Args) > 1 && os.Args[1] == "test" {
		answer := doPartOne(inputTest)
		println(answer)

		answer = doPartTwo(inputTest2)
		println(answer)
		return
	}

	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}
