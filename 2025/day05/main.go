package main

import (
	"aocli/utils"
	_ "embed"
	"fmt"
	"os"
	"slices"
	"strings"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var inputTest string

type fresh struct {
	start, end int
}

var ranges []fresh
var sections []string

func main() {
	// Check argv if we use test input or not
	if len(os.Args) > 1 && os.Args[1] == "test" {
		input = inputTest
	}

	// Parse input once
	sections = strings.Split(strings.TrimSpace(input), "\n\n")
	parseRanges()

	answer := doPartOne()
	println(answer)

	answer = doPartTwo()
	println(answer)
}

func parseRanges() {
	rangeLines := strings.Split(sections[0], "\n")
	ranges = make([]fresh, len(rangeLines))
	for i, v := range rangeLines {
		fmt.Sscanf(v, "%d-%d", &ranges[i].start, &ranges[i].end)
	}
}

func mergeRanges() []fresh {
	// Sort and merge overlapping ranges
	sorted := make([]fresh, len(ranges))
	copy(sorted, ranges)

	slices.SortFunc(sorted, func(a, b fresh) int {
		return a.start - b.start
	})

	merged := make([]fresh, 0, len(sorted))
	current := sorted[0]

	for i := 1; i < len(sorted); i++ {
		if sorted[i].start <= current.end+1 {
			// Overlapping or adjacent, merge them
			current.end = utils.Biggest(current.end, sorted[i].end)
		} else {
			// No overlap, save current and start new
			merged = append(merged, current)
			current = sorted[i]
		}
	}
	merged = append(merged, current)

	return merged
}
