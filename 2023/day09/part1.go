package main

func doPartOne(input string) int {
	var sum int
	for _, nums := range sequences {
		next, _ := extrapolate(nums)
		sum += next
	}
	return sum
}
