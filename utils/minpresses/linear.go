package minpresses

import (
	"math"
)

const (
	maxVars = 13
	eps     = 1e-8
)

// linear represents a linear equation: a_0*x_0 + a_1*x_1 + ... + a_N*x_N + b = 0
type linear struct {
	a [maxVars]float64
	b float64
}

// variable represents a button press count variable
type variable struct {
	expr linear // expression for constrained variables
	free bool   // true if this is a free variable
	val  int    // value for free variables during enumeration
	max  int    // maximum value this variable can take
}

// extract solves for variable at index in the linear equation
// Returns the expression and true if successful
func extract(lin linear, index int) (linear, bool) {
	a := -lin.a[index]
	if math.Abs(a) < eps {
		return linear{}, false
	}

	r := linear{b: lin.b / a}
	for i := 0; i < maxVars; i++ {
		if i != index {
			r.a[i] = lin.a[i] / a
		}
	}
	return r, true
}

// substitute replaces variable at index with its expression in the linear equation
func substitute(lin linear, index int, expr linear) linear {
	r := linear{}

	a := lin.a[index]
	lin.a[index] = 0

	for i := 0; i < maxVars; i++ {
		r.a[i] = lin.a[i] + a*expr.a[i]
	}
	r.b = lin.b + a*expr.b
	return r
}

// eval evaluates a variable given the current free variable values
func eval(v variable, vals [maxVars]int) float64 {
	if v.free {
		return float64(v.val)
	}

	x := v.expr.b
	for i := 0; i < maxVars; i++ {
		x += v.expr.a[i] * float64(vals[i])
	}
	return x
}

// SolveMinPressesLinear solves using Gaussian elimination with recursive enumeration
func SolveMinPressesLinear(buttonPositions [][]int, targetJoltages []int) (int, bool) {
	numButtons := len(buttonPositions)
	numPositions := len(targetJoltages)

	if numButtons > maxVars {
		return 0, false // Too many variables for this approach
	}

	// Initialize variables
	vars := make([]variable, numButtons)
	for i := range vars {
		vars[i].max = math.MaxInt
	}

	// Build linear equations: one for each position
	// Each equation: sum of button presses that affect this position = target joltage
	eqs := make([]linear, numPositions)
	for i, jolt := range targetJoltages {
		eq := linear{b: float64(-jolt)}
		for j, positions := range buttonPositions {
			// Check if button j affects position i
			affects := false
			for _, pos := range positions {
				if pos == i {
					affects = true
					break
				}
			}
			if affects {
				eq.a[j] = 1
				vars[j].max = min(vars[j].max, jolt)
			}
		}
		eqs[i] = eq
	}

	// Gaussian elimination: identify free vs constrained variables
	for i := range vars {
		vars[i].free = true

		// Try to extract this variable from an equation
		for eqIdx := range eqs {
			if expr, ok := extract(eqs[eqIdx], i); ok {
				// Variable i is constrained by this equation
				vars[i].free = false
				vars[i].expr = expr

				// Substitute this expression into all other equations
				for j := range eqs {
					eqs[j] = substitute(eqs[j], i, expr)
				}

				break
			}
		}
	}

	// Collect indices of free variables
	free := []int{}
	for i, v := range vars {
		if v.free {
			free = append(free, i)
		}
	}

	// Recursively enumerate free variables
	best, found := evalRecursive(vars, free, 0, math.MaxInt)
	return best, found
}

// evalRecursive recursively tries all combinations of free variable values
func evalRecursive(vars []variable, free []int, index int, currentBest int) (int, bool) {
	if index == len(free) {
		// Base case: all free variables assigned, compute constrained variables
		vals := [maxVars]int{}
		total := 0

		// Evaluate variables in reverse order (back-substitution)
		for i := len(vars) - 1; i >= 0; i-- {
			x := eval(vars[i], vals)

			// Check if result is valid (non-negative integer)
			if x < -eps || math.Abs(x-math.Round(x)) > eps {
				return 0, false
			}

			vals[i] = int(math.Round(x))
			total += vals[i]
		}

		return total, true
	}

	// Recursive case: try all values for current free variable
	best, found := math.MaxInt, false
	freeIdx := free[index]

	// Compute current partial sum to enable pruning
	partialSum := 0
	for i := 0; i < index; i++ {
		partialSum += vars[free[i]].val
	}

	for x := 0; x <= vars[freeIdx].max; x++ {
		// Prune if partial sum already exceeds current best
		if partialSum+x >= currentBest {
			break
		}

		vars[freeIdx].val = x
		total, ok := evalRecursive(vars, free, index+1, min(currentBest, best))

		if ok {
			found = true
			best = min(best, total)
			currentBest = min(currentBest, best)
		}
	}

	return best, found
}
