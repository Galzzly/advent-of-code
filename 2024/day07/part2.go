package main

import (
	"aocli/utils"
	"aocli/utils/reader"
	"strings"
)

func doPartTwo(input string) int {
	lines := reader.FileLineByLine(input)
	var res int
	for _, line := range lines {
		S := strings.Split(line, ": ")
		result := utils.Atoi(S[0])
		nums := []int{}
		for _, num := range strings.Split(S[1], " ") {
			nums = append(nums, utils.Atoi(num))
		}
		if checkOperation(result, nums, true) {
			res += result
		}
	}
	return res
}
