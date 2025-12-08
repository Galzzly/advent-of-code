package main

import (
	"strings"
)

func doPartTwo(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	res := 0

	for _, line := range lines {
		// Skip "Game X: " prefix
		colonIdx := strings.Index(line, ": ")
		sets := strings.Split(line[colonIdx+2:], "; ")

		maxRed, maxGreen, maxBlue := 0, 0, 0

		for _, set := range sets {
			cubes := strings.Split(set, ", ")
			for _, cube := range cubes {
				spaceIdx := strings.Index(cube, " ")
				num := 0
				for i := 0; i < spaceIdx; i++ {
					num = num*10 + int(cube[i]-'0')
				}

				color := cube[spaceIdx+1:]
				switch color[0] {
				case 'r': // red
					if num > maxRed {
						maxRed = num
					}
				case 'g': // green
					if num > maxGreen {
						maxGreen = num
					}
				case 'b': // blue
					if num > maxBlue {
						maxBlue = num
					}
				}
			}
		}

		res += maxRed * maxGreen * maxBlue
	}

	return res
}
