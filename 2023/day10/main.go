package main

import (
	_ "embed"
	"image"
	"os"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var inputTest string

type PipeMap map[image.Point]rune

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

func getPointsInLoop(mapper PipeMap, startpoint image.Point) []image.Point {
	PointsInLoop := []image.Point{startpoint}
	P := startpoint
	Dir := '>'
	for {
		var NP image.Point
		switch mapper[P] {
		case 'S': // this is the start point.
			NP = P.Add(image.Point{1, 0})
		case '|': // Vertical Pipe
			if Dir == 'U' {
				NP = P.Add(image.Point{0, -1})
			} else {
				NP = P.Add(image.Point{0, 1})
			}
		case '-': // Horizontal Pipe
			if Dir == '>' {
				NP = P.Add(image.Point{1, 0})
			} else {
				NP = P.Add(image.Point{-1, 0})
			}
		case 'L': // North to East
			if Dir == 'D' {
				NP = P.Add(image.Point{1, 0})
				Dir = '>'
			} else {
				NP = P.Add(image.Point{0, -1})
				Dir = 'U'
			}
		case 'J': // North to West
			if Dir == 'D' {
				NP = P.Add(image.Point{-1, 0})
				Dir = '<'
			} else {
				NP = P.Add(image.Point{0, -1})
				Dir = 'U'
			}
		case '7': // South to West
			if Dir == 'U' {
				NP = P.Add(image.Point{-1, 0})
				Dir = '<'
			} else {
				NP = P.Add(image.Point{0, 1})
				Dir = 'D'
			}
		case 'F': // South to East
			if Dir == '<' {
				NP = P.Add(image.Point{0, 1})
				Dir = 'D'
			} else {
				NP = P.Add(image.Point{1, 0})
				Dir = '>'
			}
		case '.': // Ground, no pipe
		}
		if NP == startpoint {
			break
		}
		P = NP
		PointsInLoop = append(PointsInLoop, P)
	}
	return PointsInLoop
}
