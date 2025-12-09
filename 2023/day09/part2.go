package main

func doPartTwo(input string) int {
	var sum int
	for _, nums := range sequences {
		_, prev := extrapolate(nums)
		sum += prev
	}
	return sum
}
