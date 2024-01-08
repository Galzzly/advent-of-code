package main

import (
	"aocli/utils"
	"aocli/utils/maps"
	"image"
	"strings"
	"unicode"
)

var Delta = []int{-1, 0, 1}

func doPartOne(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	G, R := maps.MakeImagPointMapRect(lines)
	row := R.Max.Y
	col := R.Max.X
	var res int
	for Y := R.Min.Y; Y <= row; Y++ {
		n := 0
		has_part := false
		for X := R.Min.X; X <= col+1; X++ {
			if unicode.IsDigit(G[image.Point{X, Y}]) {
				n = n*10 + utils.Atoi(string(G[image.Point{X, Y}]))
				for _, YY := range Delta {
					for _, XX := range Delta {
						if 0 <= Y+YY && Y+YY <= row && (0 <= X+XX && X+XX <= col) {
							p := image.Point{X + XX, Y + YY}
							c := G[p]
							if !unicode.IsDigit(c) && c != '.' {
								has_part = true
							}
						}
					}
				}
			} else if n > 0 {
				if has_part {
					res += n
				}
				n = 0
				has_part = false
			}
		}
	}
	return res
}
