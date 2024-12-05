package main

import (
	"aocli/utils"
	_ "embed"
	"os"
	"strings"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var inputTest string

var rules map[int][]int
var reverse map[int][]int
var incorrect [][]int

func main() {
	// Check argv if we use test input or not
	if len(os.Args) > 1 && os.Args[1] == "test" {
		input = inputTest
	}

	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo()
	println(answer)
}

func getRules(rules string) map[int][]int {
	ruleset := make(map[int][]int, 0)
	reverse = make(map[int][]int, 0)
	for _, rule := range strings.Split(rules, "\n") {
		r := strings.Split(rule, "|")
		a, b := utils.Atoi(r[0]), utils.Atoi(r[1])
		ruleset[b] = append(ruleset[b], a)
		reverse[a] = append(reverse[a], b)
	}
	return ruleset
}

func getNumbers(numbers string) (res [][]int) {
	s := strings.Split(numbers, "\n")
	res = make([][]int, 0, len(s))
	for _, nums := range s {
		s2 := strings.Split(nums, ",")
		pages := make([]int, 0, len(s2))
		for _, n := range s2 {
			pages = append(pages, utils.Atoi(n))
		}
		res = append(res, pages)
	}
	return
}
