package main

import (
	"aocli/utils"
	"strings"
	"sync"
)

func doPartOne(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n\n")
	var wg sync.WaitGroup
	wg.Add(2)
	seeds = Seeds{}
	maps = Maps{}
	go populateSeeds(lines[0], &wg)
	go populateMaps(lines[1:], &wg)
	wg.Wait()

	S := []int{}
	for _, seed := range seeds {
	nextmap:
		for _, tMap := range maps {
			for _, ttMap := range tMap {
				if ttMap.src <= seed && seed < ttMap.src+ttMap.size {
					seed = ttMap.dest + seed - ttMap.src
					continue nextmap
				}
			}
		}
		S = append(S, seed)
	}
	res, _ := utils.MinMax(S)
	return res
}
