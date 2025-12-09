package main

func doPartTwo(input string) int {
	// Find cycle length for each starting node
	cycles := make([]int, 0, len(startNodes))
	instrLen := len(instructions)

	for _, start := range startNodes {
		current := start
		steps := 0

		for current[2] != 'Z' {
			node := network[current]
			if instructions[steps%instrLen] == 'L' {
				current = node.left
			} else {
				current = node.right
			}
			steps++
		}

		cycles = append(cycles, steps)
	}

	return lcmSlice(cycles)
}
