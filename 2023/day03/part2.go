package main

import (
	"aocli/utils"
	"aocli/utils/maps"
	"image"
	"runtime"
	"strings"
	"unicode"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func doPartTwo(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	G, R := maps.MakeImagPointMapRect(lines)
	row := R.Max.Y
	col := R.Max.X
	nums := make(map[image.Point][]int, 0)
	for Y := R.Min.Y; Y <= row; Y++ {
		gears := map[image.Point]bool{}
		n := 0
		for X := R.Min.X; X <= col+1; X++ {
			if unicode.IsDigit(G[image.Point{X, Y}]) {
				n = n*10 + utils.Atoi(string(G[image.Point{X, Y}]))
				for _, YY := range Delta {
					for _, XX := range Delta {
						if 0 <= Y+YY && Y+YY <= row && (0 <= X+XX && X+XX <= col) {
							p := image.Point{X + XX, Y + YY}
							c := G[p]
							if c == '*' {
								gears[p] = true
							}
						}
					}
				}
			} else if n > 0 {
				for gear := range gears {
					nums[gear] = append(nums[gear], n)
				}
				n = 0
				gears = map[image.Point]bool{}
			}
		}
	}
	var res int
	for _, v := range nums {
		if len(v) == 2 {
			res += utils.MultiplyArray(v)
		}
	}
	return res
}
