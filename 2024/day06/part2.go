package main

import (
	"aocli/utils/reader"
	"image"
	"sync"
)

type dirVis [4]bool

func doPartTwo(input string) int {
	mapper, start := makeMap(reader.FileLineByLine(input))
	var wg sync.WaitGroup
	var ch = make(chan int, len(mapper))
	for i := range mapper {
		if mapper[i] {
			continue
		}
		wg.Add(1)
		go func(i image.Point, wg *sync.WaitGroup, ch chan int) {
			defer wg.Done()
			dir := 0
			P := start
			visited := make(map[image.Point]dirVis)
			for {
				if _, ok := visited[P]; ok && visited[P][dir] {
					ch <- 1
					break
				}
				temp := visited[P]
				temp[dir] = true
				visited[P] = temp
				NP := P.Add(delta[dir])
				if _, ok := mapper[NP]; !ok {
					ch <- 0
					break
				}
				if mapper[NP] || NP == i {
					dir = (dir + 1) % 4
					continue
				}
				P = NP
			}
		}(i, &wg, ch)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	var res int
	for out := range ch {
		res += out
	}
	return res
}
