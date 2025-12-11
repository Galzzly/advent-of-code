package main

import (
	"aocli/utils"
	"aocli/utils/minpresses"
	"strings"
	"sync"
)

func doPartTwo(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	type result struct {
		index int
		score int
		line  string
	}

	results := make(chan result, len(lines))
	var wg sync.WaitGroup

	// Process each line in parallel
	for idx, line := range lines {
		wg.Add(1)
		go func(index int, line string) {
			defer wg.Done()

			s := strings.Fields(line)

			// Parse buttons (middle section)
			buttonstring := s[1 : len(s)-1]
			buttonslen := len(buttonstring)
			buttonPositions := make([][]int, 0, buttonslen)
			for _, b := range buttonstring {
				s := strings.Split(strings.Trim(b, "()"), ",")
				positions := make([]int, 0, len(s))
				for _, n := range s {
					positions = append(positions, utils.Atoi(n))
				}
				buttonPositions = append(buttonPositions, positions)
			}

			// Parse target joltages (last section)
			targetJoltages := []int{}
			for _, c := range strings.Split(strings.Trim(s[len(s)-1], "{}"), ",") {
				targetJoltages = append(targetJoltages, utils.Atoi(c))
			}

			// Solve for minimum presses to reach joltages
			score, possible := minpresses.SolveMinPresses(buttonPositions, targetJoltages)
			if !possible {
				score = 0 // No solution
			}

			results <- result{index: index, score: score, line: line}
		}(idx, line)
	}

	// Close results channel when all goroutines are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results in order
	collected := make([]result, len(lines))
	for r := range results {
		collected[r.index] = r
	}

	// Sum results
	ans := 0
	for _, r := range collected {
		ans += r.score
	}

	return ans
}
