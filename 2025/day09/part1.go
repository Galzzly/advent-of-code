package main

import (
	"aocli/utils"
)

func doPartOne(input string) int {
	coords := parseLines(input)
	ans := 0
	for i, c := range coords {
		for _, t := range coords[i+1:] {
			w := utils.Abs(t.X-c.X) + 1
			h := utils.Abs(t.Y-c.Y) + 1
			a := w * h
			// fmt.Println(c, t, a)
			ans = utils.Ter(ans > a, ans, a)
		}
	}
	return ans
}
