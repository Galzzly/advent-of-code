package main

import (
	"aocli/utils"
	_ "embed"
	"os"
	"strconv"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var inputTest string

func main() {
	// Check argv if we use test input or not
	if len(os.Args) > 1 && os.Args[1] == "test" {
		input = inputTest
	}

	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}

func checkOperation(result int, nums []int, P2 bool) bool {
	if len(nums) == 1 {
		return result == nums[0]
	}
	if checkOperation(result, append([]int{(nums[0] + nums[1])}, nums[2:]...), P2) {
		return true
	}
	if checkOperation(result, append([]int{(nums[0] * nums[1])}, nums[2:]...), P2) {
		return true
	}
	if P2 && checkOperation(result, append([]int{(utils.Atoi(strconv.Itoa(nums[0]) + strconv.Itoa(nums[1])))}, nums[2:]...), P2) {
		return true
	}
	return false
}
