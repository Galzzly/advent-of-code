package main

import (
	"aocli/utils"
	_ "embed"
	"os"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var inputTest string

type combo struct {
	num, count int
}

var seen = make(map[combo]int, 0)
var nums = []int{}

func main() {
	// Check argv if we use test input or not
	if len(os.Args) > 1 && os.Args[1] == "test" {
		input = inputTest
	}

	input = strings.ReplaceAll(input, "\n", "")
	for _, n := range strings.Split(input, " ") {
		nums = append(nums, utils.Atoi(n))
	}

	answer := solveAll(nums, 25)
	println(answer)

	answer = solveAll(nums, 75)
	println(answer)
}

func solveAll(nums []int, count int) int {
	var res int
	for _, n := range nums {
		res += solve(n, count)
	}
	return res
}

func solve(num int, count int) int {
	if _, ok := seen[combo{num, count}]; ok {
		return seen[combo{num, count}]
	}
	if count == 0 {
		return 1
	}
	if num == 0 {
		return solve(1, count-1)
	}
	s := strconv.Itoa(num)
	if len(s)%2 == 0 {
		return solve(utils.Atoi(s[:len(s)/2]), count-1) + solve(utils.Atoi(s[len(s)/2:]), count-1)
	}
	res := solve(num*2024, count-1)
	seen[combo{num, count}] = res
	return res
}
