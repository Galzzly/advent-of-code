package main

import (
	"aocli/utils/reader"
	"slices"
)

func doPartOne(input string) int {
	sections := reader.FileDoubleLine(input)
	rules = getRules(sections[0])
	numbers := getNumbers(sections[1])
	incorrect = make([][]int, 0)
	var res int
	for _, pages := range numbers {
		ok := true
		for i, x := range pages {
			for j, y := range pages {
				if i < j && slices.Contains[[]int](rules[x], y) {
					ok = false
				}
			}
		}
		if ok {
			res += pages[len(pages)/2]
		} else {
			incorrect = append(incorrect, pages)
		}
	}
	return res
}
