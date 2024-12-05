package reader

import (
	"aocli/utils"
	"os"
	"strings"
)

func ReadFile(file string) []byte {
	f, err := os.ReadFile(file)
	utils.Check(err)
	return f
}

func FileLineByLine(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func FileSingleLine(file string) string {
	f := ReadFile(file)
	return strings.Split(strings.TrimSpace(string(f)), "\n")[0]
}

func FileDoubleLine(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n\n")
}

func IntsByLine(file string) (nums []int) {
	f := ReadFile(file)
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	nums = make([]int, 0, len(lines))
	for l := range lines {
		if len(lines[l]) == 0 {
			continue
		}
		n := utils.Atoi(lines[l])
		nums = append(nums, n)
	}
	return
}

func LineByComma(file string) []string {
	f := ReadFile(file)
	return strings.Split(strings.TrimSpace(string(f)), ",")
}

func IntsLineByComma(file string) (nums []int) {
	s := LineByComma(file)
	nums = make([]int, 0, len(s))
	for _, n := range s {
		nums = append(nums, utils.Atoi(n))
	}
	return
}

func RowIntsByLine(input string) (nums [][]int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	nums = make([][]int, len(lines))
	for l := range lines {
		if len(lines[l]) == 0 {
			continue
		}
		s := strings.Split(lines[l], " ")
		nums[l] = make([]int, 0, len(s))
		for i := range s {
			nums[l] = append(nums[l], utils.Atoi(s[i]))
		}
	}
	return
}
