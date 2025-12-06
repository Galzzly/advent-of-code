package main

import (
	"aocli/utils"
	"strings"
)

func doPartTwo(input string) int {
	lines := strings.Split(input, "\n")
	// Remove trailing empty line if present
	if len(lines) > 0 && len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	if len(lines) == 0 {
		return 0
	}

	ans := 0
	startCol := 0
	lastLine := len(lines) - 1
	lineWidth := len(lines[0])

	// Iterate through each column
	for x := 0; x < lineWidth; x++ {
		// Check if this column is blank (separator between problems)
		isBlank := true
		for y := 0; y < lastLine; y++ {
			if x < len(lines[y]) && lines[y][x] != ' ' {
				isBlank = false
				break
			}
		}

		if isBlank {
			// Process the problem from startCol to x-1
			operator := lines[lastLine][startCol]
			result := 0

			if operator == '+' {
				result = 0
			} else {
				result = 1
			}

			// Read each column in this problem
			for col := startCol; col < x; col++ {
				var numBuilder strings.Builder
				// Read vertically from top to bottom
				for y := 0; y < lastLine; y++ {
					if col < len(lines[y]) && lines[y][col] != ' ' {
						numBuilder.WriteByte(lines[y][col])
					}
				}

				if numBuilder.Len() > 0 {
					num := utils.Atoi(numBuilder.String())
					if operator == '+' {
						result += num
					} else {
						result *= num
					}
				}
			}

			ans += result
			startCol = x + 1
		}
	}

	// Handle last problem if it doesn't end with a blank column
	if startCol < lineWidth {
		operator := lines[lastLine][startCol]
		result := 0

		if operator == '+' {
			result = 0
		} else {
			result = 1
		}

		for col := startCol; col < lineWidth; col++ {
			var numBuilder strings.Builder
			for y := 0; y < lastLine; y++ {
				if col < len(lines[y]) && lines[y][col] != ' ' {
					numBuilder.WriteByte(lines[y][col])
				}
			}

			if numBuilder.Len() > 0 {
				num := utils.Atoi(numBuilder.String())
				if operator == '+' {
					result += num
				} else {
					result *= num
				}
			}
		}

		ans += result
	}

	return ans
}
