package maps

import (
	"aocli/utils"
	"image"
	"strings"
)

func ImagePointSquareBool(max int) (mapping map[image.Point]bool) {
	mapping = make(map[image.Point]bool)
	for x := 0; x < max; x++ {
		for y := 0; y < max; y++ {
			mapping[image.Point{x, y}] = false
		}
	}
	return
}

func MakeImagePointMap(lines []string) (mapping map[image.Point]rune) {
	mapping = make(map[image.Point]rune)
	for y, s := range lines {
		for x, r := range s {
			mapping[image.Point{x, y}] = r
		}
	}
	return
}

func MakeImagePointMapRect(lines []string) (mapping map[image.Point]rune, rect image.Rectangle) {
	mapping = MakeImagePointMap(lines)
	rect = image.Rect(0, 0, len(lines[0]), len(lines))
	return
}

func MakeIntImagePoint(lines []string) (mapping map[image.Point]int) {
	mapping = make(map[image.Point]int)
	for y, s := range lines {
		for x, r := range strings.Split(s, "") {
			mapping[image.Point{x, y}] = utils.Atoi(r)
		}
	}
	return
}

func Adj(p, d image.Point) image.Point {
	return p.Add(d)
}

func MapKey[K, V comparable](m map[K]V, value V) (key K, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}

func CopyMap[K comparable, V any](m map[K]V) (res map[K]V) {
	res = make(map[K]V)
	for k, v := range m {
		res[k] = v
	}
	return
}
