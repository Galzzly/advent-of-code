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

type Seeds []int
type Maps []Map
type Map []mapper
type mapper struct {
	dest, src, size int
}
type Ranges [][]int

var (
	seeds Seeds
	maps  Maps
)

func populateSeeds(line string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, seed := range strings.Fields(line)[1:] {
		seeds = append(seeds, utils.Atoi(seed))
	}
}

func populateMaps(input []string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, lines := range input {
		tMap := []mapper{}
		for _, line := range strings.Split(lines, "\n")[1:] {
			s := strings.Fields(line)
			tMap = append(tMap, mapper{utils.Atoi(s[0]), utils.Atoi(s[1]), utils.Atoi(s[2])})
		}
		maps = append(maps, tMap)
	}
}
