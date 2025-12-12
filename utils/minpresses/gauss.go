package minpresses

// SolveMinPressesOptimized uses Gaussian elimination like the reference solution
func SolveMinPressesOptimized(buttonPositions [][]int, targetJoltages []int) (int, bool) {
	n := len(buttonPositions) // number of buttons
	m := len(targetJoltages)  // number of positions

	// Build coefficient matrix: matrix[i][j] = 1 if button j affects position i
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n+1) // +1 for the constant term
		for j := 0; j < n; j++ {
			affects := false
			for _, pos := range buttonPositions[j] {
				if pos == i {
					affects = true
					break
				}
			}
			if affects {
				matrix[i][j] = 1
			}
		}
		matrix[i][n] = targetJoltages[i] // constant term
	}

	// Perform Gaussian elimination
	pivotCols, reducedMatrix := gaussianElimination(matrix)
	if reducedMatrix == nil {
		return 0, false
	}

	// Identify pivot (constrained) and free (unconstrained) variables
	pivotSet := make(map[int]bool)
	for _, col := range pivotCols {
		pivotSet[col] = true
	}

	freeVars := []int{}
	for i := 0; i < n; i++ {
		if !pivotSet[i] {
			freeVars = append(freeVars, i)
		}
	}

	bestSolution := make([]int, n)
	bestSum := -1

	// Try different values for free variables and back-substitute
	var trySolution func(freeValues []int)
	trySolution = func(freeValues []int) {
		solution := make([]int, n)

		// Set free variables
		for i, varIdx := range freeVars {
			if i < len(freeValues) {
				solution[varIdx] = freeValues[i]
			}
		}

		// Back-substitute to find pivot variables
		for i := len(pivotCols) - 1; i >= 0; i-- {
			row := i
			col := pivotCols[i]
			total := reducedMatrix[row][n] // constant term

			for j := col + 1; j < n; j++ {
				total -= reducedMatrix[row][j] * solution[j]
			}

			if reducedMatrix[row][col] == 0 {
				return // Can't solve
			}

			if total%reducedMatrix[row][col] != 0 {
				return // Not an integer solution
			}

			val := total / reducedMatrix[row][col]
			if val < 0 {
				return // Negative solution not valid
			}

			solution[col] = val
		}

		// Verify the solution satisfies all equations
		for i := 0; i < m; i++ {
			total := 0
			for j := 0; j < n; j++ {
				if solution[j] > 0 {
					for _, pos := range buttonPositions[j] {
						if pos == i {
							total += solution[j]
							break
						}
					}
				}
			}
			if total != targetJoltages[i] {
				return // Solution doesn't satisfy this equation
			}
		}

		// Calculate total presses
		totalPresses := 0
		for _, val := range solution {
			totalPresses += val
		}

		// Keep the best solution
		if bestSum == -1 || totalPresses < bestSum {
			copy(bestSolution, solution)
			bestSum = totalPresses
		}
	}

	// Enumerate small values for free variables
	maxVal := 0
	for _, j := range targetJoltages {
		if j > maxVal {
			maxVal = j
		}
	}

	if len(freeVars) == 0 {
		trySolution([]int{})
	} else if len(freeVars) == 1 {
		limit := maxVal * 3
		for val := 0; val <= limit; val++ {
			if bestSum != -1 && val > bestSum {
				break
			}
			trySolution([]int{val})
		}
	} else if len(freeVars) == 2 {
		limit := maxVal
		if limit < 200 {
			limit = 200
		}
		for v1 := 0; v1 <= limit; v1++ {
			for v2 := 0; v2 <= limit; v2++ {
				if bestSum != -1 && v1+v2 > bestSum {
					continue
				}
				trySolution([]int{v1, v2})
			}
		}
	} else if len(freeVars) == 3 {
		for v1 := 0; v1 < 250; v1++ {
			for v2 := 0; v2 < 250; v2++ {
				for v3 := 0; v3 < 250; v3++ {
					if bestSum != -1 && v1+v2+v3 > bestSum {
						continue
					}
					trySolution([]int{v1, v2, v3})
				}
			}
		}
	} else if len(freeVars) == 4 {
		for v1 := 0; v1 < 30; v1++ {
			for v2 := 0; v2 < 30; v2++ {
				for v3 := 0; v3 < 30; v3++ {
					for v4 := 0; v4 < 30; v4++ {
						if bestSum != -1 && v1+v2+v3+v4 > bestSum {
							continue
						}
						trySolution([]int{v1, v2, v3, v4})
					}
				}
			}
		}
	} else {
		// Too many free variables - try just zeros
		trySolution(make([]int, len(freeVars)))
	}

	if bestSum == -1 {
		return 0, false
	}

	return bestSum, true
}

// gaussianElimination performs Gaussian elimination and returns pivot columns and reduced matrix
func gaussianElimination(matrix [][]int) ([]int, [][]int) {
	m := len(matrix)
	if m == 0 {
		return nil, nil
	}
	n := len(matrix[0]) - 1 // Exclude constant column

	// Make a copy of the matrix
	reducedMatrix := make([][]int, m)
	for i := range reducedMatrix {
		reducedMatrix[i] = make([]int, len(matrix[i]))
		copy(reducedMatrix[i], matrix[i])
	}

	pivotCols := []int{}
	row := 0

	for col := 0; col < n && row < m; col++ {
		// Find pivot
		pivotRow := -1
		for r := row; r < m; r++ {
			if reducedMatrix[r][col] != 0 {
				pivotRow = r
				break
			}
		}

		if pivotRow == -1 {
			continue // No pivot in this column
		}

		// Swap rows
		if pivotRow != row {
			reducedMatrix[row], reducedMatrix[pivotRow] = reducedMatrix[pivotRow], reducedMatrix[row]
		}

		pivotCols = append(pivotCols, col)

		// Eliminate below pivot
		for r := row + 1; r < m; r++ {
			if reducedMatrix[r][col] != 0 {
				// Scale and subtract
				factor := reducedMatrix[r][col] / reducedMatrix[row][col]
				for c := col; c <= n; c++ {
					reducedMatrix[r][c] -= factor * reducedMatrix[row][c]
				}
			}
		}

		row++
	}

	return pivotCols, reducedMatrix
}
