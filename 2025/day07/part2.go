package main

import (
	"image"
	"strings"
)

type MapRect struct {
	m     map[image.Point]rune
	r     image.Rectangle
	cache map[image.Point]int
}

func doPartTwo(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	mapper, rect, S := makeImagePointMapRect(lines)
	m := MapRect{
		m:     mapper,
		r:     rect,
		cache: make(map[image.Point]int),
	}

	return m.score(S)
}

func (m MapRect) score(i image.Point) int {
	// Check cache first
	if val, ok := m.cache[i]; ok {
		return val
	}

	np := i.Add(image.Point{0, 1})
	if !np.In(m.r) {
		m.cache[i] = 1
		return 1
	}

	var result int
	if m.m[np] == '^' {
		result = m.score(np.Add(image.Point{-1, 0})) + m.score(np.Add(image.Point{1, 0}))
	} else {
		result = m.score(np)
	}

	m.cache[i] = result
	return result
}
