package main

import (
	"aocli/utils/reader"
)

func doPartOne(input string) int {
	numlist := reader.RowIntsByLine(input)
	var res int
	for _, nums := range numlist {
		extr := make([][]int, 0)
		extr = append(extr, nums)
		for {
			next := []int{}
			for i := 0; i < len(extr[len(extr)-1])-1; i++ {
				next = append(next, extr[len(extr)-1][i+1]-extr[len(extr)-1][i])
			}
			extr = append(extr, next)
			allzero := true
			for _, v := range next {
				if v != 0 {
					allzero = false
				}
			}
			if allzero {
				break
			}
		}
		for i := len(extr) - 2; i >= 0; i-- {
			res += extr[i][len(extr[i])-1]
		}
	}
	return res
}
