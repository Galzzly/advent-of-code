package main

import (
	"aocli/utils"
	_ "embed"
	"os"
	"strings"
	"sync"
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

type scratchcards map[int]int
type scratchcard struct {
	id, winners int
}

func parseCards(input string) scratchcards {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var wg sync.WaitGroup
	cards := make(chan scratchcard, len(lines))
	scratchcards := make(map[int]int)
	for _, line := range lines {
		wg.Add(1)
		go parseCard(line, cards, &wg)
	}

	go func() {
		wg.Wait()
		close(cards)
	}()

	for card := range cards {
		scratchcards[card.id] = card.winners
	}
	return scratchcards
}

func parseCard(line string, cards chan scratchcard, wg *sync.WaitGroup) {
	defer wg.Done()
	s := strings.Split(line, ": ")
	numbers := strings.Split(s[1], "|")
	winners := getNumbers(numbers[0])
	nums := getNumbers(numbers[1])
	var ret scratchcard
	ret.id = utils.Atoi(strings.Fields(s[0])[1])
	ret.winners = 0
	for _, n := range nums {
		if checkNum(winners, n) {
			ret.winners++
		}
	}
	cards <- ret
}

func getNumbers(nums string) []int {
	s := strings.Fields(strings.TrimSpace(nums))
	ret := make([]int, 0)
	for _, n := range s {
		ret = append(ret, utils.Atoi(n))
	}
	return ret
}

func checkNum(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}
	return false
}
