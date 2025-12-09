package main

func doPartOne() int {
	res := 0

	for y := 0; y < height; y++ {
		n := 0
		hasPart := false

		for x := 0; x <= width; x++ {
			if x < width && isDigit(grid[y][x]) {
				n = n*10 + int(grid[y][x]-'0')

				// Check neighbors for symbols
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						ny, nx := y+dy, x+dx
						if ny >= 0 && ny < height && nx >= 0 && nx < width {
							c := grid[ny][nx]
							if !isDigit(c) && c != '.' {
								hasPart = true
							}
						}
					}
				}
			} else if n > 0 {
				if hasPart {
					res += n
				}
				n = 0
				hasPart = false
			}
		}
	}

	return res
}
