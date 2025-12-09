package main

import "slices"

func doPartOne() int {
	// Create a copy for sorting
	sortedHands := make([]Hand, len(hands))
	copy(sortedHands, hands)
	
	slices.SortFunc(sortedHands, func(a, b Hand) int {
		// Compare hand strengths
		sa := getHandStrength(a, false)
		sb := getHandStrength(b, false)
		
		if sa != sb {
			return sa - sb
		}
		
		// Same strength, compare cards in order
		for i := 0; i < 5; i++ {
			if a.cards[i] != b.cards[i] {
				return a.cards[i] - b.cards[i]
			}
		}
		return 0
	})
	
	result := 0
	for i, h := range sortedHands {
		result += (i + 1) * h.bid
	}
	return result
}
