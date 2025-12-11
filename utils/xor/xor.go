package xor

// SolveMinXOR solves the minimum XOR problem using Gaussian elimination over GF(2).
// Given a set of integers (buttons) and a target, finds the minimum number of buttons
// to XOR together to reach the target value.
// Returns the minimum count and whether it's possible.
func SolveMinXOR(buttons []int, target int, numBits int) (int, bool) {
	n := len(buttons)

	// Build augmented matrix: [buttons | target]
	// Each row is a button, each column is a bit position
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, numBits+1)
		for j := 0; j < numBits; j++ {
			matrix[i][j] = (buttons[i] >> j) & 1
		}
	}

	// Target as last column
	targetRow := make([]int, numBits+1)
	for j := 0; j < numBits; j++ {
		targetRow[j] = (target >> j) & 1
	}
	targetRow[numBits] = 1 // Mark as target

	// Gaussian elimination to row echelon form
	pivot := 0
	for col := 0; col < numBits && pivot < n; col++ {
		// Find pivot
		found := -1
		for row := pivot; row < n; row++ {
			if matrix[row][col] == 1 {
				found = row
				break
			}
		}

		if found == -1 {
			continue
		}

		// Swap rows
		if found != pivot {
			matrix[pivot], matrix[found] = matrix[found], matrix[pivot]
		}

		// Eliminate
		for row := 0; row < n; row++ {
			if row != pivot && matrix[row][col] == 1 {
				for c := 0; c <= numBits; c++ {
					matrix[row][c] ^= matrix[pivot][c]
				}
			}
		}

		// Also eliminate in target row
		if targetRow[col] == 1 {
			for c := 0; c <= numBits; c++ {
				targetRow[c] ^= matrix[pivot][c]
			}
		}

		pivot++
	}

	// Check if solution exists
	for j := 0; j < numBits; j++ {
		if targetRow[j] == 1 {
			return 0, false // No solution
		}
	}

	// Now use BFS/backtracking to find minimum combination
	// We know a solution exists, find the one with minimum buttons
	return findMinCombination(buttons, target), true
}

// findMinCombination uses BFS to find minimum number of buttons to XOR to target
func findMinCombination(buttons []int, target int) int {
	if target == 0 {
		return 0
	}

	visited := make(map[int]int) // state -> min presses to reach it
	visited[0] = 0

	queue := []struct{ state, presses int }{{0, 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.state == target {
			return curr.presses
		}

		for _, btn := range buttons {
			newState := curr.state ^ btn
			newPresses := curr.presses + 1

			if prevPresses, seen := visited[newState]; !seen || newPresses < prevPresses {
				visited[newState] = newPresses
				queue = append(queue, struct{ state, presses int }{newState, newPresses})
			}
		}
	}

	return -1 // Should not reach here if solution exists
}

// SolveMinXORAuto automatically determines the number of bits needed
func SolveMinXORAuto(buttons []int, target int) (int, bool) {
	maxVal := target
	for _, b := range buttons {
		if b > maxVal {
			maxVal = b
		}
	}

	// Calculate number of bits needed
	numBits := 0
	for maxVal > 0 {
		numBits++
		maxVal >>= 1
	}

	if numBits == 0 {
		numBits = 1
	}

	return SolveMinXOR(buttons, target, numBits)
}
