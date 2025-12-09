package main

func doPartOne() int {
	minLoc := int(^uint(0) >> 1) // Max int

	for _, seed := range seeds {
		val := seed
	nextmap:
		for _, m := range maps {
			for _, mapping := range m {
				if mapping.src <= val && val < mapping.src+mapping.size {
					val = mapping.dest + val - mapping.src
					continue nextmap
				}
			}
		}
		if val < minLoc {
			minLoc = val
		}
	}

	return minLoc
}
