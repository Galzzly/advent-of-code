package main

import (
	"aocli/utils"
	"strings"
)

func doPartOne() int {
	ans := 0
	for _, v := range strings.Split(sections[1], "\n") {
		ingr := utils.Atoi(v)
		for _, r := range ranges {
			if ingr >= r.start && ingr <= r.end {
				ans++
				break
			}
		}
	}
	return ans
}
