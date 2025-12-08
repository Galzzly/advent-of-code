package main

import (
	"strings"
)

var wordDigits = []struct {
	word string
	val  int
}{
	{"one", 1}, {"two", 2}, {"three", 3}, {"four", 4}, {"five", 5},
	{"six", 6}, {"seven", 7}, {"eight", 8}, {"nine", 9},
}

func doPartTwo(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	res := 0

	for _, line := range lines {
		first, last := findFirstAndLast(line)
		res += first*10 + last
	}

	return res
}

func findFirstAndLast(line string) (int, int) {
	first, last := 0, 0
	firstPos, lastPos := len(line), -1

	// Check for digit characters
	for i := 0; i < len(line); i++ {
		if line[i] >= '0' && line[i] <= '9' {
			val := int(line[i] - '0')
			if i < firstPos {
				firstPos = i
				first = val
			}
			if i > lastPos {
				lastPos = i
				last = val
			}
		}
	}

	// Check for word digits
	for _, wd := range wordDigits {
		if idx := strings.Index(line, wd.word); idx != -1 && idx < firstPos {
			firstPos = idx
			first = wd.val
		}
		if idx := strings.LastIndex(line, wd.word); idx != -1 && idx > lastPos {
			lastPos = idx
			last = wd.val
		}
	}

	return first, last
}
