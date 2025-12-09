package main

import (
	_ "embed"
	"os"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var inputTest string

type Hand struct {
	cards [5]int
	bid   int
}

var hands []Hand

func main() {
	// Check argv if we use test input or not
	if len(os.Args) > 1 && os.Args[1] == "test" {
		input = inputTest
	}

	parseInput(input)

	answer := doPartOne()
	println(answer)

	answer = doPartTwo()
	println(answer)
}

func parseInput(input string) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	hands = make([]Hand, len(lines))
	
	for i, line := range lines {
		spaceIdx := strings.Index(line, " ")
		
		// Parse hand
		var h Hand
		for j := 0; j < 5; j++ {
			h.cards[j] = cardValue(line[j])
		}
		
		// Parse bid
		bid := 0
		for j := spaceIdx + 1; j < len(line); j++ {
			bid = bid*10 + int(line[j]-'0')
		}
		h.bid = bid
		
		hands[i] = h
	}
}

func cardValue(c byte) int {
	switch c {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		return int(c - '0')
	}
}

// Calculate hand strength from card counts
// Five of a kind: 6, Four of a kind: 5, Full house: 4, Three of a kind: 3, Two pair: 2, One pair: 1, High card: 0
func handStrength(counts [5]int) int {
	slices.Sort(counts[:])
	// counts is now sorted, highest count at the end
	switch counts[4] {
	case 5:
		return 6 // Five of a kind
	case 4:
		return 5 // Four of a kind
	case 3:
		if counts[3] == 2 {
			return 4 // Full house
		}
		return 3 // Three of a kind
	case 2:
		if counts[3] == 2 {
			return 2 // Two pair
		}
		return 1 // One pair
	default:
		return 0 // High card
	}
}

func getHandStrength(h Hand, jokerRule bool) int {
	counts := [5]int{}
	cardCounts := make(map[int]int)
	
	for _, card := range h.cards {
		cardCounts[card]++
	}
	
	// Handle jokers (value 11 in part 1, converted to 1 in part 2)
	jokers := 0
	if jokerRule {
		jokers = cardCounts[1]
		delete(cardCounts, 1)
	}
	
	// Fill counts array
	i := 0
	for _, count := range cardCounts {
		counts[i] = count
		i++
	}
	
	// Add jokers to the highest count
	if jokers > 0 {
		slices.Sort(counts[:])
		counts[4] += jokers
	}
	
	return handStrength(counts)
}
