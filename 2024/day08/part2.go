package main

import (
	"image"
)

func doPartTwo() int {
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
					if dP1.Y*dP2.X == dP2.Y*dP1.X {
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
