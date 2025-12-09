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

type Point struct {
	x, y int
}

var (
	grid     []string
	width    int
	height   int
	startX   int
	startY   int
	loopPath []Point
	inLoop   map[Point]bool
)

func parseInput(input string) {
	grid = make([]string, 0, 140)

	i := 0
	for i < len(input) {
		if input[i] == '\n' || input[i] == '\r' {
			i++
			continue
		}

		start := i
		for i < len(input) && input[i] != '\n' && input[i] != '\r' {
			i++
		}
		line := input[start:i]
		grid = append(grid, line)

		// Find start position
		for x := 0; x < len(line); x++ {
			if line[x] == 'S' {
				startX = x
				startY = len(grid) - 1
			}
		}
	}

	height = len(grid)
	width = len(grid[0])
}

func traceLoop() {
	loopPath = make([]Point, 0, 20000)
	inLoop = make(map[Point]bool, 20000)

	x, y := startX, startY
	loopPath = append(loopPath, Point{x, y})
	inLoop[Point{x, y}] = true

	// Determine initial direction from S
	dir := byte('>')
	x++

	for {
		loopPath = append(loopPath, Point{x, y})
		inLoop[Point{x, y}] = true

		pipe := grid[y][x]
		if pipe == 'S' {
			break
		}

		switch pipe {
		case '|':
			if dir == 'U' {
				y--
			} else {
				y++
			}
		case '-':
			if dir == '>' {
				x++
			} else {
				x--
			}
		case 'L':
			if dir == 'D' {
				x++
				dir = '>'
			} else {
				y--
				dir = 'U'
			}
		case 'J':
			if dir == 'D' {
				x--
				dir = '<'
			} else {
				y--
				dir = 'U'
			}
		case '7':
			if dir == 'U' {
				x--
				dir = '<'
			} else {
				y++
				dir = 'D'
			}
		case 'F':
			if dir == '<' {
				y++
				dir = 'D'
			} else {
				x++
				dir = '>'
			}
		}
	}
}

func pointInPolygon(px, py int) bool {
	// Ray casting algorithm - count intersections with edges
	count := 0
	n := len(loopPath)

	for i := 0; i < n-1; i++ {
		p1 := loopPath[i]
		p2 := loopPath[i+1]

		if p1.y <= py {
			if p2.y > py {
				// Upward crossing
				cross := (p2.x-p1.x)*(py-p1.y) - (px-p1.x)*(p2.y-p1.y)
				if cross > 0 {
					count++
				}
			}
		} else {
			if p2.y <= py {
				// Downward crossing
				cross := (p2.x-p1.x)*(py-p1.y) - (px-p1.x)*(p2.y-p1.y)
				if cross < 0 {
					count--
				}
			}
		}
	}

	return count != 0
}
