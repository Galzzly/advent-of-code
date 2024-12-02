package main

import (
	"aocli/utils"
	"aocli/utils/reader"
)

func doPartTwo(input string) int {
	numlist := reader.RowIntsByLine(input)
	var res int
	for i, nums := range numlist {
		if utils.FindInSlice(successful, i) {
			res++
			continue
		}
		checknum := 1
		if nums[0] < nums[len(nums)-1] {
			checknum = -1
		}
		if checkdampenedreport(nums, checknum) {
			res++
		}
	}
	return res
}

func checkdampenedreport(nums []int, checknum int) bool {
	for i := range len(nums) {
		if checkreport(removeId(nums, i), checknum) {
			return true
		}
	}
	return false
}

func removeId(nums []int, id int) []int {
	newnums := make([]int, 0)
	newnums = append(newnums, nums[:id]...)
	return append(newnums, nums[id+1:]...)
}
