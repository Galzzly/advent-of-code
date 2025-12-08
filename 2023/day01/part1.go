package main

import (
	"strings"
)

func doPartOne(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	res := 0

	for _, line := range lines {
		var first, last byte
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				if first == 0 {
					first = line[i]
				}
				last = line[i]
			}
		}
		if first != 0 {
			res += int(first-'0')*10 + int(last-'0')
		}
	}

	return res
}
