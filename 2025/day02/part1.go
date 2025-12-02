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
			idlen := len(idstring)

			// Must have even length
			if idlen&1 != 0 {
				continue
			}

			mid := idlen >> 1
			// Compare halves byte by byte for better performance
			match := true
			for i := 0; i < mid; i++ {
				if idstring[i] != idstring[mid+i] {
					match = false
					break
				}
			}
			if match {
				ans += id
			}
		}
	}
	return ans
}
