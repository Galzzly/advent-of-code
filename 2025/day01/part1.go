package main

import (
	"aocli/utils"
	"strings"
)

func doPartOne(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	r = r.Move(50)
	ans := 0
	for _, line := range lines {
		num := utils.Atoi(line[1:])
		if line[0] == 'L' {
			num = -num
		}
		r = r.Move(num)
		if r.Value == 0 {
			ans++
		}
	}
	return ans
}
