package main

import (
	"aocli/utils"
	_ "embed"
	"os"
	"strings"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var inputTest string

func main() {
	// Check argv if we use test input or not
	if len(os.Args) > 1 && os.Args[1] == "test" {
		input = inputTest
	}

	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}

type Maps map[rune]map[string]string

func solve(input string, suffix string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n\n")
	instr := []rune(lines[0])
	maps := buildMaps(lines[1])
	P := []string{}
	for p := range maps['L'] {
		if strings.HasSuffix(p, suffix) {
			P = append(P, p)
		}
	}

	I := []int{}
	i := 0
	for {
		NP := []string{}
		for _, p := range P {
			p = maps[instr[i%len(instr)]][p]
			if strings.HasSuffix(p, "Z") {
				I = append(I, i+1)
				if len(I) == len(P) {
					return utils.LCM(1, I[0], I[1:]...)
				}
			}
			NP = append(NP, p)
		}
		P = NP
		i++
	}
}

func buildMaps(line string) Maps {
	lines := strings.Split(line, "\n")
	maps := make(map[rune]map[string]string, 2)
	maps['L'] = make(map[string]string, len(lines))
	maps['R'] = make(map[string]string, len(lines))
	for _, line := range lines {
		s := strings.Split(line, " = ")
		I := s[0]
		lr := strings.Split(s[1], ", ")
		L := lr[0][1:]
		R := lr[1][:3]
		maps['L'][I] = L
		maps['R'][I] = R
	}
	return maps
}
