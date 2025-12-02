package main

import (
	"aocli/utils"
	"strings"
)

func doPartTwo(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// Calculate where we need to move to get to position 50
	// Since r.Value equals the position, we can calculate the offset
	currentPos := r.Value.(int)
	offset := (50 - currentPos + 100) % 100
	r = r.Move(offset)

	ans := 0
	for _, line := range lines {
		num := utils.Atoi(line[1:])
		for num > 100 {
			num -= 100
			ans++
		}
		if line[0] == 'L' {
			if r.Value.(int) > 0 && r.Value.(int)-num < 0 {
				ans++
			}
			r = r.Move(-num)
			if r.Value == 0 {
				ans++
			}
			continue
		}
		if r.Value.(int)+num > 100 {
			ans++
		}
		r = r.Move(num)
		if r.Value == 0 {
			ans++
		}
	}
	return ans
}
