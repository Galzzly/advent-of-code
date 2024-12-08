package main

import (
	"aocli/utils"
	"image"
	"strings"
)

func doPartOne(input string) int {
	lines := strings.Split(input, "\n")
	makeMap(lines)
	var res = make(map[image.Point]interface{})
	for P := range mapper {
		for _, a := range antennas {
			for i, a1 := range a {
				for j, a2 := range a {
					if i == j {
						continue
					}
					dP1 := P.Sub(a1)
					dP2 := P.Sub(a2)
					d1 := utils.Abs(dP1.X) + utils.Abs(dP1.Y)
					d2 := utils.Abs(dP2.X) + utils.Abs(dP2.Y)
					if (d1 == d2*2 || d2 == d1*2) && (dP1.Y*dP2.X == dP2.Y*dP1.X) {
						if _, ok := res[P]; !ok {
							res[P] = nil
						}
					}
				}
			}
		}
	}

	return len(res)
}
