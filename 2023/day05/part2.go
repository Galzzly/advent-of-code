package main

type Range struct {
	start, end int
}

func doPartTwo() int {
	// Parse seed ranges
	var ranges []Range
	for i := 0; i < len(seeds); i += 2 {
		ranges = append(ranges, Range{seeds[i], seeds[i] + seeds[i+1] - 1})
	}

	minLoc := int(^uint(0) >> 1) // Max int

	for _, r := range ranges {
		result := applyMapsToRange([]Range{r})
		for _, res := range result {
			if res.start < minLoc {
				minLoc = res.start
			}
		}
	}

	return minLoc
}

func applyMapsToRange(ranges []Range) []Range {
	current := ranges

	for _, m := range maps {
		var next []Range

		for _, r := range current {
			unmapped := []Range{r}

			for _, mapping := range m {
				srcEnd := mapping.src + mapping.size - 1
				var stillUnmapped []Range

				for _, u := range unmapped {
					// Split range into before, intersection, after
					if u.end < mapping.src || u.start > srcEnd {
						// No overlap
						stillUnmapped = append(stillUnmapped, u)
						continue
					}

					// Before intersection
					if u.start < mapping.src {
						stillUnmapped = append(stillUnmapped, Range{u.start, mapping.src - 1})
					}

					// Intersection - map it
					interStart := max(u.start, mapping.src)
					interEnd := min(u.end, srcEnd)
					offset := mapping.dest - mapping.src
					next = append(next, Range{interStart + offset, interEnd + offset})

					// After intersection
					if u.end > srcEnd {
						stillUnmapped = append(stillUnmapped, Range{srcEnd + 1, u.end})
					}
				}

				unmapped = stillUnmapped
			}

			// Add unmapped ranges as-is
			next = append(next, unmapped...)
		}

		current = next
	}

	return current
}
