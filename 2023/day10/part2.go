package main

func doPartTwo(input string) int {
	count := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := Point{x, y}
			if !inLoop[p] && pointInPolygon(x, y) {
				count++
			}
		}
	}

	return count
}
