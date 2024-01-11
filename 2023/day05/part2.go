package main

import (
	"aocli/utils"
	"strings"
	"sync"
)

func doPartTwo(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n\n")
	var wg sync.WaitGroup
	wg.Add(2)
	seeds = Seeds{}
	maps = Maps{}
	go populateSeeds(lines[0], &wg)
	go populateMaps(lines[1:], &wg)
	wg.Wait()

	pairs := utils.ChunkSlice[int](seeds, 2)
	var R Ranges
	for _, pair := range pairs {
		R = append(R, []int{pair[0], pair[0] + pair[1] - 1})
	}
	res, _ := utils.MinMax(apply_range(R))
	return res
}

func apply_range(input Ranges) []int {
	A := []int{}
	for _, pair := range input {
		R := [][]int{pair}
		for _, tMap := range maps {
			AR := [][]int{}
			for _, ttMap := range tMap {
				src_end := ttMap.src + ttMap.size - 1
				NR := [][]int{}
				for _, v := range R {
					before := []int{v[0], utils.Min(v[1], ttMap.src-1)}
					inter := []int{utils.Biggest(v[0], ttMap.src), utils.Min(v[1], src_end)}
					after := []int{utils.Biggest(v[0], src_end), v[1]}
					if before[1] > before[0] {
						NR = append(NR, before)
					}
					if inter[1] > inter[0] {
						diff := ttMap.dest - ttMap.src
						AR = append(AR, []int{diff + inter[0], diff + inter[1]})
					}
					if after[1] > after[0] {
						NR = append(NR, after)
					}
				}
				R = NR
			}
			R = append(AR, R...)
		}
		for _, pair := range R {
			A = append(A, pair...)
		}
	}
	return A
}
