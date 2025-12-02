package main

import (
	"aocli/utils"
	"fmt"
	"strconv"
	"strings"
)

func doPartTwo(input string) int {
	lines := strings.Split(strings.Split(strings.TrimSpace(input), "\n")[0], ",")
	ans := 0
	for _, line := range lines {
		var id1, id2 int
		fmt.Sscanf(line, "%d-%d", &id1, &id2)
	nextid:
		for id := id1; id <= id2; id++ {
			idstring := strconv.Itoa(id)
			mid := len(idstring) / 2
			for i := 1; i <= mid; i++ {
				if len(idstring)%i != 0 {
					continue
				}
				chunks := utils.ChunkSlice([]rune(idstring), i)
				match := true
				for i := 1; i < len(chunks); i++ {
					if string(chunks[i]) != string(chunks[0]) {
						match = false
						break
					}
				}
				if match {
					ans += id
					continue nextid
				}
			}
		}
	}
	return ans
}
