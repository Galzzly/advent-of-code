package main

import (
	"aocli/utils"
	"strings"
	"sync"
)

func doPartOne(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var wg sync.WaitGroup
	ch := make(chan int, len(lines))

	for _, line := range lines {
		wg.Add(1)
		go func(line string, wg *sync.WaitGroup, ch chan int) {
			defer wg.Done()
			var res int
			digits := findNums(line)
			res = utils.Atoi(digits[0].num + digits[1].num)
			ch <- res
		}(line, &wg, ch)
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
