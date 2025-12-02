package main

import (
	"fmt"
	"strconv"
	"strings"
)

func doPartOne(input string) int {
	lines := strings.Split(strings.Split(strings.TrimSpace(input), "\n")[0], ",")
	ans := 0
	for _, line := range lines {
		var id1, id2 int
		fmt.Sscanf(line, "%d-%d", &id1, &id2)
		for id := id1; id <= id2; id++ {
			idstring := strconv.Itoa(id)
			if len(idstring)%2 != 0 {
				continue
			}
			mid := len(idstring) / 2
			if idstring[:mid] == idstring[mid:] {
				ans += id
			}
		}
	}
	return ans
}
