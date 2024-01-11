package main

func doPartTwo(input string) int {
	scratchcards := parseCards(input)
	cardpile := make(map[int]int, len(scratchcards))
	for i := 1; i <= len(scratchcards); i++ {
		cardpile[i] = 1
	}
	var res int
	for i := 1; i <= len(scratchcards); i++ {
		for j := 1; j <= scratchcards[i]; j++ {
			cardpile[i+j] += cardpile[i]
		}
		res += cardpile[i]
	}

	return res
}
