package main

import (
	"strings"
)

func doPartOne(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	res := 0

	for _, line := range lines {
		// Parse game ID and sets
		colonIdx := strings.Index(line, ": ")
		gameID := 0
		for i := 5; i < colonIdx; i++ { // "Game " is 5 chars
			gameID = gameID*10 + int(line[i]-'0')
		}

		// Check if game is valid
		valid := true
		sets := strings.Split(line[colonIdx+2:], "; ")

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
					if num > 12 {
						valid = false
					}
				case 'g': // green
					if num > 13 {
						valid = false
					}
				case 'b': // blue
					if num > 14 {
						valid = false
					}
				}

				if !valid {
					break
				}
			}
			if !valid {
				break
			}
		}

		if valid {
			res += gameID
		}
	}

	return res
}
