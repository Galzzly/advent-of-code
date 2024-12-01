package main

import (
	"aocli/utils"
	"strings"
)

func doPartTwo(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	R := make(map[int]int)
	L := make([]int, len(lines))
	for i, line := range lines {
		n := strings.Split(line, "   ")
		L[i] = utils.Atoi(n[0])
		Rn := utils.Atoi(n[1])
		if _, ok := R[Rn]; !ok {
			R[Rn] = 0
		}
		R[Rn]++
	}
	var res int
	for _, n := range L {
		if v, ok := R[n]; ok {
			res += n * v
		}
	}
	return res
}
