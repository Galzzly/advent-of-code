package main

func doPartOne() int {
	result := 1
	for _, r := range races {
		result *= getWays(r)
	}
	return result
}
