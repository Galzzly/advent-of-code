package main

import (
	"aocli/utils"
	"strings"
)

func doPartOne(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	numbers := make([][]int, len(lines[0]))
	for x := 0; x < len(lines)-1; x++ {
		for y, v := range strings.Fields(lines[x]) {
			numbers[y] = append(numbers[y], utils.Atoi(v))
		}
	}
	ans := 0
	for y, v := range strings.Fields(lines[len(lines)-1]) {
		subAns := numbers[y][0]
		if v[0] == '+' {
			for i := 1; i < len(numbers[y]); i++ {
				subAns += numbers[y][i]
			}
		} else {
			for i := 1; i < len(numbers[y]); i++ {
				subAns *= numbers[y][i]
			}
		}
		ans += subAns
	}
	// mapper, rect := maps.MakeImagePointMapRect(strings.Split(strings.TrimSpace(input), "\n"))

	// ans := 0
	// for x := 0; x < rect.Max.X; x++ {
	// 	subAns := utils.Atoi(string(mapper[image.Point{x, 0}]))
	// 	if mapper[image.Point{x, rect.Max.Y - 1}] == '+' {
	// 		for y := 1; y < rect.Max.Y-1; y++ {
	// 			subAns += utils.Atoi(string(mapper[image.Point{x, y}]))
	// 		}
	// 		ans += subAns
	// 		continue
	// 	}
	// 	for y := 1; y < rect.Max.Y-1; y++ {
	// 		subAns *= utils.Atoi(string(mapper[image.Point{x, y}]))
	// 	}
	// 	ans += subAns
	// }
	return ans
}
