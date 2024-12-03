package main

import (
	"fmt"
	"regexp"
)

func doPartOne(input string) int {
	regex := regexp.MustCompile(`mul\(\d*,\d*\)`)
	var res int
	for _, f := range regex.FindAllString(input, -1) {
		var a, b int
		fmt.Sscanf(f, "mul(%d,%d)", &a, &b)
		res += a * b
	}

	return res
}
