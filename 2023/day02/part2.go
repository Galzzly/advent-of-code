package main

import (
	"aocli/utils"
	"fmt"
	"strings"
)

func doPartTwo(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var res int
	for _, line := range lines {
		subs := strings.Split(strings.Split(line, ": ")[1], "; ")
		for _, sub := range subs {
			var r, g, b int
			for _, s := range strings.Split(sub, ", ") {
				var n int
				var C string
				fmt.Sscanf(s, "%d %s", &n, &C)
				switch C {
				case "red":
					r = utils.Ter(n > r, n, r)
				case "green":
					g = utils.Ter(n > g, n, g)
				case "blue":
					b = utils.Ter(n > b, n, b)
				}
			}
			res += r * g * b
		}
	}
	return res
}
