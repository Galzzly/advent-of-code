package minpresses

import "sort"

// SolveMinPresses finds the minimum number of button presses to reach target joltages
func SolveMinPresses(buttonPositions [][]int, targetJoltages []int) (int, bool) {
	numButtons := len(buttonPositions)
	numPositions := len(targetJoltages)

	// Sort buttons by number of positions affected
	type ButtonInfo struct {
		positions []int
	}

	buttons := make([]ButtonInfo, numButtons)
	for i := 0; i < numButtons; i++ {
		buttons[i] = ButtonInfo{positions: buttonPositions[i]}
	}

	sort.Slice(buttons, func(i, j int) bool {
		return len(buttons[i].positions) < len(buttons[j].positions)
	})

	// Use a fixed-size state for faster hashing (max 10 positions should be enough)
	type State struct {
		btnIdx int
		vals   [10]int
	}

	memo := make(map[State]int)

	var dp func(btnIdx int, current []int) int
	dp = func(btnIdx int, current []int) int {
		if btnIdx >= numButtons {
			for i := 0; i < numPositions; i++ {
				if current[i] != targetJoltages[i] {
					return 1e9
				}
			}
			return 0
		}

		// Create state
		var state State
		state.btnIdx = btnIdx
		copy(state.vals[:], current)

		if cached, ok := memo[state]; ok {
			return cached
		}

		btn := buttons[btnIdx]

		// Calculate max needed
		maxNeeded := 0
		for _, pos := range btn.positions {
			if pos < numPositions {
				need := targetJoltages[pos] - current[pos]
				if need > maxNeeded {
					maxNeeded = need
				}
			}
		}

		minResult := int(1e9)

		for presses := 0; presses <= maxNeeded; presses++ {
			newCurrent := make([]int, numPositions)
			copy(newCurrent, current)
			valid := true

			for _, pos := range btn.positions {
				if pos < numPositions {
					newCurrent[pos] += presses
					if newCurrent[pos] > targetJoltages[pos] {
						valid = false
						break
					}
				}
			}

			if valid {
				result := dp(btnIdx+1, newCurrent)
				if result < 1e9 {
					total := presses + result
					if total < minResult {
						minResult = total
					}
				}
			}
		}

		memo[state] = minResult
		return minResult
	}

	initial := make([]int, numPositions)
	result := dp(0, initial)

	if result >= 1e9 {
		return 0, false
	}
	return result, true
}
