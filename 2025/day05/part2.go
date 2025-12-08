package main

func doPartTwo() int {
	merged := mergeRanges()

	ans := 0
	for _, r := range merged {
		ans += r.end - r.start + 1
	}

	return ans
}
