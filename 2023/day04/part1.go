package main

import (
	"strings"
	"sync"
)

func doPartOne(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var wg sync.WaitGroup
	cards := make(chan scratchcard, len(lines))
	for _, line := range lines {
		wg.Add(1)
		go parseCard(line, cards, &wg)
	}

	go func() {
		wg.Wait()
		close(cards)
	}()

	var res int
	for card := range cards {
		switch card.winners {
		case 0:
			continue
		case 1:
			res++
		default:
			ret := 1
			for i := 2; i <= card.winners; i++ {
				ret *= 2
			}
			res += ret
		}
	}

	return res
}
