package main

import (
	"fmt"
	"regexp"
)

func doPartTwo(input string) int {
	regex := regexp.MustCompile(`mul\(\d*,\d*\)|do\(\)|don't\(\)`)
	var res int
	do := true
	for _, m := range regex.FindAllString(input, -1) {
		switch m {
		case "do()":
			do = true
		case "don't()":
			do = false
		default:
			if do {
				var a, b int
				fmt.Sscanf(m, "mul(%d,%d)", &a, &b)
				res += a * b
			}
		}
	}
	return res
}
