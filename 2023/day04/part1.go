package main

func doPartOne() int {
	res := 0

	for _, wins := range cardWins {
		if wins > 0 {
			// Points are 2^(wins-1), which is 1 << (wins-1)
			res += 1 << (wins - 1)
		}
	}

	return res
}
