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

type Node struct {
	left  string
	right string
}

var (
	instructions string
	network      map[string]Node
	startNodes   []string
)

func parseInput(input string) {
	i := 0
	// Skip whitespace
	for i < len(input) && (input[i] == ' ' || input[i] == '\n' || input[i] == '\r') {
		i++
	}

	// Parse instructions
	start := i
	for i < len(input) && input[i] != '\n' {
		i++
	}
	instructions = input[start:i]

	// Skip blank lines
	for i < len(input) && (input[i] == '\n' || input[i] == '\r') {
		i++
	}

	// Parse network
	network = make(map[string]Node)
	startNodes = make([]string, 0, 6)

	for i < len(input) {
		if input[i] == '\n' || input[i] == '\r' {
			i++
			continue
		}

		// Parse: "AAA = (BBB, CCC)"
		name := input[i : i+3]
		i += 7 // Skip " = ("
		left := input[i : i+3]
		i += 5 // Skip ", "
		right := input[i : i+3]
		i += 4 // Skip ")\n"

		network[name] = Node{left: left, right: right}
		if name[2] == 'A' {
			startNodes = append(startNodes, name)
		}
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func lcmSlice(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result = lcm(result, nums[i])
	}
	return result
}
