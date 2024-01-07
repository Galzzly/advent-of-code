package main

import (
	"aocli/utils"
	"fmt"
	"strings"
)

type game struct {
	id  int
	rgb []RGB
}
type RGB struct {
	red   int
	green int
	blue  int
}

func doPartOne(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var total int
	var games []game
	for _, line := range lines {
		var rgb []RGB
		var id int
		var restofline string
		s := strings.Split(line, ": ")
		id = utils.Atoi(strings.Split(s[0], " ")[1])
		restofline = s[1]
		subset := strings.Split(restofline, "; ")
		for _, sub := range subset {
			var colours RGB
			s := strings.Split(sub, ", ")
			for _, s := range s {
				var num int
				var colour string
				fmt.Sscanf(s, "%d %s", &num, &colour)
				switch colour {
				case "red":
					colours.red += num
				case "green":
					colours.green += num
				case "blue":
					colours.blue += num
				}
				rgb = append(rgb, colours)
			}
		}
		total += id
		games = append(games, game{id, rgb})
	}

	var res int
	red := 12
	green := 13
	blue := 14
	res = total
nextgame:
	for _, game := range games {
		for _, rgb := range game.rgb {
			if rgb.red > red ||
				rgb.green > green ||
				rgb.blue > blue {
				res -= game.id
				continue nextgame
			}
		}
	}
	return res
}
