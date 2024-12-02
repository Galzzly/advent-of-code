package main

import (
	"aocli/utils/reader"
)

func doPartOne(input string) int {
	numlist := reader.RowIntsByLine(input)
	var res int
	for i, nums := range numlist {
		checknum := 1
		if nums[0] < nums[len(nums)-1] {
			checknum = -1
		}
		if !checkreport(nums, checknum) {
			continue
		}
		successful = append(successful, i)
		res++
	}
	return res
}

func checkreport(nums []int, checknum int) bool {
	for i := range len(nums) - 1 {
		diff := nums[i] - nums[i+1]
		diff *= checknum
		if diff > 3 || diff < 1 {
			return false
		}
	}

	return true
}
