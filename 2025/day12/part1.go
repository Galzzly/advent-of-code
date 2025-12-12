package main

import (
	"aocli/utils"
	"fmt"
	"strings"
)

func doPartOne(input string) int {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")

	// Precompute present counts (count '#' symbols in each present pattern)
	numPresents := len(parts) - 1
	presents := make([]int, numPresents)
	for i := 0; i < numPresents; i++ {
		presents[i] = strings.Count(parts[i], "#")
	}

	ans := 0
	lines := strings.Split(parts[numPresents], "\n")

	for _, line := range lines {
		// Parse dimensions and gift counts
		colonIdx := strings.Index(line, ": ")
		if colonIdx == -1 {
			continue
		}

		var x, y int
		fmt.Sscanf(line[:colonIdx], "%dx%d", &x, &y)
		area := x * y

		// Calculate total present area
		totalpresentarea := 0
		giftStart := colonIdx + 2
		for i := 0; giftStart < len(line); i++ {
			// Parse next number
			spaceIdx := strings.IndexByte(line[giftStart:], ' ')
			var numStr string
			if spaceIdx == -1 {
				numStr = line[giftStart:]
				giftStart = len(line)
			} else {
				numStr = line[giftStart : giftStart+spaceIdx]
				giftStart += spaceIdx + 1
			}
			totalpresentarea += utils.Atoi(numStr) * presents[i]
		}

		// Check if sleigh has enough space
		if totalpresentarea <= area && area > totalpresentarea+(totalpresentarea/5) {
			ans++
		}
	}

	return ans
}
