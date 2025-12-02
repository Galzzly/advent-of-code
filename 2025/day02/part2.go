package main

import (
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
		for id := id1; id <= id2; id++ {
			idstring := strconv.Itoa(id)
			idlen := len(idstring)

			// Check if the number can be split into repeating chunks
			for chunkSize := 1; chunkSize <= idlen/2; chunkSize++ {
				// Only check divisors of the length
				if idlen%chunkSize != 0 {
					continue
				}

				// Check if all chunks match the first chunk
				match := true
				for pos := chunkSize; pos < idlen; pos += chunkSize {
					for offset := 0; offset < chunkSize; offset++ {
						if idstring[offset] != idstring[pos+offset] {
							match = false
							break
						}
					}
					if !match {
						break
					}
				}

				if match {
					ans += id
					break
				}
			}
		}
	}
	return ans
}
