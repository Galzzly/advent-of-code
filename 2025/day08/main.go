package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var inputTest string

var (
	junctions []coords
	distances []distance
)

func main() {
	// Check argv if we use test input or not
	isTest := len(os.Args) > 1 && os.Args[1] == "test"
	if isTest {
		input = inputTest
	}

	// Parse input once
	parseInput(input)

	k := 1000
	if isTest {
		k = 10
	}

	answer := doPartOne(k)
	println(answer)

	answer = doPartTwo()
	println(answer)
}

func parseInput(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	junctions = make([]coords, len(lines))
	for i, line := range lines {
		fmt.Sscanf(line, "%d,%d,%d", &junctions[i].x, &junctions[i].y, &junctions[i].z)
	}

	// Calculate all pairwise distances
	n := len(junctions)
	distances = make([]distance, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		j1 := junctions[i]
		for j := i + 1; j < n; j++ {
			j2 := junctions[j]
			dx := j1.x - j2.x
			dy := j1.y - j2.y
			dz := j1.z - j2.z
			if dx < 0 {
				dx = -dx
			}
			if dy < 0 {
				dy = -dy
			}
			if dz < 0 {
				dz = -dz
			}
			d := dx*dx + dy*dy + dz*dz
			distances = append(distances, distance{d: d, a: i, b: j})
		}
	}
}
