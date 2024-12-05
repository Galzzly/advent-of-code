package main

import (
	"golang.org/x/exp/slices"
)

func doPartTwo() int {
	var res int
	for _, pages := range incorrect {
		good := []int{}
		Q := []int{}
		D := make(map[int]int)
		for i, v := range pages {
			D[v] = 0
			if _, ok := rules[v]; ok {
				for j, n := range pages {
					if i != j && slices.Contains(rules[v], n) {
						D[v]++
					}
				}
			}
		}
		for v := range D {
			if D[v] == 0 {
				Q = append(Q, v)
			}
		}
		for len(Q) > 0 {
			x := Q[0]
			Q = Q[1:]
			good = append(good, x)
			for _, y := range reverse[x] {
				if _, ok := D[y]; ok {
					D[y]--
					if D[y] == 0 {
						Q = append(Q, y)
					}
				}
			}
		}
		res += good[len(good)/2]
	}
	return res
}
