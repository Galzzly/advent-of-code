package main

import "slices"

func doPartTwo() int {
	// Create a copy for sorting with joker rules
	sortedHands := make([]Hand, len(hands))
	for i, h := range hands {
		// Convert J (11) to joker (1)
		var newHand Hand
		newHand.bid = h.bid
		for j := 0; j < 5; j++ {
			if h.cards[j] == 11 {
				newHand.cards[j] = 1
			} else {
				newHand.cards[j] = h.cards[j]
			}
		}
		sortedHands[i] = newHand
	}
	
	slices.SortFunc(sortedHands, func(a, b Hand) int {
		// Compare hand strengths with joker rule
		sa := getHandStrength(a, true)
		sb := getHandStrength(b, true)
		
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
