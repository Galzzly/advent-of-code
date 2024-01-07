package main

import (
	"aocli/utils"
	"slices"
	"strings"
	"sync"
)

func doPartTwo(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var wg sync.WaitGroup
	ch := make(chan int, len(lines))
	for _, line := range lines {
		wg.Add(1)
		go func(line string, wg *sync.WaitGroup, ch chan int) {
			defer wg.Done()
			var res int
			digits := findNums(line)
			words := findWords(line)
			if digits[0].id == -1 || (words[0].id != -1 && words[0].id < digits[0].id) {
				digits[0] = words[0]
			}
			if words[1].id != -1 && words[1].id > digits[1].id {
				digits[1] = words[1]
			}
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

func findWords(line string) Digits {
	result := [2]digit{{-1, ""}, {-1, ""}}
	charMap := make(map[int]string, 0)
	var charIdx []int

	for word, num := range digitMap {
		first := strings.Index(line, word)
		last := strings.LastIndex(line, word)
		if first != -1 {
			charIdx = append(charIdx, first)
			charMap[first] = num
		}
		if last != -1 {
			charIdx = append(charIdx, last)
			charMap[last] = num
		}
	}

	if len(charIdx) > 0 {
		slices.Sort(charIdx)
		result[0] = digit{charIdx[0], charMap[charIdx[0]]}
		result[1] = digit{charIdx[len(charIdx)-1], charMap[charIdx[len(charIdx)-1]]}
	}
	return result
}
