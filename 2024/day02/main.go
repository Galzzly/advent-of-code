package main

import (
	_ "embed"
	"os"
)

//go:embed input.txt
var input string

//go:embed input_test.txt
var inputTest string

var successful []int

func main() {
	// Check argv if we use test input or not
	if len(os.Args) > 1 && os.Args[1] == "test" {
		input = inputTest
	}

	successful = make([]int, 0)

	answer := doPartOne(input)
	println(answer)

	answer = doPartTwo(input)
	println(answer)
}

// package main

// import (
// 	"aocli/utils"
// 	"aocli/utils/reader"
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	input := reader.FileLineByLine("input.txt")

// 	getTotalSafeReportCount(input)
// }

// func isReportSafe(reportNum []int) bool {
// 	flagIncrease, flagDecrease := false, false

// 	for i := 1; i < len(reportNum); i++ {
// 		diff := reportNum[i] - reportNum[i-1]

// 		if diff > 0 {
// 			flagIncrease = true
// 		} else if diff < 0 {
// 			flagDecrease = true
// 		} else {
// 			return false
// 		}

// 		if flagDecrease && flagIncrease {
// 			return false
// 		}

// 		if diff > 3 || diff < -3 {
// 			return false
// 		}
// 	}

// 	return true
// }

// func checkReportSafetyWithDeletion(reportNum []int) bool {

// 	for i := 0; i < len(reportNum); i++ {
// 		isSafe := isReportSafeWithDeletion(reportNum, i)
// 		if isSafe {
// 			return true
// 		}
// 	}

// 	return false
// }

// func isReportSafeWithDeletion(report []int, deleteIndex int) bool {
// 	copyReport := make([]int, len(report))
// 	copy(copyReport, report)

// 	if deleteIndex == len(copyReport)-1 {
// 		copyReport = copyReport[:deleteIndex]
// 	} else {
// 		copyReport = append(copyReport[:deleteIndex], copyReport[deleteIndex+1:]...)
// 	}
// 	return isReportSafe(copyReport)
// }

// func getTotalSafeReportCount(reports []string) int {
// 	var count int
// 	var countWithDeletion int
// 	for _, report := range reports {
// 		reportNum := make([]int, 0)
// 		for _, c := range strings.Split(report, " ") {
// 			n := utils.Atoi(c)
// 			reportNum = append(reportNum, n)
// 		}

// 		if isReportSafe(reportNum) {
// 			count++
// 		} else if checkReportSafetyWithDeletion(reportNum) {
// 			countWithDeletion++
// 		}
// 	}
// 	fmt.Printf("answer for part 1: %d\nanswer for part 2: %d\n", count, count+countWithDeletion)
// 	return count
// }
