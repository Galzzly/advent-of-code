package main

func doPartTwo() int {
	n := len(cardWins)
	cardCounts := make([]int, n)

	// Initialize each card with 1 copy
	for i := 0; i < n; i++ {
		cardCounts[i] = 1
	}

	res := 0
	for i := 0; i < n; i++ {
		count := cardCounts[i]
		wins := cardWins[i]

		// Add copies of following cards
		for j := 1; j <= wins && i+j < n; j++ {
			cardCounts[i+j] += count
		}

		res += count
	}

	return res
}
