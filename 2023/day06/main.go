package main

import (
	_ "embed"
	"math"
	"os"
	"strings"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var inputTest string

type Race struct {
	time, distance int
}

var races []Race
var bigRace Race

func main() {
	// Check argv if we use test input or not
	if len(os.Args) > 1 && os.Args[1] == "test" {
		input = inputTest
	}

	parseInput(input)

	answer := doPartOne()
	println(answer)

	answer = doPartTwo()
	println(answer)
}

func parseInput(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Parse times and distances
	times := parseNumbers(lines[0])
	distances := parseNumbers(lines[1])

	// Part 1: individual races
	races = make([]Race, len(times))
	for i := range races {
		races[i] = Race{times[i], distances[i]}
	}

	// Part 2: concatenated race
	bigRace = Race{concatenateNumbers(times), concatenateNumbers(distances)}
}

func parseNumbers(line string) []int {
	var nums []int
	i := 0
	for i < len(line) {
		if line[i] >= '0' && line[i] <= '9' {
			num := 0
			for i < len(line) && line[i] >= '0' && line[i] <= '9' {
				num = num*10 + int(line[i]-'0')
				i++
			}
			nums = append(nums, num)
		} else {
			i++
		}
	}
	return nums
}

func concatenateNumbers(nums []int) int {
	result := 0
	for _, n := range nums {
		// Count digits in n
		temp := n
		multiplier := 1
		for temp > 0 {
			multiplier *= 10
			temp /= 10
		}
		result = result*multiplier + n
	}
	return result
}

// Mathematical solution using quadratic formula
// Distance d = h * (t - h) where h is hold time, t is total time
// Need: h * (t - h) > record
// Rearrange: -h² + t*h - record > 0
// Solve: h² - t*h + record < 0
// Using quadratic formula: h = (t ± sqrt(t² - 4*record)) / 2
func getWays(r Race) int {
	t := float64(r.time)
	d := float64(r.distance)

	// Discriminant
	disc := t*t - 4*d
	if disc < 0 {
		return 0
	}

	sqrtDisc := math.Sqrt(disc)

	// Two solutions from quadratic formula
	h1 := (t - sqrtDisc) / 2
	h2 := (t + sqrtDisc) / 2

	// We need integer hold times strictly greater than the record
	// Round up h1, round down h2
	minHold := int(math.Floor(h1 + 1))
	maxHold := int(math.Ceil(h2 - 1))

	// Edge case: if h1 is exactly an integer, we need strictly greater
	if h1 == math.Floor(h1) {
		minHold = int(h1) + 1
	}

	// Edge case: if h2 is exactly an integer, we need strictly less
	if h2 == math.Ceil(h2) {
		maxHold = int(h2) - 1
	}

	if maxHold < minHold {
		return 0
	}

	return maxHold - minHold + 1
}
