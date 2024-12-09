package main

import (
	"aocli/utils"
	"slices"
	"strconv"
	"strings"
)

func doPartOne(input string) int {
	mapper := make([]position, 0)
	space := make([]position, 0)
	final := []int{}
	var file bool = true
	var fileid int
	var id int
	for _, r := range strings.ReplaceAll(input, "\n", "") {
		S := utils.Atoi(string(r))
		if file {
			filename := strconv.Itoa(fileid)
			for range S {
				final = append(final, fileid)
				mapper = append(mapper, position{pos: id, char: filename, size: 1})
				id++
			}
			fileid++
			file = !file
			continue
		}
		space = append(space, position{pos: id, size: S})
		for range S {
			final = append(final, MAX)
			id++
		}
		file = !file
	}

	slices.Reverse(mapper)
	for _, P := range mapper {
		for i, S := range space {
			if S.pos < P.pos && P.size <= S.size {
				for s := range P.size {
					if final[P.pos+s] != utils.Atoi(P.char) {
						panic(strconv.Itoa(P.pos+s) + " " + P.char)
					}
					final[P.pos+s] = MAX
					final[S.pos+s] = utils.Atoi(P.char)
				}
				space[i] = position{pos: S.pos + P.size, size: S.size - P.size}
				break
			}
		}
	}
	var res int
	for i := range final {
		if final[i] == MAX {
			continue
		}
		res += i * final[i]
	}
	return res
}
