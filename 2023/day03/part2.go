package main

type point struct {
	x, y int
}

func doPartTwo() int {
	gearNums := make(map[point][]int)

	for y := 0; y < height; y++ {
		n := 0
		gears := make(map[point]bool)

		for x := 0; x <= width; x++ {
			if x < width && isDigit(grid[y][x]) {
				n = n*10 + int(grid[y][x]-'0')

				// Check neighbors for gears (*)
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						ny, nx := y+dy, x+dx
						if ny >= 0 && ny < height && nx >= 0 && nx < width {
							if grid[ny][nx] == '*' {
								gears[point{nx, ny}] = true
							}
						}
					}
				}
			} else if n > 0 {
				// Add this number to all adjacent gears
				for gear := range gears {
					gearNums[gear] = append(gearNums[gear], n)
				}
				n = 0
				gears = make(map[point]bool)
			}
		}
	}

	res := 0
	for _, nums := range gearNums {
		if len(nums) == 2 {
			res += nums[0] * nums[1]
		}
	}

	return res
}
