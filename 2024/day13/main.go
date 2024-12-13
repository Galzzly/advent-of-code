package main

import (
	"aocli/utils/reader"
	_ "embed"
	"fmt"
	"image"
	"os"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var inputTest string

type machine struct {
	A, B, P image.Point
}

func main() {
	// Check argv if we use test input or not
	if len(os.Args) > 1 && os.Args[1] == "test" {
		input = inputTest
	}

	answer := solve(input, 0)
	println(answer)

	answer = solve(input, 10000000000000)
	println(answer)
}

func solve(input string, n int) int {
	lines := reader.FileDoubleLine(input)
	var machines []machine
	for _, line := range lines {
		var A, B, P image.Point
		fmt.Sscanf(line, "Button A: X+%d, Y+%d\nButton B: X+%d, Y+%d\nPrize: X=%d, Y=%d\n", &A.X, &A.Y, &B.X, &B.Y, &P.X, &P.Y)
		machines = append(machines, machine{A, B, P.Add(image.Point{n, n})})
	}

	var res int

	for _, m := range machines {
		res += calculate(m)
	}
	return res
}

func calculate(m machine) int {
	aP := (m.B.Y*m.P.X - m.B.X*m.P.Y) / (m.A.X*m.B.Y - m.A.Y*m.B.X)
	bP := (m.A.Y*m.P.X - m.A.X*m.P.Y) / (m.A.Y*m.B.X - m.A.X*m.B.Y)

	if m.A.Mul(aP).Add(m.B.Mul(bP)) == m.P {
		return aP*3 + bP
	}
	return 0
}
