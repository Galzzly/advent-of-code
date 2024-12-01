package main

import (
	"aocli/utils"
	"sort"
	"strings"
)

func doPartOne(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	nums := make([][]int, 2)
	nums[0] = make([]int, len(lines))
	nums[1] = make([]int, len(lines))
	for i, line := range lines {
		n := strings.Split(line, "   ")
		nums[0][i] = utils.Atoi(n[0])
		nums[1][i] = utils.Atoi(n[1])
	}
	sort.Ints(nums[0])
	sort.Ints(nums[1])
	var res int
	for i := range len(lines) {
		res += utils.Abs(nums[0][i] - nums[1][i])
	}
	return res
}
