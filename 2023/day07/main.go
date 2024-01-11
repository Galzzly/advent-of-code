package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var inputTest string

func main() {
	// Check argv if we use test input or not
	if len(os.Args) > 1 && os.Args[1] == "test" {
		input = inputTest
	}

	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}

type Hands []Hand
type Hand struct {
	hand []int
	bid  int
}
type Score struct {
	hand     []int
	strength int
}

var (
	handscores = []Score{
		{[]int{5}, 10}, {[]int{1, 4}, 9}, {[]int{2, 3}, 8}, {[]int{1, 1, 3}, 7},
		{[]int{1, 2, 2}, 6}, {[]int{1, 1, 1, 2}, 5}, {[]int{1, 1, 1, 1, 1}, 4},
	}

	cardscore = map[rune]int{
		'1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7,
		'8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
	}
)

func getHands(input string, p2 bool) Hands {
	var hands Hands
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		var H string
		var B int
		fmt.Sscanf(line, "%s %d", &H, &B)
		hand := []int{}
		for _, c := range H {
			n := cardscore[c]
			if c == 'J' && p2 {
				n = 1
			}
			hand = append(hand, n)
		}
		hands = append(hands, Hand{hand, B})
	}
	return hands
}
