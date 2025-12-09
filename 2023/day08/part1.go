package main

func doPartOne(input string) int {
	current := "AAA"
	steps := 0
	instrLen := len(instructions)

	for current != "ZZZ" {
		node := network[current]
		if instructions[steps%instrLen] == 'L' {
			current = node.left
		} else {
			current = node.right
		}
		steps++
	}

	return steps
}
