package main

import (
	"aocli/utils"
	"aocli/utils/xor"
	"strings"
)

func doPartOne(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	ans := 0
	for _, line := range lines {
		s := strings.Fields(line)
		goal := 0
		for i, c := range strings.Trim(s[0], "[]") {
			goal += utils.Ter(c == '#', utils.IntPow(2, i), 0)
		}

		buttons := s[1 : len(s)-1]
		buttonscore := make([]int, 0, len(buttons))
		for _, b := range buttons {
			s := strings.Split(strings.Trim(b, "()"), ",")
			button := 0
			for _, n := range s {
				button += utils.IntPow(2, utils.Atoi(n))
			}
			buttonscore = append(buttonscore, button)
		}

		// Use optimized XOR solver
		score, possible := xor.SolveMinXORAuto(buttonscore, goal)
		if !possible {
			score = len(buttonscore) // Fallback
		}

		ans += score
	}
	return ans
}
