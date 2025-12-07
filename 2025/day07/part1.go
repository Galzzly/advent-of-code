package main

import (
	"image"
	"strings"
)

func doPartOne(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	mapper, _, S := makeImagePointMapRect(lines)

	// Pre-allocate with reasonable capacity
	seen := make(map[image.Point]bool, len(lines)*len(lines[0])/4)
	queue := make([]image.Point, 0, 256)
	queue = append(queue, S)
	ans := 0
	head := 0

	for head < len(queue) {
		P := queue[head]
		head++

		if seen[P] {
			continue
		}
		seen[P] = true

		np := P.Add(image.Point{0, 1})
		val, ok := mapper[np]
		if !ok {
			continue
		}

		if val != '^' {
			queue = append(queue, np)
			continue
		}

		ans++
		queue = append(queue, np.Add(image.Point{-1, 0}), np.Add(image.Point{1, 0}))
	}

	return ans
}

func makeImagePointMapRect(lines []string) (mapping map[image.Point]rune, rect image.Rectangle, S image.Point) {
	mapping = make(map[image.Point]rune)
	for y, s := range lines {
		for x, r := range s {
			mapping[image.Point{x, y}] = r
			if r == 'S' {
				S = image.Point{x, y}
			}
		}
	}
	rect = image.Rect(0, 0, len(lines[0]), len(lines))
	return
}
